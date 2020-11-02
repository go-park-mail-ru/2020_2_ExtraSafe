package validaton

import (
	"github.com/asaskevich/govalidator"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"regexp"
)

type Service interface {
	ValidateLogin(request models.UserInputLogin) (err error)
	ValidateRegistration(request models.UserInputReg) (err error)
	ValidateProfile(request models.UserInputProfile) (err error)
	ValidateChangePassword(request models.UserInputPassword) (err error)
	//ValidateLinks...
}

type service struct {
}

func NewService() Service {
	return &service{}
}

//TODO корректную обработку ошибок
//TODO вынести магические числа в константы


//TODO сделать функцию, которая по непровалидированному полю будет формировать нужный код ошибки (использовать ошибки govalidator в полях структуры) (пока код 103)
//https://github.com/asaskevich/govalidator#custom-error-messages

func (s *service) ValidateLogin(request models.UserInputLogin) (err error) {
	_, err = govalidator.ValidateStruct(request)
	if err != nil {
		return models.ServeError{Codes: []string{"103"}}
	}
	return nil
}

func (s *service) ValidateRegistration(request models.UserInputReg) (err error) {
	_, err = govalidator.ValidateStruct(request)
	if err != nil {
		return models.ServeError{Codes: []string{"103"}}
	}
	return nil
}

func (s *service) ValidateProfile(request models.UserInputProfile) (err error) {
	_, err = govalidator.ValidateStruct(request)
	if err != nil {
		return models.ServeError{Codes: []string{"103"}}
	}
	return nil
}

func (s *service) ValidateChangePassword(request models.UserInputPassword) (err error) {
	_, err = govalidator.ValidateStruct(request)
	if err != nil {
		return models.ServeError{Codes: []string{"103"}}
	}
	return nil
}

func IsPasswordValid(i interface{}, o interface{}) bool {
	subject, ok := i.(string)
	if !ok {
		return false
	}
	if len(subject) == 0 || len(subject) < 4 || len(subject) > 64 {
		return false
	}

	re := regexp.MustCompile( "^[a-zA-Z0-9~!@#$%^&*-_+=`|\\(){}:;\"'<>,.?/]+$")
	return re.MatchString(subject)
}

func IsFullNameValid(i interface{}, o interface{}) bool {
	subject, ok := i.(string)
	if !ok {
		return false
	}
	if len(subject) == 0 || len(subject) > 40 {
		return false
	}

	re := regexp.MustCompile( "^[a-zA-Zа-яА-Я _]+")
	return re.MatchString(subject)
}

func IsUsernameValid(i interface{}, o interface{}) bool {
	subject, ok := i.(string)
	if !ok {
		return false
	}
	if len(subject) == 0 || len(subject) < 2 || len(subject) > 40 {
		return false
	}

	re := regexp.MustCompile( "^[a-zA-Z0-9_]+$")
	return re.MatchString(subject)
}

func init() {
	govalidator.CustomTypeTagMap.Set("passwordValid", govalidator.CustomTypeValidator(IsPasswordValid))
	govalidator.CustomTypeTagMap.Set("fullNameValid", govalidator.CustomTypeValidator(IsFullNameValid))
	govalidator.CustomTypeTagMap.Set("userNameValid", govalidator.CustomTypeValidator(IsUsernameValid))
}