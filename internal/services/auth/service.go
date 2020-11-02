package auth

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
)

type Service interface {
	Auth(request models.UserInput) (response models.UserOutside, err error)
	Login(request models.UserInputLogin) (response models.UserOutside, err error)
	Registration(request models.UserInputReg) (response models.UserOutside, err error)
}

type service struct {
	userStorage userStorage
	validator validator
}

func NewService(userStorage userStorage, validator validator) Service {
	return &service{
		userStorage: userStorage,
		validator: validator,
	}
}

func (s *service)Auth(request models.UserInput) (response models.UserOutside, err error) {
	response, err = s.userStorage.GetUserProfile(request)
	return response, err
}

func (s *service)Login(request models.UserInputLogin) (response models.UserOutside, err error) {
	var user models.UserOutside

	err = s.validator.ValidateLogin(request)
	if err != nil {
		return models.UserOutside{}, err
	}

	user, err = s.userStorage.CheckUser(request)
	if err != nil {
		return models.UserOutside{}, err
	}

	return user, err
}


func (s *service)Registration(request models.UserInputReg) (response models.UserOutside, err error) {
	var user models.UserOutside

	err = s.validator.ValidateRegistration(request)
	if err != nil {
		return models.UserOutside{}, err
	}

	user, err = s.userStorage.CreateUser(request)
	if err != nil {
		return models.UserOutside{}, err
	}

	return user, err
}

