// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/go-park-mail-ru/2020_2_ExtraSafe/services/auth_service/internal/sessionsStorage (interfaces: SessionStorage)

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockSessionStorage is a mock of SessionStorage interface
type MockSessionStorage struct {
	ctrl     *gomock.Controller
	recorder *MockSessionStorageMockRecorder
}

// MockSessionStorageMockRecorder is the mock recorder for MockSessionStorage
type MockSessionStorageMockRecorder struct {
	mock *MockSessionStorage
}

// NewMockSessionStorage creates a new mock instance
func NewMockSessionStorage(ctrl *gomock.Controller) *MockSessionStorage {
	mock := &MockSessionStorage{ctrl: ctrl}
	mock.recorder = &MockSessionStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSessionStorage) EXPECT() *MockSessionStorageMockRecorder {
	return m.recorder
}

// CheckUserSession mocks base method
func (m *MockSessionStorage) CheckUserSession(arg0 string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUserSession", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckUserSession indicates an expected call of CheckUserSession
func (mr *MockSessionStorageMockRecorder) CheckUserSession(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUserSession", reflect.TypeOf((*MockSessionStorage)(nil).CheckUserSession), arg0)
}

// CreateUserSession mocks base method
func (m *MockSessionStorage) CreateUserSession(arg0 int64, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserSession", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUserSession indicates an expected call of CreateUserSession
func (mr *MockSessionStorageMockRecorder) CreateUserSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserSession", reflect.TypeOf((*MockSessionStorage)(nil).CreateUserSession), arg0, arg1)
}

// DeleteUserSession mocks base method
func (m *MockSessionStorage) DeleteUserSession(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserSession", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUserSession indicates an expected call of DeleteUserSession
func (mr *MockSessionStorageMockRecorder) DeleteUserSession(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserSession", reflect.TypeOf((*MockSessionStorage)(nil).DeleteUserSession), arg0)
}
