package boards

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
)

//go:generate mockgen -destination=./mock/mock_userStorage.go -package=mock github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/services/boards UserStorage
//go:generate mockgen -destination=./mock/mock_boardStorage.go -package=mock github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/services/boards BoardStorage
//go:generate mockgen -destination=./mock/mock_validator.go -package=mock github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/services/boards Validator

type UserStorage interface {
	GetBoardMembers(userIDs []int64) ([] models.UserOutsideShort, error) // 0 структура - админ доски
}

type BoardStorage interface {
	CreateBoard(boardInput models.BoardChangeInput) (models.BoardInternal, error)
	GetBoard(boardInput models.BoardInput) (models.BoardInternal, error)
	ChangeBoard(boardInput models.BoardChangeInput) (models.BoardInternal, error)
	DeleteBoard(boardInput models.BoardInput) error

	CreateCard(cardsInput models.CardInput) (models.CardOutside, error)
	GetCard(cardInput models.CardInput) (models.CardOutside, error)
	ChangeCard(cardInput models.CardInput) (models.CardOutside, error)
	DeleteCard(cardInput models.CardInput) error
	ChangeCardOrder(cardInput models.CardsOrderInput) error

	CreateTask(taskInput models.TaskInput) (models.TaskOutside, error)
	GetTask(taskInput models.TaskInput) (models.TaskOutside, error)
	ChangeTask(taskInput models.TaskInput) (models.TaskOutside, error)
	DeleteTask(taskInput models.TaskInput) error
	ChangeTaskOrder(taskInput models.TasksOrderInput) error
}

type Validator interface {
	ValidateLogin(request models.UserInputLogin) (err error)
	ValidateRegistration(request models.UserInputReg) (err error)
}
