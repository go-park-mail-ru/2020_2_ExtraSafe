package auth

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
)

type userStorage interface {
	LoginUser(userInput models.UserInputLogin) (models.User, error)
	RegisterUser(userInput models.UserInputReg) (models.User, error)
	GetUserProfile(userInput models.UserInput) (models.User, error)
}

type validator interface {
	ValidateLogin(request models.UserInputLogin) (err error)
	ValidateRegistration(request models.UserInputReg) (err error)
}