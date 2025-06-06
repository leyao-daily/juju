// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/client/action (interfaces: State,Model,ApplicationService,ModelInfoService)
//
// Generated by this command:
//
//	mockgen -typed -package action -destination package_mock_test.go github.com/juju/juju/apiserver/facades/client/action State,Model,ApplicationService,ModelInfoService
//

// Package action is a generated GoMock package.
package action

import (
	context "context"
	reflect "reflect"

	model "github.com/juju/juju/core/model"
	unit "github.com/juju/juju/core/unit"
	charm "github.com/juju/juju/domain/application/charm"
	charm0 "github.com/juju/juju/internal/charm"
	state "github.com/juju/juju/state"
	names "github.com/juju/names/v6"
	gomock "go.uber.org/mock/gomock"
)

// MockState is a mock of State interface.
type MockState struct {
	ctrl     *gomock.Controller
	recorder *MockStateMockRecorder
}

// MockStateMockRecorder is the mock recorder for MockState.
type MockStateMockRecorder struct {
	mock *MockState
}

// NewMockState creates a new mock instance.
func NewMockState(ctrl *gomock.Controller) *MockState {
	mock := &MockState{ctrl: ctrl}
	mock.recorder = &MockStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockState) EXPECT() *MockStateMockRecorder {
	return m.recorder
}

// ActionByTag mocks base method.
func (m *MockState) ActionByTag(arg0 names.ActionTag) (state.Action, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActionByTag", arg0)
	ret0, _ := ret[0].(state.Action)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ActionByTag indicates an expected call of ActionByTag.
func (mr *MockStateMockRecorder) ActionByTag(arg0 any) *MockStateActionByTagCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActionByTag", reflect.TypeOf((*MockState)(nil).ActionByTag), arg0)
	return &MockStateActionByTagCall{Call: call}
}

// MockStateActionByTagCall wrap *gomock.Call
type MockStateActionByTagCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateActionByTagCall) Return(arg0 state.Action, arg1 error) *MockStateActionByTagCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateActionByTagCall) Do(f func(names.ActionTag) (state.Action, error)) *MockStateActionByTagCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateActionByTagCall) DoAndReturn(f func(names.ActionTag) (state.Action, error)) *MockStateActionByTagCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// AllMachines mocks base method.
func (m *MockState) AllMachines() ([]*state.Machine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllMachines")
	ret0, _ := ret[0].([]*state.Machine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllMachines indicates an expected call of AllMachines.
func (mr *MockStateMockRecorder) AllMachines() *MockStateAllMachinesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllMachines", reflect.TypeOf((*MockState)(nil).AllMachines))
	return &MockStateAllMachinesCall{Call: call}
}

// MockStateAllMachinesCall wrap *gomock.Call
type MockStateAllMachinesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateAllMachinesCall) Return(arg0 []*state.Machine, arg1 error) *MockStateAllMachinesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateAllMachinesCall) Do(f func() ([]*state.Machine, error)) *MockStateAllMachinesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateAllMachinesCall) DoAndReturn(f func() ([]*state.Machine, error)) *MockStateAllMachinesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Model mocks base method.
func (m *MockState) Model() (Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Model")
	ret0, _ := ret[0].(Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Model indicates an expected call of Model.
func (mr *MockStateMockRecorder) Model() *MockStateModelCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Model", reflect.TypeOf((*MockState)(nil).Model))
	return &MockStateModelCall{Call: call}
}

// MockStateModelCall wrap *gomock.Call
type MockStateModelCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateModelCall) Return(arg0 Model, arg1 error) *MockStateModelCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateModelCall) Do(f func() (Model, error)) *MockStateModelCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateModelCall) DoAndReturn(f func() (Model, error)) *MockStateModelCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// WatchActionLogs mocks base method.
func (m *MockState) WatchActionLogs(arg0 string) state.StringsWatcher {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchActionLogs", arg0)
	ret0, _ := ret[0].(state.StringsWatcher)
	return ret0
}

// WatchActionLogs indicates an expected call of WatchActionLogs.
func (mr *MockStateMockRecorder) WatchActionLogs(arg0 any) *MockStateWatchActionLogsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchActionLogs", reflect.TypeOf((*MockState)(nil).WatchActionLogs), arg0)
	return &MockStateWatchActionLogsCall{Call: call}
}

// MockStateWatchActionLogsCall wrap *gomock.Call
type MockStateWatchActionLogsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateWatchActionLogsCall) Return(arg0 state.StringsWatcher) *MockStateWatchActionLogsCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateWatchActionLogsCall) Do(f func(string) state.StringsWatcher) *MockStateWatchActionLogsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateWatchActionLogsCall) DoAndReturn(f func(string) state.StringsWatcher) *MockStateWatchActionLogsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockModel is a mock of Model interface.
type MockModel struct {
	ctrl     *gomock.Controller
	recorder *MockModelMockRecorder
}

