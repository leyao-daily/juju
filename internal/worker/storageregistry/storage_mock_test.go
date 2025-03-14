// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/internal/storage (interfaces: ProviderRegistry,Provider)
//
// Generated by this command:
//
//	mockgen -typed -package storageregistry -destination storage_mock_test.go github.com/juju/juju/internal/storage ProviderRegistry,Provider
//

// Package storageregistry is a generated GoMock package.
package storageregistry

import (
	reflect "reflect"

	storage "github.com/juju/juju/internal/storage"
	gomock "go.uber.org/mock/gomock"
)

// MockProviderRegistry is a mock of ProviderRegistry interface.
type MockProviderRegistry struct {
	ctrl     *gomock.Controller
	recorder *MockProviderRegistryMockRecorder
}

// MockProviderRegistryMockRecorder is the mock recorder for MockProviderRegistry.
type MockProviderRegistryMockRecorder struct {
	mock *MockProviderRegistry
}

// NewMockProviderRegistry creates a new mock instance.
func NewMockProviderRegistry(ctrl *gomock.Controller) *MockProviderRegistry {
	mock := &MockProviderRegistry{ctrl: ctrl}
	mock.recorder = &MockProviderRegistryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProviderRegistry) EXPECT() *MockProviderRegistryMockRecorder {
	return m.recorder
}

// StorageProvider mocks base method.
func (m *MockProviderRegistry) StorageProvider(arg0 storage.ProviderType) (storage.Provider, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorageProvider", arg0)
	ret0, _ := ret[0].(storage.Provider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StorageProvider indicates an expected call of StorageProvider.
func (mr *MockProviderRegistryMockRecorder) StorageProvider(arg0 any) *MockProviderRegistryStorageProviderCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageProvider", reflect.TypeOf((*MockProviderRegistry)(nil).StorageProvider), arg0)
	return &MockProviderRegistryStorageProviderCall{Call: call}
}

// MockProviderRegistryStorageProviderCall wrap *gomock.Call
type MockProviderRegistryStorageProviderCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockProviderRegistryStorageProviderCall) Return(arg0 storage.Provider, arg1 error) *MockProviderRegistryStorageProviderCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockProviderRegistryStorageProviderCall) Do(f func(storage.ProviderType) (storage.Provider, error)) *MockProviderRegistryStorageProviderCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockProviderRegistryStorageProviderCall) DoAndReturn(f func(storage.ProviderType) (storage.Provider, error)) *MockProviderRegistryStorageProviderCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// StorageProviderTypes mocks base method.
func (m *MockProviderRegistry) StorageProviderTypes() ([]storage.ProviderType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorageProviderTypes")
	ret0, _ := ret[0].([]storage.ProviderType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StorageProviderTypes indicates an expected call of StorageProviderTypes.
func (mr *MockProviderRegistryMockRecorder) StorageProviderTypes() *MockProviderRegistryStorageProviderTypesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageProviderTypes", reflect.TypeOf((*MockProviderRegistry)(nil).StorageProviderTypes))
	return &MockProviderRegistryStorageProviderTypesCall{Call: call}
}

// MockProviderRegistryStorageProviderTypesCall wrap *gomock.Call
type MockProviderRegistryStorageProviderTypesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockProviderRegistryStorageProviderTypesCall) Return(arg0 []storage.ProviderType, arg1 error) *MockProviderRegistryStorageProviderTypesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockProviderRegistryStorageProviderTypesCall) Do(f func() ([]storage.ProviderType, error)) *MockProviderRegistryStorageProviderTypesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockProviderRegistryStorageProviderTypesCall) DoAndReturn(f func() ([]storage.ProviderType, error)) *MockProviderRegistryStorageProviderTypesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockProvider is a mock of Provider interface.
type MockProvider struct {
	ctrl     *gomock.Controller
	recorder *MockProviderMockRecorder
}

// MockProviderMockRecorder is the mock recorder for MockProvider.
type MockProviderMockRecorder struct {
	mock *MockProvider
}

