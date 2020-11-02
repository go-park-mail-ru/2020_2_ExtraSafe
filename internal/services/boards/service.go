package boards

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
)

type Service interface {
	CreateBoard(request models.BoardChangeInput) (board models.BoardOutside, err error)
	GetBoard(request models.BoardInput) (board models.BoardOutside, err error)
	ChangeBoard(request models.BoardChangeInput) (board models.BoardOutside, err error)
	DeleteBoard(request models.BoardInput) (err error)

	CreateColumn(request models.ColumnInput) (board models.ColumnOutside, err error)
	ChangeColumn(request models.ColumnInput) (board models.ColumnOutside, err error)
	DeleteColumn(request models.ColumnInput) (err error)

	CreateTask(request models.TaskInput) (board models.TaskOutside, err error)
	ChangeTask(request models.TaskInput) (board models.TaskOutside, err error)
	DeleteTask(request models.TaskInput) (err error)
}

type service struct {
	userStorage userStorage
	boardStorage boardStorage
	validator validator
}

func NewService(userStorage userStorage, boardStorage boardStorage, validator validator) Service {
	return &service{
		userStorage: userStorage,
		boardStorage: boardStorage,
		validator: validator,
	}
}

func (s *service) CreateBoard(request models.BoardChangeInput) (board models.BoardOutside, err error) {
	board, err = s.boardStorage.CreateBoard(request)
	if err != nil {
		return models.BoardOutside{}, err
	}
	
	return board, err
}

func (s *service) GetBoard(request models.BoardInput) (board models.BoardOutside, err error) {
	board, err = s.boardStorage.GetBoard(request)
	if err != nil {
		return models.BoardOutside{}, err
	}

	return board, err
}

func (s *service) ChangeBoard(request models.BoardChangeInput) (board models.BoardOutside, err error) {
	board, err = s.boardStorage.ChangeBoard(request)
	if err != nil {
		return models.BoardOutside{}, err
	}

	return board, err
}

func (s *service) DeleteBoard(request models.BoardInput) (err error) {
	err = s.boardStorage.DeleteBoard(request)
	if err != nil {
		return err
	}

	return err
}

func (s *service) CreateColumn(request models.ColumnInput) (board models.ColumnOutside, err error) {
	board, err = s.boardStorage.CreateColumn(request)
	if err != nil {
		return models.ColumnOutside{}, err
	}

	return board, err
}

func (s *service) ChangeColumn(request models.ColumnInput) (board models.ColumnOutside, err error) {
	board, err = s.boardStorage.ChangeColumn(request)
	if err != nil {
		return models.ColumnOutside{}, err
	}

	return board, err
}

func (s *service) DeleteColumn(request models.ColumnInput) (err error) {
	err = s.boardStorage.DeleteColumn(request)
	if err != nil {
		return err
	}

	return err
}

func (s *service) CreateTask(request models.TaskInput) (board models.TaskOutside, err error) {
	board, err = s.boardStorage.CreateTask(request)
	if err != nil {
		return models.TaskOutside{}, err
	}

	return board, err
}

func (s *service) ChangeTask(request models.TaskInput) (board models.TaskOutside, err error) {
	board, err = s.boardStorage.ChangeTask(request)
	if err != nil {
		return models.TaskOutside{}, err
	}

	return board, err
}

func (s *service) DeleteTask(request models.TaskInput) (err error) {
	err = s.boardStorage.DeleteTask(request)
	if err != nil {
		return err
	}

	return err
}
