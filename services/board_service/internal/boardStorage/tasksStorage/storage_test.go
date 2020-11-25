package tasksStorage

import (
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"reflect"
	"testing"
)

func TestStorage_CreateTask(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	taskInput := models.TaskInput{
		TaskID:  0,
		CardID:  1,
		Name:    "task1",
		Description: "just task",
		Order:   1,
	}

	expectTaskOutside := models.TaskOutside{
		TaskID: 1,
		Name:   taskInput.Name,
		Order:  taskInput.Order,
		Description: taskInput.Description,
	}

	mock.
		ExpectQuery("INSERT INTO tasks").
		WithArgs(taskInput.CardID, taskInput.Name, taskInput.Description, taskInput.Order).
		WillReturnRows(sqlmock.NewRows([]string{"taskID"}).AddRow(1))

	task, err := storage.CreateTask(taskInput)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(task, expectTaskOutside) {
		t.Errorf("results not match, want %v, have %v", expectTaskOutside, task)
		return
	}
}


func TestStorage_CreateTaskFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := NewStorage(db)

	taskInput := models.TaskInput{
		TaskID:  0,
		CardID:  1,
		Name:    "task1",
		Description: "just task",
		Order:   1,
	}

	mock.
		ExpectQuery("INSERT INTO tasks").
		WithArgs(taskInput.CardID, taskInput.Name, taskInput.Description, taskInput.Order).
		WillReturnError(sql.ErrNoRows)

	_, err = storage.CreateTask(taskInput)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Error("expected err")
		return
	}
}

func TestStorage_ChangeTask(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	taskInput := models.TaskInput{
		TaskID:  1,
		Name:    "taskNew",
		Description: "just updated task",
		Order:   2,
	}

	expectTaskOutside := models.TaskOutside{
		TaskID: 1,
		Name:   taskInput.Name,
		Order:  taskInput.Order,
		Description: taskInput.Description,
	}

	mock.
		ExpectExec("UPDATE tasks SET").
		WithArgs(taskInput.Name, taskInput.Description, taskInput.Order, taskInput.TaskID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	task, err := storage.ChangeTask(taskInput)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(task, expectTaskOutside) {
		t.Errorf("results not match, want %v, have %v", expectTaskOutside, task)
		return
	}
}

func TestStorage_ChangeTaskFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	taskInput := models.TaskInput{
		TaskID:  1,
		Name:    "taskNew",
		Description: "just updated task",
		Order:   2,
	}

	mock.
		ExpectExec("UPDATE tasks SET").
		WithArgs(taskInput.Name, taskInput.Description, taskInput.Order, taskInput.TaskID).
		WillReturnError(errors.New("fail update exec"))

	_, err = storage.ChangeTask(taskInput)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Error("expected err")
		return
	}
}

func TestStorage_DeleteTask(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	taskInput := models.TaskInput{ TaskID:  1 }

	mock.
		ExpectExec("DELETE FROM tasks WHERE taskID").
		WithArgs(taskInput.TaskID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = storage.DeleteTask(taskInput)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestStorage_DeleteTaskFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	taskInput := models.TaskInput{ TaskID:  1 }

	mock.
		ExpectExec("DELETE FROM tasks WHERE taskID").
		WithArgs(taskInput.TaskID).
		WillReturnError(errors.New("fail deleting exec"))

	err = storage.DeleteTask(taskInput)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Error("expected err")
		return
	}
}

func TestStorage_GetTaskByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	taskInput := models.TaskInput{ TaskID:  1}

	expectedTaskOutside := models.TaskOutside {
		TaskID: 1,
		Name:   "todo",
		Description: "just task",
		Order:  1,
	}

	rows := sqlmock.NewRows([]string{"taskName", "description", "tasksOrder"}).
					AddRow(expectedTaskOutside.Name, expectedTaskOutside.Description, expectedTaskOutside.Order)

	mock.
		ExpectQuery("SELECT taskName, description, tasksOrder FROM tasks WHERE taskID").
		WithArgs(taskInput.TaskID).
		WillReturnRows(rows)

	task, err := storage.GetTaskByID(taskInput)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(task, expectedTaskOutside) {
		t.Errorf("results not match, want %v, have %v", expectedTaskOutside, task)
		return
	}
}

