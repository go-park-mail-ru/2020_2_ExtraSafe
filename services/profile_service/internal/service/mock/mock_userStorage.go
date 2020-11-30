// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/go-park-mail-ru/2020_2_ExtraSafe/services/profile_service/internal/userStorage (interfaces: Storage)

// Package mock is a generated GoMock package.
package mock

import (
	models "github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockStorage is a mock of Storage interface
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// ChangeUserPassword mocks base method
func (m *MockStorage) ChangeUserPassword(arg0 models.UserInputPassword) (models.UserOutside, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeUserPassword", arg0)
	ret0, _ := ret[0].(models.UserOutside)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeUserPassword indicates an expected call of ChangeUserPassword
func (mr *MockStorageMockRecorder) ChangeUserPassword(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeUserPassword", reflect.TypeOf((*MockStorage)(nil).ChangeUserPassword), arg0)
}

// ChangeUserProfile mocks base method
func (m *MockStorage) ChangeUserProfile(arg0 models.UserInputProfile, arg1 models.UserAvatar) (models.UserOutside, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeUserProfile", arg0, arg1)
	ret0, _ := ret[0].(models.UserOutside)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeUserProfile indicates an expected call of ChangeUserProfile
func (mr *MockStorageMockRecorder) ChangeUserProfile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeUserProfile", reflect.TypeOf((*MockStorage)(nil).ChangeUserProfile), arg0, arg1)
}

// CheckExistingUser mocks base method
func (m *MockStorage) CheckExistingUser(arg0, arg1 string) models.MultiErrors {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckExistingUser", arg0, arg1)
	ret0, _ := ret[0].(models.MultiErrors)
	return ret0
}

// CheckExistingUser indicates an expected call of CheckExistingUser
func (mr *MockStorageMockRecorder) CheckExistingUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckExistingUser", reflect.TypeOf((*MockStorage)(nil).CheckExistingUser), arg0, arg1)
}

// CheckUser mocks base method
func (m *MockStorage) CheckUser(arg0 models.UserInputLogin) (int64, models.UserOutside, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUser", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(models.UserOutside)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CheckUser indicates an expected call of CheckUser
func (mr *MockStorageMockRecorder) CheckUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUser", reflect.TypeOf((*MockStorage)(nil).CheckUser), arg0)
}

// CreateUser mocks base method
func (m *MockStorage) CreateUser(arg0 models.UserInputReg) (int64, models.UserOutside, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(models.UserOutside)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateUser indicates an expected call of CreateUser
func (mr *MockStorageMockRecorder) CreateUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStorage)(nil).CreateUser), arg0)
}

// GetInternalUser mocks base method
func (m *MockStorage) GetInternalUser(arg0 models.UserInput) (models.UserOutside, []byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInternalUser", arg0)
	ret0, _ := ret[0].(models.UserOutside)
	ret1, _ := ret[1].([]byte)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetInternalUser indicates an expected call of GetInternalUser
func (mr *MockStorageMockRecorder) GetInternalUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInternalUser", reflect.TypeOf((*MockStorage)(nil).GetInternalUser), arg0)
}

// GetUserAvatar mocks base method
func (m *MockStorage) GetUserAvatar(arg0 models.UserInput) (models.UserAvatar, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserAvatar", arg0)
	ret0, _ := ret[0].(models.UserAvatar)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserAvatar indicates an expected call of GetUserAvatar
func (mr *MockStorageMockRecorder) GetUserAvatar(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserAvatar", reflect.TypeOf((*MockStorage)(nil).GetUserAvatar), arg0)
}

// GetUserByUsername mocks base method
func (m *MockStorage) GetUserByUsername(arg0 string) (models.UserInternal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByUsername", arg0)
	ret0, _ := ret[0].(models.UserInternal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByUsername indicates an expected call of GetUserByUsername
func (mr *MockStorageMockRecorder) GetUserByUsername(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByUsername", reflect.TypeOf((*MockStorage)(nil).GetUserByUsername), arg0)
}

// GetUserProfile mocks base method
func (m *MockStorage) GetUserProfile(arg0 models.UserInput) (models.UserOutside, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserProfile", arg0)
	ret0, _ := ret[0].(models.UserOutside)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserProfile indicates an expected call of GetUserProfile
func (mr *MockStorageMockRecorder) GetUserProfile(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserProfile", reflect.TypeOf((*MockStorage)(nil).GetUserProfile), arg0)
}

// GetUsersByIDs mocks base method
func (m *MockStorage) GetUsersByIDs(arg0 []int64) ([]models.UserOutsideShort, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsersByIDs", arg0)
	ret0, _ := ret[0].([]models.UserOutsideShort)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsersByIDs indicates an expected call of GetUsersByIDs
func (mr *MockStorageMockRecorder) GetUsersByIDs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsersByIDs", reflect.TypeOf((*MockStorage)(nil).GetUsersByIDs), arg0)
}

// checkExistingUserOnUpdate mocks base method
func (m *MockStorage) checkExistingUserOnUpdate(arg0, arg1 string, arg2 int64) models.MultiErrors {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "checkExistingUserOnUpdate", arg0, arg1, arg2)
	ret0, _ := ret[0].(models.MultiErrors)
	return ret0
}

// checkExistingUserOnUpdate indicates an expected call of checkExistingUserOnUpdate
func (mr *MockStorageMockRecorder) checkExistingUserOnUpdate(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "checkExistingUserOnUpdate", reflect.TypeOf((*MockStorage)(nil).checkExistingUserOnUpdate), arg0, arg1, arg2)
}
