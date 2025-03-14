// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/internal/services (interfaces: DomainServicesGetter)
//
// Generated by this command:
//
//	mockgen -typed -package apiserver_test -destination service_mock_test.go github.com/juju/juju/internal/services DomainServicesGetter
//

// Package apiserver_test is a generated GoMock package.
package apiserver_test

import (
	context "context"
	reflect "reflect"

	model "github.com/juju/juju/core/model"
	services "github.com/juju/juju/internal/services"
	gomock "go.uber.org/mock/gomock"
)

// MockDomainServicesGetter is a mock of DomainServicesGetter interface.
type MockDomainServicesGetter struct {
	ctrl     *gomock.Controller
	recorder *MockDomainServicesGetterMockRecorder
}

// MockDomainServicesGetterMockRecorder is the mock recorder for MockDomainServicesGetter.
type MockDomainServicesGetterMockRecorder struct {
	mock *MockDomainServicesGetter
}

// NewMockDomainServicesGetter creates a new mock instance.
func NewMockDomainServicesGetter(ctrl *gomock.Controller) *MockDomainServicesGetter {
	mock := &MockDomainServicesGetter{ctrl: ctrl}
	mock.recorder = &MockDomainServicesGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDomainServicesGetter) EXPECT() *MockDomainServicesGetterMockRecorder {
	return m.recorder
}

// ServicesForModel mocks base method.
func (m *MockDomainServicesGetter) ServicesForModel(arg0 context.Context, arg1 model.UUID) (services.DomainServices, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServicesForModel", arg0, arg1)
	ret0, _ := ret[0].(services.DomainServices)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServicesForModel indicates an expected call of ServicesForModel.
func (mr *MockDomainServicesGetterMockRecorder) ServicesForModel(arg0, arg1 any) *MockDomainServicesGetterServicesForModelCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServicesForModel", reflect.TypeOf((*MockDomainServicesGetter)(nil).ServicesForModel), arg0, arg1)
	return &MockDomainServicesGetterServicesForModelCall{Call: call}
}

// MockDomainServicesGetterServicesForModelCall wrap *gomock.Call
type MockDomainServicesGetterServicesForModelCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockDomainServicesGetterServicesForModelCall) Return(arg0 services.DomainServices, arg1 error) *MockDomainServicesGetterServicesForModelCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockDomainServicesGetterServicesForModelCall) Do(f func(context.Context, model.UUID) (services.DomainServices, error)) *MockDomainServicesGetterServicesForModelCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockDomainServicesGetterServicesForModelCall) DoAndReturn(f func(context.Context, model.UUID) (services.DomainServices, error)) *MockDomainServicesGetterServicesForModelCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
