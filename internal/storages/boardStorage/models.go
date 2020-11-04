package boardStorage

import "github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"

type cardsStorage interface {
	CreateCard(userInput models.CardInput) (models.CardOutside, error)
	ChangeCard(userInput models.CardInput) (models.CardOutside, error)
	DeleteCard(userInput models.CardInput) error

	GetCardsByBoard(boardInput models.BoardInput) ([]models.CardOutside, error)
}

type tasksStorage interface {
	CreateTask(taskInput models.TaskInput) (models.TaskOutside, error)
	ChangeTask(taskInput models.TaskInput) (models.TaskOutside, error)
	DeleteTask(taskInput models.TaskInput) error

	GetTasksByCard(cardInput models.CardInput) ([]models.TaskOutside, error)
}
