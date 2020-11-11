package boards

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"github.com/labstack/echo"
	"strconv"
)

type Transport interface {
	BoardRead(c echo.Context) (request models.BoardInput, err error)
	BoardChangeRead(c echo.Context) (request models.BoardChangeInput, err error)
	BoardWrite(board models.BoardOutside) (response models.ResponseBoard, err error)

	CardChangeRead(c echo.Context) (request models.CardInput, err error)
	CardWrite(card models.CardOutside) (response models.ResponseCard, err error)

	TaskChangeRead(c echo.Context) (request models.TaskInput, err error)
	TaskWrite(card models.TaskOutside) (response models.ResponseTask, err error)
}

type transport struct {
}

func NewTransport() Transport {
	return &transport{}
}

func (t transport) BoardRead(c echo.Context) (request models.BoardInput, err error) {
	userInput := new(models.BoardInput)

	boardID := c.Param("ID")
	userInput.BoardID, err = strconv.ParseUint(boardID, 10, 64)
	if err != nil {
		return models.BoardInput{}, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "BoardRead"}
	}

	userInput.UserID = c.Get("userId").(uint64)

	return *userInput, nil
}

func (t transport) BoardChangeRead(c echo.Context) (request models.BoardChangeInput, err error) {
	userInput := new(models.BoardChangeInput)

	boardID := c.Param("ID")
	userInput.BoardID, err = strconv.ParseUint(boardID, 10, 64)

	if err := c.Bind(userInput); err != nil {
		return models.BoardChangeInput{}, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "BoardChangeRead"}
	}

	userInput.UserID = c.Get("userId").(uint64)

	return *userInput, nil
}

func (t transport) BoardWrite(board models.BoardOutside) (response models.ResponseBoard, err error) {
	response.BoardID = board.BoardID
	response.Admin = board.Admin
	response.Theme = board.Theme
	response.Star = board.Star
	response.Users = board.Users
	response.Name = board.Name
	response.Cards = board.Cards
	response.Status = 200
	return response, err
}

func (t transport) CardChangeRead(c echo.Context) (request models.CardInput, err error) {
	userInput := new(models.CardInput)

	if err := c.Bind(userInput); err != nil {
		return models.CardInput{}, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "CardChangeRead"}
	}

	cardID := c.Param("ID")
	userInput.CardID, err = strconv.ParseUint(cardID, 10, 64)
	if err != nil {
		return models.CardInput{}, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "CardChangeRead"}
	}

	userInput.UserID = c.Get("userId").(uint64)

	return *userInput, nil
}

func (t transport) CardWrite(card models.CardOutside) (response models.ResponseCard, err error) {
	response.CardID = card.CardID
	response.Order = card.Order
	response.Name = card.Name
	response.Tasks = card.Tasks
	response.Status = 200
	return response, err
}

func (t transport) TaskChangeRead(c echo.Context) (request models.TaskInput, err error) {
	userInput := new(models.TaskInput)

	if err := c.Bind(userInput); err != nil {
		return models.TaskInput{}, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "TaskChangeRead"}
	}

	taskID := c.Param("ID")
	userInput.TaskID, err = strconv.ParseUint(taskID, 10, 64)
	if err != nil {
		return models.TaskInput{}, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "TaskChangeRead"}
	}

	userInput.UserID = c.Get("userId").(uint64)

	return *userInput, nil
}

func (t transport) TaskWrite(task models.TaskOutside) (response models.ResponseTask, err error) {
	response.TaskID = task.TaskID
	response.Description = task.Description
	response.Order = task.Order
	response.Name = task.Name
	response.Status = 200
	return response, err
}