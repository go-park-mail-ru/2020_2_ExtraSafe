package main

import (
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateUser(t *testing.T) {
	//func (h *Handlers) createUser(userInput UserInputReg) (responseUser, uint64, error) {
	someUsers := make([]User, 0)
	sessions := make(map[string]uint64, 10)

	var c echo.Context
	cc := &Handlers{c,
		&someUsers,
		&sessions,
	}

	testUser := UserInputReg{"someEmail@gmail.com", "someUsername", "somePassword"}
	_, _, err := cc.createUser(testUser)
	assert.Equal(t, err, nil)
}


func TestCheckUser(t *testing.T) {
	someUsers := make([]User, 0)
	sessions := make(map[string]uint64, 10)

	var c echo.Context
	cc := &Handlers{c,
		&someUsers,
		&sessions,
	}

	testUser := UserInputLogin{"someEmail@gmail.com", "somePassword"}

	_, _, err := cc.checkUser(testUser)
	assert.Error(t, err)
}


