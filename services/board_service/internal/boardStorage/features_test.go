package boardStorage

import (
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/boardStorage/attachmentStorage"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/boardStorage/checklistStorage"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/boardStorage/commentStorage"
	mocks "github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/boardStorage/mock"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/boardStorage/tagStorage"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestStorage_CreateTag(t *testing.T) {
	input := models.TagInput{
		UserID:  1,
		TaskID:  1,
		TagID:   0,
		BoardID: 1,
		Color:   "yellow",
		Name:    "work",
	}

	expect := models.TagOutside{
		TagID: 1,
		Color: input.Color,
		Name:  input.Name,
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db, tagStorage: tagStorage.NewStorage(db)}

	mock.
		ExpectQuery("INSERT INTO tags").
		WithArgs(input.BoardID, input.Name, input.Color).
		WillReturnRows(sqlmock.NewRows([]string{"tagID"}).AddRow(1))

	tag, err := storage.CreateTag(input)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(tag, expect) {
		t.Errorf("results not match, want %v, have %v", expect, tag)
		return
	}
}

func TestStorage_UpdateTag(t *testing.T) {
	input := models.TagInput{
		UserID:  1,
		TaskID:  1,
		TagID:   1,
		BoardID: 1,
		Color:   "yellow",
		Name:    "work",
	}

	expect := models.TagOutside{
		TagID: 1,
		Color: input.Color,
		Name:  input.Name,
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db, tagStorage: tagStorage.NewStorage(db)}

	mock.
		ExpectExec("UPDATE tags SET").
		WithArgs(input.Name, input.Color, input.TagID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	tag, err := storage.UpdateTag(input)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(tag, expect) {
		t.Errorf("results not match, want %v, have %v", expect, tag)
		return
	}
}

func TestStorage_DeleteTag(t *testing.T) {
	input := models.TagInput{
		TagID:   1,
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db, tagStorage: tagStorage.NewStorage(db)}

	mock.
		ExpectExec("DELETE FROM tags").
		WithArgs(input.TagID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = storage.DeleteTag(input)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
}

func TestStorage_AddTag(t *testing.T) {
	input := models.TaskTagInput{
		UserID:  1,
		TaskID:  1,
		TagID:   1,
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ctrlTasks := gomock.NewController(t)
	defer ctrlTasks.Finish()
	mockTasks := mocks.NewMockTasksStorage(ctrlTasks)
	storage := &storage{db: db, tagStorage: tagStorage.NewStorage(db), tasksStorage: mockTasks}

	mock.
		ExpectExec("INSERT INTO task_tags").
		WithArgs(input.TaskID, input.TagID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.
		ExpectQuery("SELECT").
		WithArgs(input.TagID).
		WillReturnRows(sqlmock.NewRows([]string{"tagID", "name", "color"}).AddRow(input.TagID, "name", "color"))

	mockTasks.EXPECT().GetCardIDByTask(input.TaskID).Return(int64(1), nil)

	_, err = storage.AddTag(input)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
}

func TestStorage_RemoveTag(t *testing.T) {
	input := models.TaskTagInput{
		TagID:   1,
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ctrlTasks := gomock.NewController(t)
	defer ctrlTasks.Finish()
	mockTasks := mocks.NewMockTasksStorage(ctrlTasks)

	storage := &storage{db: db, tagStorage: tagStorage.NewStorage(db), tasksStorage: mockTasks}

	mock.
		ExpectExec("DELETE FROM task_tags").
		WithArgs(input.TaskID, input.TagID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mockTasks.EXPECT().GetCardIDByTask(input.TaskID).Return(int64(1), nil)

	_, err = storage.RemoveTag(input)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
}

func TestStorage_CreateComment(t *testing.T) {
	input := models.CommentInput{
		UserID:  1,
		TaskID:  1,
		Order: 1,
		Message: "msg",
	}

	expect := models.CommentOutside{
		CommentID: 1,
		Message: input.Message,
		Order: input.Order,
		TaskID: input.TaskID,
		CardID: 1,
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ctrlTasks := gomock.NewController(t)
	defer ctrlTasks.Finish()
	mockTasks := mocks.NewMockTasksStorage(ctrlTasks)

	storage := &storage{db: db, commentStorage: commentStorage.NewStorage(db), tasksStorage: mockTasks}

	mock.
		ExpectQuery("INSERT INTO comments").
		WithArgs(input.Message, input.TaskID, input.Order, input.UserID).
		WillReturnRows(sqlmock.NewRows([]string{"commentID", "taskID"}).AddRow(1, input.TaskID))

	mockTasks.EXPECT().GetCardIDByTask(input.TaskID).Return(expect.CardID, nil)

	comment, err := storage.CreateComment(input)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(comment, expect) {
		t.Errorf("results not match, want %v, have %v", expect, comment)
		return
	}
}

func TestStorage_UpdateComment(t *testing.T) {
	input := models.CommentInput{
		CommentID: 1,
		Message: "msg",
	}

	expect := models.CommentOutside{
		CommentID: 1,
		Message: input.Message,
		CardID: 1,
		TaskID: 1,
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ctrlTasks := gomock.NewController(t)
	defer ctrlTasks.Finish()
	mockTasks := mocks.NewMockTasksStorage(ctrlTasks)

	storage := &storage{db: db, commentStorage: commentStorage.NewStorage(db),  tasksStorage: mockTasks}

	mock.
		ExpectQuery("UPDATE comments SET").
		WithArgs(input.Message, input.CommentID).
		WillReturnRows(sqlmock.NewRows([]string{"taskID"}).AddRow(expect.TaskID))

	mockTasks.EXPECT().GetCardIDByTask(expect.TaskID).Return(expect.CardID, nil)

	comment, err := storage.UpdateComment(input)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(comment, expect) {
		t.Errorf("results not match, want %v, have %v", expect, comment)
		return
	}
}

func TestStorage_DeleteComment(t *testing.T) {
	input := models.CommentInput{
		CommentID: 1,
	}

	expect := models.CommentOutside {
		TaskID: 1,
		CommentID: input.CommentID,
		CardID: 1,
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ctrlTasks := gomock.NewController(t)
	defer ctrlTasks.Finish()
	mockTasks := mocks.NewMockTasksStorage(ctrlTasks)

	storage := &storage{db: db, commentStorage: commentStorage.NewStorage(db),  tasksStorage: mockTasks}

	mock.
		ExpectQuery("DELETE FROM comments").
		WithArgs(input.CommentID).
		WillReturnRows(sqlmock.NewRows([]string{"taskID"}).AddRow(expect.TaskID))

	mockTasks.EXPECT().GetCardIDByTask(expect.TaskID).Return(expect.CardID, nil)

	_, err = storage.DeleteComment(input)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestStorage_CreateChecklist(t *testing.T) {
	input := models.ChecklistInput{
		UserID:      0,
		ChecklistID: 0,
		TaskID:      1,
		Name:        "check",
		Items:       json.RawMessage{},
	}

	expect := models.ChecklistOutside{
		ChecklistID: 1,
		Items:       input.Items,
		Name:        input.Name,
		TaskID:      input.TaskID,
		CardID:      1,
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ctrlTasks := gomock.NewController(t)
	defer ctrlTasks.Finish()
	mockTasks := mocks.NewMockTasksStorage(ctrlTasks)

	storage := &storage{db: db, checklistStorage: checklistStorage.NewStorage(db),  tasksStorage: mockTasks}

	mock.
		ExpectQuery("INSERT INTO checklists").
		WithArgs(input.TaskID, input.Name, input.Items).
		WillReturnRows(sqlmock.NewRows([]string{"checklistID"}).AddRow(expect.ChecklistID))

	mockTasks.EXPECT().GetCardIDByTask(expect.TaskID).Return(expect.CardID, nil)

	checklist, err := storage.CreateChecklist(input)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(checklist, expect) {
		t.Errorf("results not match, want %v, have %v", expect, checklist)
		return
	}
}

func TestStorage_UpdateChecklist(t *testing.T) {
	input := models.ChecklistInput{
		ChecklistID: 1,
		Name:        "check",
		Items:       nil,
	}

	expect := models.ChecklistOutside{
		ChecklistID: 1,
		Items:       input.Items,
		Name:        input.Name,
		TaskID:      input.TaskID,
		CardID:      1,
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ctrlTasks := gomock.NewController(t)
	defer ctrlTasks.Finish()
	mockTasks := mocks.NewMockTasksStorage(ctrlTasks)

	storage := &storage{db: db, checklistStorage: checklistStorage.NewStorage(db),  tasksStorage: mockTasks}

	mock.
		ExpectQuery("UPDATE checklists SET").
		WithArgs(input.Name, input.Items, input.ChecklistID).
		WillReturnRows(sqlmock.NewRows([]string{"taskID"}).AddRow(expect.TaskID))

	mockTasks.EXPECT().GetCardIDByTask(expect.TaskID).Return(expect.CardID, nil)

	checklist, err := storage.UpdateChecklist(input)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(checklist, expect) {
		t.Errorf("results not match, want %v, have %v", expect, checklist)
		return
	}
}

func TestStorage_DeleteChecklist(t *testing.T) {
	input := models.ChecklistInput{
		ChecklistID: 1,
	}

	expect := models.ChecklistOutside{
		ChecklistID: 1,
		TaskID:      input.TaskID,
		CardID: 1,
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ctrlTasks := gomock.NewController(t)
	defer ctrlTasks.Finish()
	mockTasks := mocks.NewMockTasksStorage(ctrlTasks)

	storage := &storage{db: db, checklistStorage: checklistStorage.NewStorage(db),  tasksStorage: mockTasks}

	mock.
		ExpectQuery("DELETE FROM checklists").
		WithArgs(input.ChecklistID).
		WillReturnRows(sqlmock.NewRows([]string{"taskID"}).AddRow(expect.TaskID))

	mockTasks.EXPECT().GetCardIDByTask(expect.TaskID).Return(expect.CardID, nil)

	_, err = storage.DeleteChecklist(input)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestStorage_AddAttachment(t *testing.T) {
	input := models.AttachmentInternal{
		TaskID:       1,
		AttachmentID: 0,
		Filename:     "file",
		Filepath:     "apkapaskampaomasppn",
	}

	expect := models.AttachmentOutside{
		AttachmentID: 1,
		Filename:     input.Filename,
		Filepath:     input.Filepath,
		TaskID: input.TaskID,
		CardID: 1,
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ctrlTasks := gomock.NewController(t)
	defer ctrlTasks.Finish()
	mockTasks := mocks.NewMockTasksStorage(ctrlTasks)

	storage := &storage{db: db, attachmentStorage: attachmentStorage.NewStorage(db), tasksStorage: mockTasks}

	mock.
		ExpectQuery("INSERT INTO attachments").
		WithArgs(input.TaskID, input.Filename, input.Filepath).
		WillReturnRows(sqlmock.NewRows([]string{"attachmentID", "taskID"}).AddRow(expect.AttachmentID, expect.TaskID))

	mockTasks.EXPECT().GetCardIDByTask(expect.TaskID).Return(expect.CardID, nil)

	attachment, err := storage.AddAttachment(input)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(attachment, expect) {
		t.Errorf("results not match, want %v, have %v", expect, attachment)
		return
	}
}

func TestStorage_RemoveAttachment(t *testing.T) {
	input := models.AttachmentInternal{
		AttachmentID: 1,
	}

	expect := models.AttachmentOutside{
		AttachmentID: 1,
		TaskID: input.TaskID,
		CardID: 1,
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ctrlTasks := gomock.NewController(t)
	defer ctrlTasks.Finish()
	mockTasks := mocks.NewMockTasksStorage(ctrlTasks)

	storage := &storage{db: db, attachmentStorage: attachmentStorage.NewStorage(db), tasksStorage: mockTasks}

	mock.
		ExpectQuery("DELETE FROM attachments").
		WithArgs(input.AttachmentID).
		WillReturnRows(sqlmock.NewRows([]string{"attachmentID", "taskID"}).AddRow(expect.AttachmentID, expect.TaskID))

	mockTasks.EXPECT().GetCardIDByTask(expect.TaskID).Return(expect.CardID, nil)

	_, err = storage.RemoveAttachment(input)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}