// Copyright 2015 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package block

import (
	"testing"

	gc "gopkg.in/check.v1"
)

//go:generate go run go.uber.org/mock/mockgen -typed -package block -destination service_mock_test.go github.com/juju/juju/apiserver/facades/client/block BlockCommandService,Authorizer

func TestAll(t *testing.T) {
	gc.TestingT(t)
}
