package profile

import (
	"fmt"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/errorWorker"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/profile_service/internal/imgStorage"
	uStorage "github.com/go-park-mail-ru/2020_2_ExtraSafe/services/profile_service/internal/userStorage"
	protoBoard "github.com/go-park-mail-ru/2020_2_ExtraSafe/services/proto/board"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/proto/profile"
	"golang.org/x/net/context"
)

type Service interface {
	CheckUser(_ context.Context, input *protoProfile.UserInputLogin) (output *protoProfile.UserID, err error)
	CreateUser(_ context.Context, input *protoProfile.UserInputReg) (output *protoProfile.UserID, err error)
	Profile(_ context.Context, input *protoProfile.UserID) (output *protoProfile.UserOutside, err error)
	Boards(c context.Context, input *protoProfile.UserID) (output *protoProfile.BoardsOutsideShort, err error)
	ProfileChange(_ context.Context, input *protoProfile.UserInputProfile) (output *protoProfile.UserOutside, err error)
	PasswordChange(_ context.Context, input *protoProfile.UserInputPassword) (output *protoProfile.UserOutside, err error)
	GetUsersByIDs(_ context.Context, input *protoProfile.UserIDS) (output *protoProfile.UsersOutsideShort, err error)
	GetUserByUsername(_ context.Context, input *protoProfile.UserName) (output *protoProfile.UserOutsideShort, err error)
}
type service struct {
	userStorage uStorage.Storage
	avatarStorage imgStorage.Storage
	boardService protoBoard.BoardClient
}

var ServiceName = "ProfileService"

func NewService(userStorage uStorage.Storage, avatarStorage imgStorage.Storage, boardService protoBoard.BoardClient) *service {
	return &service{
		userStorage: userStorage,
		avatarStorage: avatarStorage,
		boardService: boardService,
	}
}

func (s *service) CheckUser(_ context.Context, input *protoProfile.UserInputLogin) (output *protoProfile.UserID, err error) {
	userInput := models.UserInputLogin{
		Email:    input.Email,
		Password: input.Password,
	}

	userID, _, err :=  s.userStorage.CheckUser(userInput)
	if err != nil {
		return nil, errorWorker.ConvertErrorToStatus(err.(models.ServeError), ServiceName)
	}

	output = &protoProfile.UserID{ID: userID}

	return output, nil
}

func (s *service) CreateUser(_ context.Context, input *protoProfile.UserInputReg) (output *protoProfile.UserID, err error) {
	userInput := models.UserInputReg{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
	}

	userID, _, err :=  s.userStorage.CreateUser(userInput)
	if err != nil {
		return output, errorWorker.ConvertErrorToStatus(err.(models.ServeError), ServiceName)
	}

	output = &protoProfile.UserID{ID: userID}

	return output, nil
}

func (s *service) Profile(_ context.Context, input *protoProfile.UserID) (output *protoProfile.UserOutside, err error) {
	userInput := models.UserInput{
		ID: input.ID,
	}

	user, err := s.userStorage.GetUserProfile(userInput)
	if err != nil {
		return output, errorWorker.ConvertErrorToStatus(err.(models.ServeError), ServiceName)
	}

	output = &protoProfile.UserOutside{
		Email:    user.Email,
		Username: user.Username,
		FullName: user.FullName,
		Avatar:   user.Avatar,
	}

	return output, nil
}

func (s *service) Boards(c context.Context, input *protoProfile.UserID) (output *protoProfile.BoardsOutsideShort, err error) {

	output, err = s.boardService.GetBoardsList(c, input)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (s *service) ProfileChange(_ context.Context, input *protoProfile.UserInputProfile) (output *protoProfile.UserOutside, err error) {
	multiErrors := new(models.MultiErrors)

	userInput := models.UserInputProfile{
		ID:       input.ID,
		Email:    input.Email,
		Username: input.Username,
		FullName: input.FullName,
		Avatar:   input.Avatar,
	}

	// TODO - работа с аватаром
	userAvatar, errGetAvatar := s.userStorage.GetUserAvatar(models.UserInput{ID: input.ID})
	if errGetAvatar != nil {
		return output, errGetAvatar
	}

	if input.Avatar != nil {
		errAvatar := s.avatarStorage.UploadAvatar(input.Avatar, &userAvatar, false)
		if errAvatar != nil {
			multiErrors.Codes = append(multiErrors.Codes, errAvatar.(models.ServeError).Codes...)
			multiErrors.Descriptions = append(multiErrors.Descriptions, errAvatar.(models.ServeError).Descriptions...)
		}
	}

	user, errProfile := s.userStorage.ChangeUserProfile(userInput, userAvatar)
	if errProfile != nil {
		if errProfile.(models.ServeError).Codes[0] == models.ServerError {
			return output, errProfile
		}
		multiErrors.Codes = append(multiErrors.Codes, errProfile.(models.ServeError).Codes...)
		multiErrors.Descriptions = append(multiErrors.Descriptions, errProfile.(models.ServeError).Descriptions...)
	}

	if len(multiErrors.Codes) != 0 {
		return nil, errorWorker.ConvertErrorToStatus(models.ServeError{Codes: multiErrors.Codes,
			Descriptions: multiErrors.Descriptions, MethodName: "ProfileChange"}, ServiceName)
	}

	output = &protoProfile.UserOutside{
		Email:    user.Email,
		Username: user.Username,
		FullName: user.FullName,
		Avatar:   user.Avatar,
	}

	return output, nil
}

func (s *service) PasswordChange(_ context.Context, input *protoProfile.UserInputPassword) (output *protoProfile.UserOutside, err error) {
	userInput := models.UserInputPassword{
		ID:          input.ID,
		OldPassword: input.OldPassword,
		Password:    input.Password,
	}

	user, err := s.userStorage.ChangeUserPassword(userInput)
	if err != nil {
		return output, errorWorker.ConvertErrorToStatus(err.(models.ServeError), ServiceName)
	}

	output = &protoProfile.UserOutside{
		Email:    user.Email,
		Username: user.Username,
		FullName: user.FullName,
		Avatar:   user.Avatar,
	}

	return output, nil
}

func (s *service) GetUsersByIDs(_ context.Context, input *protoProfile.UserIDS) (output *protoProfile.UsersOutsideShort, err error) {
	userIDS := make([]int64, 0)

	for _, id := range input.UserIDs {
		userIDS = append(userIDS, id)
	}

	users, err := s.userStorage.GetUsersByIDs(userIDS)
	if err != nil {
		return output, errorWorker.ConvertErrorToStatus(err.(models.ServeError), ServiceName)
	}

	output = &protoProfile.UsersOutsideShort{Users: nil}

	for _, userShort := range users {
		output.Users = append(output.Users, &protoProfile.UserOutsideShort{
			ID: userShort.ID,
			Email:    userShort.Email,
			Username: userShort.Username,
			FullName: userShort.FullName,
			Avatar:   userShort.Avatar,
		})
	}

	return output, nil
}

func (s *service) GetUserByUsername(_ context.Context, input *protoProfile.UserName) (output *protoProfile.UserOutsideShort, err error) {
	fmt.Println(input)
	user, err := s.userStorage.GetUserByUsername(input.UserName)
	if err != nil {
		return output, errorWorker.ConvertErrorToStatus(err.(models.ServeError), ServiceName)
	}

	output = &protoProfile.UserOutsideShort{
		ID: user.ID,
		Email:    user.Email,
		Username: user.Username,
		FullName: user.FullName,
		Avatar:   user.Avatar,
	}

	return output, nil
}