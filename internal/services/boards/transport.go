package boards

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"github.com/labstack/echo"
	"io/ioutil"
	"strconv"
)

type Transport interface {
	BoardRead(c echo.Context) (request models.BoardInput, err error)
	BoardChangeRead(c echo.Context) (request models.BoardChangeInput, err error)
	BoardMemberRead(c echo.Context) (request models.BoardMemberInput, err error)
	BoardWrite(board models.BoardOutside) (response models.ResponseBoard, err error)
	BoardShortWrite(board models.BoardOutsideShort) (response models.ResponseBoardShort, err error)

	CardChangeRead(c echo.Context) (request models.CardInput, err error)
	CardWrite(card models.CardOutside) (response models.ResponseCard, err error)
	CardShortWrite(card models.CardOutsideShort) (response models.ResponseCard, err error)
	CardOrderRead(c echo.Context) (tasksOrder models.CardsOrderInput, err error)

	TaskChangeRead(c echo.Context) (request models.TaskInput, err error)
	TaskWrite(task models.TaskOutside) (response models.ResponseTask, err error)
	TaskSuperShortWrite(task models.TaskOutsideSuperShort) (response models.ResponseTask, err error)
	TasksOrderRead(c echo.Context) (tasksOrder models.TasksOrderInput, err error)
	TasksUserRead(c echo.Context) (request models.TaskAssignerInput, err error)

	TagChangeRead(c echo.Context) (request models.TagInput, err error)
	TagTaskRead(c echo.Context) (request models.TaskTagInput, err error)
	TagWrite(tag models.TagOutside) (response models.ResponseTag, err error)

	CommentChangeRead(c echo.Context) (request models.CommentInput, err error)
	CommentWrite(comment models.CommentOutside) (response models.ResponseComment, err error)

	ChecklistChangeRead(c echo.Context) (request models.ChecklistInput, err error)
	ChecklistWrite(checklist models.ChecklistOutside) (response models.ResponseChecklist, err error)

	AttachmentAddRead(c echo.Context) (request models.AttachmentInput, err error)
	AttachmentDeleteRead(c echo.Context) (request models.AttachmentInput, err error)
	AttachmentWrite(attachment models.AttachmentOutside) (response models.ResponseAttachment, err error)
}

type transport struct {
}

func NewTransport() Transport {
	return &transport{}
}

func (t transport) BoardRead(c echo.Context) (request models.BoardInput, err error) {
	userInput := new(models.BoardInput)

	boardID := c.Param("ID")
	userInput.BoardID, err = strconv.ParseInt(boardID, 10, 64)
	if err != nil {
		return models.BoardInput{}, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "BoardRead"}
	}

	userInput.UserID = c.Get("userId").(int64)

	return *userInput, nil
}

func (t transport) BoardChangeRead(c echo.Context) (request models.BoardChangeInput, err error) {
	userInput := new(models.BoardChangeInput)

	boardID := c.Param("ID")
	userInput.BoardID, err = strconv.ParseInt(boardID, 10, 64)

	if err := c.Bind(userInput); err != nil {
		return models.BoardChangeInput{}, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "BoardChangeRead"}
	}

	userInput.UserID = c.Get("userId").(int64)

	return *userInput, nil
}

func (t transport) BoardMemberRead(c echo.Context) (request models.BoardMemberInput, err error) {
	userInput := new(models.BoardMemberInput)

	boardID := c.Param("ID")
	userInput.BoardID, err = strconv.ParseInt(boardID, 10, 64)

	if err := c.Bind(userInput); err != nil {
		return models.BoardMemberInput{}, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "BoardChangeRead"}
	}

	userInput.UserID = c.Get("userId").(int64)

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
	response.Tags = board.Tags
	response.Status = 200
	return response, err
}

func (t transport) BoardShortWrite(board models.BoardOutsideShort) (response models.ResponseBoardShort, err error) {
	response.BoardID = board.BoardID
	response.Theme = board.Theme
	response.Star = board.Star
	response.Name = board.Name
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
	userInput.CardID, err = strconv.ParseInt(cardID, 10, 64)
	if err != nil {
		return models.CardInput{}, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "CardChangeRead"}
	}

	userInput.UserID = c.Get("userId").(int64)

	return *userInput, nil
}

func (t transport) CardOrderRead(c echo.Context) (tasksOrder models.CardsOrderInput, err error) {
	userInput := new(models.CardsOrderInput)

	if err := c.Bind(userInput); err != nil {
		return models.CardsOrderInput{}, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "CardOrderRead"}
	}

	userInput.UserID = c.Get("userId").(int64)

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

func (t transport) CardShortWrite(card models.CardOutsideShort) (response models.ResponseCard, err error) {
	response.CardID = card.CardID
	response.Name = card.Name
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
	userInput.TaskID, err = strconv.ParseInt(taskID, 10, 64)
	if err != nil {
		return models.TaskInput{}, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "TaskChangeRead"}
	}

	userInput.UserID = c.Get("userId").(int64)

	return *userInput, nil
}

