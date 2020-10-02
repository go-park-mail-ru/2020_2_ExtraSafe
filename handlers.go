package main

import (
	"errors"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func urls(e *echo.Echo) {
	e.Any("/", root)
	e.POST("/login/", login)
	e.POST("/reg/", registration)
}

func login(c echo.Context) error {
	cc := c.(*Handlers)

	//TODO убрать отсюда проверку куки
	response, err := cc.checkUserAuthorized(c)
	if err == nil {
		return c.JSON(http.StatusOK, response)
	}

	userInput := new(UserInputLogin)
	if err := c.Bind(userInput); err != nil {
		return err
	}

	var userID uint64
	response, userID, err = cc.checkUser(*userInput)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, response)
	}

	setCookie(c, userID)
	return c.JSON(http.StatusOK, response)
}

func registration(c echo.Context) error {
	cc := c.(*Handlers)

	//TODO убрать отсюда проверку куки
	response, err := cc.checkUserAuthorized(c)
	if err == nil {
		return c.JSON(http.StatusOK, response)
	}

	userInput := new(UserInputReg)
	if err := c.Bind(userInput); err != nil {
		return err
	}

	var userID uint64
	response, userID, err = cc.createUser(*userInput)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, response)
	}

	setCookie(c, userID)
	return c.JSON(http.StatusOK, response)
}

func root(c echo.Context) error {
	cc := c.(*Handlers)

	response, err := cc.checkUserAuthorized(c)
	if err == nil {
		return c.JSON(http.StatusOK, response)
	}
	return c.JSON(http.StatusTeapot, response)
}

func (h *Handlers) createUser(userInput UserInputReg) (responseUser, uint64, error) {
	for _, user := range *h.users {
		if userInput.Email == user.Email {
			fmt.Println("Email already exist ")
			return responseUser{}, 0, errors.New("Email already exist ")
		}
	}

	for _, user := range *h.users {
		if userInput.Nickname == user.Nickname {
			fmt.Println("Nickname already exist ")
			return responseUser{}, 0, errors.New("Nickname already exist ")
		}
	}

	h.mu.Lock()

	var id uint64 = 0
	if len(*h.users) > 0 {
		id = (*h.users)[len(*h.users)-1].ID + 1
	}

	newUser := User{
		ID:       id,
		Nickname: userInput.Nickname,
		Email:    userInput.Email,
		Password: userInput.Password,
	}
	*h.users = append(*h.users, newUser)
	h.mu.Unlock()

	response := new(responseUser)
	response.WriteResponse(newUser)

	return *response, id, nil
}

func (response *responseUser) WriteResponse(user User) {
	response.Status = 200
	response.Email = user.Email
	response.Nickname = user.Nickname
}

func (h *Handlers) checkUser(userInput UserInputLogin) (responseUser, uint64, error) {
	response := new(responseUser)
	for _, user := range *h.users {
		if userInput.Email == user.Email && userInput.Password == user.Password {
			response.WriteResponse(user)
			return *response, user.ID, nil
		}
	}
	return responseUser{}, 0, errors.New("No such user ")
}
