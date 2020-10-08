package tests

import (
	"../sources"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"mime/multipart"
	"testing"
)

func TestCreateUserSuccess(t *testing.T) {
	someUsers := make([]sources.User, 0)
	sessions := make(map[string]uint64, 10)

	var c echo.Context
	cc := &sources.Handlers{Context: c,
		Users:    &someUsers,
		Sessions: &sessions,
	}

	testUser := sources.UserInputReg{Email: "someEmail@gmail.com", Username: "someUsername", Password: "somePassword"}
	_, _, err := cc.CreateUser(testUser)
	assert.Equal(t, nil, err)
}

func TestCreateUserFail(t *testing.T) {
	someUsers := make([]sources.User, 0)
	someUsers = append(someUsers, sources.User{
		ID:       0,
		Username: "someUsername",
		Email:    "someEmail@gmail.com",
		Password: "somePassword",
		Links:    &sources.UserLinks{},
		Avatar:   "default/default_avatar.png",
	})

	sessions := make(map[string]uint64, 10)

	var c echo.Context
	cc := &sources.Handlers{Context: c,
		Users:    &someUsers,
		Sessions: &sessions,
	}

	testUser := sources.UserInputReg{Email: "someEmail@gmail.com", Username: "someUsername", Password: "somePassword"}
	_, _, err := cc.CreateUser(testUser)

	assert.Error(t, err)
}

func TestCheckUserSuccess(t *testing.T) {
	someUsers := make([]sources.User, 0)

	someUsers = append(someUsers, sources.User{
		ID:       0,
		Username: "someUsername",
		Email:    "someEmail@gmail.com",
		Password: "somePassword",
		Links:    &sources.UserLinks{},
		Avatar:   "default/default_avatar.png",
		FullName: "Petr",
	})

	sessions := make(map[string]uint64, 10)

	var c echo.Context
	cc := &sources.Handlers{Context: c,
		Users:    &someUsers,
		Sessions: &sessions,
	}

	testUser := sources.UserInputLogin{Email: "someEmail@gmail.com", Password: "somePassword"}

	expectResponse := sources.ResponseUser{
		Status:   200,
		Email:    "someEmail@gmail.com",
		Username: "someUsername",
		FullName: "Petr",
		Avatar:   "default/default_avatar.png",
	}

	response, _, _ := cc.CheckUser(testUser)

	assert.Equal(t, expectResponse, response)
}

func TestCheckUserFault(t *testing.T) {
	someUsers := make([]sources.User, 0)
	someUsers = append(someUsers, sources.User{
		ID:       0,
		Username: "someUsername",
		Email:    "someEmail@gmail.com",
		Password: "somePassword",
		Links:    &sources.UserLinks{},
		Avatar:   "default/default_avatar.png",
	})

	sessions := make(map[string]uint64, 10)

	var c echo.Context
	cc := &sources.Handlers{Context: c,
		Users:    &someUsers,
		Sessions: &sessions,
	}

	testUser := sources.UserInputLogin{Email: "someEmail@gmail.com", Password: "Password"}

	errorMessage := []sources.Messages{{Message: "Неверная электронная почта или пароль", ErrorName: "password"}}

	testResponse := sources.ResponseError{
		Status:        500,
		Messages:      errorMessage,
	}

	_, _, err := cc.CheckUser(testUser)

	assert.Equal(t, testResponse, err)
}

func TestChangeUserProfileSuccess(t *testing.T) {
	someUsers := make([]sources.User, 0)
	sessions := make(map[string]uint64, 10)

	someUsers = append(someUsers, sources.User{
		ID:       0,
		Username: "someUsername",
		Email:    "someEmail@gmail.com",
		Password: "somePassword",
		Links:    &sources.UserLinks{},
		Avatar:   "default/default_avatar.png",
	})

	var c echo.Context
	cc := &sources.Handlers{Context: c,
		Users:    &someUsers,
		Sessions: &sessions,
	}

	testUser := sources.UserInputProfile{Email: "someEmail@gmail.com", Username: "someUsername", FullName: "someFullName"}
	testResponse := sources.ResponseUser{Status: 200, Email: "someEmail@gmail.com", Username: "someUsername", FullName: "someFullName",  Avatar: "default/default_avatar.png"}

	userExist := someUsers[0]
	response, _ := cc.ChangeUserProfile(&testUser, &userExist)

	assert.Equal(t, testResponse, response)
}

func TestChangeUserProfileFault(t *testing.T) {
	someUsers := make([]sources.User, 0)
	sessions := make(map[string]uint64, 10)

	someUsers = append(someUsers, sources.User{
		ID:       0,
		Username: "someUsername",
		Email:    "someEmail@gmail.com",
		Password: "somePassword",
		Links:    &sources.UserLinks{},
		Avatar:   "default/default_avatar.png",
	})

	someUsers = append(someUsers, sources.User{
		ID:       1,
		Username: "anotherUsername",
		Email:    "anotherEmail@gmail.com",
		Password: "anotherPassword",
		Links:    &sources.UserLinks{},
		Avatar:   "default/default_avatar.png",
	})

	var c echo.Context
	cc := &sources.Handlers{Context: c,
		Users:    &someUsers,
		Sessions: &sessions,
	}

	testUser := sources.UserInputProfile{Email: "anotherEmail@gmail.com", Username: "someUsername", FullName: "someFullName"}
	//testResponse := responseUser{200, "someEmail@gmail.com", "someUsername", "someFullName",  "default/default_avatar.png"}

	messages := make([]sources.Messages, 0)
	messages = append(messages, sources.Messages{ErrorName: "email",  Message: "Такой адрес электронной почты уже зарегистрирован"})

	expectedResponseError := sources.ResponseError{
		Status:        500,
		Messages: messages,
	}

	userExist := someUsers[0]
	_, err := cc.ChangeUserProfile(&testUser, &userExist)

	assert.Equal(t, expectedResponseError, err)
}

