package boardsHandler

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/services/boards"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/tools/errorWorker"
	"github.com/labstack/echo"
	"net/http"
)

type Handler interface {
	BoardCreate(c echo.Context) error
	Board(c echo.Context) error
	BoardChange(c echo.Context) error
	BoardDelete(c echo.Context) error
	BoardAddMember(c echo.Context) error
	BoardRemoveMember(c echo.Context) error

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
	TaskUserAdd(c echo.Context) error
	TaskUserRemove(c echo.Context) error

	TagCreate(c echo.Context) error
	TagChange(c echo.Context) error
	TagDelete(c echo.Context) error
	TagAdd(c echo.Context) error
	TagRemove(c echo.Context) error

	CommentCreate(c echo.Context) error
	CommentChange(c echo.Context) error
	CommentDelete(c echo.Context) error

	ChecklistCreate(c echo.Context) error
	ChecklistChange(c echo.Context) error
	ChecklistDelete(c echo.Context) error

	AttachmentCreate(c echo.Context) error
	AttachmentDelete(c echo.Context) error

	SharedURL(c echo.Context) error
}

type handler struct {
	boardsService boards.ServiceBoard
	boardsTransport boards.Transport
	errorWorker errorWorker.ErrorWorker
}

func NewHandler(boardsService boards.ServiceBoard, boardsTransport boards.Transport, errorWorker errorWorker.ErrorWorker) *handler {
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

	response, err := h.boardsTransport.BoardShortWrite(board)

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

	response, err := h.boardsTransport.BoardShortWrite(board)

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

	return c.JSON(http.StatusOK, models.ResponseStatus{Status: 200})
}

func (h *handler) BoardAddMember(c echo.Context) error {
	userInput, err := h.boardsTransport.BoardMemberRead(c)
	if err != nil {
		return h.errorWorker.TransportError(c)
	}

	user, err := h.boardsService.AddMember(userInput)
	if err != nil {
		return h.errorWorker.RespError(c, err)
	}

	response, err := h.boardsTransport.UserShortWrite(user)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) BoardRemoveMember(c echo.Context) error {
	userInput, err := h.boardsTransport.BoardMemberRead(c)
	if err != nil {
		return h.errorWorker.TransportError(c)
	}

	err = h.boardsService.RemoveMember(userInput)
	if err != nil {
		return h.errorWorker.RespError(c, err)
	}

	return c.JSON(http.StatusOK, models.ResponseStatus{Status: 200})
}

func (h *handler) CardCreate(c echo.Context) error {
	userInput, err := h.boardsTransport.CardChangeRead(c)
	if err != nil {
		return h.errorWorker.TransportError(c)
	}

	card, err := h.boardsService.CreateCard(userInput)
	if err != nil {
		return h.errorWorker.RespError(c, err)
	}

	response, err := h.boardsTransport.CardShortWrite(card)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) Card(c echo.Context) error {
	userInput, err := h.boardsTransport.CardChangeRead(c)
	if err != nil {
		return h.errorWorker.TransportError(c)
	}

	card, err := h.boardsService.GetCard(userInput)
	if err != nil {
		return h.errorWorker.RespError(c, err)
	}

	response, err := h.boardsTransport.CardWrite(card)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) CardChange(c echo.Context) error {
	userInput, err := h.boardsTransport.CardChangeRead(c)
	if err != nil {
		return h.errorWorker.TransportError(c)
	}

	card, err := h.boardsService.ChangeCard(userInput)
	if err != nil {
		return h.errorWorker.RespError(c, err)
	}

	response, err := h.boardsTransport.CardShortWrite(card)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) CardDelete(c echo.Context) error {
	userInput, err := h.boardsTransport.CardChangeRead(c)
	if err != nil {
		return h.errorWorker.TransportError(c)
	}

	err = h.boardsService.DeleteCard(userInput)
	if err != nil {
		return h.errorWorker.RespError(c, err)
	}

	return c.JSON(http.StatusOK, models.ResponseStatus{Status: 200})
}

func (h *handler) CardOrder(c echo.Context) error {
	userInput, err := h.boardsTransport.CardOrderRead(c)
	if err != nil {
		return h.errorWorker.TransportError(c)
	}

	err = h.boardsService.CardOrderChange(userInput)
	if err != nil {
		return h.errorWorker.RespError(c, err)
	}

	return c.JSON(http.StatusOK, models.ResponseStatus{Status: 200})
}

