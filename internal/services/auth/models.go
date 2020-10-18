package auth

import (
	"../../../internal/models"
)

type userStorage interface {
	CheckUser(userInput models.UserInputLogin) (models.User, error)
	CreateUser(userInput models.UserInputReg) (models.User, error)
	GetUserProfile(userInput models.UserInput) (models.User, error)
}
