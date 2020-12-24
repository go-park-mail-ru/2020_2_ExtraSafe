package profile

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	mockBoards "github.com/go-park-mail-ru/2020_2_ExtraSafe/services/mock"
	//	"mime/multipart"
	"reflect"

	mocks "github.com/go-park-mail-ru/2020_2_ExtraSafe/services/profile_service/internal/service/mock"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/proto/profile"
	"github.com/golang/mock/gomock"
	//	"mime/multipart"
	"testing"
)

func TestService_CheckUser(t *testing.T) {
	request := &protoProfile.UserInputLogin{
		Email:    "epridius@yandex.ru",
		Password: "lalala",
	}

	userInput := models.UserInputLogin{
		Email:    request.Email,
		Password: request.Password,
	}

	expectedUser := &protoProfile.UserID{ID: 1}

	ctrlUser := gomock.NewController(t)
	defer ctrlUser.Finish()
	mockUserStorage := mocks.NewMockUserStorage(ctrlUser)

	service := &service{
		userStorage: mockUserStorage,
	}

	mockUserStorage.EXPECT().CheckUser(userInput).Return(int64(1), models.UserOutside{}, nil)

	user, err := service.CheckUser(context.Background(), request)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(user, expectedUser) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedUser, user)
		return
	}
}

func TestService_CheckUserFail(t *testing.T) {
	request := &protoProfile.UserInputLogin{
		Email:    "epridius@yandex.ru",
		Password: "lalala",
	}

	userInput := models.UserInputLogin{
		Email:    request.Email,
		Password: request.Password,
	}

	ctrlUser := gomock.NewController(t)
	defer ctrlUser.Finish()
	mockUserStorage := mocks.NewMockUserStorage(ctrlUser)

	service := &service{
		userStorage: mockUserStorage,
	}

	errStorage := models.ServeError{Codes: []string{"500"}, OriginalError: errors.New("")}

	mockUserStorage.EXPECT().CheckUser(userInput).Return(int64(0), models.UserOutside{}, errStorage)

	_, err := service.CheckUser(context.Background(), request)
	if err == nil {
		t.Error("expected error")
		return
	}
}

func TestService_CreateUser(t *testing.T) {
	input := &protoProfile.UserInputReg{
		Email:    "epridius@yandex.ru",
		Username: "pkaterinaa",
		Password: "lalala",
	}

	userInput := models.UserInputReg{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
	}

	expectedUser := &protoProfile.UserID{ID: 1}

	ctrlUser := gomock.NewController(t)
	defer ctrlUser.Finish()
	mockUserStorage := mocks.NewMockUserStorage(ctrlUser)

	service := &service{
		userStorage: mockUserStorage,
	}

	mockUserStorage.EXPECT().CreateUser(userInput).Return(int64(1), models.UserOutside{}, nil)

	user, err := service.CreateUser(context.Background(), input)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(user, expectedUser) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedUser, user)
		return
	}
}

func TestService_CreateUserFail(t *testing.T) {
	input := &protoProfile.UserInputReg{
		Email:    "epridius@yandex.ru",
		Username: "pkaterinaa",
		Password: "lalala",
	}

	userInput := models.UserInputReg{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
	}

	ctrlUser := gomock.NewController(t)
	defer ctrlUser.Finish()
	mockUserStorage := mocks.NewMockUserStorage(ctrlUser)

	service := &service{
		userStorage: mockUserStorage,
	}
	errStorage := models.ServeError{Codes: []string{"500"}, OriginalError: errors.New("")}

	mockUserStorage.EXPECT().CreateUser(userInput).Return(int64(1), models.UserOutside{}, errStorage)

	_, err := service.CreateUser(context.Background(), input)
	if err == nil {
		t.Error("expected error")
		return
	}
}

