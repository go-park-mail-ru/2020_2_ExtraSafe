package profile

import (
	"context"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	mockBoards "github.com/go-park-mail-ru/2020_2_ExtraSafe/services/mock"
	mocks "github.com/go-park-mail-ru/2020_2_ExtraSafe/services/profile_service/internal/service/mock"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/proto/profile"
	"github.com/golang/mock/gomock"
	"mime/multipart"
	"reflect"
	"testing"
)

func TestService_CheckUser(t *testing.T) {
	request := &protoProfile.UserInputLogin{
		Email:    "epridius@yandex.ru",
		Password: "lalala",
	}

	expectedUser:= &protoProfile.UserID{ ID: 1 }

	ctrlUser := gomock.NewController(t)
	defer ctrlUser.Finish()
	//mockUsers :=
	mockStorage := mocks.NewMockStorage(ctrlUser)

	ctrlBoards := gomock.NewController(t)
	defer ctrlBoards.Finish()
	mockBoardStorage := mockBoards.NewMockService(ctrlBoards)

	service := &service{
		userStorage: mockStorage,
		boardService: mockBoardStorage,
	}

	mockStorage.EXPECT().CheckUser(request).Return(expectedUser, nil)

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

func TestService_CreateUser(t *testing.T) {

}

func TestService_GetUserByUsername(t *testing.T) {

}

func TestService_GetUsersByIDs(t *testing.T) {

}

func TestService_Profile(t *testing.T) {
	ctrlUser := gomock.NewController(t)
	defer ctrlUser.Finish()
	mockUserStorage := mocks.NewMockUserStorage(ctrlUser)

	service := &service{
		userStorage: mockUserStorage,
	}

	userInput := models.UserInput{ID: 1}
	expectedUserOutside := models.UserOutside{
		Email:    "epridius",
		Username: "pkaterinaa",
		FullName: "",
		Links:    nil,
		Avatar:   "default/default_avatar.png",
	}

	mockUserStorage.EXPECT().GetUserProfile(userInput).Return(expectedUserOutside, nil)

	user, err := service.Profile(userInput)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(user, expectedUserOutside) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedUserOutside, user)
		return
	}
}

func TestService_Accounts(t *testing.T) {
	ctrlUser := gomock.NewController(t)
	defer ctrlUser.Finish()
	mockUserStorage := mocks.NewMockUserStorage(ctrlUser)

	service := &service{
		userStorage: mockUserStorage,
	}

	userInput := models.UserInput{ID: 1}
	expectedUserAccounts := models.UserOutside{
		Email:    "epridius",
		Username: "pkaterinaa",
		FullName: "",
		Links:    &models.UserLinks{
			Telegram:  "pkaterinaa",
			Github:    "pringleskate",
		},
		Avatar:   "default/default_avatar.png",
	}

	mockUserStorage.EXPECT().GetUserAccounts(userInput).Return(expectedUserAccounts, nil)

	user, err := service.Accounts(userInput)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(user, expectedUserAccounts) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedUserAccounts, user)
		return
	}
}

func TestService_ProfileChange(t *testing.T) {
	ctrlValid := gomock.NewController(t)
	defer ctrlValid.Finish()
	mockValidator := mocks.NewMockValidator(ctrlValid)

	ctrlUser := gomock.NewController(t)
	defer ctrlUser.Finish()
	mockUserStorage := mocks.NewMockUserStorage(ctrlUser)

	ctrlAvatar := gomock.NewController(t)
	defer ctrlAvatar.Finish()
	mockAvatarStorage := mocks.NewMockAvatarStorage(ctrlAvatar)

	service := &service{
		userStorage: mockUserStorage,
		avatarStorage: mockAvatarStorage,
		validator: mockValidator,
	}

	request := models.UserInputProfile{
		ID:       1,
		Email:    "epridius",
		Username: "pkaterinaa",
		FullName: "",
		Avatar:   &multipart.FileHeader{},
	}

	userAvatar := models.UserAvatar{
		ID:     1,
		Avatar: "default/default_avatar.png",
	}

	expectedUser := models.UserOutside{
		Email:   "epridius",
		Username: "pkaterinaa",
		FullName: "",
		Avatar:   "avatars/new_avatar.png",
	}

	mockValidator.EXPECT().ValidateProfile(request).Return(nil)
	mockUserStorage.EXPECT().GetUserAvatar(models.UserInput{ID: request.ID}).Return(userAvatar, nil)
	mockAvatarStorage.EXPECT().UploadAvatar(request.Avatar, &userAvatar, false).Return(nil)
	mockUserStorage.EXPECT().ChangeUserProfile(request, userAvatar).Return(expectedUser, nil)

	user, err := service.ProfileChange(request)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(user, expectedUser) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedUser, user)
		return
	}
}


