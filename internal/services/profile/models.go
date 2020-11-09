package profile

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"mime/multipart"
)

type userStorage interface {
	GetUserProfile(userInput models.UserInput) (models.UserOutside, error)
	GetUserAccounts(userInput models.UserInput) (models.UserOutside, error)
	GetUserAvatar(userInput models.UserInput) (models.UserAvatar, error)

	ChangeUserProfile(userInput models.UserInputProfile, userAvatar models.UserAvatar) (models.UserOutside, error)
	ChangeUserAccounts(userInput models.UserInputLinks) (models.UserOutside, error)
	ChangeUserPassword(userInput models.UserInputPassword) (models.UserOutside, error)
}

type avatarStorage interface {
	UploadAvatar(file *multipart.FileHeader, user *models.UserAvatar) error
}

type boardStorage interface {
	GetBoardsList(userInput models.UserInput) ([]models.BoardOutsideShort, error)
}

type validator interface {
	ValidateProfile(request models.UserInputProfile) (err error)
	ValidateChangePassword(request models.UserInputPassword) (err error)
	ValidateLinks(request models.UserInputLinks) (err error)
}