// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/services/boards (interfaces: BoardStorage)

// Package mock is a generated GoMock package.
package mock

import (
	models "github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockBoardStorage is a mock of BoardStorage interface
type MockBoardStorage struct {
	ctrl     *gomock.Controller
	recorder *MockBoardStorageMockRecorder
}

// MockBoardStorageMockRecorder is the mock recorder for MockBoardStorage
type MockBoardStorageMockRecorder struct {
	mock *MockBoardStorage
}

// NewMockBoardStorage creates a new mock instance
func NewMockBoardStorage(ctrl *gomock.Controller) *MockBoardStorage {
	mock := &MockBoardStorage{ctrl: ctrl}
	mock.recorder = &MockBoardStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBoardStorage) EXPECT() *MockBoardStorageMockRecorder {
	return m.recorder
}

// ChangeBoard mocks base method
func (m *MockBoardStorage) ChangeBoard(arg0 models.BoardChangeInput) (models.BoardInternal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeBoard", arg0)
	ret0, _ := ret[0].(models.BoardInternal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeBoard indicates an expected call of ChangeBoard
func (mr *MockBoardStorageMockRecorder) ChangeBoard(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeBoard", reflect.TypeOf((*MockBoardStorage)(nil).ChangeBoard), arg0)
}

// ChangeCard mocks base method
func (m *MockBoardStorage) ChangeCard(arg0 models.CardInput) (models.CardOutside, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeCard", arg0)
	ret0, _ := ret[0].(models.CardOutside)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeCard indicates an expected call of ChangeCard
func (mr *MockBoardStorageMockRecorder) ChangeCard(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeCard", reflect.TypeOf((*MockBoardStorage)(nil).ChangeCard), arg0)
}

// ChangeTask mocks base method
func (m *MockBoardStorage) ChangeTask(arg0 models.TaskInput) (models.TaskOutside, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeTask", arg0)
	ret0, _ := ret[0].(models.TaskOutside)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeTask indicates an expected call of ChangeTask
func (mr *MockBoardStorageMockRecorder) ChangeTask(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeTask", reflect.TypeOf((*MockBoardStorage)(nil).ChangeTask), arg0)
}

// CreateBoard mocks base method
func (m *MockBoardStorage) CreateBoard(arg0 models.BoardChangeInput) (models.BoardInternal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBoard", arg0)
	ret0, _ := ret[0].(models.BoardInternal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBoard indicates an expected call of CreateBoard
func (mr *MockBoardStorageMockRecorder) CreateBoard(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBoard", reflect.TypeOf((*MockBoardStorage)(nil).CreateBoard), arg0)
}

// CreateCard mocks base method
func (m *MockBoardStorage) CreateCard(arg0 models.CardInput) (models.CardOutside, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCard", arg0)
	ret0, _ := ret[0].(models.CardOutside)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCard indicates an expected call of CreateCard
func (mr *MockBoardStorageMockRecorder) CreateCard(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCard", reflect.TypeOf((*MockBoardStorage)(nil).CreateCard), arg0)
}

// CreateTask mocks base method
func (m *MockBoardStorage) CreateTask(arg0 models.TaskInput) (models.TaskOutside, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTask", arg0)
	ret0, _ := ret[0].(models.TaskOutside)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTask indicates an expected call of CreateTask
func (mr *MockBoardStorageMockRecorder) CreateTask(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTask", reflect.TypeOf((*MockBoardStorage)(nil).CreateTask), arg0)
}

// DeleteBoard mocks base method
func (m *MockBoardStorage) DeleteBoard(arg0 models.BoardInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBoard", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBoard indicates an expected call of DeleteBoard
func (mr *MockBoardStorageMockRecorder) DeleteBoard(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBoard", reflect.TypeOf((*MockBoardStorage)(nil).DeleteBoard), arg0)
}

// DeleteCard mocks base method
func (m *MockBoardStorage) DeleteCard(arg0 models.CardInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCard", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCard indicates an expected call of DeleteCard
func (mr *MockBoardStorageMockRecorder) DeleteCard(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCard", reflect.TypeOf((*MockBoardStorage)(nil).DeleteCard), arg0)
}

// DeleteTask mocks base method
func (m *MockBoardStorage) DeleteTask(arg0 models.TaskInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTask", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTask indicates an expected call of DeleteTask
func (mr *MockBoardStorageMockRecorder) DeleteTask(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTask", reflect.TypeOf((*MockBoardStorage)(nil).DeleteTask), arg0)
}

// GetBoard mocks base method
func (m *MockBoardStorage) GetBoard(arg0 models.BoardInput) (models.BoardInternal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBoard", arg0)
	ret0, _ := ret[0].(models.BoardInternal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBoard indicates an expected call of GetBoard
func (mr *MockBoardStorageMockRecorder) GetBoard(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBoard", reflect.TypeOf((*MockBoardStorage)(nil).GetBoard), arg0)
}

// GetCard mocks base method
func (m *MockBoardStorage) GetCard(arg0 models.CardInput) (models.CardOutside, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCard", arg0)
	ret0, _ := ret[0].(models.CardOutside)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCard indicates an expected call of GetCard
func (mr *MockBoardStorageMockRecorder) GetCard(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCard", reflect.TypeOf((*MockBoardStorage)(nil).GetCard), arg0)
}

// GetTask mocks base method
func (m *MockBoardStorage) GetTask(arg0 models.TaskInput) (models.TaskOutside, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTask", arg0)
	ret0, _ := ret[0].(models.TaskOutside)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTask indicates an expected call of GetTask
func (mr *MockBoardStorageMockRecorder) GetTask(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTask", reflect.TypeOf((*MockBoardStorage)(nil).GetTask), arg0)
}
