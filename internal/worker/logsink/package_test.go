// Copyright 2024 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package logsink

import (
	"testing"

	"github.com/juju/tc"
)

//go:generate go run go.uber.org/mock/mockgen -typed -package logsink -destination logger_mock_test.go github.com/juju/juju/core/logger Logger,LogWriter,LoggerContextGetter,ModelLogger,LogSink

func Test(t *testing.T) {
	tc.TestingT(t)
}
