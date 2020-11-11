package auth

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
)

type Service interface {
	Auth(request models.UserInput) (response models.UserBoardsOutside, err error)
	Login(request models.UserInputLogin) (userID uint64, response models.UserOutside, err error)
	Registration(request models.UserInputReg) (userID uint64, response models.UserOutside, err error)
}

type service struct {
	userStorage  UserStorage
	boardStorage BoardStorage
	validator    Validator
}

func NewService(userStorage UserStorage, boardStorage BoardStorage, validator Validator) Service {
	return &service{
		boardStorage: boardStorage,
		userStorage: userStorage,
		validator: validator,
	}
}

func (s *service) Auth(request models.UserInput) (response models.UserBoardsOutside, err error) {
	user, err := s.userStorage.GetUserProfile(request)
	if err != nil {
		return models.UserBoardsOutside{}, err
	}

	boards, err := s.boardStorage.GetBoardsList(request)
	if err != nil {
		return models.UserBoardsOutside{}, err
	}

	response.Boards = boards
	response.Links = user.Links
	response.Avatar = user.Avatar
	response.FullName = user.FullName
	response.Email = user.Email
	response.Username = user.Username

	return response, err
}

func (s *service) Login(request models.UserInputLogin) (userID uint64, response models.UserOutside, err error) {
	var user models.UserOutside

	err = s.validator.ValidateLogin(request)
	if err != nil {
		return 0, models.UserOutside{}, err
	}

	userID, user, err = s.userStorage.CheckUser(request)
	if err != nil {
		return 0, models.UserOutside{}, err
	}

	return userID, user, err
}


func (s *service) Registration(request models.UserInputReg) (userID uint64, response models.UserOutside, err error) {
	var user models.UserOutside

	err = s.validator.ValidateRegistration(request)
	if err != nil {
		return 0, models.UserOutside{}, err
	}

	userID, user, err = s.userStorage.CreateUser(request)
	if err != nil {
		return 0, models.UserOutside{}, err
	}

	return userID, user, err
}

