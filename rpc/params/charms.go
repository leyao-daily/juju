// Copyright 2015 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package params

import (
	"github.com/juju/juju/internal/charm"
	"github.com/juju/juju/internal/charm/assumes"
)

// ApplicationCharmResults contains a set of ApplicationCharmResults.
type ApplicationCharmResults struct {
	Results []ApplicationCharmResult `json:"results"`
}

// ApplicationCharmResult contains an ApplicationCharm or an error.
type ApplicationCharmResult struct {
	Result *ApplicationCharm `json:"result,omitempty"`
	Error  *Error            `json:"error,omitempty"`
}

// ApplicationCharm contains information about an application's charm.
type ApplicationCharm struct {
	// URL holds the URL of the charm assigned to the
	// application.
	URL string `json:"url"`

	// ForceUpgrade indicates whether or not application
	// units should upgrade to the charm even if they
	// are in an error state.
	ForceUpgrade bool `json:"force-upgrade,omitempty"`

	// SHA256 holds the SHA256 hash of the charm archive.
	SHA256 string `json:"sha256"`

	// CharmModifiedVersion increases when the charm changes in some way.
	CharmModifiedVersion int `json:"charm-modified-version"`

	// DeploymentMode is either "operator" or "workload"
	DeploymentMode string `json:"deployment-mode,omitempty"`
}

// CharmsList stores parameters for a charms.List call
type CharmsList struct {
	Names []string `json:"names"`
}

// CharmsListResult stores result from a charms.List call
type CharmsListResult struct {
	CharmURLs []string `json:"charm-urls"`
}

// CharmOption mirrors charm.Option
type CharmOption struct {
	Type        string      `json:"type"`
	Description string      `json:"description,omitempty"`
	Default     interface{} `json:"default,omitempty"`
}

// CharmRelation mirrors charm.Relation.
type CharmRelation struct {
	Name      string `json:"name"`
	Role      string `json:"role"`
	Interface string `json:"interface"`
	Optional  bool   `json:"optional"`
	Limit     int    `json:"limit"`
	Scope     string `json:"scope"`
}

// CharmStorage mirrors charm.Storage.
type CharmStorage struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Type        string   `json:"type"`
	Shared      bool     `json:"shared"`
	ReadOnly    bool     `json:"read-only"`
	CountMin    int      `json:"count-min"`
	CountMax    int      `json:"count-max"`
	MinimumSize uint64   `json:"minimum-size"`
	Location    string   `json:"location,omitempty"`
	Properties  []string `json:"properties,omitempty"`
}

// CharmDevice mirrors charm.Device.
type CharmDevice struct {
	Name        string `bson:"name"`
	Description string `bson:"description"`
	Type        string `bson:"type"`
	CountMin    int64  `bson:"count-min"`
	CountMax    int64  `bson:"count-max"`
}

// CharmResourceMeta mirrors charm.ResourceMeta.
type CharmResourceMeta struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Path        string `json:"path"`
	Description string `json:"description"`
}

// CharmMeta mirrors charm.Meta.
type CharmMeta struct {
	Name           string                       `json:"name"`
	Summary        string                       `json:"summary"`
	Description    string                       `json:"description"`
	Subordinate    bool                         `json:"subordinate"`
	Provides       map[string]CharmRelation     `json:"provides,omitempty"`
	Requires       map[string]CharmRelation     `json:"requires,omitempty"`
	Peers          map[string]CharmRelation     `json:"peers,omitempty"`
	ExtraBindings  map[string]string            `json:"extra-bindings,omitempty"`
	Categories     []string                     `json:"categories,omitempty"`
	Tags           []string                     `json:"tags,omitempty"`
	Storage        map[string]CharmStorage      `json:"storage,omitempty"`
	Devices        map[string]CharmDevice       `json:"devices,omitempty"`
	Resources      map[string]CharmResourceMeta `json:"resources,omitempty"`
	Terms          []string                     `json:"terms,omitempty"`
	MinJujuVersion string                       `json:"min-juju-version,omitempty"`
	Containers     map[string]CharmContainer    `json:"containers,omitempty"`
	AssumesExpr    *assumes.ExpressionTree      `json:"assumes-expr,omitempty"`
	CharmUser      string                       `json:"charm-user,omitempty"`
}

