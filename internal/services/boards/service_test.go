package boards

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	mocks "github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/services/boards/mock"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

//TODO исправить ошибку в тестах!!!
/*func TestService_CreateBoard(t *testing.T) {
	boardInput := models.BoardChangeInput{
		UserID:    1,
		BoardID:   0,
		BoardName: "first",
		Theme:     "dark",
		Star:      false,
	}

	boardInternal := models.BoardInternal{
		BoardID:  1,
		AdminID:  boardInput.UserID,
		Name:     boardInput.BoardName,
		Theme:    boardInput.Theme,
		Star:     boardInput.Star,
		Cards:    []models.CardOutside{},
		UsersIDs: nil,
	}

	membersIDs := make([]int64, 0)
	membersIDs = append(membersIDs, boardInternal.AdminID)

	admin := models.UserOutsideShort{
		Email:    "epridius",
		Username: "pkaterinaa",
		FullName: "",
		Avatar:   "default/default_avatar.png",
	}

	anotherMember := models.UserOutsideShort{
		Email:    "mari",
		Username: "mari",
		FullName: "",
		Avatar:  "default/default_avatar.png",
	}

	members := make([]models.UserOutsideShort, 0)
	members = append(members, admin, anotherMember)

	expectedBoardOutside := models.BoardOutside{
		BoardID: boardInternal.BoardID,
		Admin:   members[0],
		Name:    boardInternal.Name,
		Theme:   boardInternal.Theme,
		Star:    boardInternal.Star,
		Users:   members[0:],
		Cards:   boardInternal.Cards,
	}

	ctrlUser := gomock.NewController(t)
	defer ctrlUser.Finish()
	mockUserStorage := mocks.NewMockUserStorage(ctrlUser)

	ctrlBoards := gomock.NewController(t)
	defer ctrlBoards.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoards)

	service := &service{
		userStorage: mockUserStorage,
		boardStorage: mockBoardStorage,
	}

	mockBoardStorage.EXPECT().CreateBoard(boardInput).Return(boardInternal, nil)
	mockUserStorage.EXPECT().GetUsersByIDs(membersIDs).Return(members, nil)

	board, err := service.CreateBoard(boardInput)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(board, expectedBoardOutside) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedBoardOutside, board)
		return
	}
}
*/

func TestService_ChangeBoard(t *testing.T) {
	boardInput := models.BoardChangeInput{
		UserID:    1,
		BoardID:   0,
		BoardName: "first",
		Theme:     "dark",
		Star:      false,
	}

	userIDs := []int64 {2}
	boardInternal := models.BoardInternal{
		BoardID:  1,
		AdminID:  boardInput.UserID,
		Name:     boardInput.BoardName,
		Theme:    boardInput.Theme,
		Star:     boardInput.Star,
		Cards:    []models.CardOutside{},
		UsersIDs: userIDs,
	}

	membersIDs := make([]int64, 0)
	membersIDs = append(membersIDs, boardInternal.AdminID)
	membersIDs = append(membersIDs, userIDs...)

	admin := models.UserOutsideShort{
		Email:    "epridius",
		Username: "pkaterinaa",
		FullName: "",
		Avatar:   "default/default_avatar.png",
	}

	anotherMember := models.UserOutsideShort{
		Email:    "mari",
		Username: "mari",
		FullName: "",
		Avatar:  "default/default_avatar.png",
	}

	members := make([]models.UserOutsideShort, 0)
	members = append(members, admin, anotherMember)

	expectedBoardOutside := models.BoardOutside{
		BoardID: boardInternal.BoardID,
		Admin:   members[0],
		Name:    boardInternal.Name,
		Theme:   boardInternal.Theme,
		Star:    boardInternal.Star,
		Users:   members[0:],
		Cards:   boardInternal.Cards,
	}

	ctrlUser := gomock.NewController(t)
	defer ctrlUser.Finish()
	mockUserStorage := mocks.NewMockUserStorage(ctrlUser)

	ctrlBoards := gomock.NewController(t)
	defer ctrlBoards.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoards)

	service := &service{
		userStorage: mockUserStorage,
		boardStorage: mockBoardStorage,
	}

	mockBoardStorage.EXPECT().ChangeBoard(boardInput).Return(boardInternal, nil)
	mockUserStorage.EXPECT().GetBoardMembers(membersIDs).Return(members, nil)

	board, err := service.ChangeBoard(boardInput)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(board, expectedBoardOutside) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedBoardOutside, board)
		return
	}
}

