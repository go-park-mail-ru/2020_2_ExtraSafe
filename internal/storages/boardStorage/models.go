package boardStorage

import "github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"

//go:generate mockgen -destination=./mock/mock_cardsStorage.go -package=mock github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/storages/boardStorage CardsStorage
//go:generate mockgen -destination=./mock/mock_tasksStorage.go -package=mock github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/storages/boardStorage TasksStorage


type CardsStorage interface {
	CreateCard(userInput models.CardInput) (models.CardOutside, error)
	ChangeCard(userInput models.CardInput) (models.CardOutside, error)
	DeleteCard(userInput models.CardInput) error

	GetCardsByBoard(boardInput models.BoardInput) ([]models.CardOutside, error)
	GetCardByID(cardInput models.CardInput) (models.CardOutside, error)
	ChangeCardOrder(taskInput models.CardsOrderInput) error
}

type TasksStorage interface {
	CreateTask(taskInput models.TaskInput) (models.TaskOutside, error)
	ChangeTask(taskInput models.TaskInput) (models.TaskOutside, error)
	DeleteTask(taskInput models.TaskInput) error

	GetTasksByCard(cardInput models.CardInput) ([]models.TaskOutside, error)
	GetTaskByID(taskInput models.TaskInput) (models.TaskOutside, error)
	ChangeTaskOrder(taskInput models.TasksOrderInput) error
}
