package boards

import (
	"bytes"
	"context"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/tools/errorWorker"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/tools/websocket"
	protoBoard "github.com/go-park-mail-ru/2020_2_ExtraSafe/services/proto/board"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/proto/profile"
	"github.com/labstack/echo"
	"io"
)

//go:generate mockgen -destination=../../../cmd/handlers/mock/mock_boardsService.go -package=mock github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/services/boards ServiceBoard

type ServiceBoard interface {
	CreateBoard(request models.BoardChangeInput) (board models.BoardOutsideShort, err error)
	GetBoard(request models.BoardInput) (board models.BoardOutside, err error)
	WebSocketBoard(request models.BoardInput, c echo.Context) (err error)
	ChangeBoard(request models.BoardChangeInput) (board models.BoardOutsideShort, err error)
	DeleteBoard(request models.BoardInput) (err error)
	AddMember(request models.BoardMemberInput) (user models.UserOutsideShort, err error)
	RemoveMember(request models.BoardMemberInput) (err error)

	CreateCard(request models.CardInput) (card models.CardOutsideShort, err error)
	GetCard(request models.CardInput) (board models.CardOutside, err error)
	ChangeCard(request models.CardInput) (card models.CardOutsideShort, err error)
	DeleteCard(request models.CardInput) (err error)
	CardOrderChange(request models.CardsOrderInput) (err error)

	CreateTask(request models.TaskInput) (task models.TaskOutsideSuperShort, err error)
	GetTask(request models.TaskInput) (board models.TaskOutside, err error)
	ChangeTask(request models.TaskInput) (task models.TaskOutsideSuperShort, err error)
	DeleteTask(request models.TaskInput) (err error)
	TasksOrderChange(request models.TasksOrderInput) (err error)
	AssignUser(request models.TaskAssignerInput) (user models.UserOutsideShort, err error)
	DismissUser(request models.TaskAssignerInput) (err error)

	CreateTag(request models.TagInput) (task models.TagOutside, err error)
	ChangeTag(request models.TagInput) (task models.TagOutside, err error)
	DeleteTag(request models.TagInput) (err error)
	AddTag(request models.TaskTagInput) (err error)
	RemoveTag(request models.TaskTagInput) (err error)

	CreateComment(request models.CommentInput) (task models.CommentOutside, err error)
	ChangeComment(request models.CommentInput) (task models.CommentOutside, err error)
	DeleteComment(request models.CommentInput) (err error)

	CreateChecklist(request models.ChecklistInput) (task models.ChecklistOutside, err error)
	ChangeChecklist(request models.ChecklistInput) (task models.ChecklistOutside, err error)
	DeleteChecklist(request models.ChecklistInput) (err error)

	CreateAttachment(request models.AttachmentInput) (task models.AttachmentOutside, err error)
	DeleteAttachment(request models.AttachmentInput) (err error)

	GetSharedURL(request models.BoardInput) (url string, err error)
	InviteUserToBoard(request models.BoardInviteInput) (board models.BoardOutsideShort, err error)

	CheckBoardPermission(userID int64, boardID int64, ifAdmin bool) (err error)
	CheckCardPermission(userID int64, cardID int64) (boardID int64, err error)
	CheckTaskPermission(userID int64, taskID int64) (boardID int64, err error)
	CheckCommentPermission(userID int64, commentID int64, ifAdmin bool) (boardID int64, err error)
}

type service struct {
	boardService protoBoard.BoardClient
	validator    Validator
	hubs         map[int64]*websocket.Hub
}

func NewService(boardService protoBoard.BoardClient, validator Validator) ServiceBoard {
	return &service{
		boardService: boardService,
		validator:    validator,
		hubs:         make(map[int64]*websocket.Hub, 0),
	}
}

func (s *service) CreateBoard(request models.BoardChangeInput) (board models.BoardOutsideShort, err error) {
	ctx := context.Background()

	input := &protoBoard.BoardChangeInput{
		UserID:    request.UserID,
		BoardID:   request.BoardID,
		BoardName: request.BoardName,
		Theme:     request.Theme,
		Star:      request.Star,
	}

	boardInternal, err := s.boardService.CreateBoard(ctx, input)
	if err != nil {
		return models.BoardOutsideShort{}, errorWorker.ConvertStatusToError(err)
	}

	board.BoardID = boardInternal.BoardID
	board.Name = boardInternal.Name
	board.Star = boardInternal.Star
	board.Theme = boardInternal.Theme
	
	return board, err
}

