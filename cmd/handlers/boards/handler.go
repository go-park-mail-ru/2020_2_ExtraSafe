package boardsHandler

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"github.com/labstack/echo"
	"net/http"
)

type Handler interface {
	BoardCreate(c echo.Context) error
	Board(c echo.Context) error
	BoardChange(c echo.Context) error
	BoardDelete(c echo.Context) error

	CardCreate(c echo.Context) error
	Card(c echo.Context) error
	CardChange(c echo.Context) error
	CardDelete(c echo.Context) error
	CardOrder(c echo.Context) error

	TaskCreate(c echo.Context) error
	Task(c echo.Context) error
	TaskChange(c echo.Context) error
	TaskDelete(c echo.Context) error
	TaskOrder(c echo.Context) error
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
		if err := h.errorWorker.TransportError(c); err != nil {
			return err
		}
		return err
	}

	board, err := h.boardsService.CreateBoard(userInput)
	if err != nil {
		if err := h.errorWorker.RespError(c, err); err != nil {
			return err
		}
		return err
	}

	response, err := h.boardsTransport.BoardShortWrite(board)
	if err != nil {
		if err := h.errorWorker.TransportError(c); err != nil {
			return err
		}
		return err
	}

	return c.JSON(http.StatusOK, response)
}

func (h *handler) Board(c echo.Context) error {
	userInput, err := h.boardsTransport.BoardRead(c)
	if err != nil {
		if err := h.errorWorker.TransportError(c); err != nil {
			return err
		}
		return err
	}

	board, err := h.boardsService.GetBoard(userInput)
	if err != nil {
		if err := h.errorWorker.RespError(c, err); err != nil {
			return err
		}
		return err
	}

	response, err := h.boardsTransport.BoardWrite(board)
	if err != nil {
		if err := h.errorWorker.TransportError(c); err != nil {
			return err
		}
		return err
	}

	return c.JSON(http.StatusOK, response)
}

func (h *handler) BoardChange(c echo.Context) error {
	userInput, err := h.boardsTransport.BoardChangeRead(c)
	if err != nil {
		if err := h.errorWorker.TransportError(c); err != nil {
			return err
		}
		return err
	}

	board, err := h.boardsService.ChangeBoard(userInput)
	if err != nil {
		if err := h.errorWorker.RespError(c, err); err != nil {
			return err
		}
		return err
	}

	response, err := h.boardsTransport.BoardWrite(board)
	if err != nil {
		if err := h.errorWorker.TransportError(c); err != nil {
			return err
		}
		return err
	}

	return c.JSON(http.StatusOK, response)
}

func (h *handler) BoardDelete(c echo.Context) error {
	userInput, err := h.boardsTransport.BoardRead(c)
	if err != nil {
		if err := h.errorWorker.TransportError(c); err != nil {
			return err
		}
		return err
	}

	err = h.boardsService.DeleteBoard(userInput)
	if err != nil {
		if err := h.errorWorker.RespError(c, err); err != nil {
			return err
		}
		return err
	}

	return c.JSON(http.StatusOK, models.ResponseStatus{Status: 200})
}

func (h *handler) CardCreate(c echo.Context) error {
	userInput, err := h.boardsTransport.CardChangeRead(c)
	if err != nil {
		if err := h.errorWorker.TransportError(c); err != nil {
			return err
		}
		return err
	}

	card, err := h.boardsService.CreateCard(userInput)
	if err != nil {
		if err := h.errorWorker.RespError(c, err); err != nil {
			return err
		}
		return err
	}

	response, err := h.boardsTransport.CardWrite(card)
	if err != nil {
		if err := h.errorWorker.TransportError(c); err != nil {
			return err
		}
		return err
	}

	return c.JSON(http.StatusOK, response)
}

func (h *handler) Card(c echo.Context) error {
	userInput, err := h.boardsTransport.CardChangeRead(c)
	if err != nil {
		if err := h.errorWorker.TransportError(c); err != nil {
			return err
		}
		return err
	}

	card, err := h.boardsService.GetCard(userInput)
	if err != nil {
		if err := h.errorWorker.RespError(c, err); err != nil {
			return err
		}
		return err
	}

	response, err := h.boardsTransport.CardWrite(card)
	if err != nil {
		if err := h.errorWorker.TransportError(c); err != nil {
			return err
		}
		return err
	}

	return c.JSON(http.StatusOK, response)
}

