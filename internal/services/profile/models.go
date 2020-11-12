package profile

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"mime/multipart"
)

//go:generate mockgen -destination=./mock/mock_userStorage.go -package=mock github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/services/profile UserStorage
//go:generate mockgen -destination=./mock/mock_avatarStorage.go -package=mock github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/services/profile AvatarStorage
//go:generate mockgen -destination=./mock/mock_validator.go -package=mock github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/services/profile Validator
//go:generate mockgen -destination=./mock/mock_boardStorage.go -package=mock github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/services/profile BoardStorage

type UserStorage interface {
	GetUserProfile(userInput models.UserInput) (models.UserOutside, error)
	GetUserAccounts(userInput models.UserInput) (models.UserOutside, error)
	GetUserAvatar(userInput models.UserInput) (models.UserAvatar, error)

	ChangeUserProfile(userInput models.UserInputProfile, userAvatar models.UserAvatar) (models.UserOutside, error)
	ChangeUserAccounts(userInput models.UserInputLinks) (models.UserOutside, error)
	ChangeUserPassword(userInput models.UserInputPassword) (models.UserOutside, error)
}

type AvatarStorage interface {
	UploadAvatar(file *multipart.FileHeader, user *models.UserAvatar) error
}

type BoardStorage interface {
	GetBoardsList(userInput models.UserInput) ([]models.BoardOutsideShort, error)
}

type Validator interface {
	ValidateProfile(request models.UserInputProfile) (err error)
	ValidateChangePassword(request models.UserInputPassword) (err error)
	ValidateLinks(request models.UserInputLinks) (err error)
}