func (t transport) TaskWrite(task models.TaskOutside) (response models.ResponseTask, err error) {
	response.TaskID = task.TaskID
	response.Description = task.Description
	response.Order = task.Order
	response.Name = task.Name
	response.Tags = task.Tags
	response.Comments = task.Comments
	response.Checklists = task.Checklists
	response.Users = task.Users
	response.Attachments = task.Attachments
	response.Status = 200
	return response, err
}

func (t transport) TaskSuperShortWrite(task models.TaskOutsideSuperShort) (response models.ResponseTask, err error) {
	response.TaskID = task.TaskID
	response.Description = task.Description
	response.Name = task.Name
	response.Status = 200
	return response, err
}

func (t transport) TasksOrderRead(c echo.Context) (tasksOrder models.TasksOrderInput, err error) {
	userInput := new(models.TasksOrderInput)

	if err := c.Bind(userInput); err != nil {
		return models.TasksOrderInput{}, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "TasksOrderRead"}
	}

	userInput.UserID = c.Get("userId").(int64)

	return *userInput, nil
}

func (t transport) TasksUserRead(c echo.Context) (request models.TaskAssignerInput, err error) {
	userInput := new(models.TaskAssignerInput)

	if err := c.Bind(userInput); err != nil {
		return models.TaskAssignerInput{}, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "TasksUserRead"}
	}

	userInput.UserID = c.Get("userId").(int64)

	return *userInput, nil
}

func (t transport) TagChangeRead(c echo.Context) (request models.TagInput, err error) {
	userInput := new(models.TagInput)

	if err := c.Bind(userInput); err != nil {
		return models.TagInput{}, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "TagChangeRead"}
	}

	userInput.UserID = c.Get("userId").(int64)

	return *userInput, nil
}

func (t transport) TagTaskRead(c echo.Context) (request models.TaskTagInput, err error) {
	userInput := new(models.TaskTagInput)

	if err := c.Bind(userInput); err != nil {
		return models.TaskTagInput{}, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "TagTaskRead"}
	}

	userInput.UserID = c.Get("userId").(int64)

	return *userInput, nil
}

func (t transport) TagWrite(tag models.TagOutside) (response models.ResponseTag, err error) {
	response.TagID = tag.TagID
	response.Color = tag.Color
	response.TagName = tag.Name
	response.Status = 200

	return response, nil
}

func (t transport) CommentChangeRead(c echo.Context) (request models.CommentInput, err error) {
	userInput := new(models.CommentInput)

	if err := c.Bind(userInput); err != nil {
		return models.CommentInput{}, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "CommentChangeRead"}
	}

	userInput.UserID = c.Get("userId").(int64)

	return *userInput, nil
}

func (t transport) CommentWrite(comment models.CommentOutside) (response models.ResponseComment, err error) {
	response.CommentID = comment.CommentID
	response.User = comment.User
	response.Message = comment.Message
	response.Order = comment.Order
	response.Status = 200

	return response, nil
}

func (t transport) ChecklistChangeRead(c echo.Context) (request models.ChecklistInput, err error) {
	userInput := new(models.ChecklistInput)

	if err := c.Bind(userInput); err != nil {
		return models.ChecklistInput{}, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "ChecklistChangeRead"}
	}

	userInput.UserID = c.Get("userId").(int64)

	return *userInput, nil
}

func (t transport) ChecklistWrite(checklist models.ChecklistOutside) (response models.ResponseChecklist, err error) {
	response.ChecklistID = checklist.ChecklistID
	response.Name = checklist.Name
	response.Items = checklist.Items
	response.Status = 200

	return response, nil
}

func (t transport) AttachmentAddRead(c echo.Context) (request models.AttachmentInput, err error) {
	formParams, err := c.FormParams()
	if err != nil {
		return models.AttachmentInput{}, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "AttachmentRead"}
	}

	userInput := new(models.AttachmentInput)
	userInput.Filename = formParams.Get("fileName")
	userInput.TaskID, _ = strconv.ParseInt(formParams.Get("taskID"), 10, 64)

	//TODO какая то фигня
	file, err := c.FormFile("file")
	if err == nil {
		fileContent, _ := file.Open()
		var byteContainer []byte
		byteContainer = make([]byte, file.Size)
		byteContainer, _ = ioutil.ReadAll(fileContent)
		userInput.File = byteContainer
	}

	userInput.UserID = c.Get("userId").(int64)

	return *userInput, nil
}

func (t transport) AttachmentDeleteRead(c echo.Context) (request models.AttachmentInput, err error) {
	userInput := new(models.AttachmentInput)

	if err := c.Bind(userInput); err != nil {
		return models.AttachmentInput{}, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "ChecklistChangeRead"}
	}

	userInput.UserID = c.Get("userId").(int64)

	return *userInput, nil
}

func (t transport) AttachmentWrite(attachment models.AttachmentOutside) (response models.ResponseAttachment, err error) {
	response.Status = 200
	response.Filename = attachment.Filename
	response.Filepath = attachment.Filepath
	response.AttachmentID = attachment.AttachmentID

	return response, nil
}