func (h *handler) CardChange(c echo.Context) error {
	userInput, err := h.boardsTransport.CardChangeRead(c)
	if err != nil {
		if err := h.errorWorker.TransportError(c); err != nil {
			return err
		}
		return err
	}

	card, err := h.boardsService.ChangeCard(userInput)
	if err != nil {
		if err := h.errorWorker.RespError(c, err); err != nil {
			return err
		}
		return err
	}

	response, err := h.boardsTransport.CardWrite(card)
	if err != nil {
		if err := h.errorWorker.TransportError(c); err != nil {
			return err
		}
		return err
	}

	return c.JSON(http.StatusOK, response)
}

func (h *handler) CardDelete(c echo.Context) error {
	userInput, err := h.boardsTransport.CardChangeRead(c)
	if err != nil {
		if err := h.errorWorker.TransportError(c); err != nil {
			return err
		}
		return err
	}

	err = h.boardsService.DeleteCard(userInput)
	if err != nil {
		if err := h.errorWorker.RespError(c, err); err != nil {
			return err
		}
		return err
	}

	return c.JSON(http.StatusOK, models.ResponseStatus{Status: 200})
}

func (h *handler) CardOrder(c echo.Context) error {
	userInput, err := h.boardsTransport.CardOrderRead(c)
	if err != nil {
		if err := h.errorWorker.TransportError(c); err != nil {
			return err
		}
		return err
	}

	err = h.boardsService.CardOrderChange(userInput)
	if err != nil {
		if err := h.errorWorker.RespError(c, err); err != nil {
			return err
		}
		return err
	}

	return c.NoContent(http.StatusOK)
}

func (h *handler) TaskCreate(c echo.Context) error {
	userInput, err := h.boardsTransport.TaskChangeRead(c)
	if err != nil {
		if err := h.errorWorker.TransportError(c); err != nil {
			return err
		}
		return err
	}

	task, err := h.boardsService.CreateTask(userInput)
	if err != nil {
		if err := h.errorWorker.RespError(c, err); err != nil {
			return err
		}
		return err
	}

	response, err := h.boardsTransport.TaskWrite(task)
	if err != nil {
		if err := h.errorWorker.TransportError(c); err != nil {
			return err
		}
		return err
	}

	return c.JSON(http.StatusOK, response)
}

func (h *handler) Task(c echo.Context) error {
	userInput, err := h.boardsTransport.TaskChangeRead(c)
	if err != nil {
		if err := h.errorWorker.TransportError(c); err != nil {
			return err
		}
		return err
	}

	task, err := h.boardsService.GetTask(userInput)
	if err != nil {
		if err := h.errorWorker.RespError(c, err); err != nil {
			return err
		}
		return err
	}

	response, err := h.boardsTransport.TaskWrite(task)
	if err != nil {
		if err := h.errorWorker.TransportError(c); err != nil {
			return err
		}
		return err
	}

	return c.JSON(http.StatusOK, response)
}

func (h *handler) TaskChange(c echo.Context) error {
	userInput, err := h.boardsTransport.TaskChangeRead(c)
	if err != nil {
		if err := h.errorWorker.TransportError(c); err != nil {
			return err
		}
		return err
	}

	task, err := h.boardsService.ChangeTask(userInput)
	if err != nil {
		if err := h.errorWorker.RespError(c, err); err != nil {
			return err
		}
		return err
	}

	response, err := h.boardsTransport.TaskWrite(task)
	if err != nil {
		if err := h.errorWorker.TransportError(c); err != nil {
			return err
		}
		return err
	}

	return c.JSON(http.StatusOK, response)
}

func (h *handler) TaskDelete(c echo.Context) error {
	userInput, err := h.boardsTransport.TaskChangeRead(c)
	if err != nil {
		if err := h.errorWorker.TransportError(c); err != nil {
			return err
		}
		return err
	}

	err = h.boardsService.DeleteTask(userInput)
	if err != nil {
		if err := h.errorWorker.RespError(c, err); err != nil {
			return err
		}
		return err
	}

	return c.JSON(http.StatusOK, models.ResponseStatus{Status: 200})
}

func (h *handler) TaskOrder(c echo.Context) error {
	userInput, err := h.boardsTransport.TasksOrderRead(c)
	if err != nil {
		if err := h.errorWorker.TransportError(c); err != nil {
			return err
		}
		return err
	}

	err = h.boardsService.TasksOrderChange(userInput)
	if err != nil {
		if err := h.errorWorker.RespError(c, err); err != nil {
			return err
		}
		return err
	}

	return c.NoContent(http.StatusOK)
}