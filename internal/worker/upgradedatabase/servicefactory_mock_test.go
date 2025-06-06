// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/internal/services (interfaces: ControllerDomainServices)
//
// Generated by this command:
//
//	mockgen -typed -package upgradedatabase -destination servicefactory_mock_test.go github.com/juju/juju/internal/services ControllerDomainServices
//

// Package upgradedatabase is a generated GoMock package.
package upgradedatabase

import (
	reflect "reflect"

	service "github.com/juju/juju/domain/access/service"
	service0 "github.com/juju/juju/domain/agentbinary/service"
	service1 "github.com/juju/juju/domain/autocert/service"
	service2 "github.com/juju/juju/domain/cloud/service"
	service3 "github.com/juju/juju/domain/controller/service"
	service4 "github.com/juju/juju/domain/controllerconfig/service"
	service5 "github.com/juju/juju/domain/controllernode/service"
	service6 "github.com/juju/juju/domain/credential/service"
	service7 "github.com/juju/juju/domain/externalcontroller/service"
	service8 "github.com/juju/juju/domain/flag/service"
	service9 "github.com/juju/juju/domain/macaroon/service"
	service10 "github.com/juju/juju/domain/model/service"
	service11 "github.com/juju/juju/domain/modeldefaults/service"
	service12 "github.com/juju/juju/domain/secretbackend/service"
	service13 "github.com/juju/juju/domain/upgrade/service"
	gomock "go.uber.org/mock/gomock"
)

// MockControllerDomainServices is a mock of ControllerDomainServices interface.
type MockControllerDomainServices struct {
	ctrl     *gomock.Controller
	recorder *MockControllerDomainServicesMockRecorder
}

// MockControllerDomainServicesMockRecorder is the mock recorder for MockControllerDomainServices.
type MockControllerDomainServicesMockRecorder struct {
	mock *MockControllerDomainServices
}

// NewMockControllerDomainServices creates a new mock instance.
func NewMockControllerDomainServices(ctrl *gomock.Controller) *MockControllerDomainServices {
	mock := &MockControllerDomainServices{ctrl: ctrl}
	mock.recorder = &MockControllerDomainServicesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockControllerDomainServices) EXPECT() *MockControllerDomainServicesMockRecorder {
	return m.recorder
}

// Access mocks base method.
func (m *MockControllerDomainServices) Access() *service.Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Access")
	ret0, _ := ret[0].(*service.Service)
	return ret0
}

// Access indicates an expected call of Access.
func (mr *MockControllerDomainServicesMockRecorder) Access() *MockControllerDomainServicesAccessCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Access", reflect.TypeOf((*MockControllerDomainServices)(nil).Access))
	return &MockControllerDomainServicesAccessCall{Call: call}
}

// MockControllerDomainServicesAccessCall wrap *gomock.Call
type MockControllerDomainServicesAccessCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockControllerDomainServicesAccessCall) Return(arg0 *service.Service) *MockControllerDomainServicesAccessCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockControllerDomainServicesAccessCall) Do(f func() *service.Service) *MockControllerDomainServicesAccessCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockControllerDomainServicesAccessCall) DoAndReturn(f func() *service.Service) *MockControllerDomainServicesAccessCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// AutocertCache mocks base method.
func (m *MockControllerDomainServices) AutocertCache() *service1.Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AutocertCache")
	ret0, _ := ret[0].(*service1.Service)
	return ret0
}

// AutocertCache indicates an expected call of AutocertCache.
func (mr *MockControllerDomainServicesMockRecorder) AutocertCache() *MockControllerDomainServicesAutocertCacheCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AutocertCache", reflect.TypeOf((*MockControllerDomainServices)(nil).AutocertCache))
	return &MockControllerDomainServicesAutocertCacheCall{Call: call}
}

// MockControllerDomainServicesAutocertCacheCall wrap *gomock.Call
type MockControllerDomainServicesAutocertCacheCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockControllerDomainServicesAutocertCacheCall) Return(arg0 *service1.Service) *MockControllerDomainServicesAutocertCacheCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockControllerDomainServicesAutocertCacheCall) Do(f func() *service1.Service) *MockControllerDomainServicesAutocertCacheCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockControllerDomainServicesAutocertCacheCall) DoAndReturn(f func() *service1.Service) *MockControllerDomainServicesAutocertCacheCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Cloud mocks base method.
func (m *MockControllerDomainServices) Cloud() *service2.WatchableService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cloud")
	ret0, _ := ret[0].(*service2.WatchableService)
	return ret0
}

