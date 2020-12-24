package taskstorage

import (
	"database/sql"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
)

type Storage interface {
	CreateTask(taskInput models.TaskInput) (models.TaskInternalShort, error)
	ChangeTask(taskInput models.TaskInput) (models.TaskInternal, error)
	DeleteTask(taskInput models.TaskInput) (models.TaskInternalShort, error)

	GetTasksByCard(cardInput models.CardInput) ([]models.TaskInternalShort, error)
	GetTaskName(taskInput models.TaskInput) (string, error)
	GetTaskByID(taskInput models.TaskInput) (models.TaskInternal, error)
	GetTaskNameByID(taskID int64) (string, error)
	GetCardIDByTask(taskInput int64) (cardID int64, err error)
	ChangeTaskOrder(taskInput models.TasksOrderInput) error

	AssignUser(input models.TaskAssigner) (err error)
	DismissUser(input models.TaskAssigner) (err error)
	GetAssigners(input models.TaskInput) (assignerIDs []int64, err error)
}

type storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) Storage {
	return &storage{
		db: db,
	}
}

func (s *storage) CreateTask(taskInput models.TaskInput) (models.TaskInternalShort, error) {
	var taskID int64

	err := s.db.QueryRow("INSERT INTO tasks (cardID, taskName, description, tasksOrder) VALUES ($1, $2, $3, $4) RETURNING taskID",
		taskInput.CardID, taskInput.Name, taskInput.Description, taskInput.Order).Scan(&taskID)

	if err != nil {
		return models.TaskInternalShort{}, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "CreateTask"}
	}

	task := models.TaskInternalShort{
		TaskID:      taskID,
		Name:        taskInput.Name,
		Description: taskInput.Description,
		Order:       taskInput.Order,
	}

	return task, nil
}

func (s *storage) ChangeTask(taskInput models.TaskInput) (models.TaskInternal, error) {
	var cardID int64
	err := s.db.QueryRow("UPDATE tasks SET taskName = $1, description = $2 WHERE taskID = $3 RETURNING cardID",
		taskInput.Name, taskInput.Description, taskInput.TaskID).Scan(&cardID)
	if err != nil {
		return models.TaskInternal{}, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "ChangeTask"}
	}

	task := models.TaskInternal{
		TaskID:      taskInput.TaskID,
		CardID:      cardID,
		Name:        taskInput.Name,
		Description: taskInput.Description,
	}

	return task, nil
}

func (s *storage) DeleteTask(taskInput models.TaskInput) (models.TaskInternalShort, error) {
	var cardID int64
	err := s.db.QueryRow("DELETE FROM tasks WHERE taskID = $1 RETURNING cardID", taskInput.TaskID).Scan(&cardID)
	if err != nil {
		return models.TaskInternalShort{}, models.ServeError{Codes: []string{"500"}, OriginalError: err, MethodName: "DeleteTask"}
	}

	task := models.TaskInternalShort{
		TaskID: taskInput.TaskID,
		CardID: cardID,
	}

	return task, nil
}

func (s *storage) GetTasksByCard(cardInput models.CardInput) ([]models.TaskInternalShort, error) {
	tasks := make([]models.TaskInternalShort, 0)

	rows, err := s.db.Query("SELECT taskID, taskName, description, tasksOrder FROM tasks WHERE cardID = $1", cardInput.CardID)
	if err != nil {
		return []models.TaskInternalShort{}, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "GetTasksByCard"}
	}
	defer rows.Close()

	for rows.Next() {
		var task models.TaskInternalShort

		err = rows.Scan(&task.TaskID, &task.Name, &task.Description, &task.Order)
		if err != nil {
			return []models.TaskInternalShort{}, models.ServeError{Codes: []string{"500"}, OriginalError: err,
				MethodName: "GetTasksByCard"}
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (s *storage) GetTaskByID(taskInput models.TaskInput) (models.TaskInternal, error) {
	task := models.TaskInternal{}
	task.TaskID = taskInput.TaskID

	err := s.db.QueryRow("SELECT taskName, description, tasksOrder FROM tasks WHERE taskID = $1", taskInput.TaskID).
		Scan(&task.Name, &task.Description, &task.Order)

	if err != nil {
		return models.TaskInternal{}, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "GetTaskByID"}
	}

	return task, nil
}

func (s *storage) GetTaskNameByID(taskID int64) (string, error) {
	var taskName string
	err := s.db.QueryRow("SELECT taskName FROM tasks WHERE taskID = $1", taskID).
		Scan(&taskName)

	if err != nil {
		return "", models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "GetTaskNameByID"}
	}

	return taskName, nil
}

func (s *storage) GetTaskName(taskInput models.TaskInput) (string, error) {
	var taskName string

	err := s.db.QueryRow("SELECT taskName FROM tasks WHERE taskID = $1", taskInput.TaskID).Scan(&taskName)
	if err != nil {
		return "", models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "GetTaskName"}
	}

	return taskName, nil
}

func (s *storage) ChangeTaskOrder(taskInput models.TasksOrderInput) error {
	for _, card := range taskInput.Tasks {
		for _, task := range card.Tasks {
			_, err := s.db.Exec("UPDATE tasks SET cardID = $1, tasksOrder = $2 WHERE taskID = $3",
				card.CardID, task.Order, task.TaskID)

			if err != nil {
				return models.ServeError{Codes: []string{"500"}, OriginalError: err, MethodName: "ChangeTaskOrder"}
			}
		}
	}

	return nil
}

func (s *storage) AssignUser(input models.TaskAssigner) (err error) {
	_, err = s.db.Exec("INSERT INTO task_members (taskID, userID) VALUES ($1, $2)", input.TaskID, input.AssignerID)
	if err != nil {
		return models.ServeError{Codes: []string{"500"}, OriginalError: err, MethodName: "AssignUser"}
	}
	return
}
func (s *storage) DismissUser(input models.TaskAssigner) (err error) {
	_, err = s.db.Exec("DELETE FROM task_members WHERE taskID = $1 AND userID = $2", input.TaskID, input.AssignerID)
	if err != nil {
		return models.ServeError{Codes: []string{"500"}, OriginalError: err, MethodName: "DismissUser"}
	}
	return
}

func (s *storage) GetAssigners(input models.TaskInput) (assignerIDs []int64, err error) {
	rows, err := s.db.Query("SELECT userID FROM task_members WHERE taskID = $1", input.TaskID)
	if err != nil {
		return assignerIDs, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "GetAssigners"}
	}
	defer rows.Close()

	for rows.Next() {
		var assignerID int64
		err = rows.Scan(&assignerID)
		if err != nil {
			return assignerIDs, models.ServeError{Codes: []string{"500"}, OriginalError: err,
				MethodName: "GetAssigners"}
		}

		assignerIDs = append(assignerIDs, assignerID)
	}

	return
}

func (s storage) GetCardIDByTask(taskID int64) (cardID int64, err error) {
	err = s.db.QueryRow("SELECT cardID FROM tasks WHERE taskID = $1", taskID).Scan(&cardID)
	if err != nil {
		return cardID, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "GetCardIDByTask"}
	}

	return
}
