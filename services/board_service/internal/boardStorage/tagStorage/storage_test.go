package tagStorage

import (
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"reflect"
	"testing"
)

func TestStorage_CreateTagFail(t *testing.T) {
	input := models.TagInput{
		UserID:  1,
		TaskID:  1,
		TagID:   0,
		BoardID: 1,
		Color:   "yellow",
		Name:    "work",
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	mock.
		ExpectQuery("INSERT INTO tags").
		WithArgs(input.BoardID, input.Name, input.Color).
		WillReturnError(sql.ErrNoRows)

	_, err = storage.CreateTag(input)
	if err == nil {
		t.Error("expected error")
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestStorage_UpdateTagFail(t *testing.T) {
	input := models.TagInput{
		UserID:  1,
		TaskID:  1,
		TagID:   1,
		BoardID: 1,
		Color:   "yellow",
		Name:    "work",
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	mock.
		ExpectExec("UPDATE tags SET").
		WithArgs(input.Name, input.Color, input.TagID).
		WillReturnError(errors.New(""))

	_, err = storage.UpdateTag(input)
	if err == nil {
		t.Error("expected error")
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
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

	storage := &storage{db: db}

	mock.
		ExpectExec("DELETE FROM tags").
		WithArgs(input.TagID).
		WillReturnError(errors.New(""))

	err = storage.DeleteTag(input)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
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

	storage := &storage{db: db}

	mock.
		ExpectExec("INSERT INTO task_tags").
		WithArgs(input.TaskID, input.TagID).
		WillReturnError(errors.New(""))

	err = storage.AddTag(input)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
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

	storage := NewStorage(db)

	mock.
		ExpectExec("DELETE FROM task_tags").
		WithArgs(input.TaskID, input.TagID).
		WillReturnError(errors.New(""))

	err = storage.RemoveTag(input)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
}

func TestStorage_GetBoardTagsQueryFail(t *testing.T) {
	input := models.BoardInput{
		UserID:  0,
		BoardID: 1,
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := NewStorage(db)

	mock.
		ExpectQuery("SELECT tagID, name, color FROM tags").
		WithArgs(input.BoardID).
		WillReturnError(errors.New(""))

	_, err = storage.GetBoardTags(input)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
}

func TestStorage_GetBoardTagsScanFail(t *testing.T) {
	input := models.BoardInput{
		UserID:  0,
		BoardID: 1,
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := NewStorage(db)

	mock.
		ExpectQuery("SELECT tagID, name, color FROM tags").
		WithArgs(input.BoardID).
		WillReturnRows(sqlmock.NewRows([]string{"tagID", "name", "color", "taskID"}).AddRow(1, "name", "color", 1))

	_, err = storage.GetBoardTags(input)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
}

func TestStorage_GetTaskTagsQueryFail(t *testing.T) {
	input := models.TaskInput{
		UserID:  0,
		TaskID: 1,
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := NewStorage(db)

	mock.
		ExpectQuery("SELECT DISTINCT T.tagID, T.name, T.color FROM task_tags").
		WithArgs(input.TaskID).
		WillReturnError(errors.New(""))

	_, err = storage.GetTaskTags(input)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
}

func TestStorage_GetTaskTagsScanFail(t *testing.T) {
	input := models.TaskInput{
		UserID:  0,
		TaskID: 1,
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := NewStorage(db)

	mock.
		ExpectQuery("SELECT DISTINCT  T.tagID, T.name, T.color FROM task_tags").
		WithArgs(input.TaskID).
		WillReturnRows(sqlmock.NewRows([]string{"tagID", "name", "color", "taskID"}).AddRow(1, "name", "color", 1))

	_, err = storage.GetTaskTags(input)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
}
func TestStorage_GetBoardTags(t *testing.T) {
	input := models.BoardInput{
		UserID:  0,
		BoardID: 1,
	}

	expect := make([]models.TagOutside, 0)
	tag := models.TagOutside{
		TagID: 1,
		Color: "color",
		Name:  "name",
	}
	expect = append(expect, tag)

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := NewStorage(db)

	mock.
		ExpectQuery("SELECT tagID, name, color FROM tags").
		WithArgs(input.BoardID).
		WillReturnRows(sqlmock.NewRows([]string{"tagID", "name", "color"}).AddRow(tag.TagID, tag.Name, tag.Color))

	tags, err := storage.GetBoardTags(input)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(tags, expect) {
		t.Errorf("results not match, want %v, have %v", expect, tags)
		return
	}
}

func TestStorage_GetTaskTags(t *testing.T) {
	input := models.TaskInput{
		UserID:  0,
		TaskID: 1,
	}
	expect := make([]models.TagOutside, 0)
	tag := models.TagOutside{
		TagID: 1,
		Color: "color",
		Name:  "name",
	}
	expect = append(expect, tag)

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := NewStorage(db)

	mock.
		ExpectQuery("SELECT DISTINCT  T.tagID, T.name, T.color FROM task_tags").
		WithArgs(input.TaskID).
		WillReturnRows(sqlmock.NewRows([]string{"tagID", "name", "color"}).AddRow(tag.TagID, tag.Name, tag.Color))

	tags, err := storage.GetTaskTags(input)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(tags, expect) {
		t.Errorf("results not match, want %v, have %v", expect, tags)
		return
	}
}

