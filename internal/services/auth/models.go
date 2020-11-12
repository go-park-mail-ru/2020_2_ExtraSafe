package auth

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
)

type userStorage interface {
	CheckUser(userInput models.UserInputLogin) (uint64, models.UserOutside, error)
	CreateUser(userInput models.UserInputReg) (uint64, models.UserOutside, error)
	GetUserProfile(userInput models.UserInput) (models.UserOutside, error)
	GetUserAccounts(userInput models.UserInput) (models.UserOutside, error)
}

type boardStorage interface {
	GetBoardsList(userInput models.UserInput) ([]models.BoardOutsideShort, error)
}

type validator interface {
	ValidateLogin(request models.UserInputLogin) (err error)
	ValidateRegistration(request models.UserInputReg) (err error)
}