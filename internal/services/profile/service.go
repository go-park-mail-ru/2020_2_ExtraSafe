package profile

import (
	"context"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/tools/errorWorker"
	protoProfile "github.com/go-park-mail-ru/2020_2_ExtraSafe/services/proto/profile"
)

type Service interface {
	Profile(request models.UserInput) (user models.UserOutside, err error)
	Boards(request models.UserInput) (boards []models.BoardOutsideShort, err error)
	ProfileChange(request models.UserInputProfile) (user models.UserOutside, err error)
	PasswordChange(request models.UserInputPassword) (user models.UserOutside, err error)
}

type service struct {
	profileService protoProfile.ProfileClient
	validator     Validator
}


func NewService(profileService protoProfile.ProfileClient, validator Validator) Service {
	return &service{
		profileService: profileService,
		validator: validator,
	}
}

func (s *service) Profile(request models.UserInput) (user models.UserOutside, err error) {
	ctx := context.Background()

	input := &protoProfile.UserID{ID: request.ID}

	output, err := s.profileService.Profile(ctx, input)
	if err != nil {
		return models.UserOutside{}, errorWorker.ConvertStatusToError(err)
	}

	user.Username = output.Username
	user.Email = output.Email
	user.FullName = output.FullName
	user.Avatar = output.Avatar

	return user, nil
}

func (s *service) Boards(request models.UserInput) (boards []models.BoardOutsideShort, err error) {
	ctx := context.Background()

	input := &protoProfile.UserID{ID: request.ID}

	output, err := s.profileService.Boards(ctx, input)
	if err != nil {
		return nil, errorWorker.ConvertStatusToError(err)
	}

	boards = make([]models.BoardOutsideShort, 0)
	for _, board := range output.Boards{
		boards = append(boards, models.BoardOutsideShort{
			BoardID: board.BoardID,
			Name:    board.Name,
			Theme:   board.Theme,
			Star:    board.Star,
		})
	}

	return boards, nil
}

func (s *service) ProfileChange(request models.UserInputProfile) (user models.UserOutside, err error) {
	ctx := context.Background()

	err = s.validator.ValidateProfile(request)
	if err != nil {
		return models.UserOutside{}, err
	}

	input := &protoProfile.UserInputProfile{
		ID:       request.ID,
		Email:    request.Email,
		Username: request.Username,
		FullName: request.FullName,
		Avatar:   request.Avatar,
	}

	output, err := s.profileService.ProfileChange(ctx, input)
	if err != nil {
		return models.UserOutside{}, errorWorker.ConvertStatusToError(err)
	}

	user.Username = output.Username
	user.Email = output.Email
	user.FullName = output.FullName
	user.Avatar = output.Avatar

	return user, nil
}

func (s *service) PasswordChange(request models.UserInputPassword) (user models.UserOutside, err error) {
	ctx := context.Background()
	err = s.validator.ValidateChangePassword(request)
	if err != nil {
		return models.UserOutside{}, err
	}

	input := &protoProfile.UserInputPassword{
		ID:          request.ID,
		OldPassword: request.OldPassword,
		Password:    request.Password,
	}

	output, err := s.profileService.PasswordChange(ctx, input)
	if err != nil {
		return models.UserOutside{}, errorWorker.ConvertStatusToError(err)
	}

	user.Username = output.Username
	user.Email = output.Email
	user.FullName = output.FullName
	user.Avatar = output.Avatar

	return user, nil
}
