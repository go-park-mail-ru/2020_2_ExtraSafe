package profile

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
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
	validator validator
}

func NewService(userStorage userStorage, avatarStorage avatarStorage, validator validator) Service {
	return &service{
		userStorage: userStorage,
		avatarStorage: avatarStorage,
		validator: validator,
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
	errorCodes := make([]string, 0)

	err = s.validator.ValidateProfile(request)
	if err != nil {
		return models.User{}, err
	}


	//TODO здесь должны быть return по ошибкам

	user, errGetProfile := s.userStorage.GetUserProfile(models.UserInput{ID: request.ID})
	if errGetProfile != nil {
		errorCodes = append(errorCodes, errGetProfile.(models.ServeError).Codes...)
		return user, models.ServeError{Codes: errorCodes}
		//TODO сделать изменение аватара в таблице
	}

	if request.Avatar != nil {
		errAvatar, _ := s.avatarStorage.UploadAvatar(request.Avatar, &user)
		if errAvatar != nil {
			errorCodes = append(errorCodes, errAvatar.(models.ServeError).Codes...)
		}
	}

	errProfile := s.userStorage.ChangeUserProfile(&user, request)
	if errProfile != nil {
		errorCodes = append(errorCodes, errProfile.(models.ServeError).Codes...)
	}

	if len(errorCodes) != 0 {
		return models.User{}, models.ServeError{Codes: errorCodes}
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
	err = s.validator.ValidateChangePassword(request)
	if err != nil {
		return models.User{}, err
	}

	user, err = s.userStorage.ChangeUserPassword(request)
	if err != nil {
		return models.User{}, err
	}

	return user, err
}