// MockModelMockRecorder is the mock recorder for MockModel.
type MockModelMockRecorder struct {
	mock *MockModel
}

// NewMockModel creates a new mock instance.
func NewMockModel(ctrl *gomock.Controller) *MockModel {
	mock := &MockModel{ctrl: ctrl}
	mock.recorder = &MockModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModel) EXPECT() *MockModelMockRecorder {
	return m.recorder
}

// AddAction mocks base method.
func (m *MockModel) AddAction(arg0 state.ActionReceiver, arg1, arg2 string, arg3 map[string]any, arg4 *bool, arg5 *string) (state.Action, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAction", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(state.Action)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddAction indicates an expected call of AddAction.
func (mr *MockModelMockRecorder) AddAction(arg0, arg1, arg2, arg3, arg4, arg5 any) *MockModelAddActionCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAction", reflect.TypeOf((*MockModel)(nil).AddAction), arg0, arg1, arg2, arg3, arg4, arg5)
	return &MockModelAddActionCall{Call: call}
}

// MockModelAddActionCall wrap *gomock.Call
type MockModelAddActionCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelAddActionCall) Return(arg0 state.Action, arg1 error) *MockModelAddActionCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelAddActionCall) Do(f func(state.ActionReceiver, string, string, map[string]any, *bool, *string) (state.Action, error)) *MockModelAddActionCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelAddActionCall) DoAndReturn(f func(state.ActionReceiver, string, string, map[string]any, *bool, *string) (state.Action, error)) *MockModelAddActionCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// EnqueueOperation mocks base method.
func (m *MockModel) EnqueueOperation(arg0 string, arg1 int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnqueueOperation", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnqueueOperation indicates an expected call of EnqueueOperation.
func (mr *MockModelMockRecorder) EnqueueOperation(arg0, arg1 any) *MockModelEnqueueOperationCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnqueueOperation", reflect.TypeOf((*MockModel)(nil).EnqueueOperation), arg0, arg1)
	return &MockModelEnqueueOperationCall{Call: call}
}

// MockModelEnqueueOperationCall wrap *gomock.Call
type MockModelEnqueueOperationCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelEnqueueOperationCall) Return(arg0 string, arg1 error) *MockModelEnqueueOperationCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelEnqueueOperationCall) Do(f func(string, int) (string, error)) *MockModelEnqueueOperationCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelEnqueueOperationCall) DoAndReturn(f func(string, int) (string, error)) *MockModelEnqueueOperationCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// FailOperationEnqueuing mocks base method.
func (m *MockModel) FailOperationEnqueuing(arg0, arg1 string, arg2 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FailOperationEnqueuing", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// FailOperationEnqueuing indicates an expected call of FailOperationEnqueuing.
func (mr *MockModelMockRecorder) FailOperationEnqueuing(arg0, arg1, arg2 any) *MockModelFailOperationEnqueuingCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FailOperationEnqueuing", reflect.TypeOf((*MockModel)(nil).FailOperationEnqueuing), arg0, arg1, arg2)
	return &MockModelFailOperationEnqueuingCall{Call: call}
}

// MockModelFailOperationEnqueuingCall wrap *gomock.Call
type MockModelFailOperationEnqueuingCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelFailOperationEnqueuingCall) Return(arg0 error) *MockModelFailOperationEnqueuingCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelFailOperationEnqueuingCall) Do(f func(string, string, int) error) *MockModelFailOperationEnqueuingCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelFailOperationEnqueuingCall) DoAndReturn(f func(string, string, int) error) *MockModelFailOperationEnqueuingCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// FindActionsByName mocks base method.
func (m *MockModel) FindActionsByName(arg0 string) ([]state.Action, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindActionsByName", arg0)
	ret0, _ := ret[0].([]state.Action)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindActionsByName indicates an expected call of FindActionsByName.
func (mr *MockModelMockRecorder) FindActionsByName(arg0 any) *MockModelFindActionsByNameCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindActionsByName", reflect.TypeOf((*MockModel)(nil).FindActionsByName), arg0)
	return &MockModelFindActionsByNameCall{Call: call}
}

// MockModelFindActionsByNameCall wrap *gomock.Call
type MockModelFindActionsByNameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelFindActionsByNameCall) Return(arg0 []state.Action, arg1 error) *MockModelFindActionsByNameCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelFindActionsByNameCall) Do(f func(string) ([]state.Action, error)) *MockModelFindActionsByNameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelFindActionsByNameCall) DoAndReturn(f func(string) ([]state.Action, error)) *MockModelFindActionsByNameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListOperations mocks base method.
func (m *MockModel) ListOperations(arg0 []string, arg1 []names.Tag, arg2 []state.ActionStatus, arg3, arg4 int) ([]state.OperationInfo, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOperations", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]state.OperationInfo)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListOperations indicates an expected call of ListOperations.
func (mr *MockModelMockRecorder) ListOperations(arg0, arg1, arg2, arg3, arg4 any) *MockModelListOperationsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOperations", reflect.TypeOf((*MockModel)(nil).ListOperations), arg0, arg1, arg2, arg3, arg4)
	return &MockModelListOperationsCall{Call: call}
}

