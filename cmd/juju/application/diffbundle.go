// Copyright 2018 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package application

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/juju/errors"
	"github.com/juju/gnuflag"
	"gopkg.in/yaml.v2"

	"github.com/juju/juju/api/base"
	"github.com/juju/juju/api/client/annotations"
	"github.com/juju/juju/api/client/application"
	apicharms "github.com/juju/juju/api/client/charms"
	"github.com/juju/juju/api/client/modelconfig"
	commoncharm "github.com/juju/juju/api/common/charm"
	jujucmd "github.com/juju/juju/cmd"
	appbundle "github.com/juju/juju/cmd/juju/application/bundle"
	"github.com/juju/juju/cmd/juju/application/store"
	"github.com/juju/juju/cmd/juju/application/utils"
	"github.com/juju/juju/cmd/modelcmd"
	"github.com/juju/juju/core/arch"
	corebase "github.com/juju/juju/core/base"
	"github.com/juju/juju/core/constraints"
	"github.com/juju/juju/environs/config"
	bundlechanges "github.com/juju/juju/internal/bundle/changes"
	"github.com/juju/juju/internal/charm"
	"github.com/juju/juju/internal/charmhub"
	"github.com/juju/juju/internal/cmd"
	"github.com/juju/juju/rpc/params"
)

const (
	// ArchAll defines a platform that targets all architectures.
	ArchAll = "all"
)

// ModelConfigGetter defines an interface for getting model configuration.
type ModelConfigGetter interface {
	ModelGet(ctx context.Context) (map[string]interface{}, error)
}

// ModelConstraintsGetter defines an interface for getting model constraints.
type ModelConstraintsGetter interface {
	GetModelConstraints(ctx context.Context) (constraints.Value, error)
}

// ModelConfigClient represents a model config client for requesting model
// configurations.
type ModelConfigClient interface {
	ModelConfigGetter
	Close() error
}

// ModelConstraintsClient represents a client for getting the constraints from
// a model.
type ModelConstraintsClient interface {
	ModelConstraintsGetter
	Close() error
}

const bundleDiffDoc = `
Bundle can be a local bundle file or the name of a bundle in
Charmhub. The bundle can also be combined with overlays (in the
same way as the deploy command) before comparing with the model.

The map-machines option works similarly as for the deploy command, but
existing is always assumed, so it doesn't need to be specified.

Config values for comparison are always source from the "current" model.

Specifying a base will retrieve the bundle for the relevant store for
the give base.
`

const bundleDiffExamples = `
    juju diff-bundle localbundle.yaml
    juju diff-bundle charmed-kubernetes
    juju diff-bundle charmed-kubernetes --overlay local-config.yaml --overlay extra.yaml
	juju diff-bundle charmed-kubernetes --base ubuntu@22.04
    juju diff-bundle -m othermodel hadoop-spark
    juju diff-bundle localbundle.yaml --map-machines 3=4
`

// NewDiffBundleCommand returns a command to compare a bundle against
// the selected model.
func NewDiffBundleCommand() cmd.Command {
	command := &diffBundleCommand{
		arches: arch.AllArches(),
	}
	command.charmAdaptorFn = command.charmAdaptor
	command.newAPIRootFn = func(ctx context.Context) (base.APICallCloser, error) {
		return command.NewAPIRoot(ctx)
	}
	command.modelConfigClientFunc = func(api base.APICallCloser) ModelConfigClient {
		return modelconfig.NewClient(api)
	}
	command.modelConstraintsClientFunc = func(ctx context.Context) (ModelConstraintsClient, error) {
		root, err := command.NewAPIRoot(ctx)
		if err != nil {
			return nil, errors.Trace(err)
		}
		client := modelconfig.NewClient(root)
		return client, nil
	}
	return modelcmd.Wrap(command)
}

