package auth

import (
	"fmt"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/errorWorker"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	_ "github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	authStorage "github.com/go-park-mail-ru/2020_2_ExtraSafe/services/auth_servise/internal/sessionsStorage"
	protoAuth "github.com/go-park-mail-ru/2020_2_ExtraSafe/services/proto/auth"
	protoBoard "github.com/go-park-mail-ru/2020_2_ExtraSafe/services/proto/board"
	protoProfile "github.com/go-park-mail-ru/2020_2_ExtraSafe/services/proto/profile"
	"golang.org/x/net/context"
	"math/rand"
)

type service struct {
	authStorage authStorage.Storage
	profileService protoProfile.ProfileClient
	boardService protoBoard.BoardClient
}

var ServiceName = "AuthService"

func NewService(authStorage authStorage.Storage , profileService protoProfile.ProfileClient, boardService protoBoard.BoardClient) *service {
	return &service{
		authStorage: authStorage,
		profileService: profileService,
		boardService: boardService,
	}
}

func (s *service) Auth(ctx context.Context, input *protoProfile.UserID) (output *protoProfile.UserBoardsOutside, err error) {
	user, err := s.profileService.Profile(ctx, input)
	if err != nil {
		return output, err
	}

	boards, err := s.boardService.GetBoardsList(ctx, input)
	if err != nil {
		return output, err
	}

	output = &protoProfile.UserBoardsOutside{
		Email:    user.Email,
		Username: user.Username,
		FullName: user.FullName,
		Avatar:   user.Avatar,
		Boards:   boards.Boards,
	}

	return output, nil
}

func (s *service) Login(ctx context.Context, input *protoProfile.UserInputLogin) (output *protoAuth.UserSession, err error) {
	user, err := s.profileService.CheckUser(ctx, input)
	if err != nil {
		return output, err
	}

	cookieValue, err := s.SetCookie(user.ID)
	if err != nil {
		return output, err
	}

	output = &protoAuth.UserSession{SessionID: cookieValue, UserID: user.ID}

	return output, nil
}

func (s *service) Registration(ctx context.Context, input *protoProfile.UserInputReg) (output *protoAuth.UserSession, err error) {
	user, err := s.profileService.CreateUser(ctx, input)
	if err != nil {
		return output, err
	}

	cookieValue, err := s.SetCookie(user.ID)
	if err != nil {
		return output, err
	}

	output = &protoAuth.UserSession{SessionID: cookieValue, UserID: user.ID}

	return output, nil
}

func (s *service)CheckCookie(_ context.Context, input *protoAuth.UserSession) (output *protoProfile.UserID, err error) {
	userId, err := s.authStorage.CheckUserSession(input.SessionID)
	if err != nil {
		return output, errorWorker.ConvertErrorToStatus(err.(models.ServeError), ServiceName)
	}

	output = &protoProfile.UserID{ID: userId}

	return output, nil
}

func (s *service) DeleteCookie(_ context.Context, input *protoAuth.UserSession) (output *protoAuth.Nothing, err error) {
	if err := s.authStorage.DeleteUserSession(input.SessionID); err != nil {
		return &protoAuth.Nothing{Dummy: true}, errorWorker.ConvertErrorToStatus(err.(models.ServeError), ServiceName)
	}

	return &protoAuth.Nothing{Dummy: true}, nil
}

var (
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func (s *service) SetCookie(userID int64) (cookieValue string, err error) {
	cookieValue = RandStringRunes(32)
	fmt.Println(cookieValue)

	if err := s.authStorage.CreateUserSession(userID, cookieValue); err != nil {
		fmt.Println(err)
		return cookieValue, errorWorker.ConvertErrorToStatus(err.(models.ServeError), ServiceName)
	}

	return cookieValue, nil
}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

