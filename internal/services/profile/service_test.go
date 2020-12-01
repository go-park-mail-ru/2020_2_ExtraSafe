package profile

import (
	"context"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/tools/validation"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/mock"
	protoProfile "github.com/go-park-mail-ru/2020_2_ExtraSafe/services/proto/profile"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestService_Profile(t *testing.T) {
	ctrlUser := gomock.NewController(t)
	defer ctrlUser.Finish()
	mockProfileService := mock.NewMockProfileClient(ctrlUser)
	validator := validation.NewService()

	service := NewService(mockProfileService, validator)

	userInput := models.UserInput{ID: 1}
	input := &protoProfile.UserID{ID: userInput.ID}

	expect := models.UserOutside{
		Email:    "epridius",
		Username: "pkaterinaa",
		FullName: "",
		Avatar:   "default/default_avatar.png",
	}

	internal := &protoProfile.UserOutside{
		Email:    expect.Email,
		Username: expect.Username,
		FullName: expect.FullName,
		Avatar:   expect.Avatar,
	}

	mockProfileService.EXPECT().Profile(context.Background(), input).Return(internal, nil)

	output, err := service.Profile(userInput)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(output, expect) {
		t.Errorf("results not match, want %v, have %v", expect, output)
		return
	}
}
/*
func TestService_ProfileChange(t *testing.T) {
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
		Avatar:  request.Avatar,
	}

	internal := models.UserOutside{
		Email:    request.Email,
		Username: request.Username,
		FullName: request.FullName,
		Avatar: "default/default_avatar.png",
	}

	expect := &protoProfile.UserOutside{
		Email:    request.Email,
		Username: request.Username,
		FullName: request.FullName,
		Avatar:   internal.Avatar,
	}

	ctrlUser := gomock.NewController(t)
	defer ctrlUser.Finish()
	mockProfileService := mock.NewMockProfileClient(ctrlUser)
	validator := validation.NewService()

	service := NewService(mockProfileService, validator)

	mockProfileService.EXPECT().ProfileChange(context.Background(), request).Return(expect, nil)

	output, err := service.ProfileChange(input)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(output, internal) {
		t.Errorf("results not match, want %v, have %v", internal, output)
		return
	}
}*/

func TestService_Boards(t *testing.T) {

}

func TestService_PasswordChange(t *testing.T) {

}
