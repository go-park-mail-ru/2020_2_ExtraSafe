package auth

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
)

type userStorage interface {
	CheckUser(userInput models.UserInputLogin) (models.User, error)
	CreateUser(userInput models.UserInputReg) (models.User, error)
	GetUserProfile(userInput models.UserInput) (models.User, error)
}
