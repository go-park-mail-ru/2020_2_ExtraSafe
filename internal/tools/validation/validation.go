package validation

import (
	"github.com/asaskevich/govalidator"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"regexp"
)

var fullnameMaxLen = 40
var passwordMaxLen = 64
var passwordMinLen = 4
var usernameMaxLen = 40
var usernameMinLen = 2

type Service interface {
	ValidateLogin(request models.UserInputLogin) (err error)
	ValidateRegistration(request models.UserInputReg) (err error)
	ValidateProfile(request models.UserInputProfile) (err error)
	ValidateChangePassword(request models.UserInputPassword) (err error)
}

type service struct {
}

func NewService() Service {
	return &service{}
}

/*func collectErrors(err error) error {
	errorsCodes := make([]string, 0)
	errs := err.(govalidator.Errors).Errors()
	for _, e := range errs {
		errorsCodes = append(errorsCodes, e.Error())
	}
	return models.ServeError{Codes: errorsCodes}
}*/

func (s *service) ValidateLogin(request models.UserInputLogin) (err error) {
	_, err = govalidator.ValidateStruct(request)
	if err != nil {
		return models.ServeError{Codes: []string{"900"}, Descriptions: []string{"Validation error"},
			MethodName: "ValidateLogin"}
	}
	return nil
}

func (s *service) ValidateRegistration(request models.UserInputReg) (err error) {
	_, err = govalidator.ValidateStruct(request)
	if err != nil {
		return models.ServeError{Codes: []string{"900"}, Descriptions: []string{"Validation error"},
			MethodName: "ValidateRegistration"}
	}
	return nil
}

func (s *service) ValidateProfile(request models.UserInputProfile) (err error) {
	_, err = govalidator.ValidateStruct(request)
	if err != nil {
		return models.ServeError{Codes: []string{"900"}, Descriptions: []string{"Validation error"},
			MethodName: "ValidateProfile"}
	}
	return nil
}

func (s *service) ValidateChangePassword(request models.UserInputPassword) (err error) {
	_, err = govalidator.ValidateStruct(request)
	if err != nil {
		return models.ServeError{Codes: []string{"900"}, Descriptions: []string{"Validation error"},
			MethodName: "ValidateChangePassword"}
	}
	return nil
}

func IsPasswordValid(i interface{}, _ interface{}) bool {
	subject, ok := i.(string)
	if !ok {
		return false
	}
	if len(subject) == 0 || len(subject) < passwordMinLen || len(subject) > passwordMaxLen {
		return false
	}

	re := regexp.MustCompile("^[a-zA-Z0-9~!@#$%^&*-_+=`|(){}:;\"'<>,.?/]+$")
	return re.MatchString(subject)
}

func IsFullNameValid(i interface{}, _ interface{}) bool {
	subject, ok := i.(string)
	if !ok {
		return false
	}
	if len(subject) == 0 || len(subject) > fullnameMaxLen {
		return false
	}

	re := regexp.MustCompile("^[a-zA-Zа-яА-Я _]+")
	return re.MatchString(subject)
}

func IsUsernameValid(i interface{}, _ interface{}) bool {
	subject, ok := i.(string)
	if !ok {
		return false
	}
	if len(subject) == 0 || len(subject) < usernameMinLen || len(subject) > usernameMaxLen {
		return false
	}

	re := regexp.MustCompile("^[a-zA-Z0-9_]+$")
	return re.MatchString(subject)
}

func IsEmailValid(i interface{}, _ interface{}) bool {
	subject, ok := i.(string)
	if !ok {
		return false
	}
	if len(subject) == 0 {
		return false
	}

	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(subject)
}

func init() {
	govalidator.CustomTypeTagMap.Set("passwordValid", govalidator.CustomTypeValidator(IsPasswordValid))
	govalidator.CustomTypeTagMap.Set("fullNameValid", govalidator.CustomTypeValidator(IsFullNameValid))
	govalidator.CustomTypeTagMap.Set("userNameValid", govalidator.CustomTypeValidator(IsUsernameValid))
	govalidator.CustomTypeTagMap.Set("emailValid", govalidator.CustomTypeValidator(IsEmailValid))
}
