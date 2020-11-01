package profile

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"mime/multipart"
)

type userStorage interface {
	GetUserProfile(userInput models.UserInput) (models.User, error)
	GetUserAccounts(userInput models.UserInput) (models.User, error)

	ChangeUserProfile(userInput models.UserInputProfile) (models.User, error)
	ChangeUserAccounts(userInput models.UserInputLinks) (models.User, error)
	ChangeUserPassword(userInput models.UserInputPassword) (models.User, error)
}

type avatarStorage interface {
	UploadAvatar(file *multipart.FileHeader, userID uint64) (err error, filename string)
}

type validator interface {
	ValidateProfile(request models.UserInputProfile) (err error)
	ValidateChangePassword(request models.UserInputPassword) (err error)
	ValidateLinks(request models.UserInputLinks) (err error)
}