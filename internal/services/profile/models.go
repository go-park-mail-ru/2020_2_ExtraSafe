package profile

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"mime/multipart"
)

type userStorage interface {
	GetUserProfile(userInput models.UserInput) (models.UserOutside, error)
	GetUserAccounts(userInput models.UserInput) (models.UserOutside, error)

	ChangeUserProfile(userInput models.UserInputProfile) (models.UserOutside, error)
	ChangeUserAccounts(userInput models.UserInputLinks) (models.UserOutside, error)
	ChangeUserPassword(userInput models.UserInputPassword) (models.UserOutside, error)
}

type avatarStorage interface {
	UploadAvatar(file *multipart.FileHeader, userID uint64) (err error, filename string)
}

type validator interface {
	ValidateProfile(request models.UserInputProfile) (err error)
	ValidateChangePassword(request models.UserInputPassword) (err error)
	ValidateLinks(request models.UserInputLinks) (err error)
}