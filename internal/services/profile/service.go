package profile

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
)

type Service interface {
	Profile(request models.UserInput) (user models.UserOutside, err error)
	Accounts(request models.UserInput) (user models.UserOutside, err error)
	ProfileChange(request models.UserInputProfile) (user models.UserOutside, err error)
	AccountsChange(request models.UserInputLinks) (user models.UserOutside, err error)
	PasswordChange(request models.UserInputPassword) (user models.UserOutside, err error)
}

type service struct {
	userStorage userStorage
	avatarStorage avatarStorage
	validator validator
}

func NewService(userStorage userStorage, avatarStorage avatarStorage, validator validator) Service {
	return &service{
		userStorage: userStorage,
		avatarStorage: avatarStorage,
		validator: validator,
	}
}

func (s *service) Profile(request models.UserInput) (user models.UserOutside, err error) {
	user, err = s.userStorage.GetUserProfile(request)
	if err != nil {
		return models.UserOutside{}, err
	}

	return user, err
}

func (s *service) Accounts(request models.UserInput) (user models.UserOutside, err error) {
	user, err = s.userStorage.GetUserAccounts(request)
	if err != nil {
		return models.UserOutside{}, err
	}

	return user, err
}

func (s *service) ProfileChange(request models.UserInputProfile) (user models.UserOutside, err error) {
	errorCodes := make([]string, 0)

	err = s.validator.ValidateProfile(request)
	if err != nil {
		return models.UserOutside{}, err
	}

	if request.Avatar != nil {
		errAvatar, _ := s.avatarStorage.UploadAvatar(request.Avatar, request.ID)
		if errAvatar != nil {
			errorCodes = append(errorCodes, errAvatar.(models.ServeError).Codes...)
		}
	}

	user, errProfile := s.userStorage.ChangeUserProfile(request)
	if errProfile != nil {
		errorCodes = append(errorCodes, errProfile.(models.ServeError).Codes...)
	}

	if len(errorCodes) != 0 {
		return models.UserOutside{}, models.ServeError{Codes: errorCodes}
	}

	return user, err
}

func (s *service) AccountsChange(request models.UserInputLinks) (user models.UserOutside, err error) {
	err = s.validator.ValidateLinks(request)
	if err != nil {
		return models.UserOutside{}, err
	}

	user, err = s.userStorage.ChangeUserAccounts(request)
	if err != nil {
		return models.UserOutside{}, err
	}

	return user, err
}

func (s *service) PasswordChange(request models.UserInputPassword) (user models.UserOutside, err error) {
	err = s.validator.ValidateChangePassword(request)
	if err != nil {
		return models.UserOutside{}, err
	}

	user, err = s.userStorage.ChangeUserPassword(request)
	if err != nil {
		return models.UserOutside{}, err
	}

	return user, err
}
