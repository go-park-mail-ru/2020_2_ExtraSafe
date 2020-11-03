package boards

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
)

type userStorage interface {
	CheckUser(userInput models.UserInputLogin) (models.UserOutside, error)
	CreateUser(userInput models.UserInputReg) (models.UserOutside, error)
	GetUserProfile(userInput models.UserInput) (models.UserOutside, error)

	GetBoardMembers(userIDs []uint64) ([] models.UserOutsideShort, error) // 0 структура - админ доски
}

type boardStorage interface {
	CreateBoard(boardInput models.BoardChangeInput) (models.BoardInternal, error)
	GetBoard(boardInput models.BoardInput) (models.BoardInternal, error)
	ChangeBoard(boardInput models.BoardChangeInput) (models.BoardInternal, error)
	DeleteBoard(boardInput models.BoardInput) error

	CreateColumn(cardInput models.CardInput) (models.CardOutside, error)
	ChangeColumn(cardInput models.CardInput) (models.CardOutside, error)
	DeleteColumn(cardInput models.CardInput) error

	CreateTask(taskInput models.TaskInput) (models.TaskOutside, error)
	ChangeTask(taskInput models.TaskInput) (models.TaskOutside, error)
	DeleteTask(taskInput models.TaskInput) error
}

type validator interface {
	ValidateLogin(request models.UserInputLogin) (err error)
	ValidateRegistration(request models.UserInputReg) (err error)
}
