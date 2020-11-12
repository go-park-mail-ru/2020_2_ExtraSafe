package auth

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
)

//go:generate mockgen -destination=./mock/mock_userStorage.go -package=mock github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/services/auth UserStorage
//go:generate mockgen -destination=./mock/mock_boardStorage.go -package=mock github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/services/auth BoardStorage
//go:generate mockgen -destination=./mock/mock_validator.go -package=mock github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/services/auth Validator

type UserStorage interface {
	CheckUser(userInput models.UserInputLogin) (uint64, models.UserOutside, error)
	CreateUser(userInput models.UserInputReg) (uint64, models.UserOutside, error)
	GetUserProfile(userInput models.UserInput) (models.UserOutside, error)
}

type BoardStorage interface {
	GetBoardsList(userInput models.UserInput) ([]models.BoardOutsideShort, error)
}

type Validator interface {
	ValidateLogin(request models.UserInputLogin) (err error)
	ValidateRegistration(request models.UserInputReg) (err error)
}