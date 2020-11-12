package auth

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	mocks "github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/services/auth/mock"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestNewService(t *testing.T) {

}

func TestService_Auth(t *testing.T) {
	ctrlUser := gomock.NewController(t)
	defer ctrlUser.Finish()
	mockUserStorage := mocks.NewMockUserStorage(ctrlUser)

	ctrlBoards := gomock.NewController(t)
	defer ctrlBoards.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoards)

	expectedUser := models.UserOutside{
		Email:    "epridius@gmail.com",
		Username: "pkaterinaa",
		FullName: "",
		Avatar:   "default/default_avatar.png",
	}

	board1 := models.BoardOutsideShort{
		BoardID: 1,
		Name:    "board_1",
		Theme:   "dark",
		Star:    false,
	}

	board2 := models.BoardOutsideShort{
		BoardID: 2,
		Name:    "board_2",
		Theme:   "light",
		Star:    false,
	}

	expectedBoards := make([]models.BoardOutsideShort, 0)
	expectedBoards = append(expectedBoards, board1, board2)

	userInput := models.UserInput{ID: 1}

	mockUserStorage.EXPECT().GetUserProfile(userInput).Return(expectedUser, nil)
	mockBoardStorage.EXPECT().GetBoardsList(userInput).Return(expectedBoards, nil)

	service := &service{
		userStorage: mockUserStorage,
		boardStorage: mockBoardStorage,
	}

	expectedResponse := models.UserBoardsOutside{
		Email:    expectedUser.Email,
		Username: expectedUser.Username,
		FullName: expectedUser.FullName,
		Links:    expectedUser.Links,
		Avatar:   expectedUser.Avatar,
		Boards:   expectedBoards,
	}

	response, err := service.Auth(userInput)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(response, expectedResponse) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedResponse, response)
		return
	}
}

func TestService_Login(t *testing.T) {
	request := models.UserInputLogin{
		Email:    "epridius@yandex.ru",
		Password: "lalala",
	}

	expectedUser:= models.UserOutside{
		Email:    request.Email,
		Username: "pkaterinaa",
		FullName: "",
		Links:    nil,
		Avatar:   "default/default_avatar.png",
	}

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

	mockValidator.EXPECT().ValidateLogin(request).Return(nil)
	mockUserStorage.EXPECT().CheckUser(request).Return(uint64(1), expectedUser, nil)

	userID, user, err := service.Login(request)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(user, expectedUser) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedUser, user)
		return
	}
	if userID != 1 {
		t.Errorf("result ID not match, \nhave \n%v", userID)
		return
	}
}

func TestService_Registration(t *testing.T) {
	request := models.UserInputReg{
		Email:    "epridius@yandex.ru",
		Username: "pkaterinaa",
		Password: "lalala",
	}

	expectedUser:= models.UserOutside{
		Email:    request.Email,
		Username: request.Username,
		FullName: "",
		Links:    nil,
		Avatar:   "default/default_avatar.png",
	}

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

	mockValidator.EXPECT().ValidateRegistration(request).Return(nil)
	mockUserStorage.EXPECT().CreateUser(request).Return(uint64(1), expectedUser, nil)

	userID, user, err := service.Registration(request)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(user, expectedUser) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedUser, user)
		return
	}
	if userID != 1 {
		t.Errorf("result ID not match, \nhave \n%v", userID)
		return
	}
}
