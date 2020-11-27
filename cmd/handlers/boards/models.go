package boardsHandler

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"github.com/labstack/echo"
)

type boardsService interface {
	CreateBoard(request models.BoardChangeInput) (board models.BoardOutsideShort, err error)
	GetBoard(request models.BoardInput) (board models.BoardOutside, err error)
	ChangeBoard(request models.BoardChangeInput) (board models.BoardOutside, err error)
	DeleteBoard(request models.BoardInput) (err error)

	CreateCard(request models.CardInput) (card models.CardOutside, err error)
	GetCard(request models.CardInput) (card models.CardOutside, err error)
	ChangeCard(request models.CardInput) (card models.CardOutside, err error)
	DeleteCard(request models.CardInput) (err error)
	CardOrderChange(request models.CardsOrderInput) (err error)

	CreateTask(request models.TaskInput) (task models.TaskInternalShort, err error)
	GetTask(request models.TaskInput) (task models.TaskInternalShort, err error)
	ChangeTask(request models.TaskInput) (task models.TaskInternalShort, err error)
	DeleteTask(request models.TaskInput) (err error)
	TasksOrderChange(request models.TasksOrderInput) (err error)
}

type boardsTransport interface {
	BoardRead(c echo.Context) (request models.BoardInput, err error)
	BoardChangeRead(c echo.Context) (request models.BoardChangeInput, err error)
	BoardWrite(board models.BoardOutside) (response models.ResponseBoard, err error)
	BoardShortWrite(board models.BoardOutsideShort) (response models.ResponseBoardShort, err error)

	CardChangeRead(c echo.Context) (request models.CardInput, err error)
	CardWrite(card models.CardOutside) (response models.ResponseCard, err error)
	CardOrderRead(c echo.Context) (tasksOrder models.CardsOrderInput, err error)

	TaskChangeRead(c echo.Context) (request models.TaskInput, err error)
	TaskWrite(card models.TaskInternalShort) (response models.ResponseTask, err error)
	TasksOrderRead(c echo.Context) (tasksOrder models.TasksOrderInput, err error)
}

type errorWorker interface {
	RespError(c echo.Context, serveError error) (err error)
	TransportError(c echo.Context) (err error)
}
