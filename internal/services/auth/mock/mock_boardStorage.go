// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/services/auth (interfaces: BoardStorage)

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

// GetBoardsList mocks base method
func (m *MockBoardStorage) GetBoardsList(arg0 models.UserInput) ([]models.BoardOutsideShort, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBoardsList", arg0)
	ret0, _ := ret[0].([]models.BoardOutsideShort)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBoardsList indicates an expected call of GetBoardsList
func (mr *MockBoardStorageMockRecorder) GetBoardsList(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBoardsList", reflect.TypeOf((*MockBoardStorage)(nil).GetBoardsList), arg0)
}
