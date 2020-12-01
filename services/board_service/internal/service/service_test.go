package service

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	mocks "github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/service/mock"
	serviceMocks "github.com/go-park-mail-ru/2020_2_ExtraSafe/services/mock"
	protoBoard "github.com/go-park-mail-ru/2020_2_ExtraSafe/services/proto/board"
	protoProfile "github.com/go-park-mail-ru/2020_2_ExtraSafe/services/proto/profile"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

var errStorage = models.ServeError{Codes: []string{"500"}, OriginalError: errors.New("")}

func TestService_CreateBoard(t *testing.T) {
	request := &protoBoard.BoardChangeInput{
		UserID:    1,
		BoardID:   1,
		BoardName: "name",
		Theme:     "dark",
		Star:      false,
	}

	input := models.BoardChangeInput{
		UserID:    request.UserID,
		BoardID:   request.BoardID,
		BoardName: request.BoardName,
		Theme:     request.Theme,
		Star:      request.Star,
	}

	internal := models.BoardInternal{
		BoardID: input.BoardID,
		Name:    input.BoardName,
		Theme:   input.Theme,
		Star:    input.Star,
	}

	expect := &protoProfile.BoardOutsideShort{
		BoardID: internal.BoardID,
		Name:    internal.Name,
		Theme:   internal.Theme,
		Star:    internal.Star,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	service := &service{boardStorage: mockBoardStorage}

	mockBoardStorage.EXPECT().CreateBoard(input).Return(internal, nil)

	output, err := service.CreateBoard(context.Background(), request)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(output, expect) {
		t.Errorf("results not match, want %v, have %v", expect, output)
		return
	}
}

func TestService_CreateBoardFail(t *testing.T) {
	request := &protoBoard.BoardChangeInput{
		UserID:    1,
		BoardID:   1,
		BoardName: "name",
		Theme:     "dark",
		Star:      false,
	}

	input := models.BoardChangeInput{
		UserID:    request.UserID,
		BoardID:   request.BoardID,
		BoardName: request.BoardName,
		Theme:     request.Theme,
		Star:      request.Star,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	service := &service{boardStorage: mockBoardStorage}

	mockBoardStorage.EXPECT().CreateBoard(input).Return(models.BoardInternal{}, errStorage)

	_, err := service.CreateBoard(context.Background(), request)
	if err == nil {
		t.Error("expected error")
		return
	}
}

func TestService_ChangeBoard(t *testing.T) {
	request := &protoBoard.BoardChangeInput{
		UserID:    1,
		BoardID:   1,
		BoardName: "name",
		Theme:     "dark",
		Star:      false,
	}

	input := models.BoardChangeInput{
		UserID:    request.UserID,
		BoardID:   request.BoardID,
		BoardName: request.BoardName,
		Theme:     request.Theme,
		Star:      request.Star,
	}

	internal := models.BoardInternal{
		BoardID: input.BoardID,
		Name:    input.BoardName,
		Theme:   input.Theme,
		Star:    input.Star,
	}

	expect := &protoProfile.BoardOutsideShort{
		BoardID: internal.BoardID,
		Name:    internal.Name,
		Theme:   internal.Theme,
		Star:    internal.Star,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	service := &service{boardStorage: mockBoardStorage}

	mockBoardStorage.EXPECT().ChangeBoard(input).Return(internal, nil)

	output, err := service.ChangeBoard(context.Background(), request)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(output, expect) {
		t.Errorf("results not match, want %v, have %v", expect, output)
		return
	}
}

func TestService_ChangeBoardFail(t *testing.T) {
	request := &protoBoard.BoardChangeInput{
		UserID:    1,
		BoardID:   1,
		BoardName: "name",
		Theme:     "dark",
		Star:      false,
	}

	input := models.BoardChangeInput{
		UserID:    request.UserID,
		BoardID:   request.BoardID,
		BoardName: request.BoardName,
		Theme:     request.Theme,
		Star:      request.Star,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	service := &service{boardStorage: mockBoardStorage}

	mockBoardStorage.EXPECT().ChangeBoard(input).Return(models.BoardInternal{}, errStorage)

	_, err := service.ChangeBoard(context.Background(), request)
	if err == nil {
		t.Error("expected error")
		return
	}
}

func TestService_DeleteBoard(t *testing.T) {
	request := &protoBoard.BoardInput{
		UserID:    1,
		BoardID:   1,
	}

	input := models.BoardInput{
		UserID:    request.UserID,
		BoardID:   request.BoardID,
	}

	expect := &protoBoard.Nothing{Dummy: true}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	service := &service{boardStorage: mockBoardStorage}

	mockBoardStorage.EXPECT().DeleteBoard(input).Return(nil)

	output, err := service.DeleteBoard(context.Background(), request)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(output, expect) {
		t.Errorf("results not match, want %v, have %v", expect, output)
		return
	}
}

func TestService_DeleteBoardFail(t *testing.T) {
	request := &protoBoard.BoardInput{
		UserID:    1,
		BoardID:   1,
	}

	input := models.BoardInput{
		UserID:    request.UserID,
		BoardID:   request.BoardID,
	}

	expect := &protoBoard.Nothing{Dummy: true}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	service := &service{boardStorage: mockBoardStorage}

	mockBoardStorage.EXPECT().DeleteBoard(input).Return(errStorage)

	output, err := service.DeleteBoard(context.Background(), request)
	if err == nil {
		t.Error("expected error")
		return
	}
	if !reflect.DeepEqual(output, expect) {
		t.Errorf("results not match, want %v, have %v", expect, output)
		return
	}
}

func TestService_GetBoard(t *testing.T) {
	request := &protoBoard.BoardInput{UserID: 1, BoardID: 1}
	input := models.BoardInput{
		UserID:  request.UserID,
		BoardID: request.BoardID,
	}

	card := models.CardInternal{
		CardID: 1,
		Name:   "card",
		Order:  1,
		Tasks: []models.TaskInternalShort{},
	}

	cardsInternal := make([]models.CardInternal, 0)
	cardsInternal = append(cardsInternal, card)

	tag := models.TagOutside{
		TagID: 1,
		Color: "yellow",
		Name:  "work",
	}
	tagsInternal := make([]models.TagOutside, 0)
	tagsInternal = append(tagsInternal, tag)

	internal := models.BoardInternal{
		BoardID:  request.BoardID,
		AdminID:  2,
		Name:     "name",
		Theme:    "dark",
		Star:     false,
		Cards:    cardsInternal,
		UsersIDs: []int64{2},
		Tags:     tagsInternal,
	}

	users := make([]*protoProfile.UserOutsideShort,0)
	users = append(users, &protoProfile.UserOutsideShort{ID: internal.UsersIDs[0]})

	cards := make([]*protoBoard.CardOutside, 0)
	tasks := make([]*protoBoard.TaskOutsideShort, 0)

	cards = append(cards, &protoBoard.CardOutside{
		CardID: card.CardID,
		Name:   card.Name,
		Order:  card.Order,
		Tasks: tasks,
	})

	expect :=  &protoBoard.BoardOutside{
		BoardID: internal.BoardID,
		Admin:   &protoProfile.UserOutsideShort{ID: internal.AdminID},
		Name:    internal.Name,
		Theme:   internal.Theme,
		Star:    internal.Star,
		Users:   users,
		Cards:   cards,
		Tags:    convertTags(internal.Tags),
	}

	membersIDs := make([]int64, 0)
	membersIDs = append(membersIDs, internal.AdminID)
	membersIDs = append(membersIDs, internal.UsersIDs...)

	members := &protoProfile.UsersOutsideShort{Users: nil}
	members.Users = append(members.Users, &protoProfile.UserOutsideShort{ID: internal.AdminID},  &protoProfile.UserOutsideShort{ID: internal.UsersIDs[0]})

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	ctrlUser := gomock.NewController(t)
	defer ctrlUser.Finish()
	mockProfileService := serviceMocks.NewMockProfileClient(ctrlUser)

	mockBoardStorage.EXPECT().GetBoard(input).Return(internal, nil)
	mockProfileService.EXPECT().GetUsersByIDs(context.Background(),  &protoProfile.UserIDS{UserIDs: membersIDs}).Return(members, nil)

	service := &service{boardStorage: mockBoardStorage, profileService: mockProfileService}

	output, err := service.GetBoard(context.Background(), request)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(output, expect) {
		t.Errorf("results not match, \nwant %+v\n, have %+v\n", expect, output)
		return
	}

}

func TestService_GetBoardFail(t *testing.T) {
	request := &protoBoard.BoardInput{UserID: 1, BoardID: 1}
	input := models.BoardInput{
		UserID:  request.UserID,
		BoardID: request.BoardID,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	ctrlUser := gomock.NewController(t)
	defer ctrlUser.Finish()
	mockProfileService := serviceMocks.NewMockProfileClient(ctrlUser)

	mockBoardStorage.EXPECT().GetBoard(input).Return(models.BoardInternal{}, errStorage)

	service := &service{boardStorage: mockBoardStorage, profileService: mockProfileService}

	_, err := service.GetBoard(context.Background(), request)
	if err == nil {
		t.Error("expected error")
		return
	}

}

func TestService_GetBoardsList(t *testing.T) {
	request := &protoProfile.UserID{ID: 1}
	input := models.UserInput{ID: request.ID}

	internal := make([]models.BoardOutsideShort, 0)
	internal = append(internal, models.BoardOutsideShort{
		BoardID: 1,
		Name:    "",
		Theme:   "",
		Star:    false,
	})

	expect := &protoProfile.BoardsOutsideShort{Boards: nil}
	expect.Boards = append(expect.Boards, &protoProfile.BoardOutsideShort{
		BoardID: internal[0].BoardID,
		Name:    internal[0].Name,
		Theme:   internal[0].Theme,
		Star:    internal[0].Star,
	})

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	service := &service{boardStorage: mockBoardStorage}

	mockBoardStorage.EXPECT().GetBoardsList(input).Return(internal, nil)

	output, err := service.GetBoardsList(context.Background(), request)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(output, expect) {
		t.Errorf("results not match, want %v, have %v", expect, output)
		return
	}
}

func TestService_GetBoardUsersFail(t *testing.T) {
	request := &protoBoard.BoardInput{UserID: 1, BoardID: 1}
	input := models.BoardInput{
		UserID:  request.UserID,
		BoardID: request.BoardID,
	}

	card := models.CardInternal{
		CardID: 1,
		Name:   "card",
		Order:  1,
		Tasks: []models.TaskInternalShort{},
	}

	cardsInternal := make([]models.CardInternal, 0)
	cardsInternal = append(cardsInternal, card)

	tag := models.TagOutside{
		TagID: 1,
		Color: "yellow",
		Name:  "work",
	}
	tagsInternal := make([]models.TagOutside, 0)
	tagsInternal = append(tagsInternal, tag)

	internal := models.BoardInternal{
		BoardID:  request.BoardID,
		AdminID:  2,
		Name:     "name",
		Theme:    "dark",
		Star:     false,
		Cards:    cardsInternal,
		UsersIDs: []int64{2},
		Tags:     tagsInternal,
	}

	cards := make([]*protoBoard.CardOutside, 0)
	tasks := make([]*protoBoard.TaskOutsideShort, 0)

	cards = append(cards, &protoBoard.CardOutside{
		CardID: card.CardID,
		Name:   card.Name,
		Order:  card.Order,
		Tasks: tasks,
	})

	membersIDs := make([]int64, 0)
	membersIDs = append(membersIDs, internal.AdminID)
	membersIDs = append(membersIDs, internal.UsersIDs...)

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	ctrlUser := gomock.NewController(t)
	defer ctrlUser.Finish()
	mockProfileService := serviceMocks.NewMockProfileClient(ctrlUser)

	mockBoardStorage.EXPECT().GetBoard(input).Return(internal, nil)
	mockProfileService.EXPECT().GetUsersByIDs(context.Background(),  &protoProfile.UserIDS{UserIDs: membersIDs}).Return(&protoProfile.UsersOutsideShort{}, errors.New(""))

	service := &service{boardStorage: mockBoardStorage, profileService: mockProfileService}

	_, err := service.GetBoard(context.Background(), request)
	if err == nil {
		t.Error("expected error")
		return
	}

}
func TestService_GetBoardsListFail(t *testing.T) {
	request := &protoProfile.UserID{ID: 1}
	input := models.UserInput{ID: request.ID}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	service := &service{boardStorage: mockBoardStorage}

	mockBoardStorage.EXPECT().GetBoardsList(input).Return(nil, errStorage)

	_, err := service.GetBoardsList(context.Background(), request)
	if err == nil {
		t.Error("expected error")
		return
	}
}

func TestService_AddUserToBoard(t *testing.T) {
	request := &protoBoard.BoardMemberInput{
		UserID:     1,
		BoardID:    1,
		MemberName: "pkaterinaa",
	}

	input := models.BoardMember{
		UserID:    request.UserID,
		BoardID:   request.BoardID,
		MemberID:  2,
	}

	userProfile := &protoProfile.UserOutsideShort{ID: input.MemberID}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	ctrlUser := gomock.NewController(t)
	defer ctrlUser.Finish()
	mockProfileService := serviceMocks.NewMockProfileClient(ctrlUser)

	mockProfileService.EXPECT().
						GetUserByUsername(context.Background(), &protoProfile.UserName{UserName: request.MemberName}).
						Return(userProfile, nil)

	mockBoardStorage.EXPECT().AddUser(input).Return(nil)

	service := &service{boardStorage: mockBoardStorage, profileService: mockProfileService}

	output, err := service.AddUserToBoard(context.Background(), request)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(output, userProfile) {
		t.Errorf("results not match, want %v, have %v", userProfile, output)
		return
	}
}

func TestService_AddUserToBoardFail(t *testing.T) {
	request := &protoBoard.BoardMemberInput{
		UserID:     1,
		BoardID:    1,
		MemberName: "pkaterinaa",
	}

	ctrlUser := gomock.NewController(t)
	defer ctrlUser.Finish()
	mockProfileService := serviceMocks.NewMockProfileClient(ctrlUser)

	mockProfileService.EXPECT().
		GetUserByUsername(context.Background(), &protoProfile.UserName{UserName: request.MemberName}).
		Return(&protoProfile.UserOutsideShort{}, errors.New(""))

	service := &service{profileService: mockProfileService}

	_, err := service.AddUserToBoard(context.Background(), request)
	if err == nil {
		t.Error("expected error")
		return
	}
}

func TestService_AddUserToBoardFail2(t *testing.T) {
	request := &protoBoard.BoardMemberInput{
		UserID:     1,
		BoardID:    1,
		MemberName: "pkaterinaa",
	}

	input := models.BoardMember{
		UserID:    request.UserID,
		BoardID:   request.BoardID,
		MemberID:  2,
	}

	userProfile := &protoProfile.UserOutsideShort{ID: input.MemberID}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	ctrlUser := gomock.NewController(t)
	defer ctrlUser.Finish()
	mockProfileService := serviceMocks.NewMockProfileClient(ctrlUser)

	mockProfileService.EXPECT().
		GetUserByUsername(context.Background(), &protoProfile.UserName{UserName: request.MemberName}).
		Return(userProfile, nil)

	mockBoardStorage.EXPECT().AddUser(input).Return(errStorage)

	service := &service{boardStorage: mockBoardStorage, profileService: mockProfileService}

	_, err := service.AddUserToBoard(context.Background(), request)
	if err == nil {
		t.Error("expected error")
		return
	}
}

func TestService_RemoveUserToBoard(t *testing.T) {
	request := &protoBoard.BoardMemberInput{
		UserID:     1,
		BoardID:    1,
		MemberName: "pkaterinaa",
	}

	input := models.BoardMember{
		UserID:    request.UserID,
		BoardID:   request.BoardID,
		MemberID:  2,
	}

	userProfile := &protoProfile.UserOutsideShort{ID: input.MemberID}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	ctrlUser := gomock.NewController(t)
	defer ctrlUser.Finish()
	mockProfileService := serviceMocks.NewMockProfileClient(ctrlUser)

	mockProfileService.EXPECT().
		GetUserByUsername(context.Background(), &protoProfile.UserName{UserName: request.MemberName}).
		Return(userProfile, nil)

	mockBoardStorage.EXPECT().RemoveUser(input).Return(nil)

	service := &service{boardStorage: mockBoardStorage, profileService: mockProfileService}

	_, err := service.RemoveUserToBoard(context.Background(), request)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
}

func TestService_CreateCard(t *testing.T) {
	request := &protoBoard.CardInput{
		UserID:  1,
		CardID:  0,
		BoardID: 1,
		Name:    "card",
		Order:   1,
	}

	input := models.CardInput{
		UserID:  request.UserID,
		CardID:  request.CardID,
		BoardID: request.BoardID,
		Name:    request.Name,
		Order:   request.Order,
	}

	internal := models.CardOutside{
		CardID: 1,
		Name:   input.Name,
		Order:  input.Order,
		Tasks:  make([]models.TaskOutsideShort, 0),
	}

	expect := &protoBoard.CardOutsideShort{
		CardID: internal.CardID,
		Name:   internal.Name,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	service := &service{boardStorage: mockBoardStorage}

	mockBoardStorage.EXPECT().CreateCard(input).Return(internal, nil)

	output, err := service.CreateCard(context.Background(), request)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(output, expect) {
		t.Errorf("results not match, want %v, have %v", expect, output)
		return
	}
}
func TestService_CreateCardFail(t *testing.T) {
	request := &protoBoard.CardInput{
		UserID:  1,
		CardID:  0,
		BoardID: 1,
		Name:    "card",
		Order:   1,
	}

	input := models.CardInput{
		UserID:  request.UserID,
		CardID:  request.CardID,
		BoardID: request.BoardID,
		Name:    request.Name,
		Order:   request.Order,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	service := &service{boardStorage: mockBoardStorage}

	mockBoardStorage.EXPECT().CreateCard(input).Return(models.CardOutside{}, errStorage)

	_, err := service.CreateCard(context.Background(), request)
	if err == nil {
		t.Error("expected error")
		return
	}
}

func TestService_ChangeCard(t *testing.T) {
	request := &protoBoard.CardInput{
		UserID:  1,
		CardID:  0,
		BoardID: 1,
		Name:    "card",
		Order:   1,
	}

	input := models.CardInput{
		UserID:  request.UserID,
		CardID:  request.CardID,
		BoardID: request.BoardID,
		Name:    request.Name,
		Order:   request.Order,
	}

	internal := models.CardInternal{
		CardID: 1,
		Name:   input.Name,
		Order:  input.Order,
	}

	expect := &protoBoard.CardOutsideShort{
		CardID: internal.CardID,
		Name:   internal.Name,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	service := &service{boardStorage: mockBoardStorage}

	mockBoardStorage.EXPECT().ChangeCard(input).Return(internal, nil)

	output, err := service.ChangeCard(context.Background(), request)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(output, expect) {
		t.Errorf("results not match, want %v, have %v", expect, output)
		return
	}
}

func TestService_ChangeCardFail(t *testing.T) {
	request := &protoBoard.CardInput{
		UserID:  1,
		CardID:  0,
		BoardID: 1,
		Name:    "card",
		Order:   1,
	}

	input := models.CardInput{
		UserID:  request.UserID,
		CardID:  request.CardID,
		BoardID: request.BoardID,
		Name:    request.Name,
		Order:   request.Order,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	service := &service{boardStorage: mockBoardStorage}

	mockBoardStorage.EXPECT().ChangeCard(input).Return(models.CardInternal{}, errStorage)

	_, err := service.ChangeCard(context.Background(), request)
	if err == nil {
		t.Error("expected error")
		return
	}
}

func TestService_DeleteCard(t *testing.T) {
	request := &protoBoard.CardInput{
		UserID:  1,
		CardID:  0,
		BoardID: 1,
	}

	input := models.CardInput{
		UserID:  request.UserID,
		CardID:  request.CardID,
		BoardID: request.BoardID,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	service := &service{boardStorage: mockBoardStorage}

	mockBoardStorage.EXPECT().DeleteCard(input).Return(nil)

	_, err := service.DeleteCard(context.Background(), request)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
}

func TestService_DeleteCardFail(t *testing.T) {
	request := &protoBoard.CardInput{
		UserID:  1,
		CardID:  0,
		BoardID: 1,
	}

	input := models.CardInput{
		UserID:  request.UserID,
		CardID:  request.CardID,
		BoardID: request.BoardID,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	service := &service{boardStorage: mockBoardStorage}

	mockBoardStorage.EXPECT().DeleteCard(input).Return(errStorage)

	_, err := service.DeleteCard(context.Background(), request)
	if err == nil {
		t.Error("expected error")
		return
	}
}

func TestService_GetCard(t *testing.T) {
	request := &protoBoard.CardInput{
		UserID:  1,
		CardID:  1,
	}

	input := models.CardInput{
		UserID:  request.UserID,
		CardID:  request.CardID,
	}

	internal := models.CardInternal{
		CardID: input.CardID,
		Name:   "card",
		Order:  1,
	}

	expect := &protoBoard.CardOutside{
		CardID: internal.CardID,
		Name:   internal.Name,
		Order: internal.Order,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	service := &service{boardStorage: mockBoardStorage}

	mockBoardStorage.EXPECT().GetCard(input).Return(internal, nil)

	output, err := service.GetCard(context.Background(), request)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(output, expect) {
		t.Errorf("results not match, want %+v, have %+v", expect, output)
		return
	}
}

func TestService_GetCardFail(t *testing.T) {
	request := &protoBoard.CardInput{
		UserID:  1,
		CardID:  1,
	}

	input := models.CardInput{
		UserID:  request.UserID,
		CardID:  request.CardID,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	service := &service{boardStorage: mockBoardStorage}

	mockBoardStorage.EXPECT().GetCard(input).Return(models.CardInternal{}, errStorage)

	_, err := service.GetCard(context.Background(), request)
	if err == nil {
		t.Error("expected error")
		return
	}
}

func TestService_CardOrderChange(t *testing.T) {
	cards := make([]*protoBoard.CardOrder, 0)
	card :=  &protoBoard.CardOrder{Order: 1, CardID: 1}
	cards = append(cards, card)

	request := &protoBoard.CardsOrderInput{
		UserID: 1,
		Cards:  cards,
	}

	cardOrder := make([]models.CardOrder, 0)
	cardOrder = append(cardOrder, models.CardOrder{
		CardID: card.CardID,
		Order:  card.Order,
	})

	input := models.CardsOrderInput{
		UserID: request.UserID,
		Cards:  cardOrder,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	service := &service{boardStorage: mockBoardStorage}

	mockBoardStorage.EXPECT().ChangeCardOrder(input).Return(nil)
	_, err := service.CardOrderChange(context.Background(), request)
	if err != nil {
		t.Error("unexpected error")
		return
	}
}

func TestService_CardOrderChangeFail(t *testing.T) {
	cards := make([]*protoBoard.CardOrder, 0)
	card :=  &protoBoard.CardOrder{Order: 1, CardID: 1}
	cards = append(cards, card)

	request := &protoBoard.CardsOrderInput{
		UserID: 1,
		Cards:  cards,
	}

	cardOrder := make([]models.CardOrder, 0)
	cardOrder = append(cardOrder, models.CardOrder{
		CardID: card.CardID,
		Order:  card.Order,
	})

	input := models.CardsOrderInput{
		UserID: request.UserID,
		Cards:  cardOrder,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	service := &service{boardStorage: mockBoardStorage}

	mockBoardStorage.EXPECT().ChangeCardOrder(input).Return(errStorage)
	_, err := service.CardOrderChange(context.Background(), request)
	if err == nil {
		t.Error("expected error")
		return
	}
}

func TestService_CreateTask(t *testing.T) {
	request := &protoBoard.TaskInput{
		UserID:      1,
		TaskID:      0,
		CardID:      1,
		Name:        "task",
		Description: "lalala",
		Order:       1,
	}

	input := models.TaskInput{
		UserID:      request.UserID,
		CardID:      request.CardID,
		Name:        request.Name,
		Order:       request.Order,
		TaskID:      request.TaskID,
		Description: request.Description,
	}

	internal := models.TaskInternalShort{
		TaskID:      1,
		Name:        input.Name,
		Description: input.Description,
	}

	expect := &protoBoard.TaskOutsideSuperShort{
		TaskID:      internal.TaskID,
		Name:        internal.Name,
		Description: internal.Description,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	service := &service{boardStorage: mockBoardStorage}

	mockBoardStorage.EXPECT().CreateTask(input).Return(internal, nil)
	output, err := service.CreateTask(context.Background(), request)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(output, expect) {
		t.Errorf("results not match, want %v, have %v", expect, output)
		return
	}
}

func TestService_CreateTaskFail(t *testing.T) {
	request := &protoBoard.TaskInput{
		UserID:      1,
		TaskID:      0,
		CardID:      1,
		Name:        "task",
		Description: "lalala",
		Order:       1,
	}

	input := models.TaskInput{
		UserID:      request.UserID,
		CardID:      request.CardID,
		Name:        request.Name,
		Order:       request.Order,
		TaskID:      request.TaskID,
		Description: request.Description,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	service := &service{boardStorage: mockBoardStorage}

	mockBoardStorage.EXPECT().CreateTask(input).Return(models.TaskInternalShort{}, errStorage)
	_, err := service.CreateTask(context.Background(), request)
	if err == nil {
		t.Error("expected error")
		return
	}
}

func TestService_ChangeTask(t *testing.T) {
	request := &protoBoard.TaskInput{
		UserID:      1,
		TaskID:      0,
		CardID:      1,
		Name:        "task",
		Description: "lalala",
		Order:       1,
	}

	input := models.TaskInput{
		UserID:      request.UserID,
		CardID:      request.CardID,
		Name:        request.Name,
		Order:       request.Order,
		TaskID:      request.TaskID,
		Description: request.Description,
	}

	internal := models.TaskInternal{
		TaskID:      1,
		Name:        input.Name,
		Description: input.Description,
	}

	expect := &protoBoard.TaskOutsideSuperShort{
		TaskID:      internal.TaskID,
		Name:        internal.Name,
		Description: internal.Description,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	service := &service{boardStorage: mockBoardStorage}

	mockBoardStorage.EXPECT().ChangeTask(input).Return(internal, nil)
	output, err := service.ChangeTask(context.Background(), request)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(output, expect) {
		t.Errorf("results not match, want %v, have %v", expect, output)
		return
	}
}

func TestService_ChangeTaskFail(t *testing.T) {
	request := &protoBoard.TaskInput{
		UserID:      1,
		TaskID:      0,
		CardID:      1,
		Name:        "task",
		Description: "lalala",
		Order:       1,
	}

	input := models.TaskInput{
		UserID:      request.UserID,
		CardID:      request.CardID,
		Name:        request.Name,
		Order:       request.Order,
		TaskID:      request.TaskID,
		Description: request.Description,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	service := &service{boardStorage: mockBoardStorage}

	mockBoardStorage.EXPECT().ChangeTask(input).Return(models.TaskInternal{}, errStorage)
	_, err := service.ChangeTask(context.Background(), request)
	if err == nil {
		t.Error("expected error")
		return
	}
}

func TestService_DeleteTask(t *testing.T) {
	request := &protoBoard.TaskInput{
		UserID:      1,
		TaskID:      0,
		CardID:      1,
	}

	input := models.TaskInput{
		UserID:      request.UserID,
		CardID:      request.CardID,
		TaskID:      request.TaskID,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	service := &service{boardStorage: mockBoardStorage}

	mockBoardStorage.EXPECT().DeleteTask(input).Return(nil)

	_, err := service.DeleteTask(context.Background(), request)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
}

func TestService_DeleteTaskFail(t *testing.T) {
	request := &protoBoard.TaskInput{
		UserID:      1,
		TaskID:      0,
		CardID:      1,
	}

	input := models.TaskInput{
		UserID:      request.UserID,
		CardID:      request.CardID,
		TaskID:      request.TaskID,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	service := &service{boardStorage: mockBoardStorage}

	mockBoardStorage.EXPECT().DeleteTask(input).Return(errStorage)

	_, err := service.DeleteTask(context.Background(), request)
	if err == nil {
		t.Errorf("expected err: %s", err)
		return
	}
}

func TestService_GetTask(t *testing.T) {

}

func TestService_TasksOrderChange(t *testing.T) {
	tasks := make([]*protoBoard.TaskOrder, 0)
	task :=  &protoBoard.TaskOrder{Order: 1, TaskID: 1}
	tasks = append(tasks, task)
	tasksIn := &protoBoard.TasksOrder{CardID: 1, Tasks: tasks}
	tasksSlice := make([]*protoBoard.TasksOrder, 0)
	tasksSlice = append(tasksSlice, tasksIn)


	request := &protoBoard.TasksOrderInput{
		UserID: 1,
		Tasks:  tasksSlice,
	}

	input := models.CardsOrderInput{
		UserID: request.UserID,
		Cards:  cardOrder,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	service := &service{boardStorage: mockBoardStorage}

	mockBoardStorage.EXPECT().ChangeCardOrder(input).Return(errStorage)
	_, err := service.CardOrderChange(context.Background(), request)
	if err == nil {
		t.Error("expected error")
		return
	}
}

func TestService_AssignUser(t *testing.T) {

}

func TestService_DismissUser(t *testing.T) {

}

func TestService_CreateTag(t *testing.T) {

}

func TestService_ChangeTag(t *testing.T) {

}

func TestService_DeleteTag(t *testing.T) {

}

func TestService_AddTag(t *testing.T) {

}

func TestService_RemoveTag(t *testing.T) {

}

func TestService_CreateComment(t *testing.T) {

}

func TestService_ChangeComment(t *testing.T) {

}
func TestService_DeleteComment(t *testing.T) {

}

func TestService_CreateChecklist(t *testing.T) {

}
func TestService_ChangeChecklist(t *testing.T) {

}

func TestService_DeleteChecklist(t *testing.T) {

}

func TestService_AddAttachment(t *testing.T) {

}

func TestService_RemoveAttachment(t *testing.T) {

}
func TestService_CheckBoardPermission(t *testing.T) {

}

func TestService_CheckCardPermission(t *testing.T) {

}
func TestService_CheckTaskPermission(t *testing.T) {

}

func TestService_CheckCommentPermission(t *testing.T) {

}