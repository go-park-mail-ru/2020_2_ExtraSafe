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

	user := &protoProfile.UserOutsideShort{
		ID:       2,
		Email:    "lallaa",
		Username:  request.MemberName,
		FullName: "lalalla",
		Avatar:   "default",
	}

	board := models.BoardOutsideShort{
		BoardID: input.BoardID,
		Name:    "name",
		Theme:   "theme",
		Star:    false,
	}

	boardOutside := &protoProfile.BoardOutsideShort{Name: board.Name}

	initiator := &protoProfile.UsersOutsideShort{Users: nil}
	initiator.Users = append(initiator.Users, &protoProfile.UserOutsideShort{ID: request.UserID})

	userProfile := &protoBoard.BoardMemberOutside{
		Board:     boardOutside,
		User:      user,
		Initiator: initiator.Users[0].Username,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	ctrlUser := gomock.NewController(t)
	defer ctrlUser.Finish()
	mockProfileService := serviceMocks.NewMockProfileClient(ctrlUser)

	mockProfileService.EXPECT().
						GetUserByUsername(context.Background(), &protoProfile.UserName{UserName: request.MemberName}).
						Return(user, nil)

	mockProfileService.
		EXPECT().
		GetUsersByIDs(context.Background(), &protoProfile.UserIDS{UserIDs: []int64{input.UserID}}).
		Return(initiator, nil)

	mockBoardStorage.EXPECT().AddUser(input).Return(nil)

	mockBoardStorage.
		EXPECT().
		GetBoardShort(models.BoardInput{
		UserID:    input.UserID,
		BoardID:   input.BoardID,
	}).
		Return(board, nil)

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

	user := &protoProfile.UserOutsideShort{
		ID:       2,
		Email:    "lallaa",
		Username:  request.MemberName,
		FullName: "lalalla",
		Avatar:   "default",
	}

	initiator := &protoProfile.UsersOutsideShort{Users: nil}
	initiator.Users = append(initiator.Users, &protoProfile.UserOutsideShort{ID: request.UserID})

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	ctrlUser := gomock.NewController(t)
	defer ctrlUser.Finish()
	mockProfileService := serviceMocks.NewMockProfileClient(ctrlUser)

	mockProfileService.EXPECT().
		GetUserByUsername(context.Background(), &protoProfile.UserName{UserName: request.MemberName}).
		Return(user, nil)

	mockProfileService.
		EXPECT().
		GetUsersByIDs(context.Background(), &protoProfile.UserIDS{UserIDs: []int64{input.UserID}}).
		Return(initiator, nil)

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

	task := models.TaskInternalShort{TaskID: input.TaskID, CardID: input.CardID}
	output := &protoBoard.TaskOutsideSuperShort{
		TaskID: task.TaskID,
		CardID:   task.CardID,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	service := &service{boardStorage: mockBoardStorage}

	mockBoardStorage.EXPECT().DeleteTask(input).Return(task, nil)

	res, err := service.DeleteTask(context.Background(), request)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(output, res) {
		t.Errorf("results not match, want %v, have %v", output, res)
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

	mockBoardStorage.EXPECT().DeleteTask(input).Return(models.TaskInternalShort{}, errStorage)

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

	taskOrd := models.TaskOrder{
		TaskID: task.TaskID,
		Order:  task.Order,
	}
	tasksOrd := models.TasksOrder{
		CardID: tasksIn.CardID,
		Tasks:  []models.TaskOrder{taskOrd},
	}
	tasksOrder := models.TasksOrderInput{
		UserID: 1,
		Tasks:  []models.TasksOrder{tasksOrd},
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	service := &service{boardStorage: mockBoardStorage}

	mockBoardStorage.EXPECT().ChangeTaskOrder(tasksOrder).Return(nil)
	_, err := service.TasksOrderChange(context.Background(), request)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
}

func TestService_TasksOrderChangeFail(t *testing.T) {
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

	taskOrd := models.TaskOrder{
		TaskID: task.TaskID,
		Order:  task.Order,
	}
	tasksOrd := models.TasksOrder{
		CardID: tasksIn.CardID,
		Tasks:  []models.TaskOrder{taskOrd},
	}
	tasksOrder := models.TasksOrderInput{
		UserID: 1,
		Tasks:  []models.TasksOrder{tasksOrd},
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	service := &service{boardStorage: mockBoardStorage}

	mockBoardStorage.EXPECT().ChangeTaskOrder(tasksOrder).Return(errStorage)
	_, err := service.TasksOrderChange(context.Background(), request)
	if err == nil {
		t.Errorf("expected err: %s", err)
		return
	}
}

func TestService_AssignUser(t *testing.T) {
	request := &protoBoard.TaskAssignerInput{
		TaskID: 1,
		AssignerName: "pkaterina",
		UserID: 1,
	}

	task := models.TaskAssignUserOutside {
		TaskID: request.TaskID,
		CardID: 1,
		TaskName: "ala",
	}

	user := &protoProfile.UserOutsideShort{
		ID:       2,
		Email:    "lala",
		Username: "pkaterina",
		FullName: "lala",
		Avatar:   "",
	}

	initiator := &protoProfile.UsersOutsideShort{Users: nil}
	initiator.Users = append(initiator.Users, &protoProfile.UserOutsideShort{ID: request.UserID})

	output := &protoBoard.TaskAssignUserOutside{
		Assigner: user,
		TaskID:   task.TaskID,
		CardID:   task.CardID,
		Initiator: initiator.Users[0].Username,
		TaskName: task.TaskName,
	}

	input := models.TaskAssigner{
		UserID:     request.UserID,
		TaskID:     request.TaskID,
		AssignerID: user.ID,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	ctrlProfile := gomock.NewController(t)
	defer ctrlProfile.Finish()
	mockProfileService := serviceMocks.NewMockProfileClient(ctrlProfile)

	service := &service{boardStorage: mockBoardStorage, profileService: mockProfileService}

	mockProfileService.
		EXPECT().
		GetUserByUsername(context.Background(), &protoProfile.UserName{UserName: request.AssignerName}).
		Return(user, nil)

	mockBoardStorage.
		EXPECT().
		AssignUser(input).
		Return(task, nil)

	mockProfileService.
		EXPECT().
		GetUsersByIDs(context.Background(), &protoProfile.UserIDS{UserIDs: []int64{input.UserID}}).
		Return(initiator, nil)

	res, err := service.AssignUser(context.Background(), request)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(output, res) {
		t.Errorf("results not match, want %v, have %v", output, res)
		return
	}
}

func TestService_AssignUserGetUserFail(t *testing.T) {
	request := &protoBoard.TaskAssignerInput{
		TaskID: 1,
		AssignerName: "pkaterina",
		UserID: 1,
	}

	ctrlProfile := gomock.NewController(t)
	defer ctrlProfile.Finish()
	mockProfileService := serviceMocks.NewMockProfileClient(ctrlProfile)

	service := &service{profileService: mockProfileService}

	mockProfileService.
		EXPECT().
		GetUserByUsername(context.Background(), &protoProfile.UserName{UserName: request.AssignerName}).
		Return(&protoProfile.UserOutsideShort{}, errors.New(""))

	_, err := service.AssignUser(context.Background(), request)
	if err == nil {
		t.Error("expected error")
		return
	}
}

func TestService_AssignUserAssigningFail(t *testing.T) {
	request := &protoBoard.TaskAssignerInput{
		TaskID: 1,
		AssignerName: "pkaterina",
		UserID: 1,
	}

	expect := &protoProfile.UserOutsideShort{
		ID:       1,
		Email:    "lala",
		Username: "pkaterina",
		FullName: "lala",
		Avatar:   "",
	}

	input := models.TaskAssigner{
		UserID:     request.UserID,
		TaskID:     request.TaskID,
		AssignerID: expect.ID,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	ctrlProfile := gomock.NewController(t)
	defer ctrlProfile.Finish()
	mockProfileService := serviceMocks.NewMockProfileClient(ctrlProfile)

	service := &service{boardStorage: mockBoardStorage, profileService: mockProfileService}

	mockProfileService.
		EXPECT().
		GetUserByUsername(context.Background(), &protoProfile.UserName{UserName: request.AssignerName}).
		Return(expect, nil)

	mockBoardStorage.
		EXPECT().
		AssignUser(input).
		Return(models.TaskAssignUserOutside{}, errStorage)

	_, err := service.AssignUser(context.Background(), request)
	if err == nil {
		t.Error("expected error")
		return
	}
}

func TestService_DismissUser(t *testing.T) {
	request := &protoBoard.TaskAssignerInput{
		TaskID: 1,
		AssignerName: "pkaterina",
		UserID: 1,
	}

	expect := &protoProfile.UserOutsideShort{
		ID:       1,
	}

	input := models.TaskAssigner{
		UserID:     request.UserID,
		TaskID:     request.TaskID,
		AssignerID: expect.ID,
	}

	task := models.TaskAssignUserOutside {
		TaskID: request.TaskID,
		CardID: 1,
		TaskName: "ala",
	}

	output := &protoBoard.TaskAssignUserOutside{
		Assigner: expect,
		TaskID:   task.TaskID,
		CardID:   task.CardID,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	ctrlProfile := gomock.NewController(t)
	defer ctrlProfile.Finish()
	mockProfileService := serviceMocks.NewMockProfileClient(ctrlProfile)

	service := &service{boardStorage: mockBoardStorage, profileService: mockProfileService}

	mockProfileService.
		EXPECT().
		GetUserByUsername(context.Background(), &protoProfile.UserName{UserName: request.AssignerName}).
		Return(expect, nil)

	mockBoardStorage.
		EXPECT().
		DismissUser(input).
		Return(task, nil)

	res, err := service.DismissUser(context.Background(), request)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(output, res) {
		t.Errorf("results not match, want %v, have %v", output, res)
		return
	}
}

func TestService_DismissUserGetUserFail(t *testing.T) {
	request := &protoBoard.TaskAssignerInput{
		TaskID: 1,
		AssignerName: "pkaterina",
		UserID: 1,
	}

	ctrlProfile := gomock.NewController(t)
	defer ctrlProfile.Finish()
	mockProfileService := serviceMocks.NewMockProfileClient(ctrlProfile)

	service := &service{profileService: mockProfileService}

	mockProfileService.
		EXPECT().
		GetUserByUsername(context.Background(), &protoProfile.UserName{UserName: request.AssignerName}).
		Return(&protoProfile.UserOutsideShort{}, errors.New(""))

	_, err := service.DismissUser(context.Background(), request)
	if err == nil {
		t.Errorf("expected err: %s", err)
		return
	}
}

func TestService_DismissUserDismissFail(t *testing.T) {
	request := &protoBoard.TaskAssignerInput{
		TaskID: 1,
		AssignerName: "pkaterina",
		UserID: 1,
	}

	expect := &protoProfile.UserOutsideShort{
		ID:       1,
	}

	input := models.TaskAssigner{
		UserID:     request.UserID,
		TaskID:     request.TaskID,
		AssignerID: expect.ID,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	ctrlProfile := gomock.NewController(t)
	defer ctrlProfile.Finish()
	mockProfileService := serviceMocks.NewMockProfileClient(ctrlProfile)

	service := &service{boardStorage: mockBoardStorage, profileService: mockProfileService}

	mockProfileService.
		EXPECT().
		GetUserByUsername(context.Background(), &protoProfile.UserName{UserName: request.AssignerName}).
		Return(expect, nil)

	mockBoardStorage.
		EXPECT().
		DismissUser(input).
		Return(models.TaskAssignUserOutside{}, errStorage)

	_, err := service.DismissUser(context.Background(), request)
	if err == nil {
		t.Errorf("expected err")
		return
	}
}

func TestService_CreateTag(t *testing.T) {
	request := &protoBoard.TagInput{
		UserID:  1,
		TaskID:  1,
		TagID:   0,
		BoardID: 1,
		Color:   "some",
		Name:    "some",
	}

	input := models.TagInput{
		UserID:  request.UserID,
		BoardID: request.BoardID,
		Color:   request.Color,
		Name:    request.Name,
	}

	internal := models.TagOutside{
		TagID: 1,
		Color: input.Color,
		Name:  input.Name,
	}

	expect := &protoBoard.TagOutside{
		TagID: internal.TagID,
		Color: internal.Color,
		Name:  internal.Name,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	service :=  &service{boardStorage: mockBoardStorage}
	mockBoardStorage.EXPECT().CreateTag(input).Return(internal, nil)

	output, err := service.CreateTag(context.Background(), request)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(output, expect) {
		t.Errorf("results not match, want %v, have %v", expect, output)
		return
	}

}

func TestService_CreateTagFail(t *testing.T) {
	request := &protoBoard.TagInput{
		UserID:  1,
		TaskID:  1,
		TagID:   0,
		BoardID: 1,
		Color:   "some",
		Name:    "some",
	}

	input := models.TagInput{
		UserID:  request.UserID,
		BoardID: request.BoardID,
		Color:   request.Color,
		Name:    request.Name,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	service :=  &service{boardStorage: mockBoardStorage}
	mockBoardStorage.EXPECT().CreateTag(input).Return(models.TagOutside{}, errStorage)

	_, err := service.CreateTag(context.Background(), request)
	if err == nil {
		t.Error("expected error")
		return
	}
}

func TestService_ChangeTag(t *testing.T) {
	request := &protoBoard.TagInput{
		UserID:  1,
		TaskID:  1,
		TagID:   0,
		BoardID: 1,
		Color:   "some",
		Name:    "some",
	}

	input := models.TagInput{
		UserID:  request.UserID,
		TagID: request.TagID,
		BoardID: request.BoardID,
		Color:   request.Color,
		Name:    request.Name,
	}

	internal := models.TagOutside{
		TagID: 1,
		Color: input.Color,
		Name:  input.Name,
	}

	expect := &protoBoard.TagOutside{
		TagID: internal.TagID,
		Color: internal.Color,
		Name:  internal.Name,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	service :=  &service{boardStorage: mockBoardStorage}
	mockBoardStorage.EXPECT().UpdateTag(input).Return(internal, nil)

	output, err := service.ChangeTag(context.Background(), request)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(output, expect) {
		t.Errorf("results not match, want %v, have %v", expect, output)
		return
	}
}

func TestService_ChangeTagFail(t *testing.T) {
	request := &protoBoard.TagInput{
		UserID:  1,
		TaskID:  1,
		TagID:   0,
		BoardID: 1,
		Color:   "some",
		Name:    "some",
	}

	input := models.TagInput{
		UserID:  request.UserID,
		BoardID: request.BoardID,
		Color:   request.Color,
		Name:    request.Name,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	service :=  &service{boardStorage: mockBoardStorage}
	mockBoardStorage.EXPECT().UpdateTag(input).Return(models.TagOutside{}, errStorage)

	_, err := service.ChangeTag(context.Background(), request)
	if err == nil {
		t.Error("expected error")
		return
	}
}

func TestService_DeleteTag(t *testing.T) {
	request := &protoBoard.TagInput{
		UserID:  1,
		TagID:   0,
		BoardID: 1,
	}

	input := models.TagInput{
		UserID:  request.UserID,
		TagID: request.TagID,
		BoardID: request.BoardID,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	service :=  &service{boardStorage: mockBoardStorage}
	mockBoardStorage.EXPECT().DeleteTag(input).Return(nil)

	_, err := service.DeleteTag(context.Background(), request)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
}

func TestService_DeleteTagFail(t *testing.T) {
	request := &protoBoard.TagInput{
		UserID:  1,
		TagID:   0,
		BoardID: 1,
	}

	input := models.TagInput{
		UserID:  request.UserID,
		TagID: request.TagID,
		BoardID: request.BoardID,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	service :=  &service{boardStorage: mockBoardStorage}
	mockBoardStorage.EXPECT().DeleteTag(input).Return(errStorage)

	_, err := service.DeleteTag(context.Background(), request)
	if err == nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
}

func TestService_AddTag(t *testing.T) {
	request := &protoBoard.TaskTagInput{
		UserID:  1,
		TaskID:  1,
		TagID:   0,
	}

	input := models.TaskTagInput{
		UserID:  request.UserID,
		TagID: request.TagID,
		TaskID: request.TaskID,
	}

	tag := models.TagOutside{
		CardID: 1,
		TaskID: request.TaskID,
		TagID:  1,
		Color:  "lala",
		Name:   "lala",
	}

	output := &protoBoard.TagOutside{
		TaskID: tag.TaskID,
		CardID: tag.CardID,
		TagID: tag.TagID,
		Color: tag.Color,
		Name:  tag.Name,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	service :=  &service{boardStorage: mockBoardStorage}
	mockBoardStorage.EXPECT().AddTag(input).Return(tag, nil)

	res, err := service.AddTag(context.Background(), request)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(res, output) {
		t.Errorf("results not match, want %v, have %v", output, res)
		return
	}
}

func TestService_AddTagFail(t *testing.T) {
	request := &protoBoard.TaskTagInput{
		UserID:  1,
		TaskID:  1,
		TagID:   0,
	}

	input := models.TaskTagInput{
		UserID:  request.UserID,
		TagID: request.TagID,
		TaskID: request.TaskID,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	service :=  &service{boardStorage: mockBoardStorage}
	mockBoardStorage.EXPECT().AddTag(input).Return(models.TagOutside{}, errStorage)

	_, err := service.AddTag(context.Background(), request)
	if err == nil {
		t.Errorf("expected err: %s", err)
		return
	}
}
func TestService_RemoveTag(t *testing.T) {
	request := &protoBoard.TaskTagInput{
		UserID:  1,
		TaskID:  1,
		TagID:   0,
	}

	input := models.TaskTagInput{
		UserID:  request.UserID,
		TagID: request.TagID,
		TaskID: request.TaskID,
	}

	tag := models.TagOutside{
		CardID: 1,
		TaskID: request.TaskID,
		TagID:  1,
		Color:  "lala",
		Name:   "lala",
	}

	output := &protoBoard.TagOutside{
		TaskID: tag.TaskID,
		CardID: tag.CardID,
		TagID: tag.TagID,
		Color: tag.Color,
		Name:  tag.Name,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	service :=  &service{boardStorage: mockBoardStorage}
	mockBoardStorage.EXPECT().RemoveTag(input).Return(tag,nil)

	res, err := service.RemoveTag(context.Background(), request)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(res, output) {
		t.Errorf("results not match, want %v, have %v", output, res)
		return
	}
}

func TestService_RemoveTagFail(t *testing.T) {
	request := &protoBoard.TaskTagInput{
		UserID:  1,
		TaskID:  1,
		TagID:   0,
	}

	input := models.TaskTagInput{
		UserID:  request.UserID,
		TagID: request.TagID,
		TaskID: request.TaskID,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	service :=  &service{boardStorage: mockBoardStorage}
	mockBoardStorage.EXPECT().RemoveTag(input).Return(models.TagOutside{}, errStorage)

	_, err := service.RemoveTag(context.Background(), request)
	if err == nil {
		t.Errorf("expected err: %s", err)
		return
	}
}

func TestService_CreateComment(t *testing.T) {
	request := &protoBoard.CommentInput{
		CommentID: 0,
		TaskID:    1,
		Message:   "lala",
		Order:     1,
		UserID:    1,
	}

	input := models.CommentInput{
		CommentID: request.CommentID,
		TaskID:    request.TaskID,
		Message:   request.Message,
		Order:     request.Order,
		UserID:    request.UserID,
	}

	internal := models.CommentOutside{
		CommentID: request.CommentID,
		Message:   request.Message,
		Order:     request.Order,
	}

	internalUser := &protoProfile.UsersOutsideShort{Users: nil}
	internalUser.Users = append(internalUser.Users, &protoProfile.UserOutsideShort{ID: 1})

	expect := &protoBoard.CommentOutside{
		CommentID: internal.CommentID,
		Message:   internal.Message,
		Order:     internal.Order,
		User:      internalUser.Users[0],
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	ctrlProfile := gomock.NewController(t)
	defer ctrlProfile.Finish()
	mockProfileService := serviceMocks.NewMockProfileClient(ctrlProfile)

	service :=  &service{boardStorage: mockBoardStorage, profileService: mockProfileService}

	mockBoardStorage.EXPECT().CreateComment(input).Return(internal, nil)
	mockProfileService.EXPECT().GetUsersByIDs(context.Background(), &protoProfile.UserIDS{UserIDs: []int64{input.UserID}}).Return(internalUser, nil)

	output, err := service.CreateComment(context.Background(), request)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(output, expect) {
		t.Errorf("results not match, want %v, have %v", expect, output)
		return
	}
}

func TestService_CreateCommentFail(t *testing.T) {
	request := &protoBoard.CommentInput{
		CommentID: 0,
		TaskID:    1,
		Message:   "lala",
		Order:     1,
		UserID:    1,
	}

	input := models.CommentInput{
		CommentID: request.CommentID,
		TaskID:    request.TaskID,
		Message:   request.Message,
		Order:     request.Order,
		UserID:    request.UserID,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	service :=  &service{boardStorage: mockBoardStorage}

	mockBoardStorage.EXPECT().CreateComment(input).Return(models.CommentOutside{}, errStorage)

	_, err := service.CreateComment(context.Background(), request)
	if err == nil {
		t.Error("expected error")
		return
	}
}

func TestService_CreateCommentGetUserFail(t *testing.T) {
	request := &protoBoard.CommentInput{
		CommentID: 0,
		TaskID:    1,
		Message:   "lala",
		Order:     1,
		UserID:    1,
	}

	input := models.CommentInput{
		CommentID: request.CommentID,
		TaskID:    request.TaskID,
		Message:   request.Message,
		Order:     request.Order,
		UserID:    request.UserID,
	}

	internal := models.CommentOutside{
		CommentID: request.CommentID,
		Message:   request.Message,
		Order:     request.Order,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	ctrlProfile := gomock.NewController(t)
	defer ctrlProfile.Finish()
	mockProfileService := serviceMocks.NewMockProfileClient(ctrlProfile)

	service :=  &service{boardStorage: mockBoardStorage, profileService: mockProfileService}

	mockBoardStorage.EXPECT().CreateComment(input).Return(internal, nil)
	mockProfileService.EXPECT().GetUsersByIDs(context.Background(), &protoProfile.UserIDS{UserIDs: []int64{input.UserID}}).Return(&protoProfile.UsersOutsideShort{}, errors.New(""))

	_, err := service.CreateComment(context.Background(), request)
	if err == nil {
		t.Error("expected error")
		return
	}
}

func TestService_ChangeComment(t *testing.T) {
	request := &protoBoard.CommentInput{
		CommentID: 0,
		TaskID:    1,
		Message:   "lala",
		Order:     1,
		UserID:    1,
	}

	input := models.CommentInput{
		CommentID: request.CommentID,
		TaskID:    request.TaskID,
		Message:   request.Message,
		Order:     request.Order,
		UserID:    request.UserID,
	}

	internal := models.CommentOutside{
		CommentID: request.CommentID,
		Message:   request.Message,
		Order:     request.Order,
	}

	internalUser := &protoProfile.UsersOutsideShort{Users: nil}
	internalUser.Users = append(internalUser.Users, &protoProfile.UserOutsideShort{ID: 1})

	expect := &protoBoard.CommentOutside{
		CommentID: internal.CommentID,
		Message:   internal.Message,
		Order:     internal.Order,
		User:      internalUser.Users[0],
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	ctrlProfile := gomock.NewController(t)
	defer ctrlProfile.Finish()
	mockProfileService := serviceMocks.NewMockProfileClient(ctrlProfile)

	service :=  &service{boardStorage: mockBoardStorage, profileService: mockProfileService}

	mockBoardStorage.EXPECT().UpdateComment(input).Return(internal, nil)
	mockProfileService.EXPECT().GetUsersByIDs(context.Background(), &protoProfile.UserIDS{UserIDs: []int64{input.UserID}}).Return(internalUser, nil)

	output, err := service.ChangeComment(context.Background(), request)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(output, expect) {
		t.Errorf("results not match, want %v, have %v", expect, output)
		return
	}
}

func TestService_ChangeCommentFail(t *testing.T) {
	request := &protoBoard.CommentInput{
		CommentID: 0,
		TaskID:    1,
		Message:   "lala",
		Order:     1,
		UserID:    1,
	}

	input := models.CommentInput{
		CommentID: request.CommentID,
		TaskID:    request.TaskID,
		Message:   request.Message,
		Order:     request.Order,
		UserID:    request.UserID,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	service :=  &service{boardStorage: mockBoardStorage}

	mockBoardStorage.EXPECT().UpdateComment(input).Return(models.CommentOutside{}, errStorage)

	_, err := service.ChangeComment(context.Background(), request)
	if err == nil {
		t.Errorf("expected err: %s", err)
		return
	}
}

func TestService_ChangeCommentGetUserFail(t *testing.T) {
	request := &protoBoard.CommentInput{
		CommentID: 0,
		TaskID:    1,
		Message:   "lala",
		Order:     1,
		UserID:    1,
	}

	input := models.CommentInput{
		CommentID: request.CommentID,
		TaskID:    request.TaskID,
		Message:   request.Message,
		Order:     request.Order,
		UserID:    request.UserID,
	}

	internal := models.CommentOutside{
		CommentID: request.CommentID,
		Message:   request.Message,
		Order:     request.Order,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	ctrlProfile := gomock.NewController(t)
	defer ctrlProfile.Finish()
	mockProfileService := serviceMocks.NewMockProfileClient(ctrlProfile)

	service :=  &service{boardStorage: mockBoardStorage, profileService: mockProfileService}

	mockBoardStorage.EXPECT().UpdateComment(input).Return(internal, nil)
	mockProfileService.EXPECT().GetUsersByIDs(context.Background(), &protoProfile.UserIDS{UserIDs: []int64{input.UserID}}).Return(&protoProfile.UsersOutsideShort{}, errors.New(""))

	_, err := service.ChangeComment(context.Background(), request)
	if err == nil {
		t.Errorf("expected err: %s", err)
		return
	}
}

func TestService_DeleteComment(t *testing.T) {
	request := &protoBoard.CommentInput{
		CommentID: 0,
		TaskID:    1,
		Message:   "lala",
		Order:     1,
		UserID:    1,
	}

	input := models.CommentInput{
		CommentID: request.CommentID,
		TaskID:    request.TaskID,
		Message:   request.Message,
		Order:     request.Order,
		UserID:    request.UserID,
	}

	comment := models.CommentOutside{
		CommentID: request.CommentID,
		TaskID:    request.TaskID,
		CardID:    1,
	}

	output := &protoBoard.CommentOutside{
		CommentID: comment.CommentID,
		CardID: comment.CardID,
		TaskID: comment.TaskID,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	service := &service{boardStorage: mockBoardStorage}

	mockBoardStorage.EXPECT().DeleteComment(input).Return(comment, nil)

	res, err := service.DeleteComment(context.Background(), request)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(output, res) {
		t.Errorf("results not match, want %v, have %v", output, res)
		return
	}
}

func TestService_DeleteCommentFail(t *testing.T) {
	request := &protoBoard.CommentInput{
		CommentID: 0,
		TaskID:    1,
		Message:   "lala",
		Order:     1,
		UserID:    1,
	}

	input := models.CommentInput{
		CommentID: request.CommentID,
		TaskID:    request.TaskID,
		Message:   request.Message,
		Order:     request.Order,
		UserID:    request.UserID,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	service := &service{boardStorage: mockBoardStorage}

	mockBoardStorage.EXPECT().DeleteComment(input).Return(models.CommentOutside{}, errStorage)

	_, err := service.DeleteComment(context.Background(), request)
	if err == nil {
		t.Errorf("expected err")
		return
	}
}

func TestService_CreateChecklist(t *testing.T) {
	request := &protoBoard.ChecklistInput{
		UserID:      1,
		ChecklistID: 0,
		TaskID:      1,
		Name:        "lala",
		Items:       nil,
	}

	input := models.ChecklistInput{
		UserID:      request.UserID,
		ChecklistID: request.ChecklistID,
		TaskID:      request.TaskID,
		Name:        request.Name,
		Items:       request.Items,
	}

	internal := models.ChecklistOutside{
		ChecklistID: 1,
		Name:        input.Name,
		Items:       input.Items,
	}

	expect := &protoBoard.ChecklistOutside{
		ChecklistID: internal.ChecklistID,
		Name:        internal.Name,
		Items:       internal.Items,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	service := &service{boardStorage: mockBoardStorage}

	mockBoardStorage.EXPECT().CreateChecklist(input).Return(internal, nil)

	output, err := service.CreateChecklist(context.Background(), request)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(output, expect) {
		t.Errorf("results not match, want %v, have %v", expect, output)
		return
	}
}

func TestService_CreateChecklistFail(t *testing.T) {
	request := &protoBoard.ChecklistInput{
		UserID:      1,
		ChecklistID: 0,
		TaskID:      1,
		Name:        "lala",
		Items:       nil,
	}

	input := models.ChecklistInput{
		UserID:      request.UserID,
		ChecklistID: request.ChecklistID,
		TaskID:      request.TaskID,
		Name:        request.Name,
		Items:       request.Items,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	service := &service{boardStorage: mockBoardStorage}

	mockBoardStorage.EXPECT().CreateChecklist(input).Return(models.ChecklistOutside{}, errStorage)

	_, err := service.CreateChecklist(context.Background(), request)
	if err == nil {
		t.Errorf("expected err: %s", err)
		return
	}
}

func TestService_ChangeChecklist(t *testing.T) {
	request := &protoBoard.ChecklistInput{
		UserID:      1,
		ChecklistID: 0,
		TaskID:      1,
		Name:        "lala",
		Items:       nil,
	}

	input := models.ChecklistInput{
		UserID:      request.UserID,
		ChecklistID: request.ChecklistID,
		TaskID:      request.TaskID,
		Name:        request.Name,
		Items:       request.Items,
	}

	internal := models.ChecklistOutside{
		ChecklistID: 1,
		Name:        input.Name,
		Items:       input.Items,
	}

	expect := &protoBoard.ChecklistOutside{
		ChecklistID: internal.ChecklistID,
		Name:        internal.Name,
		Items:       internal.Items,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	service := &service{boardStorage: mockBoardStorage}

	mockBoardStorage.EXPECT().UpdateChecklist(input).Return(internal, nil)

	output, err := service.ChangeChecklist(context.Background(), request)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(output, expect) {
		t.Errorf("results not match, want %v, have %v", expect, output)
		return
	}
}

func TestService_ChangeChecklistFail(t *testing.T) {
	request := &protoBoard.ChecklistInput{
		UserID:      1,
		ChecklistID: 0,
		TaskID:      1,
		Name:        "lala",
		Items:       nil,
	}

	input := models.ChecklistInput{
		UserID:      request.UserID,
		ChecklistID: request.ChecklistID,
		TaskID:      request.TaskID,
		Name:        request.Name,
		Items:       request.Items,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	service := &service{boardStorage: mockBoardStorage}

	mockBoardStorage.EXPECT().UpdateChecklist(input).Return(models.ChecklistOutside{}, errStorage)

	_, err := service.ChangeChecklist(context.Background(), request)
	if err == nil {
		t.Errorf("expected err: %s", err)
		return
	}
}

func TestService_DeleteChecklist(t *testing.T) {
	request := &protoBoard.ChecklistInput{
		UserID:      1,
		ChecklistID: 0,
		TaskID:      1,
		Name:        "lala",
		Items:       nil,
	}

	input := models.ChecklistInput{
		UserID:      request.UserID,
		ChecklistID: request.ChecklistID,
		TaskID:      request.TaskID,
		Name:        request.Name,
		Items:       request.Items,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	service := &service{boardStorage: mockBoardStorage}

	mockBoardStorage.EXPECT().DeleteChecklist(input).Return(models.ChecklistOutside{}, nil)

	_, err := service.DeleteChecklist(context.Background(), request)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
}

func TestService_DeleteChecklistFail(t *testing.T) {
	request := &protoBoard.ChecklistInput{
		UserID:      1,
		ChecklistID: 0,
		TaskID:      1,
		Name:        "lala",
		Items:       nil,
	}

	input := models.ChecklistInput{
		UserID:      request.UserID,
		ChecklistID: request.ChecklistID,
		TaskID:      request.TaskID,
		Name:        request.Name,
		Items:       request.Items,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	service := &service{boardStorage: mockBoardStorage}

	mockBoardStorage.EXPECT().DeleteChecklist(input).Return(models.ChecklistOutside{}, errStorage)

	_, err := service.DeleteChecklist(context.Background(), request)
	if err == nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
}
/*
func TestService_AddAttachment(t *testing.T) {
	request := &protoBoard.AttachmentInput{
		UserID:       1,
		TaskID:       1,
		AttachmentID: 0,
		Filename:     "file",
		File:         nil,
	}

	fileInput := models.AttachmentFileInternal{
		UserID:   request.UserID,
		Filename: request.Filename,
		File:     request.File,
	}

	userInput := &models.AttachmentInternal{
		TaskID:       request.TaskID,
		Filename:     request.Filename,
	}

	internal := models.AttachmentOutside{
		AttachmentID: 1,
		Filename:     request.Filename,
		Filepath:     "../../../",
	}

	expect := &protoBoard.AttachmentOutside{
		AttachmentID: internal.AttachmentID,
		Filename:     internal.Filename,
		Filepath:     internal.Filepath,
	}

	ctrlFile := gomock.NewController(t)
	defer ctrlFile.Finish()
	mockFileStorage := mocks.NewMockFileStorage(ctrlFile)

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	ctrlProfile := gomock.NewController(t)
	defer ctrlProfile.Finish()
	mockProfileService := serviceMocks.NewMockProfileClient(ctrlProfile)

	service := NewService(mockBoardStorage, mockFileStorage, mockProfileService)

	mockFileStorage.EXPECT().UploadFile(fileInput, userInput, false).Return(nil)
	mockBoardStorage.EXPECT().AddAttachment(*userInput).Return(internal, nil)

	output, err := service.AddAttachment(context.Background(), request)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(output, expect) {
		t.Errorf("results not match, want %v, have %v", expect, output)
		return
	}
}

func TestService_AddAttachmentUploadFail(t *testing.T) {
	request := &protoBoard.AttachmentInput{
		UserID:       1,
		TaskID:       1,
		AttachmentID: 0,
		Filename:     "file",
		File:         nil,
	}

	fileInput := models.AttachmentFileInternal{
		UserID:   request.UserID,
		Filename: request.Filename,
		File:     request.File,
	}

	userInput := &models.AttachmentInternal{
		TaskID:       request.TaskID,
		Filename:     request.Filename,
	}


	ctrlFile := gomock.NewController(t)
	defer ctrlFile.Finish()
	mockFileStorage := mocks.NewMockFileStorage(ctrlFile)

	service := &service{fileStorage: mockFileStorage}

	mockFileStorage.EXPECT().UploadFile(fileInput, userInput, false).Return(errStorage)

	_, err := service.AddAttachment(context.Background(), request)
	if err == nil {
		t.Errorf("expected err: %s", err)
		return
	}
}

func TestService_AddAttachmentAttachFail(t *testing.T) {
	request := &protoBoard.AttachmentInput{
		UserID:       1,
		TaskID:       1,
		AttachmentID: 0,
		Filename:     "file",
		File:         nil,
	}

	fileInput := models.AttachmentFileInternal{
		UserID:   request.UserID,
		Filename: request.Filename,
		File:     request.File,
	}

	userInput := &models.AttachmentInternal{
		TaskID:       request.TaskID,
		Filename:     request.Filename,
	}

	ctrlFile := gomock.NewController(t)
	defer ctrlFile.Finish()
	mockFileStorage := mocks.NewMockFileStorage(ctrlFile)

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	service := &service{fileStorage: mockFileStorage, boardStorage: mockBoardStorage}

	mockFileStorage.EXPECT().UploadFile(fileInput, userInput, false).Return(nil)
	mockBoardStorage.EXPECT().AddAttachment(*userInput).Return(models.AttachmentOutside{}, errStorage)

	_, err := service.AddAttachment(context.Background(), request)
	if err == nil {
		t.Errorf("expected err: %s", err)
		return
	}
}*/

func TestService_RemoveAttachment(t *testing.T) {
	request := &protoBoard.AttachmentInfo{
		TaskID:       1,
		AttachmentID: 1,
		Filename:     "file",
	}

	input := models.AttachmentInternal{
		TaskID:       request.TaskID,
		Filename:     request.Filename,
		AttachmentID: request.AttachmentID,
	}

	ctrlFile := gomock.NewController(t)
	defer ctrlFile.Finish()
	mockFileStorage := mocks.NewMockFileStorage(ctrlFile)

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	ctrlProfile := gomock.NewController(t)
	defer ctrlProfile.Finish()
	mockProfileService := serviceMocks.NewMockProfileClient(ctrlProfile)

	service := NewService(mockBoardStorage, mockFileStorage, mockProfileService)

	mockBoardStorage.EXPECT().RemoveAttachment(input).Return(models.AttachmentOutside{}, nil)
	mockFileStorage.EXPECT().DeleteFile(input.Filename, false).Return(nil)

	_, err := service.RemoveAttachment(context.Background(), request)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
}

func TestService_RemoveAttachmentFail(t *testing.T) {
	request := &protoBoard.AttachmentInfo{
		TaskID:       1,
		AttachmentID: 1,
		Filename:     "file",
	}

	input := models.AttachmentInternal{
		TaskID:       request.TaskID,
		Filename:     request.Filename,
		AttachmentID: request.AttachmentID,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	service := &service{boardStorage: mockBoardStorage}

	mockBoardStorage.EXPECT().RemoveAttachment(input).Return(models.AttachmentOutside{}, errStorage)

	_, err := service.RemoveAttachment(context.Background(), request)
	if err == nil {
		t.Errorf("expected err: %s", err)
		return
	}
}

func TestService_RemoveAttachmentDeleteFail(t *testing.T) {
	request := &protoBoard.AttachmentInfo{
		TaskID:       1,
		AttachmentID: 1,
		Filename:     "file",
	}

	input := models.AttachmentInternal{
		TaskID:       request.TaskID,
		Filename:     request.Filename,
		AttachmentID: request.AttachmentID,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	ctrlFile := gomock.NewController(t)
	defer ctrlFile.Finish()
	mockFileStorage := mocks.NewMockFileStorage(ctrlFile)

	service := &service{boardStorage: mockBoardStorage, fileStorage: mockFileStorage}

	mockBoardStorage.EXPECT().RemoveAttachment(input).Return(models.AttachmentOutside{}, nil)
	mockFileStorage.EXPECT().DeleteFile(input.Filename, false).Return(errStorage)

	_, err := service.RemoveAttachment(context.Background(), request)
	if err == nil {
		t.Errorf("expected err: %s", err)
		return
	}
}

func TestService_CheckBoardPermission(t *testing.T) {
	input := &protoBoard.CheckPermissions{
		UserID:    1,
		ElementID: 1,
		IfAdmin:   false,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	expect := &protoBoard.Nothing{Dummy: true}
	service := &service{boardStorage: mockBoardStorage}

	mockBoardStorage.EXPECT().CheckBoardPermission(input.UserID, input.ElementID, input.IfAdmin).Return(nil)

	output, err := service.CheckBoardPermission(context.Background(), input)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(output, expect) {
		t.Errorf("results not match, want %v, have %v", expect, output)
		return
	}
}

func TestService_CheckBoardPermissionFail(t *testing.T) {
	input := &protoBoard.CheckPermissions{
		UserID:    1,
		ElementID: 1,
		IfAdmin:   false,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	expect := &protoBoard.Nothing{Dummy: true}
	service := &service{boardStorage: mockBoardStorage}

	mockBoardStorage.EXPECT().CheckBoardPermission(input.UserID, input.ElementID, input.IfAdmin).Return(errStorage)

	output, err := service.CheckBoardPermission(context.Background(), input)
	if err == nil {
		t.Errorf("expected err: %s", err)
		return
	}
	if !reflect.DeepEqual(output, expect) {
		t.Errorf("results not match, want %v, have %v", expect, output)
		return
	}
}

func TestService_CheckCardPermission(t *testing.T) {
	input := &protoBoard.CheckPermissions{
		UserID:    1,
		ElementID: 1,
		IfAdmin:   false,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	expect := &protoBoard.BoardID{BoardID: 1}
	service := &service{boardStorage: mockBoardStorage}

	mockBoardStorage.EXPECT().CheckCardPermission(input.UserID, input.ElementID).Return(int64(1), nil)

	output, err := service.CheckCardPermission(context.Background(), input)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(output, expect) {
		t.Errorf("results not match, want %v, have %v", expect, output)
		return
	}
}

func TestService_CheckCardPermissionFail(t *testing.T) {
	input := &protoBoard.CheckPermissions{
		UserID:    1,
		ElementID: 1,
		IfAdmin:   false,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	expect := &protoBoard.BoardID{}
	service := &service{boardStorage: mockBoardStorage}

	mockBoardStorage.EXPECT().CheckCardPermission(input.UserID, input.ElementID).Return(int64(0), errStorage)

	output, err := service.CheckCardPermission(context.Background(), input)
	if err == nil {
		t.Errorf("expected err: %s", err)
		return
	}
	if !reflect.DeepEqual(output, expect) {
		t.Errorf("results not match, want %v, have %v", expect, output)
		return
	}
}

func TestService_CheckTaskPermission(t *testing.T) {
	input := &protoBoard.CheckPermissions{
		UserID:    1,
		ElementID: 1,
		IfAdmin:   false,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	expect := &protoBoard.BoardID{BoardID: 1}
	service := &service{boardStorage: mockBoardStorage}

	mockBoardStorage.EXPECT().CheckTaskPermission(input.UserID, input.ElementID).Return(int64(1),nil)

	output, err := service.CheckTaskPermission(context.Background(), input)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(output, expect) {
		t.Errorf("results not match, want %v, have %v", expect, output)
		return
	}
}

func TestService_CheckTaskPermissionFail(t *testing.T) {
	input := &protoBoard.CheckPermissions{
		UserID:    1,
		ElementID: 1,
		IfAdmin:   false,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	service := &service{boardStorage: mockBoardStorage}

	mockBoardStorage.EXPECT().CheckTaskPermission(input.UserID, input.ElementID).Return(int64(0), errStorage)

	_, err := service.CheckTaskPermission(context.Background(), input)
	if err == nil {
		t.Errorf("expected err: %s", err)
		return
	}
}

func TestService_CheckCommentPermission(t *testing.T) {
	input := &protoBoard.CheckPermissions{
		UserID:    1,
		ElementID: 1,
		IfAdmin:   false,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	expect := &protoBoard.BoardID{BoardID: 1}
	service := &service{boardStorage: mockBoardStorage}

	mockBoardStorage.
		EXPECT().
		CheckCommentPermission(input.UserID, input.ElementID, input.IfAdmin).
		Return(expect.BoardID, nil)

	output, err := service.CheckCommentPermission(context.Background(), input)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(output, expect) {
		t.Errorf("results not match, want %v, have %v", expect, output)
		return
	}
}

func TestService_CheckCommentPermissionFail(t *testing.T) {
	input := &protoBoard.CheckPermissions{
		UserID:    1,
		ElementID: 1,
		IfAdmin:   false,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	expect := &protoBoard.BoardID{BoardID: 1}
	service := &service{boardStorage: mockBoardStorage}

	mockBoardStorage.EXPECT().CheckCommentPermission(input.UserID, input.ElementID, input.IfAdmin).Return(expect.BoardID, errStorage)

	_, err := service.CheckCommentPermission(context.Background(), input)
	if err == nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
}

func TestService_GetSharedURL(t *testing.T) {
	input := &protoBoard.BoardInput{
		UserID:  1,
		BoardID: 1,
	}

	userInput := models.BoardInput{
		UserID:  input.UserID,
		BoardID: input.BoardID,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	service := &service{boardStorage: mockBoardStorage}

	mockBoardStorage.EXPECT().GetSharedURL(userInput).Return("url", nil)

	_, err := service.GetSharedURL(context.Background(), input)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
}

func TestService_GetSharedURLFail(t *testing.T) {
	input := &protoBoard.BoardInput{
		UserID:  1,
		BoardID: 1,
	}

	userInput := models.BoardInput{
		UserID:  input.UserID,
		BoardID: input.BoardID,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	service := &service{boardStorage: mockBoardStorage}

	mockBoardStorage.EXPECT().GetSharedURL(userInput).Return("", errStorage)

	_, err := service.GetSharedURL(context.Background(), input)
	if err == nil {
		t.Errorf("expected err")
		return
	}
}

func TestService_InviteUserToBoard(t *testing.T) {
	input := &protoBoard.BoardInviteInput{
		UserID:  1,
		BoardID: 1,
		UrlHash: "hash",
	}

	userInput := models.BoardInviteInput{
		UserID:  input.UserID,
		BoardID: input.BoardID,
		UrlHash: input.UrlHash,
	}

	board := models.BoardOutsideShort{
		BoardID: input.BoardID,
		Name:    "name",
		Theme:   "theme",
		Star:    false,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	service := &service{boardStorage: mockBoardStorage}

	mockBoardStorage.EXPECT().GetBoardByURL(userInput).Return(board, nil)

	_, err := service.InviteUserToBoard(context.Background(), input)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
}

func TestService_InviteUserToBoardFail(t *testing.T) {
	input := &protoBoard.BoardInviteInput{
		UserID:  1,
		BoardID: 1,
		UrlHash: "hash",
	}

	userInput := models.BoardInviteInput{
		UserID:  input.UserID,
		BoardID: input.BoardID,
		UrlHash: input.UrlHash,
	}

	board := models.BoardOutsideShort{
		BoardID: input.BoardID,
		Name:    "name",
		Theme:   "theme",
		Star:    false,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoard)

	service := &service{boardStorage: mockBoardStorage}

	mockBoardStorage.EXPECT().GetBoardByURL(userInput).Return(board, errStorage)

	_, err := service.InviteUserToBoard(context.Background(), input)
	if err == nil {
		t.Error("expected err")
		return
	}
}