// diffBundleCommand compares a bundle to a model.
type diffBundleCommand struct {
	modelcmd.ModelCommandBase
	bundle         string
	bundleOverlays []string
	channelStr     string
	channel        charm.Channel
	arch           string
	arches         arch.Arches
	base           string
	annotations    bool

	bundleMachines map[string]string
	machineMap     string

	charmAdaptorFn             func(base.APICallCloser, *charm.URL) (BundleResolver, error)
	newAPIRootFn               func(ctx context.Context) (base.APICallCloser, error)
	modelConfigClientFunc      func(base.APICallCloser) ModelConfigClient
	modelConstraintsClientFunc func(ctx context.Context) (ModelConstraintsClient, error)
}

// IsSuperCommand is part of cmd.Command.
func (c *diffBundleCommand) IsSuperCommand() bool { return false }

// AllowInterspersedFlags is part of cmd.Command.
func (c *diffBundleCommand) AllowInterspersedFlags() bool { return true }

// Info is part of cmd.Command.
func (c *diffBundleCommand) Info() *cmd.Info {
	return jujucmd.Info(&cmd.Info{
		Name:     "diff-bundle",
		Args:     "<bundle file or name>",
		Purpose:  "Compare a bundle with a model and report any differences.",
		Doc:      bundleDiffDoc,
		Examples: bundleDiffExamples,
		SeeAlso: []string{
			"deploy",
		},
	})
}

// SetFlags is part of cmd.Command.
func (c *diffBundleCommand) SetFlags(f *gnuflag.FlagSet) {
	c.ModelCommandBase.SetFlags(f)

	f.StringVar(&c.arch, "arch", "", fmt.Sprintf("specify an arch <%s>", c.archArgumentList()))
	f.StringVar(&c.base, "base", "", "specify a base")
	f.StringVar(&c.channelStr, "channel", "", "Channel to use when getting the bundle from Charmhub")
	f.Var(cmd.NewAppendStringsValue(&c.bundleOverlays), "overlay", "Bundles to overlay on the primary bundle, applied in order")
	f.StringVar(&c.machineMap, "map-machines", "", "Indicates how existing machines correspond to bundle machines")
	f.BoolVar(&c.annotations, "annotations", false, "Include differences in annotations")
}

// Init is part of cmd.Command.
func (c *diffBundleCommand) Init(args []string) error {
	if len(args) < 1 {
		return errors.New("no bundle specified")
	}
	c.bundle = args[0]
	// UseExisting is assumed for diffing.
	_, mapping, err := parseMachineMap(c.machineMap)
	if err != nil {
		return errors.Annotate(err, "error in --map-machines")
	}
	c.bundleMachines = mapping
	if c.channelStr != "" {
		c.channel, err = charm.ParseChannelNormalize(c.channelStr)
		if err != nil {
			return errors.Annotate(err, "error in --channel")
		}
	}

	if c.arch != "" && !c.arches.Contains(c.arch) {
		return errors.Errorf("unexpected architecture flag value %q, expected <%s>", c.arch, c.archArgumentList())
	}
	return cmd.CheckEmpty(args[1:])
}

// Run is part of cmd.Command.
func (c *diffBundleCommand) Run(ctx *cmd.Context) error {
	var (
		base corebase.Base
		err  error
	)
	if c.base != "" {
		if base, err = corebase.ParseBaseFromString(c.base); err != nil {
			return errors.Trace(err)
		}
	}

	apiRoot, err := c.newAPIRootFn(ctx)
	if err != nil {
		return errors.Trace(err)
	}
	defer func() { _ = apiRoot.Close() }()

	// Load up the bundle data, with includes and overlays.
	baseSrc, err := c.bundleDataSource(ctx, apiRoot, base)
	if err != nil {
		return errors.Trace(err)
	}

	bundle, _, err := appbundle.ComposeAndVerifyBundle(ctx, baseSrc, c.bundleOverlays)
	if err != nil {
		return errors.Trace(err)
	}

	if err = c.warnForMissingRelationEndpoints(ctx, bundle); err != nil {
		return errors.Trace(err)
	}

	// Extract the information from the current model.
	model, err := c.readModel(ctx, apiRoot)
	if err != nil {
		return errors.Trace(err)
	}
	// Get the differences between them.
	diff, err := bundlechanges.BuildDiff(bundlechanges.DiffConfig{
		Bundle:             bundle,
		Model:              model,
		Logger:             logger,
		IncludeAnnotations: c.annotations,
	})

	if err != nil {
		return errors.Trace(err)
	}

	encoder := yaml.NewEncoder(ctx.Stdout)
	defer func() { _ = encoder.Close() }()
	err = encoder.Encode(diff)
	if err != nil {
		return errors.Trace(err)
	}
	return nil
}

