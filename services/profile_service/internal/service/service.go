package profile

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/profile_service/internal/storage"
	protoBoard "github.com/go-park-mail-ru/2020_2_ExtraSafe/services/proto/board"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/proto/profile"
	"golang.org/x/net/context"
)

type service struct {
	userStorage storage.Storage
	boardService protoBoard.BoardClient
}


func NewService(userStorage storage.Storage, boardService protoBoard.BoardClient) *service {
	return &service{
		userStorage: userStorage,
		boardService: boardService,
	}
}

func (s *service) CheckUser(c context.Context, input *protoProfile.UserInputLogin) (output *protoProfile.UserID, err error) {
	userInput := models.UserInputLogin{
		Email:    input.Email,
		Password: input.Password,
	}

	userID, _, err :=  s.userStorage.CheckUser(userInput)
	if err != nil {
		return output, err
	}

	output = &protoProfile.UserID{ID: userID}

	return output, nil
}

func (s *service) CreateUser(c context.Context, input *protoProfile.UserInputReg) (output *protoProfile.UserID, err error) {
	userInput := models.UserInputReg{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
	}

	userID, _, err :=  s.userStorage.CreateUser(userInput)
	if err != nil {
		return output, err
	}

	output = &protoProfile.UserID{ID: userID}

	return output, nil
}

func (s *service) Profile(c context.Context, input *protoProfile.UserID) (output *protoProfile.UserOutside, err error) {
	userInput := models.UserInput{
		ID: input.ID,
	}

	user, err := s.userStorage.GetUserProfile(userInput)
	if err != nil {
		return output, err
	}

	lincks := &protoProfile.UserLinks{
		Instagram: user.Links.Instagram,
		Github:    user.Links.Github,
		Bitbucket: user.Links.Bitbucket,
		Vk:        user.Links.Vk,
		Facebook:  user.Links.Facebook,
	}

	output = &protoProfile.UserOutside{
		Email:    user.Email,
		Username: user.Username,
		FullName: user.FullName,
		Links:    lincks,
		Avatar:   user.Avatar,
	}

	return output, err
}

func (s *service) Accounts(c context.Context, input *protoProfile.UserID) (output *protoProfile.UserOutside, err error) {
	userInput := models.UserInput{
		ID: input.ID,
	}

	user, err := s.userStorage.GetUserAccounts(userInput)
	if err != nil {
		return output, err
	}

	// пропущены ссылки
	output = &protoProfile.UserOutside{
		Email:    user.Email,
		Username: user.Username,
		FullName: user.FullName,
		Avatar:   user.Avatar,
	}

	return output, err
}

func (s *service) Boards(c context.Context, input *protoProfile.UserID) (output *protoProfile.BoardsOutsideShort, err error) {

	output, err = s.boardService.GetBoardsList(c, input)
	if err != nil {
		return nil, err
	}

	return output, err
}

func (s *service) ProfileChange(c context.Context, input *protoProfile.UserInputProfile) (output *protoProfile.UserOutside, err error) {
	multiErrors := new(models.MultiErrors)

	userInput := models.UserInputProfile{
		ID:       input.ID,
		Email:    input.Email,
		Username: input.Username,
		FullName: input.Username,
		Avatar:   nil,
	}

	/*err = s.validator.ValidateProfile(request)
	if err != nil {
		return nil, err
	}*/
	// TODO - работа с аватаром
	/*userAvatar, errGetAvatar := s.userStorage.GetUserAvatar(models.UserInput{ID: request.ID})
	if errGetAvatar != nil {
		return user, errGetAvatar
	}

	if request.Avatar != nil {
		errAvatar := s.avatarStorage.UploadAvatar(request.Avatar, &userAvatar)
		if errAvatar != nil {
			multiErrors.Codes = append(multiErrors.Codes, errAvatar.(models.ServeError).Codes...)
			multiErrors.Descriptions = append(multiErrors.Descriptions, errAvatar.(models.ServeError).Descriptions...)
		}
	}*/

	userAvatar := models.UserAvatar{
		ID:     input.ID,
		Avatar: "",
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
		return nil, models.ServeError{Codes: multiErrors.Codes, Descriptions: multiErrors.Descriptions,
			MethodName: "ProfileChange"}
	}

	lincks := &protoProfile.UserLinks{
		Instagram: user.Links.Instagram,
		Github:    user.Links.Github,
		Bitbucket: user.Links.Bitbucket,
		Vk:        user.Links.Vk,
		Facebook:  user.Links.Facebook,
	}

	output = &protoProfile.UserOutside{
		Email:    user.Email,
		Username: user.Username,
		FullName: user.FullName,
		Links: lincks,
		Avatar:   user.Avatar,
	}

	return output, err
}

func (s *service) AccountsChange(c context.Context, input *protoProfile.UserInputLinks) (output *protoProfile.UserOutside, err error) {
	/*err = s.validator.ValidateLinks(request)
	if err != nil {
		return output, err
	}*/

	userInput := models.UserInputLinks{
		ID:        input.ID,
		Telegram:  "",
		Instagram: input.Instagram,
		Github:    input.Github,
		Bitbucket: input.Bitbucket,
		Vk:        input.Vk,
		Facebook:  input.Facebook,
	}

	user, err := s.userStorage.ChangeUserAccounts(userInput)
	if err != nil {
		return output, err
	}

	lincks := &protoProfile.UserLinks{
		Instagram: user.Links.Instagram,
		Github:    user.Links.Github,
		Bitbucket: user.Links.Bitbucket,
		Vk:        user.Links.Vk,
		Facebook:  user.Links.Facebook,
	}

	output = &protoProfile.UserOutside{
		Email:    user.Email,
		Username: user.Username,
		FullName: user.FullName,
		Links: lincks,
		Avatar:   user.Avatar,
	}

	return output, err
}

func (s *service) PasswordChange(c context.Context, input *protoProfile.UserInputPassword) (output *protoProfile.UserOutside, err error) {
	/*err = s.validator.ValidateChangePassword(request)
	if err != nil {
		return output, err
	}*/
	userInput := models.UserInputPassword{
		ID:          input.ID,
		OldPassword: input.OldPassword,
		Password:    input.Password,
	}

	user, err := s.userStorage.ChangeUserPassword(userInput)
	if err != nil {
		return output, err
	}

	lincks := &protoProfile.UserLinks{
		Instagram: user.Links.Instagram,
		Github:    user.Links.Github,
		Bitbucket: user.Links.Bitbucket,
		Vk:        user.Links.Vk,
		Facebook:  user.Links.Facebook,
	}

	output = &protoProfile.UserOutside{
		Email:    user.Email,
		Username: user.Username,
		FullName: user.FullName,
		Links: lincks,
		Avatar:   user.Avatar,
	}

	return output, err
}

