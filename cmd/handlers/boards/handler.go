package boardsHandler

import (
	"github.com/labstack/echo"
	"net/http"
)

type Handler interface {
	BoardCreate(c echo.Context) error
	Board(c echo.Context) error
	BoardChange(c echo.Context) error
	BoardDelete(c echo.Context) error

	ColumnCreate(c echo.Context) error
	ColumnChange(c echo.Context) error
	ColumnDelete(c echo.Context) error

	TaskCreate(c echo.Context) error
	TaskChange(c echo.Context) error
	TaskDelete(c echo.Context) error
}

type handler struct {
	boardsService boardsService
	boardsTransport boardsTransport
	errorWorker errorWorker
}

func NewHandler(boardsService boardsService, boardsTransport boardsTransport, errorWorker errorWorker) *handler {
	return &handler{
		boardsService: boardsService,
		boardsTransport: boardsTransport,
		errorWorker: errorWorker,
	}
}

func (h *handler) BoardCreate(c echo.Context) error {
	userInput, err := h.boardsTransport.BoardChangeRead(c)
	if err != nil {
		return h.errorWorker.TransportError(c)
	}

	board, err := h.boardsService.CreateBoard(userInput)
	if err != nil {
		return h.errorWorker.RespError(c, err)
	}

	response, err := h.boardsTransport.BoardWrite(board)
	if err != nil {
		return h.errorWorker.TransportError(c)
	}

	return c.JSON(http.StatusOK, response)
}

func (h *handler) Board(c echo.Context) error {
	userInput, err := h.boardsTransport.BoardRead(c)
	if err != nil {
		return h.errorWorker.TransportError(c)
	}

	board, err := h.boardsService.GetBoard(userInput)
	if err != nil {
		return h.errorWorker.RespError(c, err)
	}

	response, err := h.boardsTransport.BoardWrite(board)
	if err != nil {
		return h.errorWorker.TransportError(c)
	}

	return c.JSON(http.StatusOK, response)
}

func (h *handler) BoardChange(c echo.Context) error {
	userInput, err := h.boardsTransport.BoardChangeRead(c)
	if err != nil {
		return h.errorWorker.TransportError(c)
	}

	board, err := h.boardsService.ChangeBoard(userInput)
	if err != nil {
		return h.errorWorker.RespError(c, err)
	}

	response, err := h.boardsTransport.BoardWrite(board)
	if err != nil {
		return h.errorWorker.TransportError(c)
	}

	return c.JSON(http.StatusOK, response)
}

func (h *handler) BoardDelete(c echo.Context) error {
	userInput, err := h.boardsTransport.BoardRead(c)
	if err != nil {
		return h.errorWorker.TransportError(c)
	}

	err = h.boardsService.DeleteBoard(userInput)
	if err != nil {
		return h.errorWorker.RespError(c, err)
	}

	return c.NoContent(http.StatusOK)
}

func (h *handler) ColumnCreate(c echo.Context) error {
	return nil
}

func (h *handler) ColumnChange(c echo.Context) error {
	return nil
}

func (h *handler) ColumnDelete(c echo.Context) error {
	return nil
}

func (h *handler) TaskCreate(c echo.Context) error {
	return nil
}

func (h *handler) TaskChange(c echo.Context) error {
	return nil
}

func (h *handler) TaskDelete(c echo.Context) error {
	return nil
}