func (s *service) GetBoard(request models.BoardInput) (board models.BoardOutside, err error) {
	ctx := context.Background()

	input := &protoBoard.BoardInput{
		UserID:  request.UserID,
		BoardID: request.BoardID,
	}

	boardInternal, err := s.boardService.GetBoard(ctx, input)
	if err != nil {
		return models.BoardOutside{}, errorWorker.ConvertStatusToError(err)
	}

	board.Admin = models.UserOutsideShort{
		Email:    boardInternal.Admin.Email,
		Username: boardInternal.Admin.Username,
		FullName: boardInternal.Admin.FullName,
		Avatar:   boardInternal.Admin.Avatar,
	}

	for _, user := range boardInternal.Users{
		board.Users = append(board.Users, models.UserOutsideShort{
			Email:    user.Email,
			Username: user.Username,
			FullName: user.FullName,
			Avatar:   user.Avatar,
		})
	}

	for _, card := range boardInternal.Cards{
		tasks := make([]models.TaskOutsideShort, 0)
		for _, task := range card.Tasks {
			tasks = append(tasks, models.TaskOutsideShort{
				TaskID:      task.TaskID,
				Name:        task.Name,
				Description: task.Description,
				Order:       task.Order,
				Users:       convertUsers(task.Users),
				Tags: 	   	 convertTags(task.Tags),
				Checklists:  convertChecklists(task.Checklists),
			})
		}
		board.Cards = append(board.Cards, models.CardOutside{
			CardID: card.CardID,
			Name:   card.Name,
			Order:  card.Order,
			Tasks:  tasks,
		})
	}

	board.BoardID = boardInternal.BoardID
	board.Name = boardInternal.Name
	board.Theme = boardInternal.Theme
	board.Star = boardInternal.Star
	board.Admin = models.UserOutsideShort{
		Email: boardInternal.Admin.Email,
		Username: boardInternal.Admin.Username,
		FullName:  boardInternal.Admin.FullName,
		Avatar: boardInternal.Admin.Avatar,
	}
	board.Tags = convertTags(boardInternal.Tags)
	board.Users = convertUsers(boardInternal.Users)

	return board, nil
}

func (s *service) WebSocketBoard(request models.BoardInput, c echo.Context) (err error) {
	var hub *websocket.Hub
	if h, ok := s.hubs[request.BoardID]; ok {
		hub = h
	} else {
		hub = s.createHub(request.BoardID)
	}

	websocket.ServeWs(hub, c, request.SessionID, request.UserID)
	return nil
}

func convertTags(tags []*protoBoard.TagOutside) (output []models.TagOutside) {
	output = make([]models.TagOutside, 0)
	for _, tag := range tags {
		output = append(output, models.TagOutside{
			TagID: tag.TagID,
			Color: tag.Color,
			Name:  tag.Name,
		})
	}
	return output
}

func convertUsers(users []*protoProfile.UserOutsideShort) (output []models.UserOutsideShort) {
	output = make([]models.UserOutsideShort, 0)
	for _, user := range users {
		output = append(output, models.UserOutsideShort{
			Email: user.Email,
			Username: user.Username,
			FullName:  user.FullName,
			Avatar: user.Avatar,
		})
	}
	return output
}

func (s *service) ChangeBoard(request models.BoardChangeInput) (board models.BoardOutsideShort, err error) {
	ctx := context.Background()

	input := &protoBoard.BoardChangeInput{
		UserID:    request.UserID,
		BoardID:   request.BoardID,
		BoardName: request.BoardName,
		Theme:     request.Theme,
		Star:      request.Star,
	}

	boardInternal, err := s.boardService.ChangeBoard(ctx, input)
	if err != nil {
		return models.BoardOutsideShort{}, errorWorker.ConvertStatusToError(err)
	}

	board.BoardID = boardInternal.BoardID
	board.Name = boardInternal.Name
	board.Theme = boardInternal.Theme
	board.Star = boardInternal.Star

	if hub, ok := s.hubs[request.BoardID]; ok {
		some := models.WS{
			Method: "ChangeBoard",
			Body:   board,
		}
		hub.Broadcast(some)
	}

	return board, nil
}