func TestService_GetUserByUsername(t *testing.T) {
	input := &protoProfile.UserName{UserName: "pkaterinaa"}

	expect := &protoProfile.UserOutsideShort{
		ID:       1,
		Email:    "epridius@yandex.ru",
		Username: "pkaterinaa",
		FullName: "",
		Avatar:   "default/default_avatar.png",
	}

	internal := models.UserInternal{
		ID:       expect.ID,
		Email:    expect.Email,
		Username: expect.Username,
		FullName: expect.FullName,
		Avatar:   expect.Avatar,
	}

	ctrlUser := gomock.NewController(t)
	defer ctrlUser.Finish()
	mockUserStorage := mocks.NewMockUserStorage(ctrlUser)

	service := &service{
		userStorage: mockUserStorage,
	}

	mockUserStorage.EXPECT().GetUserByUsername(input.UserName).Return(internal, nil)
	output, err := service.GetUserByUsername(context.Background(), input)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(output, expect) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expect, output)
		return
	}
}

func TestService_GetUserByUsernameFail(t *testing.T) {
	input := &protoProfile.UserName{UserName: "pkaterinaa"}

	ctrlUser := gomock.NewController(t)
	defer ctrlUser.Finish()
	mockUserStorage := mocks.NewMockUserStorage(ctrlUser)

	service := &service{
		userStorage: mockUserStorage,
	}
	errStorage := models.ServeError{Codes: []string{"500"}, OriginalError: errors.New("")}

	mockUserStorage.EXPECT().GetUserByUsername(input.UserName).Return(models.UserInternal{}, errStorage)
	_, err := service.GetUserByUsername(context.Background(), input)
	if err == nil {
		t.Error("expected error")
		return
	}
}

func TestService_GetUsersByIDs(t *testing.T) {
	userIDs := []int64{1}
	input := &protoProfile.UserIDS{UserIDs: userIDs}

	user := models.UserOutsideShort{
		Email:    "epridius@yandex.ru",
		Username: "pkaterinaa",
		FullName: "",
		Avatar:   "default/default_avatar.png",
	}
	internal := []models.UserOutsideShort{
		user,
	}

	expect := &protoProfile.UsersOutsideShort{Users: nil}
	expect.Users = append(expect.Users, &protoProfile.UserOutsideShort{
		Email:    user.Email,
		Username: user.Username,
		FullName: user.FullName,
		Avatar:   user.Avatar,
	})

	ctrlUser := gomock.NewController(t)
	defer ctrlUser.Finish()
	mockUserStorage := mocks.NewMockUserStorage(ctrlUser)

	service := &service{
		userStorage: mockUserStorage,
	}

	mockUserStorage.EXPECT().GetUsersByIDs(userIDs).Return(internal, nil)
	output, err := service.GetUsersByIDs(context.Background(), input)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(expect, output) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expect, output)
		return
	}
}

func TestService_GetUsersByIDsFail(t *testing.T) {
	userIDs := []int64{1}
	input := &protoProfile.UserIDS{UserIDs: userIDs}

	ctrlUser := gomock.NewController(t)
	defer ctrlUser.Finish()
	mockUserStorage := mocks.NewMockUserStorage(ctrlUser)

	service := &service{
		userStorage: mockUserStorage,
	}

	errStorage := models.ServeError{Codes: []string{"500"}, OriginalError: errors.New("")}

	mockUserStorage.EXPECT().GetUsersByIDs(userIDs).Return([]models.UserOutsideShort{}, errStorage)
	_, err := service.GetUsersByIDs(context.Background(), input)
	if err == nil {
		t.Error("expected error")
		return
	}
}

func TestService_Profile(t *testing.T) {
	ctrlUser := gomock.NewController(t)
	defer ctrlUser.Finish()
	mockUserStorage := mocks.NewMockUserStorage(ctrlUser)

	service := &service{
		userStorage: mockUserStorage,
	}

	userInput := models.UserInput{ID: 1}
	input := &protoProfile.UserID{ID: userInput.ID}

	internal := models.UserOutside{
		Email:    "epridius",
		Username: "pkaterinaa",
		FullName: "",
		Avatar:   "default/default_avatar.png",
	}

	expect := &protoProfile.UserOutside{
		Email:    internal.Email,
		Username: internal.Username,
		FullName: internal.FullName,
		Avatar:   internal.Avatar,
	}

	mockUserStorage.EXPECT().GetUserProfile(userInput).Return(internal, nil)

	user, err := service.Profile(context.Background(), input)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(user, expect) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expect, user)
		return
	}
}

