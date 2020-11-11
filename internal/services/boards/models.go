package boards

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
)

type userStorage interface {
	GetBoardMembers(userIDs []uint64) ([] models.UserOutsideShort, error) // 0 структура - админ доски
}

type boardStorage interface {
	CreateBoard(boardInput models.BoardChangeInput) (models.BoardInternal, error)
	GetBoard(boardInput models.BoardInput) (models.BoardInternal, error)
	ChangeBoard(boardInput models.BoardChangeInput) (models.BoardInternal, error)
	DeleteBoard(boardInput models.BoardInput) error

	CreateCard(cardsInput models.CardInput) (models.CardOutside, error)
	GetCard(cardInput models.CardInput) (models.CardOutside, error)
	ChangeCard(cardsInput models.CardInput) (models.CardOutside, error)
	DeleteCard(cardsInput models.CardInput) error

	CreateTask(taskInput models.TaskInput) (models.TaskOutside, error)
	GetTask(taskInput models.TaskInput) (models.TaskOutside, error)
	ChangeTask(taskInput models.TaskInput) (models.TaskOutside, error)
	DeleteTask(taskInput models.TaskInput) error
	ChangeTaskOrder(taskInput models.TasksOrderInput) error
}

type validator interface {
	ValidateLogin(request models.UserInputLogin) (err error)
	ValidateRegistration(request models.UserInputReg) (err error)
}
