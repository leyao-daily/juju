// Copyright 2023 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package service

import (
	"testing"

	"github.com/juju/tc"
)

//go:generate go run go.uber.org/mock/mockgen -typed -package service -destination service_mock_test.go github.com/juju/juju/domain/modeldefaults/service State
//go:generate go run go.uber.org/mock/mockgen -typed -package service -destination environs_mock_test.go github.com/juju/juju/environs ModelConfigProvider

func TestPackage(t *testing.T) {
	tc.TestingT(t)
}
