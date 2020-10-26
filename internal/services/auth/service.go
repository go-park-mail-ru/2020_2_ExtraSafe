package auth

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
)

type Service interface {
	Auth(request models.UserInput) (response models.User, err error)
	Login(request models.UserInputLogin) (response models.User, err error)
	Registration(request models.UserInputReg) (response models.User, err error)
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

func (s *service)Auth(request models.UserInput) (response models.User, err error) {
	response, err = s.userStorage.GetUserProfile(request)
	return response, err
}

func (s *service)Login(request models.UserInputLogin) (response models.User, err error) {
	var user models.User

	/*_, err = govalidator.ValidateStruct(request)
	if err != nil {
		return models.User{}, models.ServeError{Codes: []string{"103"}}
	}
*/
	err = s.validator.ValidateLogin(request)
	if err != nil {
		return models.User{}, err
	}

	user, err = s.userStorage.CheckUser(request)
	if err != nil {
		return models.User{}, err
	}

	return user, err
}

func (s *service)Registration(request models.UserInputReg) (response models.User, err error) {
	var user models.User

	/*_, err = govalidator.ValidateStruct(request)
	if err != nil {
		return models.User{}, models.ServeError{Codes: []string{"103"}}
	}*/

	err = s.validator.ValidateRegistration(request)

	user, err = s.userStorage.CreateUser(request)
	if err != nil {
		return models.User{}, err
	}

	return user, err
}