// NewMockProvider creates a new mock instance.
func NewMockProvider(ctrl *gomock.Controller) *MockProvider {
	mock := &MockProvider{ctrl: ctrl}
	mock.recorder = &MockProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProvider) EXPECT() *MockProviderMockRecorder {
	return m.recorder
}

// DefaultPools mocks base method.
func (m *MockProvider) DefaultPools() []*storage.Config {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultPools")
	ret0, _ := ret[0].([]*storage.Config)
	return ret0
}

// DefaultPools indicates an expected call of DefaultPools.
func (mr *MockProviderMockRecorder) DefaultPools() *MockProviderDefaultPoolsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultPools", reflect.TypeOf((*MockProvider)(nil).DefaultPools))
	return &MockProviderDefaultPoolsCall{Call: call}
}

// MockProviderDefaultPoolsCall wrap *gomock.Call
type MockProviderDefaultPoolsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockProviderDefaultPoolsCall) Return(arg0 []*storage.Config) *MockProviderDefaultPoolsCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockProviderDefaultPoolsCall) Do(f func() []*storage.Config) *MockProviderDefaultPoolsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockProviderDefaultPoolsCall) DoAndReturn(f func() []*storage.Config) *MockProviderDefaultPoolsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Dynamic mocks base method.
func (m *MockProvider) Dynamic() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Dynamic")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Dynamic indicates an expected call of Dynamic.
func (mr *MockProviderMockRecorder) Dynamic() *MockProviderDynamicCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dynamic", reflect.TypeOf((*MockProvider)(nil).Dynamic))
	return &MockProviderDynamicCall{Call: call}
}

// MockProviderDynamicCall wrap *gomock.Call
type MockProviderDynamicCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockProviderDynamicCall) Return(arg0 bool) *MockProviderDynamicCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockProviderDynamicCall) Do(f func() bool) *MockProviderDynamicCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockProviderDynamicCall) DoAndReturn(f func() bool) *MockProviderDynamicCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// FilesystemSource mocks base method.
func (m *MockProvider) FilesystemSource(arg0 *storage.Config) (storage.FilesystemSource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilesystemSource", arg0)
	ret0, _ := ret[0].(storage.FilesystemSource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilesystemSource indicates an expected call of FilesystemSource.
func (mr *MockProviderMockRecorder) FilesystemSource(arg0 any) *MockProviderFilesystemSourceCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilesystemSource", reflect.TypeOf((*MockProvider)(nil).FilesystemSource), arg0)
	return &MockProviderFilesystemSourceCall{Call: call}
}

// MockProviderFilesystemSourceCall wrap *gomock.Call
type MockProviderFilesystemSourceCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockProviderFilesystemSourceCall) Return(arg0 storage.FilesystemSource, arg1 error) *MockProviderFilesystemSourceCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockProviderFilesystemSourceCall) Do(f func(*storage.Config) (storage.FilesystemSource, error)) *MockProviderFilesystemSourceCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockProviderFilesystemSourceCall) DoAndReturn(f func(*storage.Config) (storage.FilesystemSource, error)) *MockProviderFilesystemSourceCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Releasable mocks base method.
func (m *MockProvider) Releasable() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Releasable")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Releasable indicates an expected call of Releasable.
func (mr *MockProviderMockRecorder) Releasable() *MockProviderReleasableCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Releasable", reflect.TypeOf((*MockProvider)(nil).Releasable))
	return &MockProviderReleasableCall{Call: call}
}

// MockProviderReleasableCall wrap *gomock.Call
type MockProviderReleasableCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockProviderReleasableCall) Return(arg0 bool) *MockProviderReleasableCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockProviderReleasableCall) Do(f func() bool) *MockProviderReleasableCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockProviderReleasableCall) DoAndReturn(f func() bool) *MockProviderReleasableCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Scope mocks base method.
func (m *MockProvider) Scope() storage.Scope {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Scope")
	ret0, _ := ret[0].(storage.Scope)
	return ret0
}

// Scope indicates an expected call of Scope.
func (mr *MockProviderMockRecorder) Scope() *MockProviderScopeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scope", reflect.TypeOf((*MockProvider)(nil).Scope))
	return &MockProviderScopeCall{Call: call}
}

