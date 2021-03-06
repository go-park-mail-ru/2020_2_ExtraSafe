// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/services/profile (interfaces: Validator)

// Package mock is a generated GoMock package.
package mock

import (
	models "github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockValidator is a mock of Validator interface
type MockValidator struct {
	ctrl     *gomock.Controller
	recorder *MockValidatorMockRecorder
}

// MockValidatorMockRecorder is the mock recorder for MockValidator
type MockValidatorMockRecorder struct {
	mock *MockValidator
}

// NewMockValidator creates a new mock instance
func NewMockValidator(ctrl *gomock.Controller) *MockValidator {
	mock := &MockValidator{ctrl: ctrl}
	mock.recorder = &MockValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockValidator) EXPECT() *MockValidatorMockRecorder {
	return m.recorder
}

// ValidateChangePassword mocks base method
func (m *MockValidator) ValidateChangePassword(arg0 models.UserInputPassword) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateChangePassword", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateChangePassword indicates an expected call of ValidateChangePassword
func (mr *MockValidatorMockRecorder) ValidateChangePassword(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateChangePassword", reflect.TypeOf((*MockValidator)(nil).ValidateChangePassword), arg0)
}

// ValidateProfile mocks base method
func (m *MockValidator) ValidateProfile(arg0 models.UserInputProfile) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateProfile", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateProfile indicates an expected call of ValidateProfile
func (mr *MockValidatorMockRecorder) ValidateProfile(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateProfile", reflect.TypeOf((*MockValidator)(nil).ValidateProfile), arg0)
}
