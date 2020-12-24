package checkliststorage

import (
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"reflect"
	"testing"
)

func TestStorage_CreateChecklistFail(t *testing.T) {
	input := models.ChecklistInput{
		UserID:      0,
		ChecklistID: 0,
		TaskID:      1,
		Name:        "check",
		Items:       nil,
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	mock.
		ExpectQuery("INSERT INTO checklists").
		WithArgs(input.TaskID, input.Name, input.Items).
		WillReturnError(errors.New(""))

	_, err = storage.CreateChecklist(input)
	if err == nil {
		t.Error("expected error")
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestStorage_UpdateChecklistFail(t *testing.T) {
	input := models.ChecklistInput{
		ChecklistID: 1,
		Name:        "check",
		Items:       nil,
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	mock.
		ExpectQuery("UPDATE checklists SET").
		WithArgs(input.Name, input.Items, input.ChecklistID).
		WillReturnError(errors.New(""))

	_, err = storage.UpdateChecklist(input)
	if err == nil {
		t.Error("expected error")
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestStorage_DeleteChecklistFail(t *testing.T) {
	input := models.ChecklistInput{
		ChecklistID: 1,
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	mock.
		ExpectQuery("DELETE FROM checklists").
		WithArgs(input.ChecklistID).
		WillReturnError(errors.New(""))

	_, err = storage.DeleteChecklist(input)
	if err == nil {
		t.Error("expected error")
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestStorage_GetChecklistsByTask(t *testing.T) {
	input := models.TaskInput{TaskID: 1}

	expect := make([]models.ChecklistOutside, 0)
	check := models.ChecklistOutside{
		ChecklistID: 1,
		Items:       nil,
		Name:        "name",
	}
	expect = append(expect, check)

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	mock.
		ExpectQuery("SELECT checklistID, name, items FROM checklists").
		WithArgs(input.TaskID).
		WillReturnRows(sqlmock.NewRows([]string{"checklistID", "name", "items"}).AddRow(check.ChecklistID, check.Name, check.Items))

	checks, err := storage.GetChecklistsByTask(input)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(checks, expect) {
		t.Errorf("results not match, want %v, have %v", expect, checks)
		return
	}
}

func TestStorage_GetChecklistsByTaskQueryFail(t *testing.T) {
	input := models.TaskInput{TaskID: 1}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	mock.
		ExpectQuery("SELECT checklistID, name, items FROM checklists").
		WithArgs(input.TaskID).
		WillReturnError(sql.ErrNoRows)

	_, err = storage.GetChecklistsByTask(input)
	if err == nil {
		t.Error("expected error")
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestStorage_GetChecklistsByTaskScanFail(t *testing.T) {
	input := models.TaskInput{TaskID: 1}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	mock.
		ExpectQuery("SELECT checklistID, name, items FROM checklists").
		WithArgs(input.TaskID).
		WillReturnRows(sqlmock.NewRows([]string{"checklistID", "name", "items", "smth"}).AddRow(1, "name", nil, "smth"))

	_, err = storage.GetChecklistsByTask(input)
	if err == nil {
		t.Error("expected error")
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}
