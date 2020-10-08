package main

import (
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"mime/multipart"
	"testing"
)

func TestCreateUserSuccess(t *testing.T) {
	someUsers := make([]User, 0)
	sessions := make(map[string]uint64, 10)

	var c echo.Context
	cc := &Handlers{c,
		&someUsers,
		&sessions,
	}

	testUser := UserInputReg{"someEmail@gmail.com", "someUsername", "somePassword"}
	_, _, err := cc.createUser(testUser)
	assert.Equal(t, nil, err)
}

func TestCreateUserFail(t *testing.T) {
	someUsers := make([]User, 0)
	someUsers = append(someUsers, User{
		ID:       0,
		Username: "someUsername",
		Email:    "someEmail@gmail.com",
		Password: "somePassword",
		Links:    &UserLinks{},
		Avatar:   "default/default_avatar.png",
	})

	sessions := make(map[string]uint64, 10)

	var c echo.Context
	cc := &Handlers{c,
		&someUsers,
		&sessions,
	}

	testUser := UserInputReg{"someEmail@gmail.com", "someUsername", "somePassword"}
	_, _, err := cc.createUser(testUser)

	assert.Error(t, err)
}

func TestCheckUserSuccess(t *testing.T) {
	someUsers := make([]User, 0)

	someUsers = append(someUsers, User{
		ID:       0,
		Username: "someUsername",
		Email:    "someEmail@gmail.com",
		Password: "somePassword",
		Links:    &UserLinks{},
		Avatar:   "default/default_avatar.png",
		FullName: "Petr",
	})

	sessions := make(map[string]uint64, 10)

	var c echo.Context
	cc := &Handlers{c,
		&someUsers,
		&sessions,
	}

	testUser := UserInputLogin{"someEmail@gmail.com", "somePassword"}

	expectResponse := responseUser{
		Status:   200,
		Email:    "someEmail@gmail.com",
		Username: "someUsername",
		FullName: "Petr",
		Avatar:   "default/default_avatar.png",
	}

	response, _, _ := cc.checkUser(testUser)

	assert.Equal(t, expectResponse, response)
}

func TestCheckUserFault(t *testing.T) {
	someUsers := make([]User, 0)
	someUsers = append(someUsers, User{
		ID:       0,
		Username: "someUsername",
		Email:    "someEmail@gmail.com",
		Password: "somePassword",
		Links:    &UserLinks{},
		Avatar:   "default/default_avatar.png",
	})

	sessions := make(map[string]uint64, 10)

	var c echo.Context
	cc := &Handlers{c,
		&someUsers,
		&sessions,
	}

	testUser := UserInputLogin{"someEmail@gmail.com", "Password"}

	errorMessage := []Messages{{Message: "Неверная электронная почта или пароль", ErrorName: "password"}}

	testResponse := responseError{
		Status:        500,
		Messages:      errorMessage,
	}

	_, _, err := cc.checkUser(testUser)

	assert.Equal(t, testResponse, err)
}

func TestChangeUserProfileSuccess(t *testing.T) {
	someUsers := make([]User, 0)
	sessions := make(map[string]uint64, 10)

	someUsers = append(someUsers, User{
		ID:       0,
		Username: "someUsername",
		Email:    "someEmail@gmail.com",
		Password: "somePassword",
		Links:    &UserLinks{},
		Avatar:   "default/default_avatar.png",
	})

	var c echo.Context
	cc := &Handlers{c,
		&someUsers,
		&sessions,
	}

	testUser := UserInputProfile{"someEmail@gmail.com", "someUsername", "someFullName"}
	testResponse := responseUser{200, "someEmail@gmail.com", "someUsername", "someFullName",  "default/default_avatar.png"}

	userExist := someUsers[0]
	response, _ := cc.changeUserProfile(&testUser, &userExist)

	assert.Equal(t, testResponse, response)
}

func TestChangeUserProfileFault(t *testing.T) {
	someUsers := make([]User, 0)
	sessions := make(map[string]uint64, 10)

	someUsers = append(someUsers, User{
		ID:       0,
		Username: "someUsername",
		Email:    "someEmail@gmail.com",
		Password: "somePassword",
		Links:    &UserLinks{},
		Avatar:   "default/default_avatar.png",
	})

	someUsers = append(someUsers, User{
		ID:       1,
		Username: "anotherUsername",
		Email:    "anotherEmail@gmail.com",
		Password: "anotherPassword",
		Links:    &UserLinks{},
		Avatar:   "default/default_avatar.png",
	})

	var c echo.Context
	cc := &Handlers{c,
		&someUsers,
		&sessions,
	}

	testUser := UserInputProfile{"anotherEmail@gmail.com", "someUsername", "someFullName"}
	//testResponse := responseUser{200, "someEmail@gmail.com", "someUsername", "someFullName",  "default/default_avatar.png"}

	messages := make([]Messages, 0)
	messages = append(messages, Messages{"email",  "Такой адрес электронной почты уже зарегистрирован"})

	expectedResponseError := responseError{
		Status:        500,
		Messages: messages,
	}

	userExist := someUsers[0]
	_, err := cc.changeUserProfile(&testUser, &userExist)

	assert.Equal(t, expectedResponseError, err)
}

