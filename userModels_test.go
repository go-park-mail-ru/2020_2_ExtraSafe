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

func TestCreateUserFault(t *testing.T) {
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
	sessions := make(map[string]uint64, 10)

	var c echo.Context
	cc := &Handlers{c,
		&someUsers,
		&sessions,
	}

	testUser := UserInputLogin{"someEmail@gmail.com", "somePassword"}

	response, _, err := cc.checkUser(testUser)

	assert.Error(t, err)
	assert.Equal(t, responseUser{}, response)
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

	testUser := UserInputLogin{"someEmail@gmail.com", "somePassword"}
	testResponse := responseUser{Email: "someEmail@gmail.com"}
	response, _, err := cc.checkUser(testUser)

	assert.Equal(t, nil, err)
	assert.Equal(t, testResponse.Email, response.Email)
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