func TestService_DeleteBoard(t *testing.T) {
	boardInput := models.BoardInput{ BoardID: 1 }

	ctrlBoards := gomock.NewController(t)
	defer ctrlBoards.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoards)

	service := &service{
		boardStorage: mockBoardStorage,
	}

	mockBoardStorage.EXPECT().DeleteBoard(boardInput).Return(nil)

	err := service.DeleteBoard(boardInput)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
}

func TestService_GetBoard(t *testing.T) {
	boardInput := models.BoardInput{
		BoardID: 1,
		UserID: 1,
	}

	userIDs := []int64 {2}
	boardInternal := models.BoardInternal{
		BoardID:  1,
		AdminID:  boardInput.UserID,
		Name: 	  "first",
		Theme:     "dark",
		Star:      false,
		Cards:    []models.CardOutside{},
		UsersIDs: userIDs,
	}

	membersIDs := make([]int64, 0)
	membersIDs = append(membersIDs, boardInternal.AdminID)
	membersIDs = append(membersIDs, userIDs...)

	admin := models.UserOutsideShort{
		Email:    "epridius",
		Username: "pkaterinaa",
		FullName: "",
		Avatar:   "default/default_avatar.png",
	}

	anotherMember := models.UserOutsideShort{
		Email:    "mari",
		Username: "mari",
		FullName: "",
		Avatar:  "default/default_avatar.png",
	}

	members := make([]models.UserOutsideShort, 0)
	members = append(members, admin, anotherMember)

	expectedBoardOutside := models.BoardOutside{
		BoardID: boardInternal.BoardID,
		Admin:   members[0],
		Name:    boardInternal.Name,
		Theme:   boardInternal.Theme,
		Star:    boardInternal.Star,
		Users:   members[0:],
		Cards:   boardInternal.Cards,
	}

	ctrlUser := gomock.NewController(t)
	defer ctrlUser.Finish()
	mockUserStorage := mocks.NewMockUserStorage(ctrlUser)

	ctrlBoards := gomock.NewController(t)
	defer ctrlBoards.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoards)

	service := &service{
		userStorage: mockUserStorage,
		boardStorage: mockBoardStorage,
	}

	mockBoardStorage.EXPECT().GetBoard(boardInput).Return(boardInternal, nil)
	mockUserStorage.EXPECT().GetBoardMembers(membersIDs).Return(members, nil)

	board, err := service.GetBoard(boardInput)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(board, expectedBoardOutside) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedBoardOutside, board)
		return
	}
}

func TestService_CreateCard(t *testing.T) {
	ctrlBoards := gomock.NewController(t)
	defer ctrlBoards.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoards)

	service := &service{
		boardStorage: mockBoardStorage,
	}

	cardInput := models.CardInput{
		UserID:  1,
		CardID:  1,
		BoardID: 1,
		Name:    "card",
		Order:   1,
	}

	expectedCardOutside := models.CardOutside{
		CardID: cardInput.CardID,
		Name:   cardInput.Name,
		Order:  cardInput.Order,
		Tasks:  []models.TaskInternalShort{},
	}

	mockBoardStorage.EXPECT().CreateCard(cardInput).Return(expectedCardOutside, nil)

	card, err := service.CreateCard(cardInput)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(card, expectedCardOutside) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedCardOutside, card)
		return
	}
}

func TestService_ChangeCard(t *testing.T) {
	ctrlBoards := gomock.NewController(t)
	defer ctrlBoards.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoards)

	service := &service{
		boardStorage: mockBoardStorage,
	}

	cardInput := models.CardInput{
		UserID:  1,
		CardID:  1,
		BoardID: 1,
		Name:    "card",
		Order:   1,
	}

	expectedCardOutside := models.CardOutside{
		CardID: cardInput.CardID,
		Name:   cardInput.Name,
		Order:  cardInput.Order,
		Tasks:  []models.TaskInternalShort{},
	}

	mockBoardStorage.EXPECT().ChangeCard(cardInput).Return(expectedCardOutside, nil)

	card, err := service.ChangeCard(cardInput)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(card, expectedCardOutside) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedCardOutside, card)
		return
	}
}

