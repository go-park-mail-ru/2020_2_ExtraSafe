package auth

import (
	"fmt"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	authStorage "github.com/go-park-mail-ru/2020_2_ExtraSafe/services/auth_servise/internal/userStorage"
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

func NewService(authStorage authStorage.Storage , profileService protoProfile.ProfileClient, boardService protoBoard.BoardClient) *service {
	return &service{
		authStorage: authStorage,
		profileService: profileService,
		boardService: boardService,
	}
}

func (s *service) Auth(ctx context.Context, input *protoProfile.UserID) (output *protoProfile.UserBoardsOutside, err error) {
	user, err := s.profileService.Accounts(ctx, input) // также берет профиль
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
		Links:    user.Links,
		Avatar:   user.Avatar,
		Boards:   boards.Boards,
	}

	return output, err
}

func (s *service) Login(ctx context.Context, input *protoProfile.UserInputLogin) (output *protoProfile.UserID, err error) {
	//err = s.validator.ValidateLogin(request)
	output, err = s.profileService.CheckUser(ctx, input)
	if err != nil {
		return output, err
	}

	return output, err
}

func (s *service) Registration(ctx context.Context, input *protoProfile.UserInputReg) (output *protoProfile.UserID, err error) {
	//err = s.validator.ValidateRegistration(request)

	output, err = s.profileService.CreateUser(ctx, input)
	if err != nil {
		return output, err
	}

	return output, err
}

var (
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func (s *service) SetCookie(userID uint64) (cookieValue string, err error) {
	cookieValue = RandStringRunes(32)

	if err := s.authStorage.CreateUserSession(userID, cookieValue); err != nil {
		fmt.Println(err)
		return cookieValue, err
	}

	return cookieValue, nil
}

func (s *service) DeleteCookie(sessionID string) error {
	if err := s.authStorage.DeleteUserSession(sessionID); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (s *service)CheckCookie(sessionID string) (uint64, error) {
	userId, err := s.authStorage.CheckUserSession(sessionID)
	if err != nil {
		return 0, models.ServeError{Codes: []string{"500"}, OriginalError: err, MethodName: "CheckCookie"}

	}
	return userId, nil
}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

