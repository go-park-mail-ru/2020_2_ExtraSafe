package boards

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"github.com/labstack/echo"
)

type Transport interface {
	BoardRead(c echo.Context) (request models.BoardInput, err error)
	BoardChangeRead(c echo.Context) (request models.BoardChangeInput, err error)

	BoardWrite(board models.BoardOutside) (response models.ResponseBoard, err error)
}

type transport struct {
}

func NewTransport() Transport {
	return &transport{}
}

func (t transport) BoardRead(c echo.Context) (request models.BoardInput, err error) {
	userInput := new(models.BoardInput)

	if err := c.Bind(userInput); err != nil {
		return models.BoardInput{}, err
	}
	return *userInput, nil
}

func (t transport) BoardChangeRead(c echo.Context) (request models.BoardChangeInput, err error) {
	userInput := new(models.BoardChangeInput)

	if err := c.Bind(userInput); err != nil {
		return models.BoardChangeInput{}, err
	}
	return *userInput, nil
}

func (t transport) BoardWrite(board models.BoardOutside) (response models.ResponseBoard, err error) {
	return response, err
}