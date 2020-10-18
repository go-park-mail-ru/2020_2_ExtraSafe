package profile

import (
	"../../../internal/models"
)

type Service interface {
	Profile(request models.UserInput) (user models.User, err error)
	Accounts(request models.UserInput) (user models.User, err error)
	ProfileChange(request models.UserInputProfile) (user models.User, err error)
	AccountsChange(request models.UserInputLinks) (user models.User, err error)
	PasswordChange(request models.UserInputPassword) (user models.User, err error)
}

type service struct {
	userStorage userStorage
	avatarStorage avatarStorage
}

func NewService(userStorage userStorage, avatarStorage avatarStorage) Service {
	return &service{
		userStorage: userStorage,
		avatarStorage: avatarStorage,
	}
}

func (s *service) Profile(request models.UserInput) (user models.User, err error) {
	user, err = s.userStorage.GetUserProfile(request)
	if err != nil {
		return models.User{}, err
	}

	return user, err
}

func (s *service) Accounts(request models.UserInput) (user models.User, err error) {
	user, err = s.userStorage.GetUserAccounts(request)
	if err != nil {
		return models.User{}, err
	}

	return user, err
}

func (s *service) ProfileChange(request models.UserInputProfile) (user models.User, err error) {
	if request.Avatar != nil {
		err, _ = s.avatarStorage.UploadAvatar(request.Avatar, request.ID)
	}

	user, err = s.userStorage.ChangeUserProfile(request)
	if err != nil {
		return models.User{}, err
	}

	return user, err
}

func (s *service) AccountsChange(request models.UserInputLinks) (user models.User, err error) {
	user, err = s.userStorage.ChangeUserAccounts(request)
	if err != nil {
		return models.User{}, err
	}

	return user, err
}

func (s *service) PasswordChange(request models.UserInputPassword) (user models.User, err error) {
	user, err = s.userStorage.ChangeUserPassword(request)
	if err != nil {
		return models.User{}, err
	}

	return user, err
}