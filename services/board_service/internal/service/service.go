package service

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/boardStorage"
	protoBoard "github.com/go-park-mail-ru/2020_2_ExtraSafe/services/proto/board"
	protoProfile "github.com/go-park-mail-ru/2020_2_ExtraSafe/services/proto/profile"
	"golang.org/x/net/context"
)

type service struct {
	boardStorage boardStorage.Storage

}


func NewService(boardStorage boardStorage.Storage) *service {
	return &service{
		boardStorage: boardStorage,
	}
}

func (s *service) GetBoardsList(c context.Context, input *protoProfile.UserID) (output *protoProfile.BoardsOutsideShort, err error) {
	userInput := models.UserInput{ID: input.ID}

	boardsList, err := s.boardStorage.GetBoardsList(userInput)
	if err != nil {
		return output, nil
	}

	output = new(protoProfile.BoardsOutsideShort)

	for _, board := range boardsList {
		outputBoard := protoProfile.BoardOutsideShort{
			BoardID: board.BoardID,
			Name:    board.Name,
			Theme:   board.Theme,
			Star:    board.Star,
		}
		output.Boards = append(output.Boards, &outputBoard)
	}
	return output, nil
}

func (s *service) CreateBoard(c context.Context, input *protoBoard.BoardChangeInput) (output *protoProfile.BoardOutsideShort, err error) {
	userInput := models.BoardChangeInput{
		UserID:    input.UserID,
		BoardID:   input.BoardID,
		BoardName: input.BoardName,
		Theme:     input.Theme,
		Star:      input.Star,
	}

	boardInternal, err := s.boardStorage.CreateBoard(userInput)
	if err != nil {
		return output, err
	}

	output = &protoProfile.BoardOutsideShort{
		BoardID: boardInternal.BoardID,
		Name:    boardInternal.Name,
		Theme:   boardInternal.Theme,
		Star:    boardInternal.Star,
	}
	return output, nil
}

func (s *service) GetBoard(c context.Context, input *protoBoard.BoardInput) (output *protoBoard.BoardOutside, err error) {
	userInput := models.BoardInput{
		UserID:    input.UserID,
		BoardID:   input.BoardID,
	}

	boardInternal, err := s.boardStorage.GetBoard(userInput)
	if err != nil {
		return output, err
	}

	cards := make([]*protoBoard.CardOutside, 0)
	for _, card := range boardInternal.Cards {
		tasks := make([]*protoBoard.TaskOutside, 0)
		for _, task := range card.Tasks {
			tasks = append(tasks, &protoBoard.TaskOutside{
				TaskID:      task.TaskID,
				Name:        task.Name,
				Description: task.Description,
				Order:       task.Order,
			})
		}
		cards = append(cards, &protoBoard.CardOutside{
			CardID: card.CardID,
			Name:   card.Name,
			Order:  card.Order,
			Tasks:  tasks,
		})
	}

	output = &protoBoard.BoardOutside{
		BoardID: boardInternal.BoardID,
		Admin:   nil,
		Name:    boardInternal.Name,
		Theme:   boardInternal.Theme,
		Star:    boardInternal.Star,
		Users:   nil,
		Cards:   cards,
	}

	return output, nil
}

func (s *service) ChangeBoard(c context.Context, input *protoBoard.BoardChangeInput) (output *protoBoard.BoardOutside, err error) {
	userInput := models.BoardChangeInput{
		UserID:    input.UserID,
		BoardID:   input.BoardID,
		BoardName: input.BoardName,
		Theme:     input.Theme,
		Star:      input.Star,
	}

	boardInternal, err := s.boardStorage.ChangeBoard(userInput)
	if err != nil {
		return output, err
	}

	cards := make([]*protoBoard.CardOutside, 0)
	for _, card := range boardInternal.Cards {
		tasks := make([]*protoBoard.TaskOutside, 0)
		for _, task := range card.Tasks {
			tasks = append(tasks, &protoBoard.TaskOutside{
				TaskID:      task.TaskID,
				Name:        task.Name,
				Description: task.Description,
				Order:       task.Order,
			})
		}
		cards = append(cards, &protoBoard.CardOutside{
			CardID: card.CardID,
			Name:   card.Name,
			Order:  card.Order,
			Tasks:  tasks,
		})
	}

	output = &protoBoard.BoardOutside{
		BoardID: boardInternal.BoardID,
		Admin:   nil,
		Name:    boardInternal.Name,
		Theme:   boardInternal.Theme,
		Star:    boardInternal.Star,
		Users:   nil,
		Cards:   cards,
	}
	return output, nil
}

func (s *service) DeleteBoard(c context.Context, input *protoBoard.BoardInput) (*protoBoard.Nothing, error) {
	userInput := models.BoardInput{
		UserID:    input.UserID,
		BoardID:   input.BoardID,
	}

	err := s.boardStorage.DeleteBoard(userInput)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *service) CreateCard(c context.Context, input *protoBoard.CardInput) (output *protoBoard.CardOutside, err error) {
	userInput := models.CardInput{
		UserID:  input.UserID,
		CardID:  input.CardID,
		BoardID: input.BoardID,
		Name:    input.Name,
		Order:   input.Order,
	}

	card, err := s.boardStorage.CreateCard(userInput)
	if err != nil {
		return output, err
	}

	output = &protoBoard.CardOutside{
		CardID: card.CardID,
		Name:   card.Name,
		Order:  card.Order,
		Tasks:  nil,
	}

	return output, nil
}