func (c *diffBundleCommand) warnForMissingRelationEndpoints(ctx *cmd.Context, bundle *charm.BundleData) error {
	var missing []string
	for _, relPair := range bundle.Relations {
		if len(relPair) != 2 {
			return errors.Errorf("malformed relation %v", relPair)
		}

		if missingRelationEndpoint(relPair[0]) || missingRelationEndpoint(relPair[1]) {
			missing = append(missing, fmt.Sprintf("[%s, %s]", relPair[0], relPair[1]))
		}
	}

	if len(missing) == 0 {
		return nil
	}

	logger.Warningf(context.TODO(),
		"The provided bundle includes relations without explicit endpoints, which may appear as extra entries in the diff output.\nTo avoid this in the future, update the endpoints for the following bundle relations:\n - %s",
		strings.Join(missing, "\n - "),
	)

	// Add an extra blank line to separate the diff output from the warning
	// and avoid confusion.
	_, _ = fmt.Fprintln(ctx.Stderr)

	return nil
}

func missingRelationEndpoint(rel string) bool {
	tokens := strings.SplitN(rel, ":", 2)
	return len(tokens) != 2 || tokens[1] == ""
}

func (c *diffBundleCommand) bundleDataSource(ctx *cmd.Context, apiRoot base.APICallCloser, base corebase.Base) (charm.BundleDataSource, error) {
	ds, err := charm.LocalBundleDataSource(c.bundle)

	// NotFound means that the provided local file is not found, and
	// therefore we should try interpreting it as a charm store bundle URL.
	if err != nil && !errors.Is(err, errors.NotFound) {
		return nil, errors.Trace(err)
	}
	if ds != nil {
		return ds, nil
	}

	modelConstraints, err := c.getModelConstraints(ctx)
	if err != nil {
		return nil, errors.Trace(err)
	}

	bURL, err := charm.ParseURL(c.bundle)
	if err != nil {
		return nil, errors.Trace(err)
	}
	platform := utils.MakePlatform(constraints.Value{
		Arch: &c.arch,
	}, base, modelConstraints)
	// We return early with local bundles, so here we know the bundle must be from charmhub
	origin, err := utils.MakeOrigin(charm.CharmHub, bURL.Revision, c.channel, platform)
	if err != nil {
		return nil, errors.Trace(err)
	}
	charmAdaptor, err := c.charmAdaptorFn(apiRoot, bURL)
	if err != nil {
		return nil, errors.Trace(err)
	}

	bundleURL, bundleOrigin, err := charmAdaptor.ResolveBundleURL(ctx, bURL, origin)
	if err != nil {
		if errors.Is(err, errors.NotValid) {
			ctx.Verbosef("%q can not be found or is not a valid bundle", c.bundle)
		}
		return nil, errors.Trace(err)
	}
	if bundleURL == nil {
		return nil, errors.Errorf("couldn't interpret %q as a local bundle", c.bundle)
	}

	// GetBundle creates the directory so we actually want to create a temp
	// directory then add a namespace (bundle name) so that charmhub get
	// bundle can create it.
	dir, err := os.MkdirTemp("", "diff-bundle-")
	if err != nil {
		return nil, errors.Trace(err)
	}
	bundlePath := filepath.Join(dir, bundleURL.Name)
	bundle, err := charmAdaptor.GetBundle(ctx, bundleURL, bundleOrigin, bundlePath)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return store.NewResolvedBundle(bundle), nil
}

func (c *diffBundleCommand) charmAdaptor(apiRoot base.APICallCloser, curl *charm.URL) (BundleResolver, error) {
	downloadClient := func(ctx context.Context) (store.DownloadBundleClient, error) {
		apiRoot, err := c.newAPIRootFn(ctx)
		if err != nil {
			return nil, errors.Trace(err)
		}

		url, err := c.getCharmHubURL(ctx, apiRoot)
		if err != nil {
			return nil, errors.Trace(err)
		}
		return charmhub.NewClient(charmhub.Config{
			URL:    url,
			Logger: logger,
		})
	}

	return store.NewCharmAdaptor(apicharms.NewClient(apiRoot), downloadClient), nil
}