func (h *handler) TaskCreate(c echo.Context) error {
	userInput, err := h.boardsTransport.TaskChangeRead(c)
	if err != nil {
		return h.errorWorker.TransportError(c)
	}

	task, err := h.boardsService.CreateTask(userInput)
	if err != nil {
		return h.errorWorker.RespError(c, err)
	}

	response, err := h.boardsTransport.TaskSuperShortWrite(task)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) Task(c echo.Context) error {
	userInput, err := h.boardsTransport.TaskChangeRead(c)
	if err != nil {
		return h.errorWorker.TransportError(c)
	}

	task, err := h.boardsService.GetTask(userInput)
	if err != nil {
		return h.errorWorker.RespError(c, err)
	}

	response, err := h.boardsTransport.TaskWrite(task)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) TaskChange(c echo.Context) error {
	userInput, err := h.boardsTransport.TaskChangeRead(c)
	if err != nil {
		return h.errorWorker.TransportError(c)
	}

	task, err := h.boardsService.ChangeTask(userInput)
	if err != nil {
		return h.errorWorker.RespError(c, err)
	}

	response, err := h.boardsTransport.TaskSuperShortWrite(task)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) TaskDelete(c echo.Context) error {
	userInput, err := h.boardsTransport.TaskChangeRead(c)
	if err != nil {
		return h.errorWorker.TransportError(c)
	}

	err = h.boardsService.DeleteTask(userInput)
	if err != nil {
		return h.errorWorker.RespError(c, err)
	}

	return c.JSON(http.StatusOK, models.ResponseStatus{Status: 200})
}

func (h *handler) TaskOrder(c echo.Context) error {
	userInput, err := h.boardsTransport.TasksOrderRead(c)
	if err != nil {
		return h.errorWorker.TransportError(c)
	}

	err = h.boardsService.TasksOrderChange(userInput)
	if err != nil {
		return h.errorWorker.RespError(c, err)
	}

	return c.NoContent(http.StatusOK)
}

func (h *handler) TaskUserAdd(c echo.Context) error {
	userInput, err := h.boardsTransport.TasksUserRead(c)
	if err != nil {
		return h.errorWorker.TransportError(c)
	}

	user, err := h.boardsService.AssignUser(userInput)
	if err != nil {
		return h.errorWorker.RespError(c, err)
	}

	response, err := h.boardsTransport.UserShortWrite(user)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) TaskUserRemove(c echo.Context) error {
	userInput, err := h.boardsTransport.TasksUserRead(c)
	if err != nil {
		return h.errorWorker.TransportError(c)
	}

	err = h.boardsService.DismissUser(userInput)
	if err != nil {
		return h.errorWorker.RespError(c, err)
	}

	return c.JSON(http.StatusOK, models.ResponseStatus{Status: 200})
}

func (h *handler) TagCreate(c echo.Context) error {
	userInput, err := h.boardsTransport.TagChangeRead(c)
	if err != nil {
		return h.errorWorker.TransportError(c)
	}

	tag, err := h.boardsService.CreateTag(userInput)
	if err != nil {
		return h.errorWorker.RespError(c, err)
	}

	response, err := h.boardsTransport.TagWrite(tag)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) TagChange(c echo.Context) error {
	userInput, err := h.boardsTransport.TagChangeRead(c)
	if err != nil {
		return h.errorWorker.TransportError(c)
	}

	tag, err := h.boardsService.ChangeTag(userInput)
	if err != nil {
		return h.errorWorker.RespError(c, err)
	}

	response, err := h.boardsTransport.TagWrite(tag)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) TagDelete(c echo.Context) error {
	userInput, err := h.boardsTransport.TagChangeRead(c)
	if err != nil {
		return h.errorWorker.TransportError(c)
	}

	err = h.boardsService.DeleteTag(userInput)
	if err != nil {
		return h.errorWorker.RespError(c, err)
	}

	return c.JSON(http.StatusOK, models.ResponseStatus{Status: 200})
}

func (h *handler) TagAdd(c echo.Context) error {
	userInput, err := h.boardsTransport.TagTaskRead(c)
	if err != nil {
		return h.errorWorker.TransportError(c)
	}

	err = h.boardsService.AddTag(userInput)
	if err != nil {
		return h.errorWorker.RespError(c, err)
	}

	return c.JSON(http.StatusOK, models.ResponseStatus{Status: 200})
}

