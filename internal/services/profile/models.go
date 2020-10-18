package profile

import (
	"../../../internal/models"
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