func TestStorage_GetTaskByIDFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	taskInput := models.TaskInput{ TaskID:  1}

	mock.
		ExpectQuery("SELECT taskName, description, tasksOrder FROM tasks WHERE taskID").
		WithArgs(taskInput.TaskID).
		WillReturnError(sql.ErrNoRows)

	_, err = storage.GetTaskByID(taskInput)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Error("expected err")
		return
	}
}

func TestStorage_GetTasksByCard(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	cardInput := models.CardInput{ CardID:  1}

	expectedTasks := make([]models.TaskOutside, 0)
	task1 := models.TaskOutside{
		TaskID: 1,
		Name:   "task 1",
		Description: "first task ever",
		Order:  1,
	}

	task2 := models.TaskOutside{
		TaskID: 2,
		Name:   "task 2",
		Description: "second task",
		Order:  2,
	}

	expectedTasks = append(expectedTasks, task1, task2)

	rows := sqlmock.NewRows([]string{"taskID", "taskName", "description", "tasksOrder"})
	rows.AddRow(task1.TaskID, task1.Name, task1.Description, task1.Order).
		 AddRow(task2.TaskID, task2.Name, task2.Description, task2.Order)

	mock.
		ExpectQuery("SELECT taskID, taskName, description, tasksOrder FROM tasks WHERE cardID").
		WithArgs(cardInput.CardID).
		WillReturnRows(rows)

	tasks, err := storage.GetTasksByCard(cardInput)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(tasks, expectedTasks) {
		t.Errorf("results not match, want %v, have %v", expectedTasks, tasks)
		return
	}
}

func TestStorage_GetTasksByCardFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	cardInput := models.CardInput{ CardID:  1}

	mock.
		ExpectQuery("SELECT taskID, taskName, description, tasksOrder FROM tasks WHERE cardID").
		WithArgs(cardInput.CardID).
		WillReturnError(sql.ErrNoRows)

	_, err = storage.GetTasksByCard(cardInput)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Error("expected err")
		return
	}
}

func TestStorage_GetTasksByCardFailScan(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	cardInput := models.CardInput{ CardID:  1}

	expectedTasks := make([]models.TaskOutside, 0)
	task1 := models.TaskOutside{
		TaskID: 1,
		Name:   "task 1",
		Description: "first task ever",
		Order:  1,
	}

	expectedTasks = append(expectedTasks, task1)

	rows := sqlmock.NewRows([]string{"taskID", "taskName", "description"})
	rows.AddRow(task1.TaskID, task1.Name, task1.Description)

	mock.
		ExpectQuery("SELECT taskID, taskName, description, tasksOrder FROM tasks WHERE cardID").
		WithArgs(cardInput.CardID).
		WillReturnRows(rows)

	_, err = storage.GetTasksByCard(cardInput)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Error("expected err")
		return
	}
}

func TestStorage_ChangeTaskOrder(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	tasks := models.TaskOrder{
		TaskID: 1,
		Order:  2,
	}
	card := models.TasksOrder{CardID: 1}
	card.Tasks = append(card.Tasks, tasks)
	input := models.TasksOrderInput{ UserID: 1 }
	input.Tasks = append(input.Tasks, card)

	mock.
		ExpectExec("UPDATE tasks SET").
		WithArgs(input.Tasks[0].CardID, input.Tasks[0].Tasks[0].Order, input.Tasks[0].Tasks[0].TaskID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = storage.ChangeTaskOrder(input)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
}

func TestStorage_ChangeTaskOrderFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	tasks := models.TaskOrder{
		TaskID: 1,
		Order:  2,
	}
	card := models.TasksOrder{CardID: 1}
	card.Tasks = append(card.Tasks, tasks)
	input := models.TasksOrderInput{ UserID: 1 }
	input.Tasks = append(input.Tasks, card)

	mock.
		ExpectExec("UPDATE tasks SET").
		WithArgs(input.Tasks[0].CardID, input.Tasks[0].Tasks[0].Order, input.Tasks[0].Tasks[0].TaskID).
		WillReturnError(errors.New("update exec error"))

	err = storage.ChangeTaskOrder(input)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Error("expected err")
		return
	}
}