// Cloud indicates an expected call of Cloud.
func (mr *MockControllerDomainServicesMockRecorder) Cloud() *MockControllerDomainServicesCloudCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cloud", reflect.TypeOf((*MockControllerDomainServices)(nil).Cloud))
	return &MockControllerDomainServicesCloudCall{Call: call}
}

// MockControllerDomainServicesCloudCall wrap *gomock.Call
type MockControllerDomainServicesCloudCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockControllerDomainServicesCloudCall) Return(arg0 *service2.WatchableService) *MockControllerDomainServicesCloudCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockControllerDomainServicesCloudCall) Do(f func() *service2.WatchableService) *MockControllerDomainServicesCloudCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockControllerDomainServicesCloudCall) DoAndReturn(f func() *service2.WatchableService) *MockControllerDomainServicesCloudCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Controller mocks base method.
func (m *MockControllerDomainServices) Controller() *service3.Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Controller")
	ret0, _ := ret[0].(*service3.Service)
	return ret0
}

// Controller indicates an expected call of Controller.
func (mr *MockControllerDomainServicesMockRecorder) Controller() *MockControllerDomainServicesControllerCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Controller", reflect.TypeOf((*MockControllerDomainServices)(nil).Controller))
	return &MockControllerDomainServicesControllerCall{Call: call}
}

// MockControllerDomainServicesControllerCall wrap *gomock.Call
type MockControllerDomainServicesControllerCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockControllerDomainServicesControllerCall) Return(arg0 *service3.Service) *MockControllerDomainServicesControllerCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockControllerDomainServicesControllerCall) Do(f func() *service3.Service) *MockControllerDomainServicesControllerCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockControllerDomainServicesControllerCall) DoAndReturn(f func() *service3.Service) *MockControllerDomainServicesControllerCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ControllerAgentBinaryStore mocks base method.
func (m *MockControllerDomainServices) ControllerAgentBinaryStore() *service0.AgentBinaryStore {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerAgentBinaryStore")
	ret0, _ := ret[0].(*service0.AgentBinaryStore)
	return ret0
}

// ControllerAgentBinaryStore indicates an expected call of ControllerAgentBinaryStore.
func (mr *MockControllerDomainServicesMockRecorder) ControllerAgentBinaryStore() *MockControllerDomainServicesControllerAgentBinaryStoreCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerAgentBinaryStore", reflect.TypeOf((*MockControllerDomainServices)(nil).ControllerAgentBinaryStore))
	return &MockControllerDomainServicesControllerAgentBinaryStoreCall{Call: call}
}

// MockControllerDomainServicesControllerAgentBinaryStoreCall wrap *gomock.Call
type MockControllerDomainServicesControllerAgentBinaryStoreCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockControllerDomainServicesControllerAgentBinaryStoreCall) Return(arg0 *service0.AgentBinaryStore) *MockControllerDomainServicesControllerAgentBinaryStoreCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockControllerDomainServicesControllerAgentBinaryStoreCall) Do(f func() *service0.AgentBinaryStore) *MockControllerDomainServicesControllerAgentBinaryStoreCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockControllerDomainServicesControllerAgentBinaryStoreCall) DoAndReturn(f func() *service0.AgentBinaryStore) *MockControllerDomainServicesControllerAgentBinaryStoreCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ControllerConfig mocks base method.
func (m *MockControllerDomainServices) ControllerConfig() *service4.WatchableService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerConfig")
	ret0, _ := ret[0].(*service4.WatchableService)
	return ret0
}

// ControllerConfig indicates an expected call of ControllerConfig.
func (mr *MockControllerDomainServicesMockRecorder) ControllerConfig() *MockControllerDomainServicesControllerConfigCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerConfig", reflect.TypeOf((*MockControllerDomainServices)(nil).ControllerConfig))
	return &MockControllerDomainServicesControllerConfigCall{Call: call}
}

// MockControllerDomainServicesControllerConfigCall wrap *gomock.Call
type MockControllerDomainServicesControllerConfigCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockControllerDomainServicesControllerConfigCall) Return(arg0 *service4.WatchableService) *MockControllerDomainServicesControllerConfigCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockControllerDomainServicesControllerConfigCall) Do(f func() *service4.WatchableService) *MockControllerDomainServicesControllerConfigCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockControllerDomainServicesControllerConfigCall) DoAndReturn(f func() *service4.WatchableService) *MockControllerDomainServicesControllerConfigCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ControllerNode mocks base method.
func (m *MockControllerDomainServices) ControllerNode() *service5.WatchableService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerNode")
	ret0, _ := ret[0].(*service5.WatchableService)
	return ret0
}

