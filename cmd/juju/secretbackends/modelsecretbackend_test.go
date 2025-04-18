// Copyright 2023 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package secretbackends_test

import (
	"github.com/juju/errors"
	jujutesting "github.com/juju/testing"
	jc "github.com/juju/testing/checkers"
	"go.uber.org/mock/gomock"
	gc "gopkg.in/check.v1"

	"github.com/juju/juju/cmd/juju/secretbackends"
	secretbackenderrors "github.com/juju/juju/domain/secretbackend/errors"
	"github.com/juju/juju/internal/cmd/cmdtesting"
	"github.com/juju/juju/jujuclient"
)

type modelSecretBackendCommandSuite struct {
	jujutesting.IsolationSuite
	store      *jujuclient.MemStore
	secretsAPI *secretbackends.MockModelSecretBackendAPI
}

var _ = gc.Suite(&modelSecretBackendCommandSuite{})

func (s *modelSecretBackendCommandSuite) SetUpTest(c *gc.C) {
	s.IsolationSuite.SetUpTest(c)
	store := jujuclient.NewMemStore()
	store.Controllers["mycontroller"] = jujuclient.ControllerDetails{}
	store.CurrentControllerName = "mycontroller"
	s.store = store
}

func (s *modelSecretBackendCommandSuite) setup(c *gc.C) *gomock.Controller {
	ctrl := gomock.NewController(c)
	s.secretsAPI = secretbackends.NewMockModelSecretBackendAPI(ctrl)
	return ctrl
}

func (s *modelSecretBackendCommandSuite) TestGet(c *gc.C) {
	defer s.setup(c).Finish()

	s.secretsAPI.EXPECT().GetModelSecretBackend(gomock.Any()).Return("myVault", nil)
	s.secretsAPI.EXPECT().Close().Return(nil)

	ctx, err := cmdtesting.RunCommand(c, secretbackends.NewModelCredentialCommandForTest(s.store, s.secretsAPI))
	c.Assert(err, jc.ErrorIsNil)
	out := cmdtesting.Stdout(ctx)
	c.Assert(out, gc.Equals, "myVault"+"\n")
}

func (s *modelSecretBackendCommandSuite) TestGetNotSupported(c *gc.C) {
	defer s.setup(c).Finish()

	s.secretsAPI.EXPECT().GetModelSecretBackend(gomock.Any()).Return("", errors.NotSupportedf("model secret backend"))
	s.secretsAPI.EXPECT().Close().Return(nil)

	_, err := cmdtesting.RunCommand(c, secretbackends.NewModelCredentialCommandForTest(s.store, s.secretsAPI))
	c.Assert(err, gc.ErrorMatches, `"model-secret-backend" has not been implemented on the controller, use the "model-config" command instead`)
}

func (s *modelSecretBackendCommandSuite) TestSet(c *gc.C) {
	defer s.setup(c).Finish()

	s.secretsAPI.EXPECT().SetModelSecretBackend(gomock.Any(), "myVault").Return(nil)
	s.secretsAPI.EXPECT().Close().Return(nil)

	_, err := cmdtesting.RunCommand(c, secretbackends.NewModelCredentialCommandForTest(s.store, s.secretsAPI), "myVault")
	c.Assert(err, jc.ErrorIsNil)
}

func (s *modelSecretBackendCommandSuite) TestSetNotSupported(c *gc.C) {
	defer s.setup(c).Finish()

	s.secretsAPI.EXPECT().SetModelSecretBackend(gomock.Any(), "myVault").Return(errors.NotSupportedf("model secret backend"))
	s.secretsAPI.EXPECT().Close().Return(nil)

	_, err := cmdtesting.RunCommand(c, secretbackends.NewModelCredentialCommandForTest(s.store, s.secretsAPI), "myVault")
	c.Assert(err, gc.ErrorMatches, `"model-secret-backend" has not been implemented on the controller, use the "model-config" command instead`)
}

func (s *modelSecretBackendCommandSuite) TestSetSecretBackendNotFound(c *gc.C) {
	defer s.setup(c).Finish()

	s.secretsAPI.EXPECT().SetModelSecretBackend(gomock.Any(), "myVault").Return(secretbackenderrors.NotFound)
	s.secretsAPI.EXPECT().Close().Return(nil)

	_, err := cmdtesting.RunCommand(c, secretbackends.NewModelCredentialCommandForTest(s.store, s.secretsAPI), "myVault")
	c.Assert(err, gc.ErrorMatches, `secret backend not found: please use "add-secret-backend" to add "myVault" to the controller first`)
}

func (s *modelSecretBackendCommandSuite) TestSetSecretBackendNotValid(c *gc.C) {
	defer s.setup(c).Finish()

	s.secretsAPI.EXPECT().SetModelSecretBackend(gomock.Any(), "internal").Return(secretbackenderrors.NotValid)
	s.secretsAPI.EXPECT().Close().Return(nil)

	_, err := cmdtesting.RunCommand(c, secretbackends.NewModelCredentialCommandForTest(s.store, s.secretsAPI), "internal")
	c.Assert(err, gc.ErrorMatches, `secret backend not valid: please use "auto" instead`)
}

func (s *modelSecretBackendCommandSuite) TestSetFailedMoreThanOneArgs(c *gc.C) {
	defer s.setup(c).Finish()

	_, err := cmdtesting.RunCommand(c, secretbackends.NewModelCredentialCommandForTest(s.store, s.secretsAPI), "foo", "bar")
	c.Assert(err, gc.ErrorMatches, "cannot specify multiple secret backend names")
}
