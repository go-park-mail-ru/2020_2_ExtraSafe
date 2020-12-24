// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/services/boards (interfaces: ServiceBoard)

// Package mock is a generated GoMock package.
package mock

import (
	models "github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	gomock "github.com/golang/mock/gomock"
	echo "github.com/labstack/echo"
	reflect "reflect"
)

// MockServiceBoard is a mock of ServiceBoard interface
type MockServiceBoard struct {
	ctrl     *gomock.Controller
	recorder *MockServiceBoardMockRecorder
}

// MockServiceBoardMockRecorder is the mock recorder for MockServiceBoard
type MockServiceBoardMockRecorder struct {
	mock *MockServiceBoard
}

// NewMockServiceBoard creates a new mock instance
func NewMockServiceBoard(ctrl *gomock.Controller) *MockServiceBoard {
	mock := &MockServiceBoard{ctrl: ctrl}
	mock.recorder = &MockServiceBoardMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServiceBoard) EXPECT() *MockServiceBoardMockRecorder {
	return m.recorder
}

// AddMember mocks base method
func (m *MockServiceBoard) AddMember(arg0 models.BoardMemberInput) (models.UserOutsideShort, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddMember", arg0)
	ret0, _ := ret[0].(models.UserOutsideShort)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddMember indicates an expected call of AddMember
func (mr *MockServiceBoardMockRecorder) AddMember(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMember", reflect.TypeOf((*MockServiceBoard)(nil).AddMember), arg0)
}

// AddTag mocks base method
func (m *MockServiceBoard) AddTag(arg0 models.TaskTagInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTag", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddTag indicates an expected call of AddTag
func (mr *MockServiceBoardMockRecorder) AddTag(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTag", reflect.TypeOf((*MockServiceBoard)(nil).AddTag), arg0)
}

// AssignUser mocks base method
func (m *MockServiceBoard) AssignUser(arg0 models.TaskAssignerInput) (models.UserOutsideShort, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssignUser", arg0)
	ret0, _ := ret[0].(models.UserOutsideShort)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssignUser indicates an expected call of AssignUser
func (mr *MockServiceBoardMockRecorder) AssignUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignUser", reflect.TypeOf((*MockServiceBoard)(nil).AssignUser), arg0)
}

// CardOrderChange mocks base method
func (m *MockServiceBoard) CardOrderChange(arg0 models.CardsOrderInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CardOrderChange", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CardOrderChange indicates an expected call of CardOrderChange
func (mr *MockServiceBoardMockRecorder) CardOrderChange(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CardOrderChange", reflect.TypeOf((*MockServiceBoard)(nil).CardOrderChange), arg0)
}

// ChangeBoard mocks base method
func (m *MockServiceBoard) ChangeBoard(arg0 models.BoardChangeInput) (models.BoardOutsideShort, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeBoard", arg0)
	ret0, _ := ret[0].(models.BoardOutsideShort)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeBoard indicates an expected call of ChangeBoard
func (mr *MockServiceBoardMockRecorder) ChangeBoard(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeBoard", reflect.TypeOf((*MockServiceBoard)(nil).ChangeBoard), arg0)
}

// ChangeCard mocks base method
func (m *MockServiceBoard) ChangeCard(arg0 models.CardInput) (models.CardOutsideShort, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeCard", arg0)
	ret0, _ := ret[0].(models.CardOutsideShort)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeCard indicates an expected call of ChangeCard
func (mr *MockServiceBoardMockRecorder) ChangeCard(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeCard", reflect.TypeOf((*MockServiceBoard)(nil).ChangeCard), arg0)
}

// ChangeChecklist mocks base method
func (m *MockServiceBoard) ChangeChecklist(arg0 models.ChecklistInput) (models.ChecklistOutside, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeChecklist", arg0)
	ret0, _ := ret[0].(models.ChecklistOutside)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeChecklist indicates an expected call of ChangeChecklist
func (mr *MockServiceBoardMockRecorder) ChangeChecklist(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeChecklist", reflect.TypeOf((*MockServiceBoard)(nil).ChangeChecklist), arg0)
}

// ChangeComment mocks base method
func (m *MockServiceBoard) ChangeComment(arg0 models.CommentInput) (models.CommentOutside, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeComment", arg0)
	ret0, _ := ret[0].(models.CommentOutside)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeComment indicates an expected call of ChangeComment
func (mr *MockServiceBoardMockRecorder) ChangeComment(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeComment", reflect.TypeOf((*MockServiceBoard)(nil).ChangeComment), arg0)
}

// ChangeTag mocks base method
func (m *MockServiceBoard) ChangeTag(arg0 models.TagInput) (models.TagOutside, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeTag", arg0)
	ret0, _ := ret[0].(models.TagOutside)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeTag indicates an expected call of ChangeTag
func (mr *MockServiceBoardMockRecorder) ChangeTag(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeTag", reflect.TypeOf((*MockServiceBoard)(nil).ChangeTag), arg0)
}

// ChangeTask mocks base method
func (m *MockServiceBoard) ChangeTask(arg0 models.TaskInput) (models.TaskOutsideSuperShort, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeTask", arg0)
	ret0, _ := ret[0].(models.TaskOutsideSuperShort)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeTask indicates an expected call of ChangeTask
func (mr *MockServiceBoardMockRecorder) ChangeTask(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeTask", reflect.TypeOf((*MockServiceBoard)(nil).ChangeTask), arg0)
}

// CheckBoardPermission mocks base method
func (m *MockServiceBoard) CheckBoardPermission(arg0, arg1 int64, arg2 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckBoardPermission", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckBoardPermission indicates an expected call of CheckBoardPermission
func (mr *MockServiceBoardMockRecorder) CheckBoardPermission(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckBoardPermission", reflect.TypeOf((*MockServiceBoard)(nil).CheckBoardPermission), arg0, arg1, arg2)
}

// CheckCardPermission mocks base method
func (m *MockServiceBoard) CheckCardPermission(arg0, arg1 int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckCardPermission", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckCardPermission indicates an expected call of CheckCardPermission
func (mr *MockServiceBoardMockRecorder) CheckCardPermission(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckCardPermission", reflect.TypeOf((*MockServiceBoard)(nil).CheckCardPermission), arg0, arg1)
}

// CheckCommentPermission mocks base method
func (m *MockServiceBoard) CheckCommentPermission(arg0, arg1 int64, arg2 bool) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckCommentPermission", arg0, arg1, arg2)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckCommentPermission indicates an expected call of CheckCommentPermission
func (mr *MockServiceBoardMockRecorder) CheckCommentPermission(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckCommentPermission", reflect.TypeOf((*MockServiceBoard)(nil).CheckCommentPermission), arg0, arg1, arg2)
}

// CheckTaskPermission mocks base method
func (m *MockServiceBoard) CheckTaskPermission(arg0, arg1 int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckTaskPermission", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckTaskPermission indicates an expected call of CheckTaskPermission
func (mr *MockServiceBoardMockRecorder) CheckTaskPermission(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckTaskPermission", reflect.TypeOf((*MockServiceBoard)(nil).CheckTaskPermission), arg0, arg1)
}

// CreateAttachment mocks base method
func (m *MockServiceBoard) CreateAttachment(arg0 models.AttachmentInput) (models.AttachmentOutside, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAttachment", arg0)
	ret0, _ := ret[0].(models.AttachmentOutside)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAttachment indicates an expected call of CreateAttachment
func (mr *MockServiceBoardMockRecorder) CreateAttachment(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAttachment", reflect.TypeOf((*MockServiceBoard)(nil).CreateAttachment), arg0)
}

// CreateBoard mocks base method
func (m *MockServiceBoard) CreateBoard(arg0 models.BoardChangeInput) (models.BoardOutsideShort, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBoard", arg0)
	ret0, _ := ret[0].(models.BoardOutsideShort)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBoard indicates an expected call of CreateBoard
func (mr *MockServiceBoardMockRecorder) CreateBoard(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBoard", reflect.TypeOf((*MockServiceBoard)(nil).CreateBoard), arg0)
}

// CreateCard mocks base method
func (m *MockServiceBoard) CreateCard(arg0 models.CardInput) (models.CardOutsideShort, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCard", arg0)
	ret0, _ := ret[0].(models.CardOutsideShort)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCard indicates an expected call of CreateCard
func (mr *MockServiceBoardMockRecorder) CreateCard(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCard", reflect.TypeOf((*MockServiceBoard)(nil).CreateCard), arg0)
}

// CreateChecklist mocks base method
func (m *MockServiceBoard) CreateChecklist(arg0 models.ChecklistInput) (models.ChecklistOutside, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateChecklist", arg0)
	ret0, _ := ret[0].(models.ChecklistOutside)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateChecklist indicates an expected call of CreateChecklist
func (mr *MockServiceBoardMockRecorder) CreateChecklist(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateChecklist", reflect.TypeOf((*MockServiceBoard)(nil).CreateChecklist), arg0)
}

// CreateComment mocks base method
func (m *MockServiceBoard) CreateComment(arg0 models.CommentInput) (models.CommentOutside, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateComment", arg0)
	ret0, _ := ret[0].(models.CommentOutside)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateComment indicates an expected call of CreateComment
func (mr *MockServiceBoardMockRecorder) CreateComment(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateComment", reflect.TypeOf((*MockServiceBoard)(nil).CreateComment), arg0)
}

// CreateTag mocks base method
func (m *MockServiceBoard) CreateTag(arg0 models.TagInput) (models.TagOutside, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTag", arg0)
	ret0, _ := ret[0].(models.TagOutside)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTag indicates an expected call of CreateTag
func (mr *MockServiceBoardMockRecorder) CreateTag(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTag", reflect.TypeOf((*MockServiceBoard)(nil).CreateTag), arg0)
}

// CreateTask mocks base method
func (m *MockServiceBoard) CreateTask(arg0 models.TaskInput) (models.TaskOutsideSuperShort, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTask", arg0)
	ret0, _ := ret[0].(models.TaskOutsideSuperShort)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTask indicates an expected call of CreateTask
func (mr *MockServiceBoardMockRecorder) CreateTask(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTask", reflect.TypeOf((*MockServiceBoard)(nil).CreateTask), arg0)
}

// DeleteAttachment mocks base method
func (m *MockServiceBoard) DeleteAttachment(arg0 models.AttachmentInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAttachment", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAttachment indicates an expected call of DeleteAttachment
func (mr *MockServiceBoardMockRecorder) DeleteAttachment(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAttachment", reflect.TypeOf((*MockServiceBoard)(nil).DeleteAttachment), arg0)
}

// DeleteBoard mocks base method
func (m *MockServiceBoard) DeleteBoard(arg0 models.BoardInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBoard", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBoard indicates an expected call of DeleteBoard
func (mr *MockServiceBoardMockRecorder) DeleteBoard(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBoard", reflect.TypeOf((*MockServiceBoard)(nil).DeleteBoard), arg0)
}

// DeleteCard mocks base method
func (m *MockServiceBoard) DeleteCard(arg0 models.CardInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCard", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCard indicates an expected call of DeleteCard
func (mr *MockServiceBoardMockRecorder) DeleteCard(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCard", reflect.TypeOf((*MockServiceBoard)(nil).DeleteCard), arg0)
}

// DeleteChecklist mocks base method
func (m *MockServiceBoard) DeleteChecklist(arg0 models.ChecklistInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteChecklist", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteChecklist indicates an expected call of DeleteChecklist
func (mr *MockServiceBoardMockRecorder) DeleteChecklist(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteChecklist", reflect.TypeOf((*MockServiceBoard)(nil).DeleteChecklist), arg0)
}

// DeleteComment mocks base method
func (m *MockServiceBoard) DeleteComment(arg0 models.CommentInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteComment", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteComment indicates an expected call of DeleteComment
func (mr *MockServiceBoardMockRecorder) DeleteComment(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteComment", reflect.TypeOf((*MockServiceBoard)(nil).DeleteComment), arg0)
}

// DeleteTag mocks base method
func (m *MockServiceBoard) DeleteTag(arg0 models.TagInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTag", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTag indicates an expected call of DeleteTag
func (mr *MockServiceBoardMockRecorder) DeleteTag(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTag", reflect.TypeOf((*MockServiceBoard)(nil).DeleteTag), arg0)
}

// DeleteTask mocks base method
func (m *MockServiceBoard) DeleteTask(arg0 models.TaskInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTask", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTask indicates an expected call of DeleteTask
func (mr *MockServiceBoardMockRecorder) DeleteTask(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTask", reflect.TypeOf((*MockServiceBoard)(nil).DeleteTask), arg0)
}

// DismissUser mocks base method
func (m *MockServiceBoard) DismissUser(arg0 models.TaskAssignerInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DismissUser", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DismissUser indicates an expected call of DismissUser
func (mr *MockServiceBoardMockRecorder) DismissUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DismissUser", reflect.TypeOf((*MockServiceBoard)(nil).DismissUser), arg0)
}

// GetBoard mocks base method
func (m *MockServiceBoard) GetBoard(arg0 models.BoardInput) (models.BoardOutside, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBoard", arg0)
	ret0, _ := ret[0].(models.BoardOutside)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBoard indicates an expected call of GetBoard
func (mr *MockServiceBoardMockRecorder) GetBoard(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBoard", reflect.TypeOf((*MockServiceBoard)(nil).GetBoard), arg0)
}

// GetCard mocks base method
func (m *MockServiceBoard) GetCard(arg0 models.CardInput) (models.CardOutside, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCard", arg0)
	ret0, _ := ret[0].(models.CardOutside)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCard indicates an expected call of GetCard
func (mr *MockServiceBoardMockRecorder) GetCard(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCard", reflect.TypeOf((*MockServiceBoard)(nil).GetCard), arg0)
}

// GetSharedURL mocks base method
func (m *MockServiceBoard) GetSharedURL(arg0 models.BoardInput) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSharedURL", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSharedURL indicates an expected call of GetSharedURL
func (mr *MockServiceBoardMockRecorder) GetSharedURL(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSharedURL", reflect.TypeOf((*MockServiceBoard)(nil).GetSharedURL), arg0)
}

// GetTask mocks base method
func (m *MockServiceBoard) GetTask(arg0 models.TaskInput) (models.TaskOutside, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTask", arg0)
	ret0, _ := ret[0].(models.TaskOutside)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTask indicates an expected call of GetTask
func (mr *MockServiceBoardMockRecorder) GetTask(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTask", reflect.TypeOf((*MockServiceBoard)(nil).GetTask), arg0)
}

// InviteUserToBoard mocks base method
func (m *MockServiceBoard) InviteUserToBoard(arg0 models.BoardInviteInput) (models.BoardOutsideShort, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InviteUserToBoard", arg0)
	ret0, _ := ret[0].(models.BoardOutsideShort)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InviteUserToBoard indicates an expected call of InviteUserToBoard
func (mr *MockServiceBoardMockRecorder) InviteUserToBoard(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InviteUserToBoard", reflect.TypeOf((*MockServiceBoard)(nil).InviteUserToBoard), arg0)
}

// RemoveMember mocks base method
func (m *MockServiceBoard) RemoveMember(arg0 models.BoardMemberInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveMember", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveMember indicates an expected call of RemoveMember
func (mr *MockServiceBoardMockRecorder) RemoveMember(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveMember", reflect.TypeOf((*MockServiceBoard)(nil).RemoveMember), arg0)
}

// RemoveTag mocks base method
func (m *MockServiceBoard) RemoveTag(arg0 models.TaskTagInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveTag", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveTag indicates an expected call of RemoveTag
func (mr *MockServiceBoardMockRecorder) RemoveTag(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveTag", reflect.TypeOf((*MockServiceBoard)(nil).RemoveTag), arg0)
}

// TasksOrderChange mocks base method
func (m *MockServiceBoard) TasksOrderChange(arg0 models.TasksOrderInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TasksOrderChange", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// TasksOrderChange indicates an expected call of TasksOrderChange
func (mr *MockServiceBoardMockRecorder) TasksOrderChange(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TasksOrderChange", reflect.TypeOf((*MockServiceBoard)(nil).TasksOrderChange), arg0)
}

// WebSocketBoard mocks base method
func (m *MockServiceBoard) WebSocketBoard(arg0 models.BoardInput, arg1 echo.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WebSocketBoard", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WebSocketBoard indicates an expected call of WebSocketBoard
func (mr *MockServiceBoardMockRecorder) WebSocketBoard(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WebSocketBoard", reflect.TypeOf((*MockServiceBoard)(nil).WebSocketBoard), arg0, arg1)
}

// WebSocketNotification mocks base method
func (m *MockServiceBoard) WebSocketNotification(arg0 models.UserInput, arg1 echo.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WebSocketNotification", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WebSocketNotification indicates an expected call of WebSocketNotification
func (mr *MockServiceBoardMockRecorder) WebSocketNotification(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WebSocketNotification", reflect.TypeOf((*MockServiceBoard)(nil).WebSocketNotification), arg0, arg1)
}
