// Copyright 2011, 2012, 2013 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package ec2

import (
	"context"
	"fmt"

	"github.com/juju/schema"

	"github.com/juju/juju/environs/config"
	"github.com/juju/juju/internal/configschema"
)

var configSchema = configschema.Fields{
	"vpc-id": {
		Description: "Use a specific AWS VPC ID (optional). When not specified, Juju requires a default VPC or EC2-Classic features to be available for the account/region.",
		Example:     "vpc-a1b2c3d4",
		Type:        configschema.Tstring,
		Group:       configschema.AccountGroup,
		Immutable:   true,
	},
	"vpc-id-force": {
		Description: "Force Juju to use the AWS VPC ID specified with vpc-id, when it fails the minimum validation criteria. Not accepted without vpc-id",
		Type:        configschema.Tbool,
		Group:       configschema.AccountGroup,
		Immutable:   true,
	},
}

var configFields = func() schema.Fields {
	fs, _, err := configSchema.ValidationSchema()
	if err != nil {
		panic(err)
	}
	return fs
}()

var configDefaults = schema.Defaults{
	"vpc-id":       "",
	"vpc-id-force": false,
}

type environConfig struct {
	*config.Config
	attrs map[string]interface{}
}

func (c *environConfig) vpcID() string {
	return c.attrs["vpc-id"].(string)
}

func (c *environConfig) forceVPCID() bool {
	return c.attrs["vpc-id-force"].(bool)
}

func (p environProvider) newConfig(ctx context.Context, cfg *config.Config) (*environConfig, error) {
	valid, err := p.Validate(ctx, cfg, nil)
	if err != nil {
		return nil, err
	}
	return &environConfig{Config: valid, attrs: valid.UnknownAttrs()}, nil
}

// ModelConfigDefaults provides a set of default model config attributes that
// should be set on a models config if they have not been specified by the user.
func (p environProvider) ModelConfigDefaults(_ context.Context) (map[string]any, error) {
	return map[string]any{
		config.StorageDefaultBlockSourceKey: string(EBS_ProviderType),
	}, nil
}

// Schema returns the configuration schema for an environment.
func (environProvider) Schema() configschema.Fields {
	fields, err := config.Schema(configSchema)
	if err != nil {
		panic(err)
	}
	return fields
}

// ConfigSchema returns extra config attributes specific
// to this provider only.
func (p environProvider) ConfigSchema() schema.Fields {
	return configFields
}

// ConfigDefaults returns the default values for the
// provider specific config attributes.
func (p environProvider) ConfigDefaults() schema.Defaults {
	return configDefaults
}

func validateConfig(ctx context.Context, cfg, old *config.Config) (*environConfig, error) {
	// Check for valid changes for the base config values.
	if err := config.Validate(ctx, cfg, old); err != nil {
		return nil, err
	}
	validated, err := cfg.ValidateUnknownAttrs(configFields, configDefaults)
	if err != nil {
		return nil, err
	}
	ecfg := &environConfig{Config: cfg, attrs: validated}

	if vpcID := ecfg.vpcID(); isVPCIDSetButInvalid(vpcID) {
		return nil, fmt.Errorf("vpc-id: %q is not a valid AWS VPC ID", vpcID)
	} else if !isVPCIDSet(vpcID) && ecfg.forceVPCID() {
		return nil, fmt.Errorf("cannot use vpc-id-force without specifying vpc-id as well")
	}

	if old != nil {
		attrs := old.UnknownAttrs()

		if vpcID, _ := attrs["vpc-id"].(string); vpcID != ecfg.vpcID() {
			return nil, fmt.Errorf("cannot change vpc-id from %q to %q", vpcID, ecfg.vpcID())
		}

		if forceVPCID, _ := attrs["vpc-id-force"].(bool); forceVPCID != ecfg.forceVPCID() {
			return nil, fmt.Errorf("cannot change vpc-id-force from %v to %v", forceVPCID, ecfg.forceVPCID())
		}
	}

	// ssl-hostname-verification cannot be disabled
	if !ecfg.SSLHostnameVerification() {
		return nil, fmt.Errorf("disabling ssh-hostname-verification is not supported")
	}
	return ecfg, nil
}
