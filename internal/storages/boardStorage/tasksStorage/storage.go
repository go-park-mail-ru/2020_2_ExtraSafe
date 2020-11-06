package tasksStorage

import (
	"database/sql"
	"fmt"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
)

type Storage interface {
	CreateTask(taskInput models.TaskInput) (models.TaskOutside, error)
	ChangeTask(taskInput models.TaskInput) (models.TaskOutside, error)
	DeleteTask(taskInput models.TaskInput) error

	GetTasksByCard(cardInput models.CardInput) ([]models.TaskOutside, error)
	GetTaskByID(taskInput models.TaskInput) (models.TaskOutside, error)

	CheckTaskAccessory(taskID uint64) (cardID uint64, err error)
}

type storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) Storage {
	return &storage{
		db: db,
	}
}

func (s *storage) CreateTask(taskInput models.TaskInput) (models.TaskOutside, error) {
	var taskID uint64 = 0
	var quantityTasks uint64 = 0

	//FIXME сделать исправление ID
	err := s.db.QueryRow("SELECT COUNT(*) FROM tasks").Scan(&quantityTasks)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err)
		return models.TaskOutside{}, models.ServeError{Codes: []string{"500"}}
	}

	taskID = quantityTasks + 1

	_, err = s.db.Exec("INSERT INTO tasks (taskID, cardID, taskName, description, tasksOrder) VALUES ($1, $2, $3, $4, $5)",
								taskID, taskInput.CardID, taskInput.Name, taskInput.Description, taskInput.Order)

	if err != nil {
		//TODO error
		fmt.Println(err)
		return models.TaskOutside{} ,models.ServeError{Codes: []string{"500"}}
	}

	task := models.TaskOutside{
		TaskID:      taskID,
		Name:        taskInput.Name,
		Description: taskInput.Description,
		Order:       taskInput.Order,
	}

	return task, nil
}

func (s *storage) ChangeTask(taskInput models.TaskInput) (models.TaskOutside, error) {
	_, err := s.db.Exec("UPDATE tasks SET taskName = $1, description = $2, tasksOrder = $3 WHERE taskID = $4",
								taskInput.Name, taskInput.Description, taskInput.Order, taskInput.TaskID)
	if err != nil {
		fmt.Println(err)
		return models.TaskOutside{}, models.ServeError{Codes: []string{"500"}}
	}

	task := models.TaskOutside{
		TaskID:      taskInput.TaskID,
		Name:        taskInput.Name,
		Description: taskInput.Description,
		Order:       taskInput.Order,
	}

	return task, nil
}

func (s *storage) DeleteTask(taskInput models.TaskInput) error {
	_, err := s.db.Exec("DELETE FROM tasks WHERE taskID = $1", taskInput.TaskID)
	if err != nil {
		fmt.Println(err)
		return models.ServeError{Codes: []string{"500"}}
	}

	return nil
}

func (s *storage) GetTasksByCard(cardInput models.CardInput) ([]models.TaskOutside, error) {
	tasks := make([]models.TaskOutside, 0)

	rows, err := s.db.Query("SELECT taskID, taskName, description, tasksOrder FROM tasks WHERE cardID = $1", cardInput.CardID)
	if err != nil {
		return []models.TaskOutside{}, models.ServeError{Codes: []string{"500"}}
	}
	defer rows.Close()

	for rows.Next() {
		var task models.TaskOutside

		err = rows.Scan(&task.TaskID, &task.Name, &task.Description, &task.Order)
		if err != nil {
			return []models.TaskOutside{}, models.ServeError{Codes: []string{"500"}}
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (s *storage) GetTaskByID(taskInput models.TaskInput) (models.TaskOutside, error) {
	task := models.TaskOutside{}
	task.TaskID = taskInput.TaskID

	err := s.db.QueryRow("SELECT taskName, description, tasksOrder FROM tasks WHERE taskID = $1", taskInput.TaskID).
				Scan(&task.Name, &task.Description, &task.Order)

	if err != nil {
		fmt.Println(err)
		return models.TaskOutside{}, models.ServeError{Codes: []string{"500"}}
	}

	return task, nil
}

func (s *storage) CheckTaskAccessory(taskID uint64) (cardID uint64, err error) {
	err = s.db.QueryRow("SELECT cardID FROM tasks WHERE taskID = $1", taskID).Scan(&cardID)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return cardID, nil
}