func (s *service) DeleteBoard(request models.BoardInput) (err error) {
	ctx := context.Background()

	input := &protoBoard.BoardInput{
		UserID:  request.UserID,
		BoardID: request.BoardID,
	}

	_, err = s.boardService.DeleteBoard(ctx, input)
	if err != nil {
		return errorWorker.ConvertStatusToError(err)
	}

	if hub, ok := s.hubs[request.BoardID]; ok {
		some := models.WS{
			Method: "DeleteBoard",
			Body:   request,
		}
		hub.Broadcast(some)
	}

	return nil
}

func (s *service) AddMember(request models.BoardMemberInput) (user models.UserOutsideShort, err error) {
	ctx := context.Background()

	input := &protoBoard.BoardMemberInput{
		UserID:  request.UserID,
		BoardID: request.BoardID,
		MemberName: request.MemberName,
	}

	output, err := s.boardService.AddUserToBoard(ctx, input)
	if err != nil {
		return user, errorWorker.ConvertStatusToError(err)
	}

	user.Username = output.Username
	user.FullName = output.FullName
	user.Avatar = output.Avatar
	user.Email = output.Email

	if hub, ok := s.hubs[request.BoardID]; ok {
		some := models.WS{
			Method: "AddMember",
			Body:   user,
		}
		hub.Broadcast(some)
	}

	return user,nil
}

func (s *service) RemoveMember(request models.BoardMemberInput) (err error) {
	ctx := context.Background()

	input := &protoBoard.BoardMemberInput{
		UserID:  request.UserID,
		BoardID: request.BoardID,
		MemberName: request.MemberName,
	}

	_, err = s.boardService.RemoveUserToBoard(ctx, input)
	if err != nil {
		return errorWorker.ConvertStatusToError(err)
	}

	if hub, ok := s.hubs[request.BoardID]; ok {
		some := models.WS{
			Method: "RemoveMember",
			Body:   request,
		}
		hub.Broadcast(some)
	}

	return nil
}

func (s *service) CreateCard(request models.CardInput) (card models.CardOutsideShort, err error) {
	ctx := context.Background()

	input := &protoBoard.CardInput{
		UserID:  request.UserID,
		CardID:  request.CardID,
		BoardID: request.BoardID,
		Name:    request.Name,
		Order:   request.Order,
	}

	output, err := s.boardService.CreateCard(ctx, input)
	if err != nil {
		return models.CardOutsideShort{}, errorWorker.ConvertStatusToError(err)
	}

	card.CardID = output.CardID
	card.Name = output.Name

	if hub, ok := s.hubs[request.BoardID]; ok {
		some := models.WS{
			Method: "CreateCard",
			Body:   card,
		}
		hub.Broadcast(some)
	}

	return card, nil
}

func (s *service) GetCard(request models.CardInput) (card models.CardOutside, err error) {
	ctx := context.Background()

	input := &protoBoard.CardInput{
		UserID:  request.UserID,
		CardID:  request.CardID,
		BoardID: request.BoardID,
		Name:    request.Name,
		Order:   request.Order,
	}

	output, err := s.boardService.GetCard(ctx, input)
	if err != nil {
		return models.CardOutside{}, errorWorker.ConvertStatusToError(err)
	}

	for _, task := range output.Tasks{
		card.Tasks = append(card.Tasks, models.TaskOutsideShort{
			TaskID:      task.TaskID,
			Name:        task.Name,
			Description: task.Description,
			Order:       task.Order,
			Users:       []models.UserOutsideShort{},
		})
	}

	card.CardID = output.CardID
	card.Name = output.Name
	card.Order = output.Order

	return card, nil
}

func (s *service) ChangeCard(request models.CardInput) (card models.CardOutsideShort, err error) {
	ctx := context.Background()

	input := &protoBoard.CardInput{
		UserID:  request.UserID,
		CardID:  request.CardID,
		BoardID: request.BoardID,
		Name:    request.Name,
		Order:   request.Order,
	}

	output, err := s.boardService.ChangeCard(ctx, input)
	if err != nil {
		return models.CardOutsideShort{}, errorWorker.ConvertStatusToError(err)
	}

	card.CardID = output.CardID
	card.Name = output.Name

	if hub, ok := s.hubs[request.BoardID]; ok {
		some := models.WS{
			Method: "ChangeCard",
			Body:   card,
		}
		hub.Broadcast(some)
	}

	return card, nil
}

