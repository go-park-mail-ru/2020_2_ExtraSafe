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

func TestService_ProfileChange(t *testing.T) {
	input := models.UserInputProfile{
		ID:       int64(1),
		Email:    "epridius@mail.ru",
		Username: "pkaterinaa",
		FullName: "",
	}

	request := &protoProfile.UserInputProfile{
		ID:       input.ID,
		Email:    input.Email,
		Username: input.Username,
		FullName: input.FullName,
	}

	internal := models.UserOutside{
		Email:    input.Email,
		Username: input.Username,
		FullName: input.FullName,
		Avatar:   "default/default_avatar.png",
	}

	expect := &protoProfile.UserOutside{
		Email:    input.Email,
		Username: input.Username,
		FullName: input.FullName,
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
}

func TestService_Boards(t *testing.T) {
	input := models.UserInput{ID: 1}
	userInput := &protoProfile.UserID{ID: input.ID}

	board := &protoProfile.BoardOutsideShort{
		BoardID: 1,
		Name:    "name",
		Theme:   "dark",
		Star:    false,
	}
	expected := &protoProfile.BoardsOutsideShort{Boards: nil}
	expected.Boards = append(expected.Boards, board)

	output := make([]models.BoardOutsideShort, 0)
	output = append(output, models.BoardOutsideShort{
		BoardID: board.BoardID,
		Name:    board.Name,
		Theme:   board.Theme,
		Star:    board.Star,
	})

	ctrlUser := gomock.NewController(t)
	defer ctrlUser.Finish()
	mockProfileService := mock.NewMockProfileClient(ctrlUser)
	validator := validation.NewService()

	service := NewService(mockProfileService, validator)

	c := context.Background()
	mockProfileService.EXPECT().Boards(c, userInput).Return(expected, nil)

	boards, err := service.Boards(input)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(boards, output) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", output, boards)
		return
	}

}

func TestService_PasswordChange(t *testing.T) {
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
		Email:    "epridius@mail.ru",
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

	ctrlUser := gomock.NewController(t)
	defer ctrlUser.Finish()
	mockProfileService := mock.NewMockProfileClient(ctrlUser)
	validator := validation.NewService()

	service := NewService(mockProfileService, validator)

	mockProfileService.EXPECT().PasswordChange(context.Background(), request).Return(expected, nil)

	output, err := service.PasswordChange(input)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(output, internal) {
		t.Errorf("results not match, want %v, have %v", internal, output)
		return
	}
}
