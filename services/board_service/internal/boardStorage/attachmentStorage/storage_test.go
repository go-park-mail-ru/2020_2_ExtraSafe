package attachmentStorage

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"reflect"
	"testing"
)

func TestStorage_AddAttachmentFail(t *testing.T) {
	input := models.AttachmentInternal{
		TaskID:       1,
		AttachmentID: 0,
		Filename:     "file",
		Filepath:     "apkapaskampaomasppn",
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	mock.
		ExpectQuery("INSERT INTO attachments").
		WithArgs(input.TaskID, input.Filename, input.Filepath).
		WillReturnError(errors.New(""))


	_, err = storage.AddAttachment(input)
	if err == nil {
		t.Error("expected error")
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestStorage_RemoveAttachment(t *testing.T) {
	input := models.AttachmentInternal{
		AttachmentID: 1,
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	mock.
		ExpectExec("DELETE FROM attachments").
		WithArgs(input.AttachmentID).
		WillReturnError(errors.New(""))

	_, err = storage.RemoveAttachment(input)
	if err == nil {
		t.Error("expected error")
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestStorage_GetAttachmentsByTask(t *testing.T) {
	input := models.TaskInput{
		TaskID: 1,
	}

	expect := make([]models.AttachmentOutside, 0)
	attach := models.AttachmentOutside{
		AttachmentID: 1,
		Filename:     "filename",
		Filepath:     "filepath",
	}

	expect = append(expect, attach)

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	mock.
		ExpectQuery("SELECT attachmentID, filename, filepath FROM attachments").
		WithArgs(input.TaskID).
		WillReturnRows(sqlmock.NewRows([]string{"attachmentID", "filename", "filepath"}).AddRow(attach.AttachmentID, attach.Filename, attach.Filepath))

	attachs, err := storage.GetAttachmentsByTask(input)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(attachs, expect) {
		t.Errorf("results not match, want %v, have %v", expect, attachs)
		return
	}
}

func TestStorage_GetAttachmentsByTaskFail(t *testing.T) {
	input := models.TaskInput{
		TaskID: 1,
	}

	expect := make([]models.AttachmentOutside, 0)
	attach := models.AttachmentOutside{
		AttachmentID: 1,
		Filename:     "filename",
		Filepath:     "filepath",
	}

	expect = append(expect, attach)

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	mock.
		ExpectQuery("SELECT attachmentID, filename, filepath FROM attachments").
		WithArgs(input.TaskID).
		WillReturnRows(sqlmock.NewRows([]string{"attachmentID", "filename", "filepath", "smth"}).AddRow(attach.AttachmentID, attach.Filename, attach.Filepath, "smth"))

	_, err = storage.GetAttachmentsByTask(input)
	if err == nil {
		t.Error("expected error")
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

}