func (s *service) DeleteCard(request models.CardInput) (err error) {
	ctx := context.Background()
	input := &protoBoard.CardInput{
		UserID:  request.UserID,
		CardID:  request.CardID,
		BoardID: request.BoardID,
		Name:    request.Name,
		Order:   request.Order,
	}

	_, err = s.boardService.DeleteCard(ctx, input)
	if err != nil {
		return errorWorker.ConvertStatusToError(err)
	}

	if hub, ok := s.hubs[request.BoardID]; ok {
		some := models.WS{
			Method: "DeleteCard",
			Body:   request,
		}
		hub.Broadcast(some)
	}

	return nil
}

func (s *service) CardOrderChange(request models.CardsOrderInput) (err error) {
	ctx := context.Background()

	input := &protoBoard.CardsOrderInput{
		UserID: request.UserID,
		Cards:  nil,
	}

	for _, card := range request.Cards {
		input.Cards = append(input.Cards, &protoBoard.CardOrder{
			CardID: card.CardID,
			Order:  card.Order,
		})
	}
	
	_, err = s.boardService.CardOrderChange(ctx, input)
	if err != nil {
		return errorWorker.ConvertStatusToError(err)
	}

	if hub, ok := s.hubs[request.BoardID]; ok {
		some := models.WS{
			Method: "CardOrderChange",
			Body:   request,
		}
		hub.Broadcast(some)
	}

	return nil
}

func (s *service) CreateTask(request models.TaskInput) (task models.TaskOutsideSuperShort, err error) {
	ctx := context.Background()

	input := &protoBoard.TaskInput{
		UserID:      request.UserID,
		TaskID:      request.TaskID,
		CardID:      request.CardID,
		Name:        request.Name,
		Description: request.Description,
		Order:       request.Order,
	}

	output, err := s.boardService.CreateTask(ctx, input)
	if err != nil {
		return models.TaskOutsideSuperShort{}, errorWorker.ConvertStatusToError(err)
	}

	task.TaskID = output.TaskID
	task.Description = output.Description
	task.Name = output.Name

	if hub, ok := s.hubs[request.BoardID]; ok {
		some := models.WS{
			Method: "CreateTask",
			Body:   task,
		}
		hub.Broadcast(some)
	}

	return task, nil
}

func (s *service) GetTask(request models.TaskInput) (task models.TaskOutside, err error) {
	ctx := context.Background()

	input := &protoBoard.TaskInput{
		UserID:      request.UserID,
		TaskID:      request.TaskID,
		CardID:      request.CardID,
		Name:        request.Name,
		Description: request.Description,
		Order:       request.Order,
	}

	output, err := s.boardService.GetTask(ctx, input)
	if err != nil {
		return models.TaskOutside{}, errorWorker.ConvertStatusToError(err)
	}

	comments := make([]models.CommentOutside, 0)
	for _, comment := range output.Comments{
		comments = append(comments, models.CommentOutside{
			CommentID: comment.CommentID,
			Message:   comment.Message,
			Order:     comment.Order,
			User:      models.UserOutsideShort{
							Email: comment.User.Email,
							Username: comment.User.Username,
							FullName:  comment.User.FullName,
							Avatar: comment.User.Avatar,
			},
		})
	}

	attachments := make([]models.AttachmentOutside, 0)
	for _, attachment := range output.Attachments{
		attachments = append(attachments, models.AttachmentOutside{
			AttachmentID: attachment.AttachmentID,
			Filename:   attachment.Filename,
			Filepath:   attachment.Filepath,
		})
	}

	task.TaskID = output.TaskID
	task.Description = output.Description
	task.Name = output.Name
	task.Order = output.Order
	task.Users = convertUsers(output.Users)
	task.Checklists = convertChecklists(output.Checklists)
	task.Tags = convertTags(output.Tags)
	task.Comments = comments
	task.Attachments = attachments

	return task, nil
}

func convertChecklists(checklists []*protoBoard.ChecklistOutside) []models.ChecklistOutside {
	output := make([]models.ChecklistOutside, 0)
	for _, checklist := range checklists{
		output = append(output, models.ChecklistOutside{
			ChecklistID: checklist.ChecklistID,
			Name:   checklist.Name,
			Items:  checklist.Items,
		})
	}
	return output
}

