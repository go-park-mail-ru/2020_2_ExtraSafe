package boards

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
)

type userStorage interface {
	CheckUser(userInput models.UserInputLogin) (models.UserOutside, error)
	CreateUser(userInput models.UserInputReg) (models.UserOutside, error)
	GetUserProfile(userInput models.UserInput) (models.UserOutside, error)
}

type boardStorage interface {
	CreateBoard(userInput models.BoardChangeInput) (models.BoardOutside, error)
	GetBoard(userInput models.BoardInput) (models.BoardOutside, error)
	ChangeBoard(userInput models.BoardChangeInput) (models.BoardOutside, error)
	DeleteBoard(userInput models.BoardInput) error

	CreateColumn(userInput models.ColumnInput) (models.ColumnOutside, error)
	ChangeColumn(userInput models.ColumnInput) (models.ColumnOutside, error)
	DeleteColumn(userInput models.ColumnInput) error

	CreateTask(userInput models.TaskInput) (models.TaskOutside, error)
	ChangeTask(userInput models.TaskInput) (models.TaskOutside, error)
	DeleteTask(userInput models.TaskInput) error
}

type validator interface {
	ValidateLogin(request models.UserInputLogin) (err error)
	ValidateRegistration(request models.UserInputReg) (err error)
}
