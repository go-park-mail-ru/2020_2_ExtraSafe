// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/go-park-mail-ru/2020_2_ExtraSafe/cmd/handlers/auth (interfaces: AuthService)

// Package mock is a generated GoMock package.
package mock

import (
	models "github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAuthService is a mock of AuthService interface
type MockAuthService struct {
	ctrl     *gomock.Controller
	recorder *MockAuthServiceMockRecorder
}

// MockAuthServiceMockRecorder is the mock recorder for MockAuthService
type MockAuthServiceMockRecorder struct {
	mock *MockAuthService
}

// NewMockAuthService creates a new mock instance
func NewMockAuthService(ctrl *gomock.Controller) *MockAuthService {
	mock := &MockAuthService{ctrl: ctrl}
	mock.recorder = &MockAuthServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthService) EXPECT() *MockAuthServiceMockRecorder {
	return m.recorder
}

// Auth mocks base method
func (m *MockAuthService) Auth(arg0 models.UserInput) (models.UserBoardsOutside, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Auth", arg0)
	ret0, _ := ret[0].(models.UserBoardsOutside)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Auth indicates an expected call of Auth
func (mr *MockAuthServiceMockRecorder) Auth(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Auth", reflect.TypeOf((*MockAuthService)(nil).Auth), arg0)
}

// Login mocks base method
func (m *MockAuthService) Login(arg0 models.UserInputLogin) (uint64, models.UserOutside, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", arg0)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(models.UserOutside)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Login indicates an expected call of Login
func (mr *MockAuthServiceMockRecorder) Login(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockAuthService)(nil).Login), arg0)
}

// Registration mocks base method
func (m *MockAuthService) Registration(arg0 models.UserInputReg) (uint64, models.UserOutside, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Registration", arg0)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(models.UserOutside)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Registration indicates an expected call of Registration
func (mr *MockAuthServiceMockRecorder) Registration(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Registration", reflect.TypeOf((*MockAuthService)(nil).Registration), arg0)
}