// Charm holds all the charm data that the client needs.
// To be honest, it probably returns way more than what is actually needed.
type Charm struct {
	Revision   int                    `json:"revision"`
	URL        string                 `json:"url"`
	Config     map[string]CharmOption `json:"config"`
	Meta       *CharmMeta             `json:"meta,omitempty"`
	Actions    *CharmActions          `json:"actions,omitempty"`
	Manifest   *CharmManifest         `json:"manifest,omitempty"`
	LXDProfile *CharmLXDProfile       `json:"lxd-profile,omitempty"`
	Version    string                 `json:"version,omitempty"`
}

// CharmActions mirrors charm.Actions.
type CharmActions struct {
	ActionSpecs map[string]CharmActionSpec `json:"specs,omitempty"`
}

// CharmActionSpec mirrors charm.ActionSpec.
type CharmActionSpec struct {
	Description    string                 `json:"description"`
	Parallel       bool                   `json:"parallel,omitempty"`
	Params         map[string]interface{} `json:"params"`
	ExecutionGroup string                 `json:"execution-group,omitempty"`
}

// CharmMetric mirrors charm.Metric.
type CharmMetric struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}

// CharmPlan mirrors charm.Plan
type CharmPlan struct {
	Required bool `json:"required"`
}

// CharmManifest mirrors charm.Manifest
type CharmManifest struct {
	Bases []CharmBase `json:"bases,omitempty"`
}

// CharmBase mirrors charm.Base
type CharmBase struct {
	Name          string   `json:"name,omitempty"`
	Channel       string   `json:"channel,omitempty"`
	Architectures []string `json:"architectures,omitempty"`
}

// CharmContainer mirrors charm.Container
type CharmContainer struct {
	Resource string       `json:"resource,omitempty"`
	Mounts   []CharmMount `json:"mounts,omitempty"`
	Uid      *int         `json:"uid,omitempty"`
	Gid      *int         `json:"gid,omitempty"`
}

// CharmMount mirrors charm.Mount
type CharmMount struct {
	Storage  string `json:"storage,omitempty"`
	Location string `json:"location,omitempty"`
}

// CharmLXDProfile mirrors charm.LXDProfile
type CharmLXDProfile struct {
	Config      map[string]string            `json:"config"`
	Description string                       `json:"description"`
	Devices     map[string]map[string]string `json:"devices"`
}

// CharmLXDProfileResult returns the result of finding the CharmLXDProfile
type CharmLXDProfileResult struct {
	LXDProfile *CharmLXDProfile `json:"lxd-profile"`
}

// ContainerLXDProfile contains the charm.LXDProfile information in addition to
// the name of the profile.
type ContainerLXDProfile struct {
	Profile CharmLXDProfile `json:"profile" yaml:"profile"`
	Name    string          `json:"name" yaml:"name"`
}

// ContainerProfileResult returns the result of finding the CharmLXDProfile and name of
// the lxd profile to be used for 1 unit on the container
type ContainerProfileResult struct {
	Error       *Error                 `json:"error,omitempty"`
	LXDProfiles []*ContainerLXDProfile `json:"lxd-profiles,omitempty"`
}

// ContainerProfileResults returns the ContainerProfileResult for each unit to be placed
// on the container.
type ContainerProfileResults struct {
	Results []ContainerProfileResult `json:"results"`
}

func ToCharmOptionMap(config *charm.Config) map[string]CharmOption {
	if config == nil {
		return nil
	}
	result := make(map[string]CharmOption)
	for key, value := range config.Options {
		result[key] = toParamsCharmOption(value)
	}
	return result
}

func toParamsCharmOption(opt charm.Option) CharmOption {
	return CharmOption{
		Type:        opt.Type,
		Description: opt.Description,
		Default:     opt.Default,
	}
}

func FromCharmOptionMap(config map[string]CharmOption) *charm.Config {
	if len(config) == 0 {
		return nil
	}
	result := &charm.Config{
		Options: make(map[string]charm.Option),
	}
	for key, value := range config {
		result.Options[key] = fromParamsCharmOption(value)
	}
	return result
}

func fromParamsCharmOption(opt CharmOption) charm.Option {
	return charm.Option{
		Type:        opt.Type,
		Description: opt.Description,
		Default:     opt.Default,
	}
}

type ApplicationCharmPlacements struct {
	Placements []ApplicationCharmPlacement `json:"placements"`
}

type ApplicationCharmPlacement struct {
	Application string `json:"application"`
	CharmURL    string `json:"charm-url"`
}