// MockModelListOperationsCall wrap *gomock.Call
type MockModelListOperationsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelListOperationsCall) Return(arg0 []state.OperationInfo, arg1 bool, arg2 error) *MockModelListOperationsCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelListOperationsCall) Do(f func([]string, []names.Tag, []state.ActionStatus, int, int) ([]state.OperationInfo, bool, error)) *MockModelListOperationsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelListOperationsCall) DoAndReturn(f func([]string, []names.Tag, []state.ActionStatus, int, int) ([]state.OperationInfo, bool, error)) *MockModelListOperationsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// OperationWithActions mocks base method.
func (m *MockModel) OperationWithActions(arg0 string) (*state.OperationInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OperationWithActions", arg0)
	ret0, _ := ret[0].(*state.OperationInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OperationWithActions indicates an expected call of OperationWithActions.
func (mr *MockModelMockRecorder) OperationWithActions(arg0 any) *MockModelOperationWithActionsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OperationWithActions", reflect.TypeOf((*MockModel)(nil).OperationWithActions), arg0)
	return &MockModelOperationWithActionsCall{Call: call}
}

// MockModelOperationWithActionsCall wrap *gomock.Call
type MockModelOperationWithActionsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelOperationWithActionsCall) Return(arg0 *state.OperationInfo, arg1 error) *MockModelOperationWithActionsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelOperationWithActionsCall) Do(f func(string) (*state.OperationInfo, error)) *MockModelOperationWithActionsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelOperationWithActionsCall) DoAndReturn(f func(string) (*state.OperationInfo, error)) *MockModelOperationWithActionsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

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

// GetAllUnitNames mocks base method.
func (m *MockApplicationService) GetAllUnitNames(arg0 context.Context) ([]unit.Name, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUnitNames", arg0)
	ret0, _ := ret[0].([]unit.Name)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUnitNames indicates an expected call of GetAllUnitNames.
func (mr *MockApplicationServiceMockRecorder) GetAllUnitNames(arg0 any) *MockApplicationServiceGetAllUnitNamesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUnitNames", reflect.TypeOf((*MockApplicationService)(nil).GetAllUnitNames), arg0)
	return &MockApplicationServiceGetAllUnitNamesCall{Call: call}
}

