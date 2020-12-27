package profile

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"github.com/labstack/echo"
	"io/ioutil"
)

type Transport interface {
	ProfileRead(c echo.Context) (request models.UserInput, err error)

	ProfileChangeRead(c echo.Context) (request models.UserInputProfile, err error)
	PasswordChangeRead(c echo.Context) (request models.UserInputPassword, err error)

	BoardsWrite(boards models.BoardsOutside) (response models.ResponseBoards, err error)
	ProfileWrite(user models.UserOutside) (response models.ResponseUser, err error)
}

type transport struct {
}

func NewTransport() Transport {
	return &transport{}
}

func (t transport) ProfileRead(c echo.Context) (request models.UserInput, err error) {
	userInput := new(models.UserInput)
	userInput.ID = c.Get("userId").(int64)
	return *userInput, nil
}

func (t transport) ProfileChangeRead(c echo.Context) (request models.UserInputProfile, err error) {
	formParams, err := c.FormParams()
	if err != nil {
		return models.UserInputProfile{}, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "ProfileChangeRead"}
	}

	userInput := new(models.UserInputProfile)
	userInput.Username = formParams.Get("username")
	userInput.Email = formParams.Get("email")
	userInput.FullName = formParams.Get("fullName")

	file, err := c.FormFile("avatar")
	if err == nil {
		fileContent, _ := file.Open()
		var byteContainer []byte
		byteContainer = make([]byte, file.Size)
		byteContainer, _ = ioutil.ReadAll(fileContent)
		userInput.Avatar = byteContainer
	}

	userInput.ID = c.Get("userId").(int64)

	return *userInput, nil
}

func (t transport) PasswordChangeRead(c echo.Context) (request models.UserInputPassword, err error) {
	userInput := new(models.UserInputPassword)

	if err := c.Bind(userInput); err != nil {
		return models.UserInputPassword{}, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "PasswordChangeRead"}
	}

	userInput.ID = c.Get("userId").(int64)

	return *userInput, nil
}

func (t transport) BoardsWrite(boards models.BoardsOutside) (response models.ResponseBoards, err error) {
	response.Status = 200
	response.Boards = boards.Boards
	response.Templates = boards.Templates
	return response, nil
}

func (t transport) ProfileWrite(user models.UserOutside) (response models.ResponseUser, err error) {
	response.Status = 200
	response.Email = user.Email
	response.Username = user.Username
	response.FullName = user.FullName
	response.Avatar = user.Avatar
	return response, err
}