func (c *diffBundleCommand) readModel(ctx context.Context, apiRoot base.APICallCloser) (*bundlechanges.Model, error) {
	status, err := c.getStatus(ctx, apiRoot)
	if err != nil {
		return nil, errors.Annotate(err, "getting model status")
	}
	model, err := appbundle.BuildModelRepresentation(ctx, status, c.makeModelExtractor(apiRoot), true, c.bundleMachines)
	return model, errors.Trace(err)
}

func (c *diffBundleCommand) getStatus(ctx context.Context, apiRoot base.APICallCloser) (*params.FullStatus, error) {
	// Ported from api.Client which is nigh impossible to test without
	// a real api.Connection.
	_, facade := base.NewClientFacade(apiRoot, "Client")
	var result params.FullStatus
	if err := facade.FacadeCall(ctx, "FullStatus", params.StatusParams{}, &result); err != nil {
		return nil, errors.Trace(err)
	}
	// We don't care about model type.
	return &result, nil
}

func (c *diffBundleCommand) makeModelExtractor(apiRoot base.APICallCloser) appbundle.ModelExtractor {
	return &extractorImpl{
		application: application.NewClient(apiRoot),
		annotations: annotations.NewClient(apiRoot),
		modelConfig: modelconfig.NewClient(apiRoot),
	}
}

func (c *diffBundleCommand) archArgumentList() string {
	archList := strings.Join(c.arches.StringList(), "|")
	return fmt.Sprintf("%s|%s", ArchAll, archList)
}

func (c *diffBundleCommand) getCharmHubURL(ctx context.Context, apiRoot base.APICallCloser) (string, error) {
	modelConfigClient := c.modelConfigClientFunc(apiRoot)
	defer func() { _ = modelConfigClient.Close() }()

	attrs, err := modelConfigClient.ModelGet(ctx)
	if err != nil {
		return "", errors.Trace(err)
	}

	conf, err := config.New(config.NoDefaults, attrs)
	if err != nil {
		return "", errors.Trace(err)
	}

	charmHubURL, _ := conf.CharmHubURL()
	return charmHubURL, nil
}

func (c *diffBundleCommand) getModelConstraints(ctx context.Context) (constraints.Value, error) {
	modelConsClient, err := c.modelConstraintsClientFunc(ctx)
	if err != nil {
		return constraints.Value{}, errors.Trace(err)
	}
	defer func() { _ = modelConsClient.Close() }()
	return modelConsClient.GetModelConstraints(ctx)
}

type extractorImpl struct {
	application *application.Client
	annotations *annotations.Client
	modelConfig *modelconfig.Client
}

// GetAnnotations is part of ModelExtractor.
func (e *extractorImpl) GetAnnotations(ctx context.Context, tags []string) ([]params.AnnotationsGetResult, error) {
	return e.annotations.Get(ctx, tags)
}

// GetConstraints is part of ModelExtractor.
func (e *extractorImpl) GetConstraints(ctx context.Context, applications ...string) ([]constraints.Value, error) {
	return e.application.GetConstraints(ctx, applications...)
}

// GetConfig is part of ModelExtractor.
func (e *extractorImpl) GetConfig(ctx context.Context, applications ...string) ([]map[string]interface{}, error) {
	return e.application.GetConfig(ctx, applications...)
}

// Sequences is part of ModelExtractor.
func (e *extractorImpl) Sequences(ctx context.Context) (map[string]int, error) {
	return e.modelConfig.Sequences(ctx)
}

// BundleResolver defines what we need from a charm store to resolve a
// bundle and read the bundle data.
type BundleResolver interface {
	ResolveBundleURL(context.Context, *charm.URL, commoncharm.Origin) (*charm.URL, commoncharm.Origin, error)
	GetBundle(context.Context, *charm.URL, commoncharm.Origin, string) (charm.Bundle, error)
}