// MockApplicationServiceGetAllUnitNamesCall wrap *gomock.Call
type MockApplicationServiceGetAllUnitNamesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationServiceGetAllUnitNamesCall) Return(arg0 []unit.Name, arg1 error) *MockApplicationServiceGetAllUnitNamesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationServiceGetAllUnitNamesCall) Do(f func(context.Context) ([]unit.Name, error)) *MockApplicationServiceGetAllUnitNamesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationServiceGetAllUnitNamesCall) DoAndReturn(f func(context.Context) ([]unit.Name, error)) *MockApplicationServiceGetAllUnitNamesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetCharmActions mocks base method.
func (m *MockApplicationService) GetCharmActions(arg0 context.Context, arg1 charm.CharmLocator) (charm0.Actions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCharmActions", arg0, arg1)
	ret0, _ := ret[0].(charm0.Actions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCharmActions indicates an expected call of GetCharmActions.
func (mr *MockApplicationServiceMockRecorder) GetCharmActions(arg0, arg1 any) *MockApplicationServiceGetCharmActionsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCharmActions", reflect.TypeOf((*MockApplicationService)(nil).GetCharmActions), arg0, arg1)
	return &MockApplicationServiceGetCharmActionsCall{Call: call}
}

// MockApplicationServiceGetCharmActionsCall wrap *gomock.Call
type MockApplicationServiceGetCharmActionsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationServiceGetCharmActionsCall) Return(arg0 charm0.Actions, arg1 error) *MockApplicationServiceGetCharmActionsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationServiceGetCharmActionsCall) Do(f func(context.Context, charm.CharmLocator) (charm0.Actions, error)) *MockApplicationServiceGetCharmActionsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationServiceGetCharmActionsCall) DoAndReturn(f func(context.Context, charm.CharmLocator) (charm0.Actions, error)) *MockApplicationServiceGetCharmActionsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetCharmLocatorByApplicationName mocks base method.
func (m *MockApplicationService) GetCharmLocatorByApplicationName(arg0 context.Context, arg1 string) (charm.CharmLocator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCharmLocatorByApplicationName", arg0, arg1)
	ret0, _ := ret[0].(charm.CharmLocator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCharmLocatorByApplicationName indicates an expected call of GetCharmLocatorByApplicationName.
func (mr *MockApplicationServiceMockRecorder) GetCharmLocatorByApplicationName(arg0, arg1 any) *MockApplicationServiceGetCharmLocatorByApplicationNameCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCharmLocatorByApplicationName", reflect.TypeOf((*MockApplicationService)(nil).GetCharmLocatorByApplicationName), arg0, arg1)
	return &MockApplicationServiceGetCharmLocatorByApplicationNameCall{Call: call}
}

// MockApplicationServiceGetCharmLocatorByApplicationNameCall wrap *gomock.Call
type MockApplicationServiceGetCharmLocatorByApplicationNameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationServiceGetCharmLocatorByApplicationNameCall) Return(arg0 charm.CharmLocator, arg1 error) *MockApplicationServiceGetCharmLocatorByApplicationNameCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationServiceGetCharmLocatorByApplicationNameCall) Do(f func(context.Context, string) (charm.CharmLocator, error)) *MockApplicationServiceGetCharmLocatorByApplicationNameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationServiceGetCharmLocatorByApplicationNameCall) DoAndReturn(f func(context.Context, string) (charm.CharmLocator, error)) *MockApplicationServiceGetCharmLocatorByApplicationNameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetUnitNamesForApplication mocks base method.
func (m *MockApplicationService) GetUnitNamesForApplication(arg0 context.Context, arg1 string) ([]unit.Name, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUnitNamesForApplication", arg0, arg1)
	ret0, _ := ret[0].([]unit.Name)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUnitNamesForApplication indicates an expected call of GetUnitNamesForApplication.
func (mr *MockApplicationServiceMockRecorder) GetUnitNamesForApplication(arg0, arg1 any) *MockApplicationServiceGetUnitNamesForApplicationCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUnitNamesForApplication", reflect.TypeOf((*MockApplicationService)(nil).GetUnitNamesForApplication), arg0, arg1)
	return &MockApplicationServiceGetUnitNamesForApplicationCall{Call: call}
}

// MockApplicationServiceGetUnitNamesForApplicationCall wrap *gomock.Call
type MockApplicationServiceGetUnitNamesForApplicationCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationServiceGetUnitNamesForApplicationCall) Return(arg0 []unit.Name, arg1 error) *MockApplicationServiceGetUnitNamesForApplicationCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationServiceGetUnitNamesForApplicationCall) Do(f func(context.Context, string) ([]unit.Name, error)) *MockApplicationServiceGetUnitNamesForApplicationCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationServiceGetUnitNamesForApplicationCall) DoAndReturn(f func(context.Context, string) ([]unit.Name, error)) *MockApplicationServiceGetUnitNamesForApplicationCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockModelInfoService is a mock of ModelInfoService interface.
type MockModelInfoService struct {
	ctrl     *gomock.Controller
	recorder *MockModelInfoServiceMockRecorder
}

// MockModelInfoServiceMockRecorder is the mock recorder for MockModelInfoService.
type MockModelInfoServiceMockRecorder struct {
	mock *MockModelInfoService
}

// NewMockModelInfoService creates a new mock instance.
func NewMockModelInfoService(ctrl *gomock.Controller) *MockModelInfoService {
	mock := &MockModelInfoService{ctrl: ctrl}
	mock.recorder = &MockModelInfoServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModelInfoService) EXPECT() *MockModelInfoServiceMockRecorder {
	return m.recorder
}

// GetModelInfo mocks base method.
func (m *MockModelInfoService) GetModelInfo(arg0 context.Context) (model.ModelInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModelInfo", arg0)
	ret0, _ := ret[0].(model.ModelInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetModelInfo indicates an expected call of GetModelInfo.
func (mr *MockModelInfoServiceMockRecorder) GetModelInfo(arg0 any) *MockModelInfoServiceGetModelInfoCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModelInfo", reflect.TypeOf((*MockModelInfoService)(nil).GetModelInfo), arg0)
	return &MockModelInfoServiceGetModelInfoCall{Call: call}
}

// MockModelInfoServiceGetModelInfoCall wrap *gomock.Call
type MockModelInfoServiceGetModelInfoCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelInfoServiceGetModelInfoCall) Return(arg0 model.ModelInfo, arg1 error) *MockModelInfoServiceGetModelInfoCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelInfoServiceGetModelInfoCall) Do(f func(context.Context) (model.ModelInfo, error)) *MockModelInfoServiceGetModelInfoCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelInfoServiceGetModelInfoCall) DoAndReturn(f func(context.Context) (model.ModelInfo, error)) *MockModelInfoServiceGetModelInfoCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
