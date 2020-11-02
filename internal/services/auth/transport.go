package auth

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"github.com/labstack/echo"
)

type Transport interface {
	AuthRead(c echo.Context) (request models.UserInput, err error)
	LoginRead(c echo.Context) (request models.UserInputLogin, err error)
	RegRead(c echo.Context) (request models.UserInputReg, err error)

	AuthWrite(user models.UserOutside) (response models.ResponseUserAuth, err error)
	LoginWrite() (response models.ResponseStatus, err error)
	RegWrite() (response models.ResponseStatus, err error)
}

type transport struct {
}

func NewTransport() Transport {
	return &transport{}
}

func (t transport)AuthRead(c echo.Context) (request models.UserInput, err error)  {
	userInput := new(models.UserInput)
	userInput.ID = c.Get("userId").(uint64)
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

func (t transport)AuthWrite(user models.UserOutside) (response models.ResponseUserAuth, err error)  {
	response.Status = 200
	response.Email = user.Email
	response.Username = user.Username
	response.FullName = user.FullName
	response.Avatar = user.Avatar
	response.Telegram = user.Links.Telegram
	response.Instagram = user.Links.Instagram
	response.Github = user.Links.Github
	response.Bitbucket = user.Links.Bitbucket
	response.Vk = user.Links.Vk
	response.Facebook = user.Links.Facebook
	return response, err
}

func (t transport)LoginWrite() (response models.ResponseStatus, err error)  {
	response.Status = 200
	return response, err
}

func (t transport)RegWrite() (response models.ResponseStatus, err error)  {
	response.Status = 200
	return response, err
}