func (s *service) ChangeTask(request models.TaskInput) (task models.TaskOutsideSuperShort, err error) {
	ctx := context.Background()

	input := &protoBoard.TaskInput{
		UserID:      request.UserID,
		TaskID:      request.TaskID,
		CardID:      request.CardID,
		Name:        request.Name,
		Description: request.Description,
		Order:       request.Order,
	}

	output, err := s.boardService.ChangeTask(ctx, input)
	if err != nil {
		return models.TaskOutsideSuperShort{}, errorWorker.ConvertStatusToError(err)
	}

	task.TaskID = output.TaskID
	task.Description = output.Description
	task.Name = output.Name

	if hub, ok := s.hubs[request.BoardID]; ok {
		some := models.WS{
			Method: "ChangeTask",
			Body:   task,
		}
		hub.Broadcast(some)
	}

	return task, nil
}

func (s *service) DeleteTask(request models.TaskInput) (err error) {
	ctx := context.Background()

	input := &protoBoard.TaskInput{
		UserID:      request.UserID,
		TaskID:      request.TaskID,
		CardID:      request.CardID,
		Name:        request.Name,
		Description: request.Description,
		Order:       request.Order,
	}

	_, err = s.boardService.DeleteTask(ctx, input)
	if err != nil {
		return errorWorker.ConvertStatusToError(err)
	}

	if hub, ok := s.hubs[request.BoardID]; ok {
		some := models.WS{
			Method: "DeleteTask",
			Body:   request,
		}
		hub.Broadcast(some)
	}

	return nil
}

func (s *service) TasksOrderChange(request models.TasksOrderInput) (err error) {
	ctx := context.Background()

	input := &protoBoard.TasksOrderInput{
		UserID: request.UserID,
		Tasks:  nil,
	}

	for _, task := range request.Tasks {
		tasks := make([]*protoBoard.TaskOrder, 0)
		for _, t := range task.Tasks {
			tasks = append(tasks, &protoBoard.TaskOrder{
				TaskID: t.TaskID,
				Order:  t.Order,
			})
		}
		input.Tasks = append(input.Tasks, &protoBoard.TasksOrder{
			CardID: task.CardID,
			Tasks:  tasks,
		})
	}

	_, err = s.boardService.TasksOrderChange(ctx, input)
	if err != nil {
		return errorWorker.ConvertStatusToError(err)
	}

	if hub, ok := s.hubs[request.BoardID]; ok {
		some := models.WS{
			Method: "TasksOrderChange",
			Body:   request,
		}
		hub.Broadcast(some)
	}

	return nil
}

func (s *service) AssignUser(request models.TaskAssignerInput) (user models.UserOutsideShort, err error) {
	ctx := context.Background()

	userInput := &protoBoard.TaskAssignerInput{
		UserID:     request.UserID,
		TaskID:     request.TaskID,
		AssignerName: request.AssignerName,
	}

	output, err := s.boardService.AssignUser(ctx, userInput)
	if err != nil {
		return user, errorWorker.ConvertStatusToError(err)
	}

	user.Username = output.Username
	user.FullName = output.FullName
	user.Avatar = output.Avatar
	user.Email = output.Email

	if hub, ok := s.hubs[request.BoardID]; ok {
		some := models.WS{
			Method: "AssignUser",
			Body:   user,
		}
		hub.Broadcast(some)
	}

	return user, nil
}

func (s *service) DismissUser(request models.TaskAssignerInput) (err error) {
	ctx := context.Background()

	userInput := &protoBoard.TaskAssignerInput{
		UserID:     request.UserID,
		TaskID:     request.TaskID,
		AssignerName: request.AssignerName,
	}

	_, err = s.boardService.DismissUser(ctx, userInput)
	if err != nil {
		return errorWorker.ConvertStatusToError(err)
	}

	if hub, ok := s.hubs[request.BoardID]; ok {
		some := models.WS{
			Method: "DismissUser",
			Body:   request,
		}
		hub.Broadcast(some)
	}

	return nil
}

func (s *service) CreateTag(request models.TagInput) (tag models.TagOutside, err error) {
	ctx := context.Background()

	input := &protoBoard.TagInput{
		UserID:      request.UserID,
		TaskID:      request.TaskID,
		TagID:      request.TagID,
		BoardID: request.BoardID,
		Name:        request.Name,
		Color: request.Color,
	}

	output, err := s.boardService.CreateTag(ctx, input)
	if err != nil {
		return models.TagOutside{}, errorWorker.ConvertStatusToError(err)
	}

	tag.TagID = output.TagID
	tag.Color = output.Color
	tag.Name = output.Name

	if hub, ok := s.hubs[request.BoardID]; ok {
		some := models.WS{
			Method: "CreateTag",
			Body:   tag,
		}
		hub.Broadcast(some)
	}

	return tag, nil
}

