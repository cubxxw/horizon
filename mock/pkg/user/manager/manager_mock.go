// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/user/manager/manager.go

// Package mock_manager is a generated GoMock package.
package mock_manager

import (
	context "context"
	q "g.hz.netease.com/horizon/lib/q"
	models "g.hz.netease.com/horizon/pkg/user/models"
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
func (m *MockManager) Create(ctx context.Context, user *models.User) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, user)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockManagerMockRecorder) Create(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockManager)(nil).Create), ctx, user)
}

// GetByOIDCMeta mocks base method
func (m *MockManager) GetByOIDCMeta(ctx context.Context, oidcType, email string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByOIDCMeta", ctx, oidcType, email)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByOIDCMeta indicates an expected call of GetByOIDCMeta
func (mr *MockManagerMockRecorder) GetByOIDCMeta(ctx, oidcType, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByOIDCMeta", reflect.TypeOf((*MockManager)(nil).GetByOIDCMeta), ctx, oidcType, email)
}

// SearchUser mocks base method
func (m *MockManager) SearchUser(ctx context.Context, filter string, query *q.Query) (int, []models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchUser", ctx, filter, query)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].([]models.User)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SearchUser indicates an expected call of SearchUser
func (mr *MockManagerMockRecorder) SearchUser(ctx, filter, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchUser", reflect.TypeOf((*MockManager)(nil).SearchUser), ctx, filter, query)
}

// GetUserByEmail mocks base method
func (m *MockManager) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", ctx, email)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail
func (mr *MockManagerMockRecorder) GetUserByEmail(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockManager)(nil).GetUserByEmail), ctx, email)
}

// GetUserByID mocks base method
func (m *MockManager) GetUserByID(ctx context.Context, userID uint) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", ctx, userID)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID
func (mr *MockManagerMockRecorder) GetUserByID(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockManager)(nil).GetUserByID), ctx, userID)
}

// GetUserByIDs mocks base method
func (m *MockManager) GetUserByIDs(ctx context.Context, userIDs []uint) ([]models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByIDs", ctx, userIDs)
	ret0, _ := ret[0].([]models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByIDs indicates an expected call of GetUserByIDs
func (mr *MockManagerMockRecorder) GetUserByIDs(ctx, userIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByIDs", reflect.TypeOf((*MockManager)(nil).GetUserByIDs), ctx, userIDs)
}

// GetUserMapByIDs mocks base method
func (m *MockManager) GetUserMapByIDs(ctx context.Context, userIDs []uint) (map[uint]*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserMapByIDs", ctx, userIDs)
	ret0, _ := ret[0].(map[uint]*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserMapByIDs indicates an expected call of GetUserMapByIDs
func (mr *MockManagerMockRecorder) GetUserMapByIDs(ctx, userIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserMapByIDs", reflect.TypeOf((*MockManager)(nil).GetUserMapByIDs), ctx, userIDs)
}

// ListByEmail mocks base method
func (m *MockManager) ListByEmail(ctx context.Context, emails []string) ([]*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByEmail", ctx, emails)
	ret0, _ := ret[0].([]*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByEmail indicates an expected call of ListByEmail
func (mr *MockManagerMockRecorder) ListByEmail(ctx, emails interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByEmail", reflect.TypeOf((*MockManager)(nil).ListByEmail), ctx, emails)
}
