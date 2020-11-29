package boardStorage

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/boardStorage/commentStorage"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/boardStorage/tagStorage"
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

	storage := &storage{db: db, tagStorage: tagStorage.NewStorage(db)}

	mock.
		ExpectExec("INSERT INTO task_tags").
		WithArgs(input.TaskID, input.TagID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = storage.AddTag(input)
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

	storage := &storage{db: db, tagStorage: tagStorage.NewStorage(db)}

	mock.
		ExpectExec("DELETE FROM task_tags").
		WithArgs(input.TaskID, input.TagID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = storage.RemoveTag(input)
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
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db, commentStorage: commentStorage.NewStorage(db)}

	mock.
		ExpectQuery("INSERT INTO comments").
		WithArgs(input.Message, input.TaskID, input.Order, input.UserID).
		WillReturnRows(sqlmock.NewRows([]string{"commentID"}).AddRow(1))

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
		UserID:  1,
		TaskID:  1,
		Order: 1,
		Message: "msg",
	}

	expect := models.CommentOutside{
		CommentID: 1,
		Message: input.Message,
		Order: input.Order,
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db, commentStorage: commentStorage.NewStorage(db)}

	mock.
		ExpectQuery("UPDATE comments SET").
		WithArgs(input.Message, input.TaskID, input.Order, input.UserID).
		WillReturnRows(sqlmock.NewRows([]string{"commentID"}).AddRow(1))

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