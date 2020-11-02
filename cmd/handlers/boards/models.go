package boardsHandler

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"github.com/labstack/echo"
)

type boardsService interface {
	CreateBoard(request models.BoardChangeInput) (board models.Board, err error)
	GetBoard(request models.BoardInput) (board models.Board, err error)
	ChangeBoard(request models.BoardChangeInput) (board models.Board, err error)
	DeleteBoard(request models.BoardInput) (err error)

	CreateColumn(request models.ColumnInput) (board models.Column, err error)
	ChangeColumn(request models.ColumnInput) (board models.Column, err error)
	DeleteColumn(request models.ColumnInput) (err error)

	CreateTask(request models.TaskInput) (board models.Tasks, err error)
	ChangeTask(request models.TaskInput) (board models.Tasks, err error)
	DeleteTask(request models.TaskInput) (err error)
}

type boardsTransport interface {
	BoardRead(c echo.Context) (request models.BoardInput, err error)
	BoardChangeRead(c echo.Context) (request models.BoardChangeInput, err error)

	BoardWrite(board models.Board) (response models.ResponseBoard, err error)
}

type errorWorker interface {
	RespError(c echo.Context, serveError error) (err error)
	TransportError(c echo.Context) (err error)
}