// ControllerNode indicates an expected call of ControllerNode.
func (mr *MockControllerDomainServicesMockRecorder) ControllerNode() *MockControllerDomainServicesControllerNodeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerNode", reflect.TypeOf((*MockControllerDomainServices)(nil).ControllerNode))
	return &MockControllerDomainServicesControllerNodeCall{Call: call}
}

// MockControllerDomainServicesControllerNodeCall wrap *gomock.Call
type MockControllerDomainServicesControllerNodeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockControllerDomainServicesControllerNodeCall) Return(arg0 *service5.WatchableService) *MockControllerDomainServicesControllerNodeCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockControllerDomainServicesControllerNodeCall) Do(f func() *service5.WatchableService) *MockControllerDomainServicesControllerNodeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockControllerDomainServicesControllerNodeCall) DoAndReturn(f func() *service5.WatchableService) *MockControllerDomainServicesControllerNodeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Credential mocks base method.
func (m *MockControllerDomainServices) Credential() *service6.WatchableService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Credential")
	ret0, _ := ret[0].(*service6.WatchableService)
	return ret0
}

// Credential indicates an expected call of Credential.
func (mr *MockControllerDomainServicesMockRecorder) Credential() *MockControllerDomainServicesCredentialCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Credential", reflect.TypeOf((*MockControllerDomainServices)(nil).Credential))
	return &MockControllerDomainServicesCredentialCall{Call: call}
}

// MockControllerDomainServicesCredentialCall wrap *gomock.Call
type MockControllerDomainServicesCredentialCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockControllerDomainServicesCredentialCall) Return(arg0 *service6.WatchableService) *MockControllerDomainServicesCredentialCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockControllerDomainServicesCredentialCall) Do(f func() *service6.WatchableService) *MockControllerDomainServicesCredentialCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockControllerDomainServicesCredentialCall) DoAndReturn(f func() *service6.WatchableService) *MockControllerDomainServicesCredentialCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ExternalController mocks base method.
func (m *MockControllerDomainServices) ExternalController() *service7.WatchableService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExternalController")
	ret0, _ := ret[0].(*service7.WatchableService)
	return ret0
}

// ExternalController indicates an expected call of ExternalController.
func (mr *MockControllerDomainServicesMockRecorder) ExternalController() *MockControllerDomainServicesExternalControllerCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExternalController", reflect.TypeOf((*MockControllerDomainServices)(nil).ExternalController))
	return &MockControllerDomainServicesExternalControllerCall{Call: call}
}

// MockControllerDomainServicesExternalControllerCall wrap *gomock.Call
type MockControllerDomainServicesExternalControllerCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockControllerDomainServicesExternalControllerCall) Return(arg0 *service7.WatchableService) *MockControllerDomainServicesExternalControllerCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockControllerDomainServicesExternalControllerCall) Do(f func() *service7.WatchableService) *MockControllerDomainServicesExternalControllerCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockControllerDomainServicesExternalControllerCall) DoAndReturn(f func() *service7.WatchableService) *MockControllerDomainServicesExternalControllerCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Flag mocks base method.
func (m *MockControllerDomainServices) Flag() *service8.Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Flag")
	ret0, _ := ret[0].(*service8.Service)
	return ret0
}

// Flag indicates an expected call of Flag.
func (mr *MockControllerDomainServicesMockRecorder) Flag() *MockControllerDomainServicesFlagCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Flag", reflect.TypeOf((*MockControllerDomainServices)(nil).Flag))
	return &MockControllerDomainServicesFlagCall{Call: call}
}

// MockControllerDomainServicesFlagCall wrap *gomock.Call
type MockControllerDomainServicesFlagCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockControllerDomainServicesFlagCall) Return(arg0 *service8.Service) *MockControllerDomainServicesFlagCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockControllerDomainServicesFlagCall) Do(f func() *service8.Service) *MockControllerDomainServicesFlagCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockControllerDomainServicesFlagCall) DoAndReturn(f func() *service8.Service) *MockControllerDomainServicesFlagCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Macaroon mocks base method.
func (m *MockControllerDomainServices) Macaroon() *service9.Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Macaroon")
	ret0, _ := ret[0].(*service9.Service)
	return ret0
}

// Macaroon indicates an expected call of Macaroon.
func (mr *MockControllerDomainServicesMockRecorder) Macaroon() *MockControllerDomainServicesMacaroonCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Macaroon", reflect.TypeOf((*MockControllerDomainServices)(nil).Macaroon))
	return &MockControllerDomainServicesMacaroonCall{Call: call}
}

