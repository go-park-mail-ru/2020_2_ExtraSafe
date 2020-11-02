package auth

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
)

type userStorage interface {
	CheckUser(userInput models.UserInputLogin) (models.UserOutside, error)
	CreateUser(userInput models.UserInputReg) (models.UserOutside, error)
	GetUserProfile(userInput models.UserInput) (models.UserOutside, error)
}

type validator interface {
	ValidateLogin(request models.UserInputLogin) (err error)
	ValidateRegistration(request models.UserInputReg) (err error)
}