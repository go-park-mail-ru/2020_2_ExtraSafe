package profile

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	mocks "github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/services/profile/mock"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

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

//FIXME
func TestService_ProfileChange(t *testing.T) {

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