func TestService_ProfileChangeError(t *testing.T) {
	ctrlValid := gomock.NewController(t)
	defer ctrlValid.Finish()
	mockValidator := mocks.NewMockValidator(ctrlValid)

	ctrlUser := gomock.NewController(t)
	defer ctrlUser.Finish()
	mockUserStorage := mocks.NewMockUserStorage(ctrlUser)

	ctrlAvatar := gomock.NewController(t)
	defer ctrlAvatar.Finish()
	mockAvatarStorage := mocks.NewMockAvatarStorage(ctrlAvatar)

	service := &service{
		userStorage: mockUserStorage,
		avatarStorage: mockAvatarStorage,
		validator: mockValidator,
	}

	request := models.UserInputProfile{
		ID:       1,
		Email:    "epridius",
		Username: "pkaterinaa",
		FullName: "",
		Avatar:   &multipart.FileHeader{},
	}

	userAvatar := models.UserAvatar{
		ID:     1,
		Avatar: "default/default_avatar.png",
	}

	expectedUser := models.UserOutside{}

	errAvatar := models.ServeError{Codes: []string{"600"}, Descriptions: []string{"File error"}, MethodName: "UploadAvatar"}

	multiErrors := new(models.MultiErrors)
	multiErrors.Codes = append(multiErrors.Codes, errAvatar.Codes...)
	multiErrors.Descriptions = append(multiErrors.Descriptions, errAvatar.Descriptions...)
	expectedErr := models.ServeError{Codes: multiErrors.Codes, Descriptions: multiErrors.Descriptions,
		MethodName: "ProfileChange"}

	mockValidator.EXPECT().ValidateProfile(request).Return(nil)
	mockUserStorage.EXPECT().GetUserAvatar(models.UserInput{ID: request.ID}).Return(userAvatar, nil)
	mockAvatarStorage.
		EXPECT().
		UploadAvatar(request.Avatar, &userAvatar, false).
		Return(errAvatar)
	mockUserStorage.EXPECT().ChangeUserProfile(request, userAvatar).Return(expectedUser, nil)

	user, err := service.ProfileChange(request)
	if !reflect.DeepEqual(err, expectedErr) {
		t.Errorf("result errors not match, want \n%v, \nhave \n%v", expectedErr, err)
		return
	}
	if !reflect.DeepEqual(user, expectedUser) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedUser, user)
		return
	}
	if err == nil {
		t.Errorf("wanted error, but have nil")
		return
	}
}

func TestService_AccountsChange(t *testing.T) {
	ctrlUser := gomock.NewController(t)
	defer ctrlUser.Finish()
	mockUserStorage := mocks.NewMockUserStorage(ctrlUser)

	ctrlValid := gomock.NewController(t)
	defer ctrlValid.Finish()
	mockValidator := mocks.NewMockValidator(ctrlValid)

	service := &service{
		userStorage: mockUserStorage,
		validator: mockValidator,
	}

	request := models.UserInputLinks {
		ID : 				1,
		Telegram:  "pkaterinaa",
		Github:    "pringleskate",
	}

	expectedUserAccounts := models.UserOutside{
		Email:    "epridius",
		Username: "pkaterinaa",
		FullName: "",
		Links:    &models.UserLinks{
			Telegram:  "pkaterinaa",
			Github:    "pringleskate",
		},
		Avatar:   "default/default_avatar.png",
	}

	mockValidator.EXPECT().ValidateLinks(request).Return(nil)
	mockUserStorage.EXPECT().ChangeUserAccounts(request).Return(expectedUserAccounts, nil)

	user, err := service.AccountsChange(request)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(user, expectedUserAccounts) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedUserAccounts, user)
		return
	}
}

func TestService_PasswordChange(t *testing.T) {
	ctrlUser := gomock.NewController(t)
	defer ctrlUser.Finish()
	mockUserStorage := mocks.NewMockUserStorage(ctrlUser)

	ctrlValid := gomock.NewController(t)
	defer ctrlValid.Finish()
	mockValidator := mocks.NewMockValidator(ctrlValid)

	service := &service{
		userStorage: mockUserStorage,
		validator: mockValidator,
	}

	userInput := models.UserInputPassword{
		ID: 1,
		OldPassword: "lalala",
		Password: "nanana",
	}

	expectedUserOutside := models.UserOutside{
		Email:    "epridius",
		Username: "pkaterinaa",
		FullName: "",
		Links:    nil,
		Avatar:   "default/default_avatar.png",
	}

	mockValidator.EXPECT().ValidateChangePassword(userInput).Return(nil)
	mockUserStorage.EXPECT().ChangeUserPassword(userInput).Return(expectedUserOutside, nil)

	user, err := service.PasswordChange(userInput)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(user, expectedUserOutside) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedUserOutside, user)
		return
	}
}

func TestService_Boards(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	service := &service{
		boardStorage: mockBoardStorage,
	}

	userInput := models.UserInput{ID: 1}
	expectedBoards := make([]models.BoardOutsideShort, 0)

	board1 := models.BoardOutsideShort{
		BoardID: 1,
		Name:    "first",
		Theme:   "dark",
		Star:    false,
	}

	board2 := models.BoardOutsideShort{
		BoardID: 2,
		Name:    "second",
		Theme:   "dark",
		Star:    false,
	}

	expectedBoards = append(expectedBoards, board1, board2)

	mockBoardStorage.EXPECT().GetBoardsList(userInput).Return(expectedBoards, nil)

	boards, err := service.Boards(userInput)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(boards, expectedBoards) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedBoards, boards)
		return
	}
}