// MockControllerDomainServicesMacaroonCall wrap *gomock.Call
type MockControllerDomainServicesMacaroonCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockControllerDomainServicesMacaroonCall) Return(arg0 *service9.Service) *MockControllerDomainServicesMacaroonCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockControllerDomainServicesMacaroonCall) Do(f func() *service9.Service) *MockControllerDomainServicesMacaroonCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockControllerDomainServicesMacaroonCall) DoAndReturn(f func() *service9.Service) *MockControllerDomainServicesMacaroonCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Model mocks base method.
func (m *MockControllerDomainServices) Model() *service10.WatchableService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Model")
	ret0, _ := ret[0].(*service10.WatchableService)
	return ret0
}

// Model indicates an expected call of Model.
func (mr *MockControllerDomainServicesMockRecorder) Model() *MockControllerDomainServicesModelCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Model", reflect.TypeOf((*MockControllerDomainServices)(nil).Model))
	return &MockControllerDomainServicesModelCall{Call: call}
}

// MockControllerDomainServicesModelCall wrap *gomock.Call
type MockControllerDomainServicesModelCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockControllerDomainServicesModelCall) Return(arg0 *service10.WatchableService) *MockControllerDomainServicesModelCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockControllerDomainServicesModelCall) Do(f func() *service10.WatchableService) *MockControllerDomainServicesModelCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockControllerDomainServicesModelCall) DoAndReturn(f func() *service10.WatchableService) *MockControllerDomainServicesModelCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ModelDefaults mocks base method.
func (m *MockControllerDomainServices) ModelDefaults() *service11.Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelDefaults")
	ret0, _ := ret[0].(*service11.Service)
	return ret0
}

// ModelDefaults indicates an expected call of ModelDefaults.
func (mr *MockControllerDomainServicesMockRecorder) ModelDefaults() *MockControllerDomainServicesModelDefaultsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelDefaults", reflect.TypeOf((*MockControllerDomainServices)(nil).ModelDefaults))
	return &MockControllerDomainServicesModelDefaultsCall{Call: call}
}

// MockControllerDomainServicesModelDefaultsCall wrap *gomock.Call
type MockControllerDomainServicesModelDefaultsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockControllerDomainServicesModelDefaultsCall) Return(arg0 *service11.Service) *MockControllerDomainServicesModelDefaultsCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockControllerDomainServicesModelDefaultsCall) Do(f func() *service11.Service) *MockControllerDomainServicesModelDefaultsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockControllerDomainServicesModelDefaultsCall) DoAndReturn(f func() *service11.Service) *MockControllerDomainServicesModelDefaultsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SecretBackend mocks base method.
func (m *MockControllerDomainServices) SecretBackend() *service12.WatchableService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SecretBackend")
	ret0, _ := ret[0].(*service12.WatchableService)
	return ret0
}

// SecretBackend indicates an expected call of SecretBackend.
func (mr *MockControllerDomainServicesMockRecorder) SecretBackend() *MockControllerDomainServicesSecretBackendCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecretBackend", reflect.TypeOf((*MockControllerDomainServices)(nil).SecretBackend))
	return &MockControllerDomainServicesSecretBackendCall{Call: call}
}

// MockControllerDomainServicesSecretBackendCall wrap *gomock.Call
type MockControllerDomainServicesSecretBackendCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockControllerDomainServicesSecretBackendCall) Return(arg0 *service12.WatchableService) *MockControllerDomainServicesSecretBackendCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockControllerDomainServicesSecretBackendCall) Do(f func() *service12.WatchableService) *MockControllerDomainServicesSecretBackendCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockControllerDomainServicesSecretBackendCall) DoAndReturn(f func() *service12.WatchableService) *MockControllerDomainServicesSecretBackendCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Upgrade mocks base method.
func (m *MockControllerDomainServices) Upgrade() *service13.WatchableService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upgrade")
	ret0, _ := ret[0].(*service13.WatchableService)
	return ret0
}

// Upgrade indicates an expected call of Upgrade.
func (mr *MockControllerDomainServicesMockRecorder) Upgrade() *MockControllerDomainServicesUpgradeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upgrade", reflect.TypeOf((*MockControllerDomainServices)(nil).Upgrade))
	return &MockControllerDomainServicesUpgradeCall{Call: call}
}

// MockControllerDomainServicesUpgradeCall wrap *gomock.Call
type MockControllerDomainServicesUpgradeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockControllerDomainServicesUpgradeCall) Return(arg0 *service13.WatchableService) *MockControllerDomainServicesUpgradeCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockControllerDomainServicesUpgradeCall) Do(f func() *service13.WatchableService) *MockControllerDomainServicesUpgradeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockControllerDomainServicesUpgradeCall) DoAndReturn(f func() *service13.WatchableService) *MockControllerDomainServicesUpgradeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