func TestChangeUserProfileFault2(t *testing.T) {
	someUsers := make([]User, 0)
	sessions := make(map[string]uint64, 10)

	someUsers = append(someUsers, User{
		ID:       0,
		Username: "someUsername",
		Email:    "someEmail@gmail.com",
		Password: "somePassword",
		Links:    &UserLinks{},
		Avatar:   "default/default_avatar.png",
	})

	someUsers = append(someUsers, User{
		ID:       1,
		Username: "anotherUsername",
		Email:    "anotherEmail@gmail.com",
		Password: "anotherPassword",
		Links:    &UserLinks{},
		Avatar:   "default/default_avatar.png",
	})

	var c echo.Context
	cc := &Handlers{c,
		&someUsers,
		&sessions,
	}

	testUser := UserInputProfile{"someEmail@gmail.com", "anotherUsername", "someFullName"}

	messages := make([]Messages, 0)
	messages = append(messages, Messages{"username",  "Такое имя пользователя уже существует"})

	expectedResponseError := responseError{
		Status:        500,
		Messages: messages,
	}

	userExist := someUsers[0]
	_, err := cc.changeUserProfile(&testUser, &userExist)

	assert.Equal(t, expectedResponseError, err)
}

func TestChangeUserAccountsSuccess(t *testing.T)  {
	someUsers := make([]User, 0)
	sessions := make(map[string]uint64, 10)

	someUsers = append(someUsers, User{
		ID:       0,
		Username: "someUsername",
		Email:    "someEmail@gmail.com",
		Password: "somePassword",
		Links:    &UserLinks{},
		Avatar:   "default/default_avatar.png",
	})

	var c echo.Context
	cc := &Handlers{c,
		&someUsers,
		&sessions,
	}

	testUser := UserLinks{
		Telegram:  "@telegram",
		Instagram: "@keith",
		Github:    "github/bab",
		Bitbucket: "bitbucket/ket",
		Vk:        "",
		Facebook:  "facebook",
	}

	expectedResponse := responseUserLinks{
		Status:    200,
		Username:  "someUsername",
		Telegram:  "@telegram",
		Instagram: "@keith",
		Github:    "github/bab",
		Bitbucket: "bitbucket/ket",
		Vk:        "",
		Facebook:  "facebook",
		Avatar:    "default/default_avatar.png",
	}

	userExist := someUsers[0]
	response, _ := cc.changeUserAccounts(&testUser, &userExist)

	assert.Equal(t, expectedResponse, response)
}

func TestUploadAvatarFault(t *testing.T) {
	someUsers := make([]User, 0)
	sessions := make(map[string]uint64, 10)
	var c echo.Context

	cc := &Handlers{c,
		&someUsers,
		&sessions,
	}

	file := multipart.FileHeader{}

	err, _ := cc.uploadAvatar(&file, 0)
	assert.Error(t, err)
}

func TestUserPasswordSuccess(t *testing.T) {
	someUsers := make([]User, 0)
	sessions := make(map[string]uint64, 10)

	someUsers = append(someUsers, User{
		ID:       0,
		Username: "someUsername",
		Email:    "someEmail@gmail.com",
		Password: "somePassword",
		Links:    &UserLinks{},
		Avatar:   "default/default_avatar.png",
	})

	var c echo.Context
	cc := &Handlers{c,
		&someUsers,
		&sessions,
	}

	testUser := UserInputPassword{"somePassword", "newPassword"}
	testResponse := responseUser{200, "someEmail@gmail.com", "someUsername", "",  "default/default_avatar.png"}

	userExist := someUsers[0]
	response, _ := cc.changeUserPassword(&testUser, &userExist)

	assert.Equal(t, testResponse, response)
}

func TestUserPasswordFault(t *testing.T) {
	someUsers := make([]User, 0)
	sessions := make(map[string]uint64, 10)

	someUsers = append(someUsers, User{
		ID:       0,
		Username: "someUsername",
		Email:    "someEmail@gmail.com",
		Password: "somePassword",
		Links:    &UserLinks{},
		Avatar:   "default/default_avatar.png",
	})

	var c echo.Context
	cc := &Handlers{c,
		&someUsers,
		&sessions,
	}

	testUser := UserInputPassword{"wrongPassword", "newPassword"}

	messages := []Messages{{"oldPassword",  "Неверный пароль"}}

	testResponse := responseError{
		OriginalError: nil,
		Status:        500,
		Messages:      messages,
	}

	userExist := someUsers[0]
	_, err := cc.changeUserPassword(&testUser, &userExist)

	assert.Equal(t, testResponse, err)
}

