// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/storages/boardStorage (interfaces: TasksStorage)

// Package mock is a generated GoMock package.
package mock

import (
	models "github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockTasksStorage is a mock of TasksStorage interface
type MockTasksStorage struct {
	ctrl     *gomock.Controller
	recorder *MockTasksStorageMockRecorder
}

// MockTasksStorageMockRecorder is the mock recorder for MockTasksStorage
type MockTasksStorageMockRecorder struct {
	mock *MockTasksStorage
}

// NewMockTasksStorage creates a new mock instance
func NewMockTasksStorage(ctrl *gomock.Controller) *MockTasksStorage {
	mock := &MockTasksStorage{ctrl: ctrl}
	mock.recorder = &MockTasksStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTasksStorage) EXPECT() *MockTasksStorageMockRecorder {
	return m.recorder
}

// ChangeTask mocks base method
func (m *MockTasksStorage) ChangeTask(arg0 models.TaskInput) (models.TaskOutside, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeTask", arg0)
	ret0, _ := ret[0].(models.TaskOutside)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeTask indicates an expected call of ChangeTask
func (mr *MockTasksStorageMockRecorder) ChangeTask(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeTask", reflect.TypeOf((*MockTasksStorage)(nil).ChangeTask), arg0)
}

// CreateTask mocks base method
func (m *MockTasksStorage) CreateTask(arg0 models.TaskInput) (models.TaskOutside, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTask", arg0)
	ret0, _ := ret[0].(models.TaskOutside)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTask indicates an expected call of CreateTask
func (mr *MockTasksStorageMockRecorder) CreateTask(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTask", reflect.TypeOf((*MockTasksStorage)(nil).CreateTask), arg0)
}

// DeleteTask mocks base method
func (m *MockTasksStorage) DeleteTask(arg0 models.TaskInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTask", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTask indicates an expected call of DeleteTask
func (mr *MockTasksStorageMockRecorder) DeleteTask(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTask", reflect.TypeOf((*MockTasksStorage)(nil).DeleteTask), arg0)
}

// GetTaskByID mocks base method
func (m *MockTasksStorage) GetTaskByID(arg0 models.TaskInput) (models.TaskOutside, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTaskByID", arg0)
	ret0, _ := ret[0].(models.TaskOutside)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTaskByID indicates an expected call of GetTaskByID
func (mr *MockTasksStorageMockRecorder) GetTaskByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTaskByID", reflect.TypeOf((*MockTasksStorage)(nil).GetTaskByID), arg0)
}

// GetTasksByCard mocks base method
func (m *MockTasksStorage) GetTasksByCard(arg0 models.CardInput) ([]models.TaskOutside, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTasksByCard", arg0)
	ret0, _ := ret[0].([]models.TaskOutside)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTasksByCard indicates an expected call of GetTasksByCard
func (mr *MockTasksStorageMockRecorder) GetTasksByCard(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTasksByCard", reflect.TypeOf((*MockTasksStorage)(nil).GetTasksByCard), arg0)
}
