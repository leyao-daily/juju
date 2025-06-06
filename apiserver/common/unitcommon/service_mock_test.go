// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/common/unitcommon (interfaces: ApplicationService)
//
// Generated by this command:
//
//	mockgen -typed -package unitcommon -destination service_mock_test.go github.com/juju/juju/apiserver/common/unitcommon ApplicationService
//

// Package unitcommon is a generated GoMock package.
package unitcommon

import (
	context "context"
	reflect "reflect"

	application "github.com/juju/juju/core/application"
	gomock "go.uber.org/mock/gomock"
)

// MockApplicationService is a mock of ApplicationService interface.
type MockApplicationService struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationServiceMockRecorder
}

// MockApplicationServiceMockRecorder is the mock recorder for MockApplicationService.
type MockApplicationServiceMockRecorder struct {
	mock *MockApplicationService
}

// NewMockApplicationService creates a new mock instance.
func NewMockApplicationService(ctrl *gomock.Controller) *MockApplicationService {
	mock := &MockApplicationService{ctrl: ctrl}
	mock.recorder = &MockApplicationServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApplicationService) EXPECT() *MockApplicationServiceMockRecorder {
	return m.recorder
}

// GetApplicationIDByName mocks base method.
func (m *MockApplicationService) GetApplicationIDByName(arg0 context.Context, arg1 string) (application.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetApplicationIDByName", arg0, arg1)
	ret0, _ := ret[0].(application.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetApplicationIDByName indicates an expected call of GetApplicationIDByName.
func (mr *MockApplicationServiceMockRecorder) GetApplicationIDByName(arg0, arg1 any) *MockApplicationServiceGetApplicationIDByNameCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApplicationIDByName", reflect.TypeOf((*MockApplicationService)(nil).GetApplicationIDByName), arg0, arg1)
	return &MockApplicationServiceGetApplicationIDByNameCall{Call: call}
}

// MockApplicationServiceGetApplicationIDByNameCall wrap *gomock.Call
type MockApplicationServiceGetApplicationIDByNameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationServiceGetApplicationIDByNameCall) Return(arg0 application.ID, arg1 error) *MockApplicationServiceGetApplicationIDByNameCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationServiceGetApplicationIDByNameCall) Do(f func(context.Context, string) (application.ID, error)) *MockApplicationServiceGetApplicationIDByNameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationServiceGetApplicationIDByNameCall) DoAndReturn(f func(context.Context, string) (application.ID, error)) *MockApplicationServiceGetApplicationIDByNameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
