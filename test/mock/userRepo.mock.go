// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/fxmbx/go_practice_pr/repository (interfaces: UserRepository)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	reflect "reflect"

	models "github.com/fxmbx/go_practice_pr/models"
	utils "github.com/fxmbx/go_practice_pr/utils"
	gomock "github.com/golang/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// DeleteByUserID mocks base method.
func (m *MockUserRepository) DeleteByUserID(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByUserID", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByUserID indicates an expected call of DeleteByUserID.
func (mr *MockUserRepositoryMockRecorder) DeleteByUserID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByUserID", reflect.TypeOf((*MockUserRepository)(nil).DeleteByUserID), arg0)
}

// FindUserByEmail mocks base method.
func (m *MockUserRepository) FindUserByEmail(arg0 string) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByEmail", arg0)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByEmail indicates an expected call of FindUserByEmail.
func (mr *MockUserRepositoryMockRecorder) FindUserByEmail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByEmail", reflect.TypeOf((*MockUserRepository)(nil).FindUserByEmail), arg0)
}

// FindUserByUserID mocks base method.
func (m *MockUserRepository) FindUserByUserID(arg0 string) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByUserID", arg0)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByUserID indicates an expected call of FindUserByUserID.
func (mr *MockUserRepositoryMockRecorder) FindUserByUserID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByUserID", reflect.TypeOf((*MockUserRepository)(nil).FindUserByUserID), arg0)
}

// InsertUser mocks base method.
func (m *MockUserRepository) InsertUser(arg0 models.User) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertUser", arg0)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertUser indicates an expected call of InsertUser.
func (mr *MockUserRepositoryMockRecorder) InsertUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertUser", reflect.TypeOf((*MockUserRepository)(nil).InsertUser), arg0)
}

// ListAllUsers mocks base method.
func (m *MockUserRepository) ListAllUsers(arg0 utils.Pagination) (*utils.Pagination, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllUsers", arg0)
	ret0, _ := ret[0].(*utils.Pagination)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllUsers indicates an expected call of ListAllUsers.
func (mr *MockUserRepositoryMockRecorder) ListAllUsers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllUsers", reflect.TypeOf((*MockUserRepository)(nil).ListAllUsers), arg0)
}

// ListAllUsersIncludeDeleted mocks base method.
func (m *MockUserRepository) ListAllUsersIncludeDeleted(arg0 utils.Pagination) (*utils.Pagination, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllUsersIncludeDeleted", arg0)
	ret0, _ := ret[0].(*utils.Pagination)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllUsersIncludeDeleted indicates an expected call of ListAllUsersIncludeDeleted.
func (mr *MockUserRepositoryMockRecorder) ListAllUsersIncludeDeleted(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllUsersIncludeDeleted", reflect.TypeOf((*MockUserRepository)(nil).ListAllUsersIncludeDeleted), arg0)
}

// UpdatetUser mocks base method.
func (m *MockUserRepository) UpdatetUser(arg0 models.User) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatetUser", arg0)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatetUser indicates an expected call of UpdatetUser.
func (mr *MockUserRepositoryMockRecorder) UpdatetUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatetUser", reflect.TypeOf((*MockUserRepository)(nil).UpdatetUser), arg0)
}