func TestService_DeleteCard(t *testing.T) {
	ctrlBoards := gomock.NewController(t)
	defer ctrlBoards.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoards)

	service := &service{
		boardStorage: mockBoardStorage,
	}

	cardInput := models.CardInput{ CardID:  1 }

	mockBoardStorage.EXPECT().DeleteCard(cardInput).Return(nil)
	err := service.DeleteCard(cardInput)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
}

func TestService_GetCard(t *testing.T) {
	ctrlBoards := gomock.NewController(t)
	defer ctrlBoards.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoards)

	service := &service{
		boardStorage: mockBoardStorage,
	}

	cardInput := models.CardInput{ CardID: 1 }

	expectedCardOutside := models.CardOutside{
		CardID: cardInput.CardID,
		Name:   cardInput.Name,
		Order:  cardInput.Order,
		Tasks:  []models.TaskInternalShort{},
	}

	mockBoardStorage.EXPECT().GetCard(cardInput).Return(expectedCardOutside, nil)

	card, err := service.GetCard(cardInput)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(card, expectedCardOutside) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedCardOutside, card)
		return
	}
}

func TestService_CreateTask(t *testing.T) {
	ctrlBoards := gomock.NewController(t)
	defer ctrlBoards.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoards)

	service := &service{
		boardStorage: mockBoardStorage,
	}

	taskInput := models.TaskInput{
		UserID:  1,
		CardID:  1,
		TaskID: 1,
		Name:    "task",
		Order:   1,
	}

	expectedTaskOutside := models.TaskInternalShort{
		TaskID: taskInput.TaskID,
		Name:   taskInput.Name,
		Order:  taskInput.Order,
		Description: taskInput.Description,
	}

	mockBoardStorage.EXPECT().CreateTask(taskInput).Return(expectedTaskOutside, nil)

	task, err := service.CreateTask(taskInput)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(task, expectedTaskOutside) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedTaskOutside, task)
		return
	}
}

func TestService_ChangeTask(t *testing.T) {
	ctrlBoards := gomock.NewController(t)
	defer ctrlBoards.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoards)

	service := &service{
		boardStorage: mockBoardStorage,
	}

	taskInput := models.TaskInput{
		UserID:  1,
		CardID:  1,
		TaskID: 1,
		Name:    "task",
		Order:   1,
	}

	expectedTaskOutside := models.TaskInternalShort{
		TaskID: taskInput.TaskID,
		Name:   taskInput.Name,
		Order:  taskInput.Order,
		Description: taskInput.Description,
	}

	mockBoardStorage.EXPECT().ChangeTask(taskInput).Return(expectedTaskOutside, nil)

	task, err := service.ChangeTask(taskInput)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(task, expectedTaskOutside) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedTaskOutside, task)
		return
	}
}

func TestService_DeleteTask(t *testing.T) {
	ctrlBoards := gomock.NewController(t)
	defer ctrlBoards.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoards)

	service := &service{
		boardStorage: mockBoardStorage,
	}

	taskInput := models.TaskInput{ TaskID: 1 }

	mockBoardStorage.EXPECT().DeleteTask(taskInput)

	err := service.DeleteTask(taskInput)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
}

func TestService_GetTask(t *testing.T) {
	ctrlBoards := gomock.NewController(t)
	defer ctrlBoards.Finish()
	mockBoardStorage := mocks.NewMockBoardStorage(ctrlBoards)

	service := &service{
		boardStorage: mockBoardStorage,
	}

	taskInput := models.TaskInput{ TaskID: 1 }

	expectedTaskOutside := models.TaskInternalShort{
		TaskID: taskInput.TaskID,
		Name:   "task",
		Order:  1,
		Description: "lalala",
	}

	mockBoardStorage.EXPECT().GetTask(taskInput).Return(expectedTaskOutside, nil)

	task, err := service.GetTask(taskInput)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(task, expectedTaskOutside) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedTaskOutside, task)
		return
	}
}
