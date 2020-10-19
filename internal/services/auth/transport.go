package auth

import (
	"../../../internal/models"
	"github.com/labstack/echo"
)

type Transport interface {
	AuthRead(c echo.Context) (request models.UserInput, err error)
	LoginRead(c echo.Context) (request models.UserInputLogin, err error)
	RegRead(c echo.Context) (request models.UserInputReg, err error)

	AuthWrite(user models.User) (response models.ResponseUser, err error)
}

type transport struct {
}

func NewTransport() Transport {
	return &transport{}
}

func (t transport)AuthRead(c echo.Context) (request models.UserInput, err error)  {
	cc := c.(*models.CustomContext)
	userInput := new(models.UserInput)
	userInput.ID = cc.UserId
	return *userInput, nil
}

func (t transport)RegRead(c echo.Context) (request models.UserInputReg, err error)  {
	userInput := new(models.UserInputReg)

	if err := c.Bind(userInput); err != nil {
		return models.UserInputReg{}, err
	}
	return *userInput, nil
}

func (t transport)LoginRead(c echo.Context) (request models.UserInputLogin, err error)  {
	userInput := new(models.UserInputLogin)

	if err := c.Bind(userInput); err != nil {
		return models.UserInputLogin{}, err
	}
	return *userInput, nil
}

func (t transport)AuthWrite(user models.User) (response models.ResponseUser, err error)  {
	response.Status = 200
	response.Email = user.Email
	response.Username = user.Username
	response.FullName = user.FullName
	response.Avatar = user.Avatar
	return response, err
}