func TestService_ProfileFail(t *testing.T) {
	ctrlUser := gomock.NewController(t)
	defer ctrlUser.Finish()
	mockUserStorage := mocks.NewMockUserStorage(ctrlUser)

	service := &service{
		userStorage: mockUserStorage,
	}

	userInput := models.UserInput{ID: 1}
	input := &protoProfile.UserID{ID: userInput.ID}

	internal := models.UserOutside{
		Email:    "epridius",
		Username: "pkaterinaa",
		FullName: "",
		Avatar:   "default/default_avatar.png",
	}
	errStorage := models.ServeError{Codes: []string{"500"}, OriginalError: errors.New("")}

	mockUserStorage.EXPECT().GetUserProfile(userInput).Return(internal, errStorage)

	_, err := service.Profile(context.Background(), input)
	if err == nil {
		t.Error("expected error")
		return
	}
}

func TestService_ProfileChange(t *testing.T) {
	ctrlUser := gomock.NewController(t)
	defer ctrlUser.Finish()
	mockUserStorage := mocks.NewMockUserStorage(ctrlUser)

	ctrlAvatar := gomock.NewController(t)
	defer ctrlAvatar.Finish()
	mockAvatarStorage := mocks.NewMockAvatarStorage(ctrlAvatar)

	service := &service{
		userStorage:   mockUserStorage,
		avatarStorage: mockAvatarStorage,
	}

	request := &protoProfile.UserInputProfile{
		ID:       1,
		Email:    "epridius",
		Username: "pkaterinaa",
		FullName: "",
		Avatar:   []byte{},
	}

	input := models.UserInputProfile{
		ID:       request.ID,
		Email:    request.Email,
		Username: request.Username,
		FullName: request.FullName,
		Avatar:   request.Avatar,
	}
	userAvatar := models.UserAvatar{
		ID:     1,
		Avatar: "default/default_avatar.png",
	}

	internal := models.UserOutside{
		Email:    request.Email,
		Username: request.Username,
		FullName: request.FullName,
		Avatar:   "default/default_avatar.png",
	}

	expect := &protoProfile.UserOutside{
		Email:    request.Email,
		Username: request.Username,
		FullName: request.FullName,
		Avatar:   internal.Avatar,
	}

	mockUserStorage.EXPECT().GetUserAvatar(models.UserInput{ID: request.ID}).Return(userAvatar, nil)
	mockAvatarStorage.EXPECT().UploadAvatar(request.Avatar, &userAvatar, false).Return(nil)
	mockUserStorage.EXPECT().ChangeUserProfile(input, userAvatar).Return(internal, nil)

	output, err := service.ProfileChange(context.Background(), request)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(output, expect) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expect, output)
		return
	}
}

func TestService_ProfileChangeAvatarFail(t *testing.T) {
	ctrlUser := gomock.NewController(t)
	defer ctrlUser.Finish()
	mockUserStorage := mocks.NewMockUserStorage(ctrlUser)

	ctrlAvatar := gomock.NewController(t)
	defer ctrlAvatar.Finish()
	mockAvatarStorage := mocks.NewMockAvatarStorage(ctrlAvatar)

	service := &service{
		userStorage:   mockUserStorage,
		avatarStorage: mockAvatarStorage,
	}

	request := &protoProfile.UserInputProfile{
		ID:       1,
		Email:    "epridius",
		Username: "pkaterinaa",
		FullName: "",
		Avatar:   []byte{},
	}

	userAvatar := models.UserAvatar{}
	errStorage := models.ServeError{Codes: []string{"500"}, OriginalError: errors.New("")}

	mockUserStorage.EXPECT().GetUserAvatar(models.UserInput{ID: request.ID}).Return(userAvatar, errStorage)

	_, err := service.ProfileChange(context.Background(), request)
	if err == nil {
		t.Error("expected error")
		return
	}
}

