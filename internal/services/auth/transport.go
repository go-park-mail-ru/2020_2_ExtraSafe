package auth

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"github.com/labstack/echo"
)

type Transport interface {
	AuthRead(c echo.Context) (request models.UserInput, err error)
	LoginRead(c echo.Context) (request models.UserInputLogin, err error)
	RegRead(c echo.Context) (request models.UserInputReg, err error)

	AuthWrite(user models.UserBoardsOutside, token string) (response models.ResponseUserAuth, err error)
	LoginWrite(token string) (response models.ResponseToken, err error)
	RegWrite() (response models.ResponseStatus, err error)
}

type transport struct {
}

func NewTransport() Transport {
	return &transport{}
}

func (t transport) AuthRead(c echo.Context) (request models.UserInput, err error)  {
	userInput := new(models.UserInput)
	userInput.ID = c.Get("userId").(int64)
	return *userInput, nil
}

func (t transport) RegRead(c echo.Context) (request models.UserInputReg, err error)  {
	userInput := new(models.UserInputReg)

	if err := c.Bind(userInput); err != nil {
		return models.UserInputReg{}, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "RegRead"}
	}
	return *userInput, nil
}

func (t transport) LoginRead(c echo.Context) (request models.UserInputLogin, err error)  {
	userInput := new(models.UserInputLogin)

	if err := c.Bind(userInput); err != nil {
		return models.UserInputLogin{}, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "LoginRead"}
	}
	return *userInput, nil
}

func (t transport)AuthWrite(user models.UserBoardsOutside, token string) (response models.ResponseUserAuth, err error)  {
	response.Status = 200
	response.Token = token
	response.Email = user.Email
	response.Username = user.Username
	response.FullName = user.FullName
	response.Avatar = user.Avatar
	response.Boards = user.Boards
	return response, nil
}

func (t transport)LoginWrite(token string) (response models.ResponseToken, err error)  {
	response.Status = 200
	response.Token = token
	return response, nil
}

func (t transport) RegWrite() (response models.ResponseStatus, err error)  {
	response.Status = 200
	return response, nil
}