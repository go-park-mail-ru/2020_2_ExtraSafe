package main

import (
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
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

//func (h *Handlers) checkUserAuthorized(c echo.Context) (responseUser, error) {
//func (h *Handlers) changeUserProfile(userInput *UserInputProfile, userExist *User) (responseUser, error) {
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
	response, err := cc.changeUserProfile(&testUser, &userExist)

	assert.Equal(t, nil, err)
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
		ID:       0,
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

	userExist := someUsers[0]
	_, err := cc.changeUserProfile(&testUser, &userExist)

	assert.Error(t, err)
	//assert.Equal(t, testResponse, response)
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
		ID:       0,
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
	//assert.Error(t, err)
	//assert.Equal(t, testResponse, response)
}