func TestService_ProfileChangeError(t *testing.T) {
	ctrlUser := gomock.NewController(t)
	defer ctrlUser.Finish()
	mockUserStorage := mocks.NewMockUserStorage(ctrlUser)

	ctrlAvatar := gomock.NewController(t)
	defer ctrlAvatar.Finish()
	mockAvatarStorage := mocks.NewMockAvatarStorage(ctrlAvatar)

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardService := mockBoards.NewMockBoardClient(ctrlBoard)

	service := NewService(mockUserStorage, mockAvatarStorage, mockBoardService)

	request := &protoProfile.UserInputProfile{
		ID:       1,
		Email:    "epridius",
		Username: "pkaterinaa",
		FullName: "",
		Avatar:   []byte{},
	}

	input := models.UserInputProfile{
		ID:       request.ID,
		Email:    request.Email,
		Username: request.Username,
		FullName: request.FullName,
		Avatar:   request.Avatar,
	}
	userAvatar := models.UserAvatar{
		ID:     1,
		Avatar: "default/default_avatar.png",
	}

	expectedUser := models.UserOutside{}

	errAvatar := models.ServeError{Codes: []string{"600"}, Descriptions: []string{"File error"}, MethodName: "UploadAvatar"}

	mockUserStorage.EXPECT().GetUserAvatar(models.UserInput{ID: request.ID}).Return(userAvatar, nil)
	mockAvatarStorage.
		EXPECT().
		UploadAvatar(request.Avatar, &userAvatar, false).
		Return(errAvatar)
	mockUserStorage.EXPECT().ChangeUserProfile(input, userAvatar).Return(expectedUser, nil)

	_, err := service.ProfileChange(context.Background(), request)

	if err == nil {
		t.Errorf("wanted error, but have nil")
		return
	}
}

func TestService_ProfileChangeMultiErrors(t *testing.T) {
	ctrlUser := gomock.NewController(t)
	defer ctrlUser.Finish()
	mockUserStorage := mocks.NewMockUserStorage(ctrlUser)

	ctrlAvatar := gomock.NewController(t)
	defer ctrlAvatar.Finish()
	mockAvatarStorage := mocks.NewMockAvatarStorage(ctrlAvatar)

	service := &service{
		userStorage:   mockUserStorage,
		avatarStorage: mockAvatarStorage,
	}

	request := &protoProfile.UserInputProfile{
		ID:       1,
		Email:    "epridius",
		Username: "pkaterinaa",
		FullName: "",
		Avatar:   []byte{},
	}

	input := models.UserInputProfile{
		ID:       request.ID,
		Email:    request.Email,
		Username: request.Username,
		FullName: request.FullName,
		Avatar:   request.Avatar,
	}
	userAvatar := models.UserAvatar{
		ID:     1,
		Avatar: "default/default_avatar.png",
	}

	multiErrors := new(models.MultiErrors)
	multiErrors.Codes = append(multiErrors.Codes, "101")
	multiErrors.Descriptions = append(multiErrors.Descriptions, "No such user")
	expectedErr := models.ServeError{Codes: multiErrors.Codes, Descriptions: multiErrors.Descriptions,
		MethodName: "ProfileChange"}

	mockUserStorage.EXPECT().GetUserAvatar(models.UserInput{ID: request.ID}).Return(userAvatar, nil)
	mockAvatarStorage.EXPECT().UploadAvatar(request.Avatar, &userAvatar, false).Return(nil)
	mockUserStorage.EXPECT().ChangeUserProfile(input, userAvatar).Return(models.UserOutside{}, expectedErr)

	_, err := service.ProfileChange(context.Background(), request)
	if err == nil {
		t.Error("expected error")
		return
	}
}

