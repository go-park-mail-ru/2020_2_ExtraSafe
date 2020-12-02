package commentStorage

import (
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"reflect"
	"testing"
)

func TestStorage_CreateComment(t *testing.T) {
	input := models.CommentInput{
		UserID:  1,
		TaskID:  1,
		Order: 1,
		Message: "msg",
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	mock.
		ExpectQuery("INSERT INTO comments").
		WithArgs(input.Message, input.TaskID, input.Order, input.UserID).
		WillReturnError(sql.ErrNoRows)

	_, err = storage.CreateComment(input)
	if err == nil {
		t.Error("expected error")
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestStorage_UpdateComment(t *testing.T) {
	input := models.CommentInput{
		CommentID: 1,
		Message: "msg",
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	mock.
		ExpectExec("UPDATE comments SET").
		WithArgs(input.Message, input.CommentID).
		WillReturnError(errors.New(""))

	_, err = storage.UpdateComment(input)
	if err == nil {
		t.Error("expected error")
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestStorage_DeleteComment(t *testing.T) {
	input := models.CommentInput{
		CommentID: 1,
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	mock.
		ExpectExec("DELETE FROM comments").
		WithArgs(input.CommentID).
		WillReturnError(errors.New(""))

	err = storage.DeleteComment(input)
	if err == nil {
		t.Error("expected error")
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestStorage_GetCommentsByTask(t *testing.T) {
	input := models.TaskInput{TaskID: 1}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	comment := models.CommentOutside{
		CommentID: 1,
		Message:   "lala",
		Order:     1,
	}

	expect := make([]models.CommentOutside, 0)
	expect = append(expect, comment)

	mock.ExpectQuery("SELECT commentID, message, commentOrder, userID").
		WithArgs(input.TaskID).
		WillReturnRows(sqlmock.NewRows([]string{"commentID", "message", "commentOrder", "userID"}).
								AddRow(comment.CommentID, comment.Message, comment.Order, 1))

	comments, _, err := storage.GetCommentsByTask(input)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(comments, expect) {
		t.Errorf("results not match, want %v, have %v", expect, comments)
		return
	}
}

func TestStorage_GetCommentsByTaskScanFail(t *testing.T) {
	input := models.TaskInput{TaskID: 1}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	comment := models.CommentOutside{
		CommentID: 1,
		Message:   "lala",
		Order:     1,
	}

	expect := make([]models.CommentOutside, 0)
	expect = append(expect, comment)

	mock.ExpectQuery("SELECT commentID, message, commentOrder, userID").
		WithArgs(input.TaskID).
		WillReturnRows(sqlmock.NewRows([]string{"commentID", "message", "commentOrder"}).
			AddRow(comment.CommentID, comment.Message, comment.Order))

	_, _, err = storage.GetCommentsByTask(input)
	if err == nil {
		t.Error("expected error")
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}