package boardStorage

import "github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"

//go:generate mockgen -destination=./mock/mock_cardsStorage.go -package=mock github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/boardStorage CardsStorage
//go:generate mockgen -destination=./mock/mock_tasksStorage.go -package=mock github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/boardStorage TasksStorage

type CardsStorage interface {
	CreateCard(cardInput models.CardInput) (models.CardOutside, error)
	ChangeCard(userInput models.CardInput) (models.CardInternal, error)
	DeleteCard(userInput models.CardInput) error

	GetCardsByBoard(boardInput models.BoardInput) ([]models.CardInternal, error)
	GetCardByID(cardInput models.CardInput) (models.CardInternal, error)
	ChangeCardOrder(taskInput models.CardsOrderInput) error
}

type TasksStorage interface {
	CreateTask(taskInput models.TaskInput) (models.TaskInternalShort, error)
	ChangeTask(taskInput models.TaskInput) (models.TaskInternal, error)
	DeleteTask(taskInput models.TaskInput) (models.TaskInternalShort, error)

	GetTasksByCard(cardInput models.CardInput) ([]models.TaskInternalShort, error)
	GetTaskByID(taskInput models.TaskInput) (models.TaskInternal, error)
	GetTaskName(taskInput models.TaskInput) (string, error)
	ChangeTaskOrder(taskInput models.TasksOrderInput) error

	AssignUser(input models.TaskAssigner) (err error)
	DismissUser(input models.TaskAssigner) (err error)
	GetAssigners(input models.TaskInput) (assignerIDs []int64, err error)

	GetCardIDByTask(taskInput int64) (cardID int64, err error)
}
