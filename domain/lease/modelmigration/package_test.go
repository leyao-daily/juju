// Copyright 2023 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package modelmigration

//go:generate go run go.uber.org/mock/mockgen -typed -package modelmigration -destination migrations_mock_test.go github.com/juju/juju/domain/lease/modelmigration Coordinator,ImportService,ExportService
//go:generate go run go.uber.org/mock/mockgen -typed -package modelmigration -destination database_mock_test.go github.com/juju/juju/core/database TxnRunner
