package boards

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/errorWorker"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	protoBoard "github.com/go-park-mail-ru/2020_2_ExtraSafe/services/proto/board"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/proto/profile"
)

type Service interface {
	CreateBoard(request models.BoardChangeInput) (board models.BoardOutsideShort, err error)
	GetBoard(request models.BoardInput) (board models.BoardOutside, err error)
	ChangeBoard(request models.BoardChangeInput) (board models.BoardOutsideShort, err error)
	DeleteBoard(request models.BoardInput) (err error)
	AddMember(request models.BoardMemberInput) (err error)
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
	AssignUser(request models.TaskAssignerInput) (err error)
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

	CheckBoardPermission(userID int64, boardID int64, ifAdmin bool) (err error)
	CheckCardPermission(userID int64, cardID int64) (err error)
	CheckTaskPermission(userID int64, taskID int64) (err error)
	CheckCommentPermission(userID int64, commentID int64, ifAdmin bool) (err error)
}

type service struct {
	boardService protoBoard.BoardClient
	validator    Validator
}

func NewService(boardService protoBoard.BoardClient, validator Validator) Service {
	return &service{
		boardService: boardService,
		validator: validator,
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

	return nil
}

func (s *service) AddMember(request models.BoardMemberInput) (err error) {
	ctx := context.Background()

	input := &protoBoard.BoardMemberInput{
		UserID:  request.UserID,
		BoardID: request.BoardID,
		MemberName: request.MemberName,
	}

	_, err = s.boardService.AddUserToBoard(ctx, input)
	if err != nil {
		return errorWorker.ConvertStatusToError(err)
	}

	return nil
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

	return nil
}

func (s *service) AssignUser(request models.TaskAssignerInput) (err error) {
	ctx := context.Background()

	userInput := &protoBoard.TaskAssignerInput{
		UserID:     request.UserID,
		TaskID:     request.TaskID,
		AssignerName: request.AssignerName,
	}

	fmt.Println(userInput)

	_, err = s.boardService.AssignUser(ctx, userInput)
	if err != nil {
		return errorWorker.ConvertStatusToError(err)
	}

	return nil
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

	return nil
}

func (s *service) CreateAttachment(request models.AttachmentInput) (attachment models.AttachmentOutside, err error) {
	ctx := context.Background()

	input := &protoBoard.AttachmentInput{
		UserID: request.UserID,
		TaskID: request.TaskID,
		Filename: request.Filename,
		File: request.File,
	}

	output, err := s.boardService.AddAttachment(ctx, input)
	if err != nil {
		return models.AttachmentOutside{}, errorWorker.ConvertStatusToError(err)
	}

	attachment.AttachmentID = output.AttachmentID
	attachment.Filename = output.Filename
	attachment.Filepath = output.Filepath

	return attachment, nil
}

func (s *service) DeleteAttachment(request models.AttachmentInput) (err error) {
	ctx := context.Background()

	input := &protoBoard.AttachmentInput{
		UserID: request.UserID,
		TaskID: request.TaskID,
		AttachmentID: request.AttachmentID,
		Filename: request.Filename,
		File: request.File,
	}
	fmt.Println(input)

	_, err = s.boardService.RemoveAttachment(ctx, input)
	if err != nil {
		return errorWorker.ConvertStatusToError(err)
	}

	return nil
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

func (s *service) CheckCardPermission(userID int64, cardID int64) (err error) {
	ctx := context.Background()

	input := &protoBoard.CheckPermissions{
		UserID:    userID,
		ElementID: cardID,
	}

	_, err = s.boardService.CheckCardPermission(ctx, input)
	if err != nil {
		return errorWorker.ConvertStatusToError(err)
	}

	return nil
}

func (s *service) CheckTaskPermission(userID int64, taskID int64) (err error) {
	ctx := context.Background()

	input := &protoBoard.CheckPermissions{
		UserID:    userID,
		ElementID: taskID,
	}

	_, err = s.boardService.CheckTaskPermission(ctx, input)
	if err != nil {
		return errorWorker.ConvertStatusToError(err)
	}

	return nil
}

func (s *service) CheckCommentPermission(userID int64, commentID int64, ifAdmin bool) (err error) {
	ctx := context.Background()

	input := &protoBoard.CheckPermissions{
		UserID:    userID,
		ElementID: commentID,
		IfAdmin: ifAdmin,
	}

	_, err = s.boardService.CheckCommentPermission(ctx, input)
	if err != nil {
		return errorWorker.ConvertStatusToError(err)
	}

	return nil
}