func (h *handler) TagRemove(c echo.Context) error {
	userInput, err := h.boardsTransport.TagTaskRead(c)
	if err != nil {
		return h.errorWorker.TransportError(c)
	}

	err = h.boardsService.RemoveTag(userInput)
	if err != nil {
		return h.errorWorker.RespError(c, err)
	}

	return c.JSON(http.StatusOK, models.ResponseStatus{Status: 200})
}

func (h *handler) CommentCreate(c echo.Context) error {
	userInput, err := h.boardsTransport.CommentChangeRead(c)
	if err != nil {
		return h.errorWorker.TransportError(c)
	}

	tag, err := h.boardsService.CreateComment(userInput)
	if err != nil {
		return h.errorWorker.RespError(c, err)
	}

	response, err := h.boardsTransport.CommentWrite(tag)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) CommentChange(c echo.Context) error {
	userInput, err := h.boardsTransport.CommentChangeRead(c)
	if err != nil {
		return h.errorWorker.TransportError(c)
	}

	tag, err := h.boardsService.ChangeComment(userInput)
	if err != nil {
		return h.errorWorker.RespError(c, err)
	}

	response, err := h.boardsTransport.CommentWrite(tag)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) CommentDelete(c echo.Context) error {
	userInput, err := h.boardsTransport.CommentChangeRead(c)
	if err != nil {
		return h.errorWorker.TransportError(c)
	}

	err = h.boardsService.DeleteComment(userInput)
	if err != nil {
		return h.errorWorker.RespError(c, err)
	}

	return c.JSON(http.StatusOK, models.ResponseStatus{Status: 200})
}

func (h *handler) ChecklistCreate(c echo.Context) error {
	userInput, err := h.boardsTransport.ChecklistChangeRead(c)
	if err != nil {
		return h.errorWorker.TransportError(c)
	}

	tag, err := h.boardsService.CreateChecklist(userInput)
	if err != nil {
		return h.errorWorker.RespError(c, err)
	}

	response, err := h.boardsTransport.ChecklistWrite(tag)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) ChecklistChange(c echo.Context) error {
	userInput, err := h.boardsTransport.ChecklistChangeRead(c)
	if err != nil {
		return h.errorWorker.TransportError(c)
	}

	tag, err := h.boardsService.ChangeChecklist(userInput)
	if err != nil {
		return h.errorWorker.RespError(c, err)
	}

	response, err := h.boardsTransport.ChecklistWrite(tag)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) ChecklistDelete(c echo.Context) error {
	userInput, err := h.boardsTransport.ChecklistChangeRead(c)
	if err != nil {
		return h.errorWorker.TransportError(c)
	}

	err = h.boardsService.DeleteChecklist(userInput)
	if err != nil {
		return h.errorWorker.RespError(c, err)
	}

	return c.JSON(http.StatusOK, models.ResponseStatus{Status: 200})
}

func (h *handler) AttachmentCreate(c echo.Context) error {
	userInput, err := h.boardsTransport.AttachmentAddRead(c)
	if err != nil {
		return h.errorWorker.TransportError(c)
	}

	tag, err := h.boardsService.CreateAttachment(userInput)
	if err != nil {
		return h.errorWorker.RespError(c, err)
	}

	response, err := h.boardsTransport.AttachmentWrite(tag)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) AttachmentDelete(c echo.Context) error {
	userInput, err := h.boardsTransport.AttachmentDeleteRead(c)
	if err != nil {
		return h.errorWorker.TransportError(c)
	}

	err = h.boardsService.DeleteAttachment(userInput)
	if err != nil {
		return h.errorWorker.RespError(c, err)
	}

	return c.JSON(http.StatusOK, models.ResponseStatus{Status: 200})
}

func (h *handler) SharedURL(c echo.Context) error {
	userInput, err := h.boardsTransport.BoardRead(c)
	if err != nil {
		return h.errorWorker.TransportError(c)
	}

	tag, err := h.boardsService.GetSharedURL(userInput)
	if err != nil {
		return h.errorWorker.RespError(c, err)
	}

	response, err := h.boardsTransport.URLWrite(tag)

	return c.JSON(http.StatusOK, response)
}