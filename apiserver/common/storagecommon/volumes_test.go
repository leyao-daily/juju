// Copyright 2015 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package storagecommon_test

import (
	"context"

	"github.com/juju/names/v6"
	jc "github.com/juju/testing/checkers"
	gc "gopkg.in/check.v1"

	"github.com/juju/juju/apiserver/common/storagecommon"
	"github.com/juju/juju/environs/tags"
	"github.com/juju/juju/internal/storage/provider"
	"github.com/juju/juju/internal/testing"
	"github.com/juju/juju/rpc/params"
	"github.com/juju/juju/state"
)

type volumesSuite struct{}

var _ = gc.Suite(&volumesSuite{})

func (s *volumesSuite) TestVolumeParams(c *gc.C) {
	s.testVolumeParams(c, &state.VolumeParams{
		Pool: "loop",
		Size: 1024,
	}, nil)
}

func (s *volumesSuite) TestVolumeParamsAlreadyProvisioned(c *gc.C) {
	s.testVolumeParams(c, nil, &state.VolumeInfo{
		Pool: "loop",
		Size: 1024,
	})
}

func (*volumesSuite) testVolumeParams(c *gc.C, volumeParams *state.VolumeParams, info *state.VolumeInfo) {
	tag := names.NewVolumeTag("100")
	p, err := storagecommon.VolumeParams(
		context.Background(),
		&fakeVolume{tag: tag, params: volumeParams, info: info},
		nil, // StorageInstance
		testing.ModelTag.Id(),
		testing.ControllerTag.Id(),
		testing.CustomModelConfig(c, testing.Attrs{
			"resource-tags": "a=b c=",
		}),
		&fakeStoragePoolGetter{},
		provider.CommonStorageProviders(),
	)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(p, jc.DeepEquals, params.VolumeParams{
		VolumeTag: "volume-100",
		Provider:  "loop",
		Size:      1024,
		Tags: map[string]string{
			tags.JujuController: testing.ControllerTag.Id(),
			tags.JujuModel:      testing.ModelTag.Id(),
			"a":                 "b",
			"c":                 "",
		},
	})
}

func (*volumesSuite) TestVolumeParamsStorageTags(c *gc.C) {
	volumeTag := names.NewVolumeTag("100")
	storageTag := names.NewStorageTag("mystore/0")
	unitTag := names.NewUnitTag("mysql/123")
	p, err := storagecommon.VolumeParams(
		context.Background(),
		&fakeVolume{tag: volumeTag, params: &state.VolumeParams{
			Pool: "loop", Size: 1024,
		}},
		&fakeStorageInstance{tag: storageTag, owner: unitTag},
		testing.ModelTag.Id(),
		testing.ControllerTag.Id(),
		testing.CustomModelConfig(c, nil),
		&fakeStoragePoolGetter{},
		provider.CommonStorageProviders(),
	)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(p, jc.DeepEquals, params.VolumeParams{
		VolumeTag: "volume-100",
		Provider:  "loop",
		Size:      1024,
		Tags: map[string]string{
			tags.JujuController:      testing.ControllerTag.Id(),
			tags.JujuModel:           testing.ModelTag.Id(),
			tags.JujuStorageInstance: "mystore/0",
			tags.JujuStorageOwner:    "mysql/123",
		},
	})
}