func (s *service) ChangeTag(request models.TagInput) (tag models.TagOutside, err error) {
	ctx := context.Background()

	input := &protoBoard.TagInput{
		UserID:      request.UserID,
		TaskID:      request.TaskID,
		TagID:      request.TagID,
		BoardID: request.BoardID,
		Name:        request.Name,
		Color: request.Color,
	}

	output, err := s.boardService.ChangeTag(ctx, input)
	if err != nil {
		return models.TagOutside{}, errorWorker.ConvertStatusToError(err)
	}

	tag.TagID = output.TagID
	tag.Color = output.Color
	tag.Name = output.Name

	if hub, ok := s.hubs[request.BoardID]; ok {
		some := models.WS{
			Method: "ChangeTag",
			Body:   tag,
		}
		hub.Broadcast(some)
	}

	return tag, nil
}

func (s *service) DeleteTag(request models.TagInput) (err error) {
	ctx := context.Background()

	input := &protoBoard.TagInput{
		UserID:      request.UserID,
		TaskID:      request.TaskID,
		TagID:      request.TagID,
		BoardID: request.BoardID,
		Name:        request.Name,
		Color: request.Color,
	}

	_, err = s.boardService.DeleteTag(ctx, input)
	if err != nil {
		return errorWorker.ConvertStatusToError(err)
	}

	if hub, ok := s.hubs[request.BoardID]; ok {
		some := models.WS{
			Method: "DeleteTag",
			Body:   request,
		}
		hub.Broadcast(some)
	}

	return nil
}

func (s *service) AddTag(request models.TaskTagInput) (err error) {
	ctx := context.Background()

	input := &protoBoard.TaskTagInput{
		UserID:      request.UserID,
		TaskID:      request.TaskID,
		TagID:      request.TagID,
	}

	_, err = s.boardService.AddTag(ctx, input)
	if err != nil {
		return errorWorker.ConvertStatusToError(err)
	}

	if hub, ok := s.hubs[request.BoardID]; ok {
		some := models.WS{
			Method: "AddTag",
			Body:   request,
		}
		hub.Broadcast(some)
	}

	return nil
}

func (s *service) RemoveTag(request models.TaskTagInput) (err error) {
	ctx := context.Background()

	input := &protoBoard.TaskTagInput{
		UserID:      request.UserID,
		TaskID:      request.TaskID,
		TagID:      request.TagID,
	}

	_, err = s.boardService.RemoveTag(ctx, input)
	if err != nil {
		return errorWorker.ConvertStatusToError(err)
	}

	if hub, ok := s.hubs[request.BoardID]; ok {
		some := models.WS{
			Method: "RemoveTag",
			Body:   request,
		}
		hub.Broadcast(some)
	}

	return nil
}

func (s *service) CreateComment(request models.CommentInput) (comment models.CommentOutside, err error) {
	ctx := context.Background()

	input := &protoBoard.CommentInput{
		UserID:      request.UserID,
		TaskID:      request.TaskID,
		CommentID:      request.CommentID,
		Message: request.Message,
		Order: request.Order,
	}

	output, err := s.boardService.CreateComment(ctx, input)
	if err != nil {
		return models.CommentOutside{}, errorWorker.ConvertStatusToError(err)
	}

	user := models.UserOutsideShort{
		Email:    output.User.Email,
		Username: output.User.Username,
		FullName: output.User.FullName,
		Avatar:   output.User.Avatar,
	}

	comment.CommentID = output.CommentID
	comment.Message = output.Message
	comment.Order = output.Order
	comment.User = user

	if hub, ok := s.hubs[request.BoardID]; ok {
		some := models.WS{
			Method: "CreateComment",
			Body:   comment,
		}
		hub.Broadcast(some)
	}

	return comment, nil
}

func (s *service) ChangeComment(request models.CommentInput) (comment models.CommentOutside, err error) {
	ctx := context.Background()

	input := &protoBoard.CommentInput{
		UserID:      request.UserID,
		TaskID:      request.TaskID,
		CommentID:      request.CommentID,
		Message: request.Message,
		Order: request.Order,
	}

	output, err := s.boardService.ChangeComment(ctx, input)
	if err != nil {
		return models.CommentOutside{}, errorWorker.ConvertStatusToError(err)
	}

	user := models.UserOutsideShort{
		Email:    output.User.Email,
		Username: output.User.Username,
		FullName: output.User.FullName,
		Avatar:   output.User.Avatar,
	}

	comment.CommentID = output.CommentID
	comment.Message = output.Message
	comment.Order = output.Order
	comment.User = user

	if hub, ok := s.hubs[request.BoardID]; ok {
		some := models.WS{
			Method: "ChangeComment",
			Body:   comment,
		}
		hub.Broadcast(some)
	}

	return comment, nil
}

