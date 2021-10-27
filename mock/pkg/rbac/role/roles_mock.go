// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/rbac/role/roles.go

// Package mock_role is a generated GoMock package.
package mock_role

import (
	context "context"
	reflect "reflect"

	role "g.hz.netease.com/horizon/pkg/rbac/role"
	types "g.hz.netease.com/horizon/pkg/rbac/types"
	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// GetRole mocks base method.
func (m *MockService) GetRole(ctx context.Context, roleName string) (*types.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRole", ctx, roleName)
	ret0, _ := ret[0].(*types.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRole indicates an expected call of GetRole.
func (mr *MockServiceMockRecorder) GetRole(ctx, roleName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRole", reflect.TypeOf((*MockService)(nil).GetRole), ctx, roleName)
}

// ListRole mocks base method.
func (m *MockService) ListRole(ctx context.Context) ([]types.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRole", ctx)
	ret0, _ := ret[0].([]types.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRole indicates an expected call of ListRole.
func (mr *MockServiceMockRecorder) ListRole(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRole", reflect.TypeOf((*MockService)(nil).ListRole), ctx)
}

// RoleCompare mocks base method.
func (m *MockService) RoleCompare(ctx context.Context, role1, role2 string) (role.RoleCompResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RoleCompare", ctx, role1, role2)
	ret0, _ := ret[0].(role.RoleCompResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RoleCompare indicates an expected call of RoleCompare.
func (mr *MockServiceMockRecorder) RoleCompare(ctx, role1, role2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RoleCompare", reflect.TypeOf((*MockService)(nil).RoleCompare), ctx, role1, role2)
}
