package boardStorage

import "github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"

type cardsStorage interface {
	CreateCard(userInput models.CardInput) (models.CardOutside, error)
	ChangeCard(userInput models.CardInput) (models.CardOutside, error)
	DeleteCard(userInput models.CardInput) error

	GetCardsByBoard(boardInput models.BoardInput) ([]models.CardOutside, error)
	GetCardByID(cardInput models.CardInput) (models.CardOutside, error)
	CheckCardAccessory(cardID uint64) (boardID uint64, err error)
}

type tasksStorage interface {
	CreateTask(taskInput models.TaskInput) (models.TaskOutside, error)
	ChangeTask(taskInput models.TaskInput) (models.TaskOutside, error)
	DeleteTask(taskInput models.TaskInput) error

	GetTasksByCard(cardInput models.CardInput) ([]models.TaskOutside, error)
	GetTaskByID(taskInput models.TaskInput) (models.TaskOutside, error)
	CheckTaskAccessory(taskID uint64) (cardID uint64, err error)
}