func (s *service) DeleteComment(request models.CommentInput) (err error) {
	ctx := context.Background()

	input := &protoBoard.CommentInput{
		UserID:      request.UserID,
		TaskID:      request.TaskID,
		CommentID:      request.CommentID,
		Message: request.Message,
		Order: request.Order,
	}

	_, err = s.boardService.DeleteComment(ctx, input)
	if err != nil {
		return errorWorker.ConvertStatusToError(err)
	}

	if hub, ok := s.hubs[request.BoardID]; ok {
		some := models.WS{
			Method: "DeleteComment",
			Body:   request,
		}
		hub.Broadcast(some)
	}

	return nil
}

func (s *service) CreateChecklist(request models.ChecklistInput) (checklist models.ChecklistOutside, err error) {
	ctx := context.Background()

	input := &protoBoard.ChecklistInput{
		UserID: request.UserID,
		TaskID: request.TaskID,
		ChecklistID: request.ChecklistID,
		Name: request.Name,
		Items: request.Items,
	}

	output, err := s.boardService.CreateChecklist(ctx, input)
	if err != nil {
		return models.ChecklistOutside{}, errorWorker.ConvertStatusToError(err)
	}

	checklist.ChecklistID = output.ChecklistID
	checklist.Name = output.Name
	checklist.Items = output.Items

	if hub, ok := s.hubs[request.BoardID]; ok {
		some := models.WS{
			Method: "CreateChecklist",
			Body:   checklist,
		}
		hub.Broadcast(some)
	}

	return checklist, nil
}

func (s *service) ChangeChecklist(request models.ChecklistInput) (checklist models.ChecklistOutside, err error) {
	ctx := context.Background()

	input := &protoBoard.ChecklistInput{
		UserID: request.UserID,
		TaskID: request.TaskID,
		ChecklistID: request.ChecklistID,
		Name: request.Name,
		Items: request.Items,
	}

	output, err := s.boardService.ChangeChecklist(ctx, input)
	if err != nil {
		return models.ChecklistOutside{}, errorWorker.ConvertStatusToError(err)
	}

	checklist.ChecklistID = output.ChecklistID
	checklist.Name = output.Name
	checklist.Items = output.Items

	if hub, ok := s.hubs[request.BoardID]; ok {
		some := models.WS{
			Method: "ChangeChecklist",
			Body:   checklist,
		}
		hub.Broadcast(some)
	}

	return checklist, nil
}

func (s *service) DeleteChecklist(request models.ChecklistInput) (err error) {
	ctx := context.Background()

	input := &protoBoard.ChecklistInput{
		UserID: request.UserID,
		TaskID: request.TaskID,
		ChecklistID: request.ChecklistID,
		Name: request.Name,
		Items: request.Items,
	}

	_, err = s.boardService.DeleteChecklist(ctx, input)
	if err != nil {
		return errorWorker.ConvertStatusToError(err)
	}

	if hub, ok := s.hubs[request.BoardID]; ok {
		some := models.WS{
			Method: "DeleteChecklist",
			Body:   request,
		}
		hub.Broadcast(some)
	}

	return nil
}
func (s *service) CreateAttachment(request models.AttachmentInput) (attachment models.AttachmentOutside, err error) {
	ctx := context.Background()

	stream, err := s.boardService.AddAttachment(ctx)
	if err != nil {
		return attachment, errorWorker.ConvertStatusToError(err)
	}

	req := &protoBoard.AttachmentInput{
		Data: &protoBoard.AttachmentInput_Info{
			Info: &protoBoard.AttachmentInfo{
				UserID: request.UserID,
				TaskID: request.TaskID,
				Filename: request.Filename,
			},
		},
	}

	err = stream.Send(req)
	if err != nil {
		stream.RecvMsg(nil)
		return attachment, errorWorker.ConvertStatusToError(err)
	}

	reader := bytes.NewReader(request.File)
	buffer := make([]byte, 1024)

	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return attachment, errorWorker.ConvertStatusToError(err)
		}

		req := &protoBoard.AttachmentInput{
			Data: &protoBoard.AttachmentInput_File{
				File: buffer[:n],
			},
		}

		err = stream.Send(req)
		if err != nil {
			return attachment, errorWorker.ConvertStatusToError(err)
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		return attachment, errorWorker.ConvertStatusToError(err)
	}


	attachment.AttachmentID = res.AttachmentID
	attachment.Filename = res.Filename
	attachment.Filepath = res.Filepath

	if hub, ok := s.hubs[request.BoardID]; ok {
		some := models.WS{
			Method: "CreateAttachment",
			Body:   attachment,
		}
		hub.Broadcast(some)
	}

	return attachment, nil
}