// MockProviderScopeCall wrap *gomock.Call
type MockProviderScopeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockProviderScopeCall) Return(arg0 storage.Scope) *MockProviderScopeCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockProviderScopeCall) Do(f func() storage.Scope) *MockProviderScopeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockProviderScopeCall) DoAndReturn(f func() storage.Scope) *MockProviderScopeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Supports mocks base method.
func (m *MockProvider) Supports(arg0 storage.StorageKind) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Supports", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Supports indicates an expected call of Supports.
func (mr *MockProviderMockRecorder) Supports(arg0 any) *MockProviderSupportsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Supports", reflect.TypeOf((*MockProvider)(nil).Supports), arg0)
	return &MockProviderSupportsCall{Call: call}
}

// MockProviderSupportsCall wrap *gomock.Call
type MockProviderSupportsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockProviderSupportsCall) Return(arg0 bool) *MockProviderSupportsCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockProviderSupportsCall) Do(f func(storage.StorageKind) bool) *MockProviderSupportsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockProviderSupportsCall) DoAndReturn(f func(storage.StorageKind) bool) *MockProviderSupportsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ValidateConfig mocks base method.
func (m *MockProvider) ValidateConfig(arg0 *storage.Config) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateConfig", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateConfig indicates an expected call of ValidateConfig.
func (mr *MockProviderMockRecorder) ValidateConfig(arg0 any) *MockProviderValidateConfigCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateConfig", reflect.TypeOf((*MockProvider)(nil).ValidateConfig), arg0)
	return &MockProviderValidateConfigCall{Call: call}
}

// MockProviderValidateConfigCall wrap *gomock.Call
type MockProviderValidateConfigCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockProviderValidateConfigCall) Return(arg0 error) *MockProviderValidateConfigCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockProviderValidateConfigCall) Do(f func(*storage.Config) error) *MockProviderValidateConfigCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockProviderValidateConfigCall) DoAndReturn(f func(*storage.Config) error) *MockProviderValidateConfigCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ValidateForK8s mocks base method.
func (m *MockProvider) ValidateForK8s(arg0 map[string]any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateForK8s", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateForK8s indicates an expected call of ValidateForK8s.
func (mr *MockProviderMockRecorder) ValidateForK8s(arg0 any) *MockProviderValidateForK8sCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateForK8s", reflect.TypeOf((*MockProvider)(nil).ValidateForK8s), arg0)
	return &MockProviderValidateForK8sCall{Call: call}
}

// MockProviderValidateForK8sCall wrap *gomock.Call
type MockProviderValidateForK8sCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockProviderValidateForK8sCall) Return(arg0 error) *MockProviderValidateForK8sCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockProviderValidateForK8sCall) Do(f func(map[string]any) error) *MockProviderValidateForK8sCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockProviderValidateForK8sCall) DoAndReturn(f func(map[string]any) error) *MockProviderValidateForK8sCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// VolumeSource mocks base method.
func (m *MockProvider) VolumeSource(arg0 *storage.Config) (storage.VolumeSource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VolumeSource", arg0)
	ret0, _ := ret[0].(storage.VolumeSource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VolumeSource indicates an expected call of VolumeSource.
func (mr *MockProviderMockRecorder) VolumeSource(arg0 any) *MockProviderVolumeSourceCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VolumeSource", reflect.TypeOf((*MockProvider)(nil).VolumeSource), arg0)
	return &MockProviderVolumeSourceCall{Call: call}
}

// MockProviderVolumeSourceCall wrap *gomock.Call
type MockProviderVolumeSourceCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockProviderVolumeSourceCall) Return(arg0 storage.VolumeSource, arg1 error) *MockProviderVolumeSourceCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockProviderVolumeSourceCall) Do(f func(*storage.Config) (storage.VolumeSource, error)) *MockProviderVolumeSourceCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockProviderVolumeSourceCall) DoAndReturn(f func(*storage.Config) (storage.VolumeSource, error)) *MockProviderVolumeSourceCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
