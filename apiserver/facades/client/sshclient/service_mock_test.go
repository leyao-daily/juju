// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/client/sshclient (interfaces: ModelConfigService,ModelProviderService)
//
// Generated by this command:
//
//	mockgen -typed -package sshclient_test -destination service_mock_test.go github.com/juju/juju/apiserver/facades/client/sshclient ModelConfigService,ModelProviderService
//

// Package sshclient_test is a generated GoMock package.
package sshclient_test

import (
	context "context"
	reflect "reflect"

	cloudspec "github.com/juju/juju/environs/cloudspec"
	config "github.com/juju/juju/environs/config"
	gomock "go.uber.org/mock/gomock"
)

// MockModelConfigService is a mock of ModelConfigService interface.
type MockModelConfigService struct {
	ctrl     *gomock.Controller
	recorder *MockModelConfigServiceMockRecorder
}

// MockModelConfigServiceMockRecorder is the mock recorder for MockModelConfigService.
type MockModelConfigServiceMockRecorder struct {
	mock *MockModelConfigService
}

// NewMockModelConfigService creates a new mock instance.
func NewMockModelConfigService(ctrl *gomock.Controller) *MockModelConfigService {
	mock := &MockModelConfigService{ctrl: ctrl}
	mock.recorder = &MockModelConfigServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModelConfigService) EXPECT() *MockModelConfigServiceMockRecorder {
	return m.recorder
}

// ModelConfig mocks base method.
func (m *MockModelConfigService) ModelConfig(arg0 context.Context) (*config.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelConfig", arg0)
	ret0, _ := ret[0].(*config.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelConfig indicates an expected call of ModelConfig.
func (mr *MockModelConfigServiceMockRecorder) ModelConfig(arg0 any) *MockModelConfigServiceModelConfigCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelConfig", reflect.TypeOf((*MockModelConfigService)(nil).ModelConfig), arg0)
	return &MockModelConfigServiceModelConfigCall{Call: call}
}

// MockModelConfigServiceModelConfigCall wrap *gomock.Call
type MockModelConfigServiceModelConfigCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelConfigServiceModelConfigCall) Return(arg0 *config.Config, arg1 error) *MockModelConfigServiceModelConfigCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelConfigServiceModelConfigCall) Do(f func(context.Context) (*config.Config, error)) *MockModelConfigServiceModelConfigCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelConfigServiceModelConfigCall) DoAndReturn(f func(context.Context) (*config.Config, error)) *MockModelConfigServiceModelConfigCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockModelProviderService is a mock of ModelProviderService interface.
type MockModelProviderService struct {
	ctrl     *gomock.Controller
	recorder *MockModelProviderServiceMockRecorder
}

// MockModelProviderServiceMockRecorder is the mock recorder for MockModelProviderService.
type MockModelProviderServiceMockRecorder struct {
	mock *MockModelProviderService
}

// NewMockModelProviderService creates a new mock instance.
func NewMockModelProviderService(ctrl *gomock.Controller) *MockModelProviderService {
	mock := &MockModelProviderService{ctrl: ctrl}
	mock.recorder = &MockModelProviderServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModelProviderService) EXPECT() *MockModelProviderServiceMockRecorder {
	return m.recorder
}

// GetCloudSpecForSSH mocks base method.
func (m *MockModelProviderService) GetCloudSpecForSSH(arg0 context.Context) (cloudspec.CloudSpec, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCloudSpecForSSH", arg0)
	ret0, _ := ret[0].(cloudspec.CloudSpec)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCloudSpecForSSH indicates an expected call of GetCloudSpecForSSH.
func (mr *MockModelProviderServiceMockRecorder) GetCloudSpecForSSH(arg0 any) *MockModelProviderServiceGetCloudSpecForSSHCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCloudSpecForSSH", reflect.TypeOf((*MockModelProviderService)(nil).GetCloudSpecForSSH), arg0)
	return &MockModelProviderServiceGetCloudSpecForSSHCall{Call: call}
}

// MockModelProviderServiceGetCloudSpecForSSHCall wrap *gomock.Call
type MockModelProviderServiceGetCloudSpecForSSHCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelProviderServiceGetCloudSpecForSSHCall) Return(arg0 cloudspec.CloudSpec, arg1 error) *MockModelProviderServiceGetCloudSpecForSSHCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelProviderServiceGetCloudSpecForSSHCall) Do(f func(context.Context) (cloudspec.CloudSpec, error)) *MockModelProviderServiceGetCloudSpecForSSHCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelProviderServiceGetCloudSpecForSSHCall) DoAndReturn(f func(context.Context) (cloudspec.CloudSpec, error)) *MockModelProviderServiceGetCloudSpecForSSHCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