func (s *service) GetCard(c context.Context, input *protoBoard.CardInput) (output *protoBoard.CardOutside, err error) {
	userInput := models.CardInput{
		UserID:  input.UserID,
		CardID:  input.CardID,
		BoardID: input.BoardID,
		Name:    input.Name,
		Order:   input.Order,
	}

	card, err := s.boardStorage.GetCard(userInput)
	if err != nil {
		return output, err
	}

	output = &protoBoard.CardOutside{
		CardID: card.CardID,
		Name:   card.Name,
		Order:  card.Order,
		Tasks:  nil,
	}

	return output, nil
}

func (s *service) ChangeCard(c context.Context, input *protoBoard.CardInput) (output *protoBoard.CardOutside, err error) {
	userInput := models.CardInput{
		UserID:  input.UserID,
		CardID:  input.CardID,
		BoardID: input.BoardID,
		Name:    input.Name,
		Order:   input.Order,
	}

	card, err := s.boardStorage.ChangeCard(userInput)
	if err != nil {
		return output, err
	}

	output = &protoBoard.CardOutside{
		CardID: card.CardID,
		Name:   card.Name,
		Order:  card.Order,
		Tasks:  nil,
	}

	return output, nil
}

func (s *service) DeleteCard(c context.Context, input *protoBoard.CardInput) (*protoBoard.Nothing, error) {
	userInput := models.CardInput{
		UserID:    input.UserID,
		CardID:   input.BoardID,
		BoardID: input.BoardID,
	}

	err := s.boardStorage.DeleteCard(userInput)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *service) CardOrderChange(c context.Context, input *protoBoard.CardsOrderInput) (*protoBoard.Nothing, error) {
	cardOrder := make([]models.CardOrder, 0)
	for _, card := range input.Cards {
		cardOrder = append(cardOrder, models.CardOrder{
			CardID: card.CardID,
			Order:  card.Order,
		})
	}

	userInput := models.CardsOrderInput{
		UserID:  input.UserID,
		Cards: cardOrder,
	}

	err := s.boardStorage.ChangeCardOrder(userInput)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *service) CreateTask(c context.Context, input *protoBoard.TaskInput) (output *protoBoard.TaskOutside, err error) {
	userInput := models.TaskInput{
		UserID:  input.UserID,
		CardID:  input.CardID,
		Name:    input.Name,
		Order:   input.Order,
		TaskID: input.TaskID,
		Description: input.Description,
	}

	task, err := s.boardStorage.CreateTask(userInput)
	if err != nil {
		return output, err
	}

	output = &protoBoard.TaskOutside{
		TaskID: task.TaskID,
		Name:   task.Name,
		Order:  task.Order,
		Description:  task.Description,
	}

	return output, nil
}
func (s *service) GetTask(c context.Context, input *protoBoard.TaskInput) (output *protoBoard.TaskOutside, err error) {
	userInput := models.TaskInput{
		UserID:  input.UserID,
		CardID:  input.CardID,
		Name:    input.Name,
		Order:   input.Order,
		TaskID: input.TaskID,
		Description: input.Description,
	}

	task, err := s.boardStorage.GetTask(userInput)
	if err != nil {
		return output, err
	}

	output = &protoBoard.TaskOutside{
		TaskID: task.TaskID,
		Name:   task.Name,
		Order:  task.Order,
		Description:  task.Description,
	}

	return output, nil
}

func (s *service) ChangeTask(c context.Context, input *protoBoard.TaskInput) (output *protoBoard.TaskOutside, err error) {
	userInput := models.TaskInput{
		UserID:  input.UserID,
		CardID:  input.CardID,
		Name:    input.Name,
		Order:   input.Order,
		TaskID: input.TaskID,
		Description: input.Description,
	}

	task, _, err := s.boardStorage.ChangeTask(userInput)
	if err != nil {
		return output, err
	}

	output = &protoBoard.TaskOutside{
		TaskID: task.TaskID,
		Name:   task.Name,
		Order:  task.Order,
		Description:  task.Description,
	}

	return output, nil
}

func (s *service) DeleteTask(c context.Context, input *protoBoard.TaskInput) (*protoBoard.Nothing, error) {
	userInput := models.TaskInput{
		UserID:    input.UserID,
		TaskID:   input.TaskID,
		CardID: input.CardID,
	}

	err := s.boardStorage.DeleteTask(userInput)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *service) TasksOrderChange(c context.Context, input *protoBoard.TasksOrderInput) (*protoBoard.Nothing, error) {
	tasksOrder := make([]models.TasksOrder, 0)
	for _, tasks := range input.Tasks {
		taskOrder := make([]models.TaskOrder, 0)
		for _, task := range tasks.Tasks {
			taskOrder = append(taskOrder, models.TaskOrder{
				TaskID: task.TaskID,
				Order:  task.Order,
			})
		}
		tasksOrder = append(tasksOrder, models.TasksOrder{
			CardID: tasks.CardID,
			Tasks:  taskOrder,
		})
	}

	userInput := models.TasksOrderInput{
		UserID:  input.UserID,
		Tasks: tasksOrder,
	}

	err := s.boardStorage.ChangeTaskOrder(userInput)
	if err != nil {
		return nil, err
	}

	return nil, nil
}