func (s *service) DeleteAttachment(request models.AttachmentInput) (err error) {
	ctx := context.Background()

	input := &protoBoard.AttachmentInfo{
		UserID: request.UserID,
		TaskID: request.TaskID,
		AttachmentID: request.AttachmentID,
		Filename: request.Filename,
	}

	_, err = s.boardService.RemoveAttachment(ctx, input)
	if err != nil {
		return errorWorker.ConvertStatusToError(err)
	}

	if hub, ok := s.hubs[request.BoardID]; ok {
		some := models.WS{
			Method: "DeleteAttachment",
			Body:   request,
		}
		hub.Broadcast(some)
	}

	return nil
}

func (s *service) GetSharedURL(request models.BoardInput) (url string, err error) {
	ctx := context.Background()

	input := &protoBoard.BoardInput{
		UserID:  request.UserID,
		BoardID: request.BoardID,
	}

	output, err := s.boardService.GetSharedURL(ctx, input)
	if err != nil {
		return url, errorWorker.ConvertStatusToError(err)
	}

	return output.Url, nil
}

func (s *service) InviteUserToBoard(request models.BoardInviteInput) (board models.BoardOutsideShort, err error) {
	ctx := context.Background()

	input := &protoBoard.BoardInviteInput{
		UserID:  request.UserID,
		BoardID: request.BoardID,
		UrlHash: request.UrlHash,
	}

	output, err := s.boardService.InviteUserToBoard(ctx, input)
	if err != nil {
		return board, errorWorker.ConvertStatusToError(err)
	}

	board = models.BoardOutsideShort{
		BoardID: output.BoardID,
		Name:    output.Name,
		Theme:   output.Theme,
		Star:    output.Star,
	}

	return board, nil
}

func (s *service) CheckBoardPermission(userID int64, boardID int64, ifAdmin bool) (err error) {
	ctx := context.Background()

	input := &protoBoard.CheckPermissions{
		UserID:    userID,
		ElementID: boardID,
		IfAdmin:   ifAdmin,
	}

	_, err = s.boardService.CheckBoardPermission(ctx, input)
	if err != nil {
		return errorWorker.ConvertStatusToError(err)
	}

	return nil
}

func (s *service) CheckCardPermission(userID int64, cardID int64) (boardID int64, err error) {
	ctx := context.Background()

	input := &protoBoard.CheckPermissions{
		UserID:    userID,
		ElementID: cardID,
	}

	output, err := s.boardService.CheckCardPermission(ctx, input)
	if err != nil {
		return 0, errorWorker.ConvertStatusToError(err)
	}

	return output.BoardID, nil
}

func (s *service) CheckTaskPermission(userID int64, taskID int64) (boardID int64, err error) {
	ctx := context.Background()

	input := &protoBoard.CheckPermissions{
		UserID:    userID,
		ElementID: taskID,
	}

	output, err := s.boardService.CheckTaskPermission(ctx, input)
	if err != nil {
		return 0, errorWorker.ConvertStatusToError(err)
	}

	return output.BoardID, nil
}

func (s *service) CheckCommentPermission(userID int64, commentID int64, ifAdmin bool) (boardID int64, err error) {
	ctx := context.Background()

	input := &protoBoard.CheckPermissions{
		UserID:    userID,
		ElementID: commentID,
		IfAdmin: ifAdmin,
	}

	output, err := s.boardService.CheckCommentPermission(ctx, input)
	if err != nil {
		return 0, errorWorker.ConvertStatusToError(err)
	}

	return output.BoardID, nil
}

func (s *service) createHub(boardID int64) *websocket.Hub {
	hub := websocket.NewHub(boardID, &s.hubs)
	s.hubs[boardID] = hub
	go hub.Run()
	return hub
}

func (s *service) deleteHub(boardID int64) {
	s.hubs[boardID].StopHub()
	delete(s.hubs, boardID)
}
