package boards

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
)

type Service interface {
	CreateBoard(request models.BoardChangeInput) (board models.BoardOutside, err error)
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
}

type service struct {
	userStorage  UserStorage
	boardStorage BoardStorage
	validator    Validator
}

func NewService(userStorage UserStorage, boardStorage BoardStorage, validator Validator) Service {
	return &service{
		userStorage: userStorage,
		boardStorage: boardStorage,
		validator: validator,
	}
}

func (s *service) CreateBoard(request models.BoardChangeInput) (board models.BoardOutside, err error) {
	boardInternal, err := s.boardStorage.CreateBoard(request)
	if err != nil {
		return models.BoardOutside{}, err
	}

	membersIDs := make([]uint64, 0)
	membersIDs = append(membersIDs, boardInternal.AdminID)

	members, err := s.userStorage.GetBoardMembers(membersIDs)

	writeBoardOutside(boardInternal, &board, members)
	
	return board, err
}

func (s *service) GetBoard(request models.BoardInput) (board models.BoardOutside, err error) {
	boardInternal, err := s.boardStorage.GetBoard(request)
	if err != nil {
		return models.BoardOutside{}, err
	}

	membersIDs := make([]uint64, 0)
	membersIDs = append(membersIDs, boardInternal.AdminID)
	membersIDs = append(membersIDs, boardInternal.UsersIDs...)

	members, err := s.userStorage.GetBoardMembers(membersIDs)
	if err != nil {
		return models.BoardOutside{}, err
	}

	writeBoardOutside(boardInternal, &board, members)

	return board, err
}

func (s *service) ChangeBoard(request models.BoardChangeInput) (board models.BoardOutside, err error) {
	boardInternal, err := s.boardStorage.ChangeBoard(request)
	if err != nil {
		return models.BoardOutside{}, err
	}

	membersIDs := make([]uint64, 0)
	membersIDs = append(membersIDs, boardInternal.AdminID)
	membersIDs = append(membersIDs, boardInternal.UsersIDs...)

	members, err := s.userStorage.GetBoardMembers(membersIDs)
	if err != nil {
		return models.BoardOutside{}, err
	}

	writeBoardOutside(boardInternal, &board, members)

	return board, err
}

func writeBoardOutside(boardInternal models.BoardInternal, board *models.BoardOutside, members []models.UserOutsideShort)  {
	board.BoardID = boardInternal.BoardID
	board.Name = boardInternal.Name
	board.Star = boardInternal.Star
	board.Theme = boardInternal.Theme
	board.Cards = boardInternal.Cards
	board.Admin = members[0]
	if len(members) > 1 {
		board.Users = members[0:]
	}
}

func (s *service) DeleteBoard(request models.BoardInput) (err error) {
	err = s.boardStorage.DeleteBoard(request)
	if err != nil {
		return err
	}

	return err
}

func (s *service) CreateCard(request models.CardInput) (card models.CardOutside, err error) {
	card, err = s.boardStorage.CreateCard(request)
	if err != nil {
		return models.CardOutside{}, err
	}

	return card, err
}

func (s *service) GetCard(request models.CardInput) (card models.CardOutside, err error) {
	card, err = s.boardStorage.GetCard(request)
	if err != nil {
		return models.CardOutside{}, err
	}

	return card, err
}

func (s *service) ChangeCard(request models.CardInput) (card models.CardOutside, err error) {
	card, err = s.boardStorage.ChangeCard(request)
	if err != nil {
		return models.CardOutside{}, err
	}

	return card, err
}

func (s *service) DeleteCard(request models.CardInput) (err error) {
	err = s.boardStorage.DeleteCard(request)
	if err != nil {
		return err
	}

	return err
}

func (s *service) CardOrderChange(request models.CardsOrderInput) (err error) {
	err = s.boardStorage.ChangeCardOrder(request)
	if err != nil {
		return err
	}

	return err
}

func (s *service) CreateTask(request models.TaskInput) (task models.TaskOutside, err error) {
	task, err = s.boardStorage.CreateTask(request)
	if err != nil {
		return models.TaskOutside{}, err
	}

	return task, err
}

func (s *service) GetTask(request models.TaskInput) (task models.TaskOutside, err error) {
	task, err = s.boardStorage.GetTask(request)
	if err != nil {
		return models.TaskOutside{}, err
	}

	return task, err
}

func (s *service) ChangeTask(request models.TaskInput) (task models.TaskOutside, err error) {
	task, err = s.boardStorage.ChangeTask(request)
	if err != nil {
		return models.TaskOutside{}, err
	}

	return task, err
}

func (s *service) DeleteTask(request models.TaskInput) (err error) {
	err = s.boardStorage.DeleteTask(request)
	if err != nil {
		return err
	}

	return err
}

func (s *service) TasksOrderChange(request models.TasksOrderInput) (err error) {
	err = s.boardStorage.ChangeTaskOrder(request)
	if err != nil {
		return err
	}

	return err
}