func TestService_ProfileChangeFail(t *testing.T) {
	ctrlUser := gomock.NewController(t)
	defer ctrlUser.Finish()
	mockUserStorage := mocks.NewMockUserStorage(ctrlUser)

	ctrlAvatar := gomock.NewController(t)
	defer ctrlAvatar.Finish()
	mockAvatarStorage := mocks.NewMockAvatarStorage(ctrlAvatar)

	service := &service{
		userStorage:   mockUserStorage,
		avatarStorage: mockAvatarStorage,
	}

	request := &protoProfile.UserInputProfile{
		ID:       1,
		Email:    "epridius",
		Username: "pkaterinaa",
		FullName: "",
		Avatar:   []byte{},
	}

	input := models.UserInputProfile{
		ID:       request.ID,
		Email:    request.Email,
		Username: request.Username,
		FullName: request.FullName,
		Avatar:   request.Avatar,
	}
	userAvatar := models.UserAvatar{
		ID:     1,
		Avatar: "default/default_avatar.png",
	}
	errStorage := models.ServeError{Codes: []string{"500"}, OriginalError: errors.New("")}

	mockUserStorage.EXPECT().GetUserAvatar(models.UserInput{ID: request.ID}).Return(userAvatar, nil)
	mockAvatarStorage.EXPECT().UploadAvatar(request.Avatar, &userAvatar, false).Return(nil)
	mockUserStorage.EXPECT().ChangeUserProfile(input, userAvatar).Return(models.UserOutside{}, errStorage)

	_, err := service.ProfileChange(context.Background(), request)
	if err == nil {
		t.Error("expected error")
		return
	}
}

func TestService_PasswordChange(t *testing.T) {
	ctrlUser := gomock.NewController(t)
	defer ctrlUser.Finish()
	mockUserStorage := mocks.NewMockUserStorage(ctrlUser)

	service := &service{
		userStorage: mockUserStorage,
	}

	request := &protoProfile.UserInputPassword{
		ID:          1,
		OldPassword: "lalala",
		Password:    "nanana",
	}

	input := models.UserInputPassword{
		ID:          request.ID,
		OldPassword: request.OldPassword,
		Password:    request.Password,
	}

	internal := models.UserOutside{
		Email:    "epridius",
		Username: "pkaterinaa",
		FullName: "",
		Avatar:   "default/default_avatar.png",
	}

	expected := &protoProfile.UserOutside{
		Email:    internal.Email,
		Username: internal.Username,
		FullName: internal.FullName,
		Avatar:   internal.Avatar,
	}

	mockUserStorage.EXPECT().ChangeUserPassword(input).Return(internal, nil)

	output, err := service.PasswordChange(context.Background(), request)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expected, output)
		return
	}
}

func TestService_PasswordChangeFail(t *testing.T) {
	ctrlUser := gomock.NewController(t)
	defer ctrlUser.Finish()
	mockUserStorage := mocks.NewMockUserStorage(ctrlUser)

	service := &service{
		userStorage: mockUserStorage,
	}

	request := &protoProfile.UserInputPassword{
		ID:          1,
		OldPassword: "lalala",
		Password:    "nanana",
	}

	input := models.UserInputPassword{
		ID:          request.ID,
		OldPassword: request.OldPassword,
		Password:    request.Password,
	}
	errStorage := models.ServeError{Codes: []string{"500"}, OriginalError: errors.New("")}

	mockUserStorage.EXPECT().ChangeUserPassword(input).Return(models.UserOutside{}, errStorage)

	_, err := service.PasswordChange(context.Background(), request)
	if err == nil {
		t.Error("expected error")
		return
	}
}

func TestService_Boards(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardService := mockBoards.NewMockBoardClient(ctrlBoard)

	service := &service{
		boardService: mockBoardService,
	}

	userInput := &protoProfile.UserID{ID: 1}

	expected := &protoProfile.BoardsOutsideShort{Boards: nil}
	expected.Boards = append(expected.Boards, &protoProfile.BoardOutsideShort{})

	c := context.Background()
	mockBoardService.EXPECT().GetBoardsList(c, userInput).Return(expected, nil)

	boards, err := service.Boards(c, userInput)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(boards, expected) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expected, boards)
		return
	}
}

func TestService_BoardsFail(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardService := mockBoards.NewMockBoardClient(ctrlBoard)

	service := &service{
		boardService: mockBoardService,
	}

	userInput := &protoProfile.UserID{ID: 1}

	expected := &protoProfile.BoardsOutsideShort{Boards: nil}

	c := context.Background()
	errStorage := models.ServeError{Codes: []string{"500"}, OriginalError: errors.New("")}

	mockBoardService.EXPECT().GetBoardsList(c, userInput).Return(expected, errStorage)

	_, err := service.Boards(c, userInput)
	if err == nil {
		t.Error("expected error")
		return
	}
}
