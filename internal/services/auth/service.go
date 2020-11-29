package auth

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/errorWorker"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/proto/auth"
	protoProfile "github.com/go-park-mail-ru/2020_2_ExtraSafe/services/proto/profile"
	"github.com/labstack/echo"
)

type Service interface {
	CheckCookie(c echo.Context) (int64, error)
	Logout(c echo.Context) (err error)
	Auth(request models.UserInput) (response models.UserBoardsOutside, err error)
	Login(request models.UserInputLogin) (userSession models.UserSession, err error)
	Registration(request models.UserInputReg) (userSession models.UserSession, err error)
}

type service struct {
	authService protoAuth.AuthClient
	validator   Validator
}

func NewService(authService protoAuth.AuthClient, validator Validator) Service {
	return &service{
		authService: authService,
		validator: validator,
	}
}

func (s *service)CheckCookie(c echo.Context) (int64, error) {
	session, err := c.Cookie("tabutask_id")
	if err != nil {
		fmt.Println(err)
		return -1, models.ServeError{Codes: []string{"500"}, OriginalError: err, MethodName: "CheckCookie"}
	}
	ctx := context.Background()

	input := &protoAuth.UserSession{
		SessionID: session.Value,
		UserID:    -1,
	}

	user, err := s.authService.CheckCookie(ctx, input)
	if err != nil {
		return -1, errorWorker.ConvertStatusToError(err)
	}
	return user.ID, nil
}

func (s *service) Auth(request models.UserInput) (response models.UserBoardsOutside, err error) {
	ctx := context.Background()
	input := &protoProfile.UserID{ID: request.ID}

	user, err := s.authService.Auth(ctx, input)
	if err != nil {
		return models.UserBoardsOutside{}, errorWorker.ConvertStatusToError(err)
	}

	boards := make([]models.BoardOutsideShort, 0)
	for _, board := range user.Boards {
		boards = append(boards, models.BoardOutsideShort{
			BoardID: board.BoardID,
			Name:    board.Name,
			Theme:   board.Theme,
			Star:    board.Star,
		})
	}

	response.Boards = boards
	response.Avatar = user.Avatar
	response.FullName = user.FullName
	response.Email = user.Email
	response.Username = user.Username

	return response, err
}

func (s *service) Login(request models.UserInputLogin) (userSession models.UserSession, err error) {
	ctx := context.Background()

	err = s.validator.ValidateLogin(request)
	if err != nil {
		return models.UserSession{}, err
	}

	input := &protoProfile.UserInputLogin{
		Email:    request.Email,
		Password: request.Password,
	}

	user, err := s.authService.Login(ctx, input)
	if err != nil {
		return models.UserSession{}, errorWorker.ConvertStatusToError(err)
	}

	userSession.UserID = user.UserID
	userSession.SessionID = user.SessionID

	return userSession, nil
}


func (s *service) Registration(request models.UserInputReg) (userSession models.UserSession, err error) {
	ctx := context.Background()

	err = s.validator.ValidateRegistration(request)
	if err != nil {
		return models.UserSession{}, err
	}

	input := &protoProfile.UserInputReg{
		Email:    request.Email,
		Username: request.Username,
		Password: request.Password,
	}

	user, err := s.authService.Registration(ctx, input)
	if err != nil {
		return models.UserSession{}, errorWorker.ConvertStatusToError(err)
	}

	userSession.UserID = user.UserID
	userSession.SessionID = user.SessionID

	return userSession, nil
}

func (s *service) Logout(c echo.Context) (err error) {
	session, err := c.Cookie("tabutask_id")
	if err != nil {
		fmt.Println(err)
		return models.ServeError{Codes: []string{"500"}, OriginalError: err, MethodName: "CheckCookie"}
	}
	ctx := context.Background()

	input := &protoAuth.UserSession{
		SessionID: session.Value,
		UserID:    -1,
	}

	_, err = s.authService.DeleteCookie(ctx, input)
	if err != nil {
		return err
	}
	return nil
}

