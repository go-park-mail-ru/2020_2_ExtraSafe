package tasksStorage

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"reflect"
	"testing"
)

//TODO TESTS сделать тесты на ошибки
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

/*	err := s.db.QueryRow("INSERT INTO tasks (cardID, taskName, description, tasksOrder) VALUES ($1, $2, $3, $4) RETURNING taskID",
		taskInput.CardID, taskInput.Name, taskInput.Description, taskInput.Order).Scan(&taskID)
*/
	//ok query
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

	/*_, err := s.db.Exec("UPDATE tasks SET taskName = $1, description = $2, tasksOrder = $3 WHERE taskID = $4",
		taskInput.Name, taskInput.Description, taskInput.Order, taskInput.TaskID)*/

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

func TestStorage_GetTaskByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	taskInput := models.TaskInput{ TaskID:  1}

/*	err := s.db.QueryRow("SELECT taskName, description, tasksOrder FROM tasks WHERE taskID = $1", taskInput.TaskID).
		Scan(&task.Name, &task.Description, &task.Order)*/

	expectedTaskOutside := models.TaskOutside {
		TaskID: 1,
		Name:   "todo",
		Description: "just task",
		Order:  1,
	}

	rows := sqlmock.NewRows([]string{"taskName", "description", "tasksOrder"}).
					AddRow(expectedTaskOutside.Name, expectedTaskOutside.Description, expectedTaskOutside.Order)
	//ok query
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

func TestStorage_GetTasksByCard(t *testing.T) {
	//rows, err := s.db.Query("SELECT taskID, taskName, description, tasksOrder FROM tasks WHERE cardID = $1", cardInput.CardID)

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