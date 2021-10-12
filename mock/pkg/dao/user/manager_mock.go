// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/dao/userdao/manager.go

// Package mock_user is a generated GoMock package.
package mock_user

import (
	context "context"
	userdao "g.hz.netease.com/horizon/pkg/dao/user"
	q "g.hz.netease.com/horizon/pkg/lib/q"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockManager is a mock of Manager interface
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockManager) Create(ctx context.Context, user *userdao.User) (*userdao.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, user)
	ret0, _ := ret[0].(*userdao.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockManagerMockRecorder) Create(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockManager)(nil).Create), ctx, user)
}

// GetByOIDCMeta mocks base method
func (m *MockManager) GetByOIDCMeta(ctx context.Context, oidcID, oidcType string) (*userdao.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByOIDCMeta", ctx, oidcID, oidcType)
	ret0, _ := ret[0].(*userdao.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByOIDCMeta indicates an expected call of GetByOIDCMeta
func (mr *MockManagerMockRecorder) GetByOIDCMeta(ctx, oidcID, oidcType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByOIDCMeta", reflect.TypeOf((*MockManager)(nil).GetByOIDCMeta), ctx, oidcID, oidcType)
}

// SearchUser mocks base method
func (m *MockManager) SearchUser(ctx context.Context, filter string, query *q.Query) (int, []userdao.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchUser", ctx, filter, query)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].([]userdao.User)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SearchUser indicates an expected call of SearchUser
func (mr *MockManagerMockRecorder) SearchUser(ctx, filter, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchUser", reflect.TypeOf((*MockManager)(nil).SearchUser), ctx, filter, query)
}