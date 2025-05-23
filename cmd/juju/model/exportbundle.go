// Copyright 2018 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package model

import (
	"context"
	"fmt"
	"os"

	"github.com/juju/errors"
	"github.com/juju/gnuflag"

	"github.com/juju/juju/api/client/bundle"
	jujucmd "github.com/juju/juju/cmd"
	"github.com/juju/juju/cmd/modelcmd"
	"github.com/juju/juju/internal/cmd"
)

// NewExportBundleCommand returns a fully constructed export bundle command.
func NewExportBundleCommand() cmd.Command {
	command := &exportBundleCommand{}
	command.newAPIFunc = func(ctx context.Context) (ExportBundleAPI, error) {
		return command.getAPIs(ctx)
	}
	return modelcmd.Wrap(command)
}

type exportBundleCommand struct {
	modelcmd.ModelCommandBase
	newAPIFunc           func(ctx context.Context) (ExportBundleAPI, error)
	Filename             string
	includeCharmDefaults bool
}

const exportBundleHelpDoc = `
Exports the current model configuration as a reusable bundle.

If --filename is not used, the configuration is printed to stdout.
 --filename specifies an output file.
`

const exportBundleHelpExamples = `
    juju export-bundle
    juju export-bundle --filename mymodel.yaml
    juju export-bundle --include-charm-defaults
`

// Info implements Command.
func (c *exportBundleCommand) Info() *cmd.Info {
	return jujucmd.Info(&cmd.Info{
		Name:     "export-bundle",
		Purpose:  "Exports the current model configuration as a reusable bundle.",
		Doc:      exportBundleHelpDoc,
		Examples: exportBundleHelpExamples,
	})
}

// SetFlags implements Command.
func (c *exportBundleCommand) SetFlags(f *gnuflag.FlagSet) {
	c.ModelCommandBase.SetFlags(f)
	f.StringVar(&c.Filename, "filename", "", "Bundle file")
	f.BoolVar(&c.includeCharmDefaults, "include-charm-defaults", false, "Whether to include charm config default values in the exported bundle")
}

// Init implements Command.
func (c *exportBundleCommand) Init(args []string) error {
	return cmd.CheckEmpty(args)
}

// ExportBundleAPI specifies the used function calls of the BundleFacade.
type ExportBundleAPI interface {
	Close() error
	ExportBundle(ctx context.Context, includeCharmDefaults bool) (string, error)
}

func (c *exportBundleCommand) getAPIs(ctx context.Context) (ExportBundleAPI, error) {
	api, err := c.NewAPIRoot(ctx)
	if err != nil {
		return nil, err
	}

	return bundle.NewClient(api), nil
}

// Run implements Command.
func (c *exportBundleCommand) Run(ctx *cmd.Context) error {
	bundleClient, err := c.newAPIFunc(ctx)
	if err != nil {
		return err
	}
	defer bundleClient.Close()

	result, err := bundleClient.ExportBundle(ctx, c.includeCharmDefaults)
	if err != nil {
		return err
	}

	if c.Filename == "" {
		_, err := fmt.Fprintf(ctx.Stdout, "%v", result)
		return err
	}
	filename := c.Filename
	file, err := c.Filesystem().OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0600)
	if err != nil {
		return errors.Annotate(err, "while creating local file")
	}
	defer file.Close()

	_, err = file.WriteString(result)
	if err != nil {
		return errors.Annotate(err, "while copying in local file")
	}

	fmt.Fprintln(ctx.Stdout, "Bundle successfully exported to", filename)

	return nil
}
