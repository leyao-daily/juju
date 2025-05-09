// Copyright 2025 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package caasagent_test

import (
	"context"

	"github.com/juju/testing"
	jc "github.com/juju/testing/checkers"
	"github.com/juju/worker/v4/workertest"
	gc "gopkg.in/check.v1"

	"github.com/juju/juju/api/agent/caasagent"
	apitesting "github.com/juju/juju/api/base/testing"
	internaltesting "github.com/juju/juju/internal/testing"
	"github.com/juju/juju/rpc/params"
)

var _ = gc.Suite(&ClientSuite{})

type ClientSuite struct {
	testing.IsolationSuite
}

func (s *ClientSuite) TestWatchCloudSpecChanges(c *gc.C) {
	called := false
	apiCaller := apitesting.APICallerFunc(func(objType string, version int, id, request string, arg, result interface{}) error {
		// We might get a second call to "Next" but
		// we don't care.
		if called {
			return nil
		}
		c.Check(objType, gc.Equals, "CAASAgent")
		c.Check(version, gc.Equals, 0)
		c.Check(id, gc.Equals, "")
		c.Check(request, gc.Equals, "WatchCloudSpecsChanges")
		c.Check(arg, gc.DeepEquals, params.Entities{
			Entities: []params.Entity{{
				Tag: internaltesting.ModelTag.String(),
			}},
		})
		c.Assert(result, gc.FitsTypeOf, &params.NotifyWatchResults{})
		*(result.(*params.NotifyWatchResults)) = params.NotifyWatchResults{
			Results: []params.NotifyWatchResult{{
				NotifyWatcherId: "666",
			}},
		}
		called = true
		return nil
	})

	api, err := caasagent.NewClient(apiCaller)
	c.Assert(err, jc.ErrorIsNil)
	w, err := api.WatchCloudSpecChanges(context.Background())
	c.Assert(err, jc.ErrorIsNil)
	c.Check(called, jc.IsTrue)
	workertest.CleanKill(c, w)
}
