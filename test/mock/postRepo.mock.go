// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/fxmbx/go_practice_pr/repository (interfaces: PostRepository)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	reflect "reflect"

	models "github.com/fxmbx/go_practice_pr/models"
	utils "github.com/fxmbx/go_practice_pr/utils"
	gomock "github.com/golang/mock/gomock"
)

// MockPostRepository is a mock of PostRepository interface.
type MockPostRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPostRepositoryMockRecorder
}

// MockPostRepositoryMockRecorder is the mock recorder for MockPostRepository.
type MockPostRepositoryMockRecorder struct {
	mock *MockPostRepository
}

// NewMockPostRepository creates a new mock instance.
func NewMockPostRepository(ctrl *gomock.Controller) *MockPostRepository {
	mock := &MockPostRepository{ctrl: ctrl}
	mock.recorder = &MockPostRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPostRepository) EXPECT() *MockPostRepositoryMockRecorder {
	return m.recorder
}

// DeletePost mocks base method.
func (m *MockPostRepository) DeletePost(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePost", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePost indicates an expected call of DeletePost.
func (mr *MockPostRepositoryMockRecorder) DeletePost(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePost", reflect.TypeOf((*MockPostRepository)(nil).DeletePost), arg0)
}

// FindPostByPostID mocks base method.
func (m *MockPostRepository) FindPostByPostID(arg0 string) (models.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindPostByPostID", arg0)
	ret0, _ := ret[0].(models.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindPostByPostID indicates an expected call of FindPostByPostID.
func (mr *MockPostRepositoryMockRecorder) FindPostByPostID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindPostByPostID", reflect.TypeOf((*MockPostRepository)(nil).FindPostByPostID), arg0)
}

// FindPostByTitle mocks base method.
func (m *MockPostRepository) FindPostByTitle(arg0 string) (models.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindPostByTitle", arg0)
	ret0, _ := ret[0].(models.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindPostByTitle indicates an expected call of FindPostByTitle.
func (mr *MockPostRepositoryMockRecorder) FindPostByTitle(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindPostByTitle", reflect.TypeOf((*MockPostRepository)(nil).FindPostByTitle), arg0)
}

// InserPost mocks base method.
func (m *MockPostRepository) InserPost(arg0 models.Post) (models.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InserPost", arg0)
	ret0, _ := ret[0].(models.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InserPost indicates an expected call of InserPost.
func (mr *MockPostRepositoryMockRecorder) InserPost(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InserPost", reflect.TypeOf((*MockPostRepository)(nil).InserPost), arg0)
}

// ListPost mocks base method.
func (m *MockPostRepository) ListPost(arg0 utils.Pagination) (*utils.Pagination, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPost", arg0)
	ret0, _ := ret[0].(*utils.Pagination)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPost indicates an expected call of ListPost.
func (mr *MockPostRepositoryMockRecorder) ListPost(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPost", reflect.TypeOf((*MockPostRepository)(nil).ListPost), arg0)
}

// UpdatePost mocks base method.
func (m *MockPostRepository) UpdatePost(arg0 models.Post) (models.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePost", arg0)
	ret0, _ := ret[0].(models.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePost indicates an expected call of UpdatePost.
func (mr *MockPostRepositoryMockRecorder) UpdatePost(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePost", reflect.TypeOf((*MockPostRepository)(nil).UpdatePost), arg0)
}