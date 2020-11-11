package profile

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
)

type Service interface {
	Profile(request models.UserInput) (user models.UserOutside, err error)
	Accounts(request models.UserInput) (user models.UserOutside, err error)
	Boards(request models.UserInput) (boards []models.BoardOutsideShort, err error)
	ProfileChange(request models.UserInputProfile) (user models.UserOutside, err error)
	AccountsChange(request models.UserInputLinks) (user models.UserOutside, err error)
	PasswordChange(request models.UserInputPassword) (user models.UserOutside, err error)
}

type service struct {
	userStorage userStorage
	avatarStorage avatarStorage
	boardStorage boardStorage
	validator validator
}

func NewService(userStorage userStorage, avatarStorage avatarStorage, boardStorage boardStorage,
	validator validator) Service {
	return &service{
		userStorage: userStorage,
		avatarStorage: avatarStorage,
		boardStorage: boardStorage,
		validator: validator,
	}
}

func (s *service) Profile(request models.UserInput) (user models.UserOutside, err error) {
	user, err = s.userStorage.GetUserProfile(request)
	if err != nil {
		return models.UserOutside{}, err
	}

	return user, err
}

func (s *service) Accounts(request models.UserInput) (user models.UserOutside, err error) {
	user, err = s.userStorage.GetUserAccounts(request)
	if err != nil {
		return models.UserOutside{}, err
	}

	return user, err
}

func (s *service) Boards(request models.UserInput) (boards []models.BoardOutsideShort, err error) {
	boards, err = s.boardStorage.GetBoardsList(request)
	if err != nil {
		return nil, err
	}

	return boards, err
}

func (s *service) ProfileChange(request models.UserInputProfile) (user models.UserOutside, err error) {
	multiErrors := new(models.MultiErrors)

	err = s.validator.ValidateProfile(request)
	if err != nil {
		return models.UserOutside{}, err
	}

	userAvatar, errGetAvatar := s.userStorage.GetUserAvatar(models.UserInput{ID: request.ID})
	if errGetAvatar != nil {
		return user, errGetAvatar
	}

	if request.Avatar != nil {
		errAvatar := s.avatarStorage.UploadAvatar(request.Avatar, &userAvatar)
		if errAvatar != nil {
			multiErrors.Codes = append(multiErrors.Codes, errAvatar.(models.ServeError).Codes...)
			multiErrors.Descriptions = append(multiErrors.Descriptions, errAvatar.(models.ServeError).Descriptions...)
		}
	}

	user, errProfile := s.userStorage.ChangeUserProfile(request, userAvatar)
	if errProfile != nil {
		if errProfile.(models.ServeError).Codes[0] == models.ServerError {
			return user, errProfile
		}
		multiErrors.Codes = append(multiErrors.Codes, errProfile.(models.ServeError).Codes...)
		multiErrors.Descriptions = append(multiErrors.Descriptions, errProfile.(models.ServeError).Descriptions...)
	}

	if len(multiErrors.Codes) != 0 {
		return models.UserOutside{}, models.ServeError{Codes: multiErrors.Codes, Descriptions: multiErrors.Descriptions,
			MethodName: "ProfileChange"}
	}

	return user, err
}

func (s *service) AccountsChange(request models.UserInputLinks) (user models.UserOutside, err error) {
	err = s.validator.ValidateLinks(request)
	if err != nil {
		return models.UserOutside{}, err
	}

	user, err = s.userStorage.ChangeUserAccounts(request)
	if err != nil {
		return models.UserOutside{}, err
	}

	return user, err
}

func (s *service) PasswordChange(request models.UserInputPassword) (user models.UserOutside, err error) {
	err = s.validator.ValidateChangePassword(request)
	if err != nil {
		return models.UserOutside{}, err
	}

	user, err = s.userStorage.ChangeUserPassword(request)
	if err != nil {
		return models.UserOutside{}, err
	}

	return user, err
}