func TestChangeUserProfileFault2(t *testing.T) {
	someUsers := make([]sources.User, 0)
	sessions := make(map[string]uint64, 10)

	someUsers = append(someUsers, sources.User{
		ID:       0,
		Username: "someUsername",
		Email:    "someEmail@gmail.com",
		Password: "somePassword",
		Links:    &sources.UserLinks{},
		Avatar:   "default/default_avatar.png",
	})

	someUsers = append(someUsers, sources.User{
		ID:       1,
		Username: "anotherUsername",
		Email:    "anotherEmail@gmail.com",
		Password: "anotherPassword",
		Links:    &sources.UserLinks{},
		Avatar:   "default/default_avatar.png",
	})

	var c echo.Context
	cc := &sources.Handlers{Context: c,
		Users:    &someUsers,
		Sessions: &sessions,
	}

	testUser := sources.UserInputProfile{Email: "someEmail@gmail.com", Username: "anotherUsername", FullName: "someFullName"}

	messages := make([]sources.Messages, 0)
	messages = append(messages, sources.Messages{ErrorName: "username",  Message: "Такое имя пользователя уже существует"})

	expectedResponseError := sources.ResponseError{
		Status:        500,
		Messages: messages,
	}

	userExist := someUsers[0]
	_, err := cc.ChangeUserProfile(&testUser, &userExist)

	assert.Equal(t, expectedResponseError, err)
}

func TestChangeUserAccountsSuccess(t *testing.T)  {
	someUsers := make([]sources.User, 0)
	sessions := make(map[string]uint64, 10)

	someUsers = append(someUsers, sources.User{
		ID:       0,
		Username: "someUsername",
		Email:    "someEmail@gmail.com",
		Password: "somePassword",
		Links:    &sources.UserLinks{},
		Avatar:   "default/default_avatar.png",
	})

	var c echo.Context
	cc := &sources.Handlers{Context: c,
		Users:    &someUsers,
		Sessions: &sessions,
	}

	testUser := sources.UserLinks{
		Telegram:  "@telegram",
		Instagram: "@keith",
		Github:    "github/bab",
		Bitbucket: "bitbucket/ket",
		Vk:        "",
		Facebook:  "facebook",
	}

	expectedResponse := sources.ResponseUserLinks{
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
	response, _ := cc.ChangeUserAccounts(&testUser, &userExist)

	assert.Equal(t, expectedResponse, response)
}

func TestUploadAvatarFault(t *testing.T) {
	someUsers := make([]sources.User, 0)
	sessions := make(map[string]uint64, 10)
	var c echo.Context

	cc := &sources.Handlers{Context: c,
		Users:    &someUsers,
		Sessions: &sessions,
	}

	file := multipart.FileHeader{}

	err, _ := cc.UploadAvatar(&file, 0)
	assert.Error(t, err)
}

func TestUserPasswordSuccess(t *testing.T) {
	someUsers := make([]sources.User, 0)
	sessions := make(map[string]uint64, 10)

	someUsers = append(someUsers, sources.User{
		ID:       0,
		Username: "someUsername",
		Email:    "someEmail@gmail.com",
		Password: "somePassword",
		Links:    &sources.UserLinks{},
		Avatar:   "default/default_avatar.png",
	})

	var c echo.Context
	cc := &sources.Handlers{Context: c,
		Users:    &someUsers,
		Sessions: &sessions,
	}

	testUser := sources.UserInputPassword{OldPassword: "somePassword", Password: "newPassword"}
	testResponse := sources.ResponseUser{Status: 200, Email: "someEmail@gmail.com", Username: "someUsername",  Avatar: "default/default_avatar.png"}

	userExist := someUsers[0]
	response, _ := cc.ChangeUserPassword(&testUser, &userExist)

	assert.Equal(t, testResponse, response)
}

func TestUserPasswordFault(t *testing.T) {
	someUsers := make([]sources.User, 0)
	sessions := make(map[string]uint64, 10)

	someUsers = append(someUsers, sources.User{
		ID:       0,
		Username: "someUsername",
		Email:    "someEmail@gmail.com",
		Password: "somePassword",
		Links:    &sources.UserLinks{},
		Avatar:   "default/default_avatar.png",
	})

	var c echo.Context
	cc := &sources.Handlers{Context: c,
		Users:    &someUsers,
		Sessions: &sessions,
	}

	testUser := sources.UserInputPassword{OldPassword: "wrongPassword", Password: "newPassword"}

	messages := []sources.Messages{{"oldPassword",  "Неверный пароль"}}

	testResponse := sources.ResponseError{
		OriginalError: nil,
		Status:        500,
		Messages:      messages,
	}

	userExist := someUsers[0]
	_, err := cc.ChangeUserPassword(&testUser, &userExist)

	assert.Equal(t, testResponse, err)
}

