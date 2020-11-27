package boards

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/errorWorker"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	protoBoard "github.com/go-park-mail-ru/2020_2_ExtraSafe/services/proto/board"
)

type Service interface {
	CreateBoard(request models.BoardChangeInput) (board models.BoardOutsideShort, err error)
	GetBoard(request models.BoardInput) (board models.BoardOutside, err error)
	ChangeBoard(request models.BoardChangeInput) (board models.BoardOutside, err error)
	DeleteBoard(request models.BoardInput) (err error)

	CreateCard(request models.CardInput) (card models.CardOutside, err error)
	GetCard(request models.CardInput) (board models.CardOutside, err error)
	ChangeCard(request models.CardInput) (card models.CardOutside, err error)
	DeleteCard(request models.CardInput) (err error)
	CardOrderChange(request models.CardsOrderInput) (err error)

	CreateTask(request models.TaskInput) (task models.TaskOutside, err error)
	GetTask(request models.TaskInput) (board models.TaskOutside, err error)
	ChangeTask(request models.TaskInput) (task models.TaskOutside, err error)
	DeleteTask(request models.TaskInput) (err error)
	TasksOrderChange(request models.TasksOrderInput) (err error)

	CheckBoardPermission(userID int64, boardID int64, ifAdmin bool) (err error)
	CheckCardPermission(userID int64, cardID int64) (err error)
	CheckTaskPermission(userID int64, taskID int64) (err error)
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
		tasks := make([]models.TaskOutside, 0)
		for _, task := range card.Tasks {
			tasks = append(tasks, models.TaskOutside{
				TaskID:      task.TaskID,
				Name:        task.Name,
				Description: task.Description,
				Order:       task.Order,
				Users:       models.UserOutsideShort{},
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

	return board, nil
}

func (s *service) ChangeBoard(request models.BoardChangeInput) (board models.BoardOutside, err error) {
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
		tasks := make([]models.TaskOutside, 0)
		for _, task := range card.Tasks {
			tasks = append(tasks, models.TaskOutside{
				TaskID:      task.TaskID,
				Name:        task.Name,
				Description: task.Description,
				Order:       task.Order,
				Users:       models.UserOutsideShort{},
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

	return board, nil
}

func writeBoardOutside(boardInternal models.BoardInternal, members []models.UserOutsideShort) (board models.BoardOutside) {
	board.BoardID = boardInternal.BoardID
	board.Name = boardInternal.Name
	board.Star = boardInternal.Star
	board.Theme = boardInternal.Theme
	board.Cards = boardInternal.Cards
	board.Admin = members[0]
	if len(members) > 1 {
		board.Users = members[0:]
	}
	return
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

func (s *service) CreateCard(request models.CardInput) (card models.CardOutside, err error) {
	ctx := context.Background()

	fmt.Println(request)
	input := &protoBoard.CardInput{
		UserID:  request.UserID,
		CardID:  request.CardID,
		BoardID: request.BoardID,
		Name:    request.Name,
		Order:   request.Order,
	}

	output, err := s.boardService.CreateCard(ctx, input)
	if err != nil {
		return models.CardOutside{}, errorWorker.ConvertStatusToError(err)
	}

	for _, task := range output.Tasks{
		card.Tasks = append(card.Tasks, models.TaskOutside{
			TaskID:      task.TaskID,
			Name:        task.Name,
			Description: task.Description,
			Order:       task.Order,
			Users:       models.UserOutsideShort{},
		})
	}

	card.CardID = output.CardID
	card.Name = output.Name
	card.Order = output.Order

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
		card.Tasks = append(card.Tasks, models.TaskOutside{
			TaskID:      task.TaskID,
			Name:        task.Name,
			Description: task.Description,
			Order:       task.Order,
			Users:       models.UserOutsideShort{},
		})
	}

	card.CardID = output.CardID
	card.Name = output.Name
	card.Order = output.Order

	return card, nil
}

func (s *service) ChangeCard(request models.CardInput) (card models.CardOutside, err error) {
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
		return models.CardOutside{}, errorWorker.ConvertStatusToError(err)
	}

	for _, task := range output.Tasks{
		card.Tasks = append(card.Tasks, models.TaskOutside{
			TaskID:      task.TaskID,
			Name:        task.Name,
			Description: task.Description,
			Order:       task.Order,
			Users:       models.UserOutsideShort{},
		})
	}

	card.CardID = output.CardID
	card.Name = output.Name
	card.Order = output.Order

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

func (s *service) CreateTask(request models.TaskInput) (task models.TaskOutside, err error) {
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
		return models.TaskOutside{}, errorWorker.ConvertStatusToError(err)
	}

	task.TaskID = output.TaskID
	task.Description = output.Description
	task.Name = output.Name
	task.Order = output.Order
	task.Users = models.UserOutsideShort{}

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

	task.TaskID = output.TaskID
	task.Description = output.Description
	task.Name = output.Name
	task.Order = output.Order
	task.Users = models.UserOutsideShort{}

	return task, nil
}

func (s *service) ChangeTask(request models.TaskInput) (task models.TaskOutside, err error) {
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
		return models.TaskOutside{}, errorWorker.ConvertStatusToError(err)
	}

	task.TaskID = output.TaskID
	task.Description = output.Description
	task.Name = output.Name
	task.Order = output.Order
	task.Users = models.UserOutsideShort{}

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