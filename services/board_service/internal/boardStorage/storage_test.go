package boardStorage

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	mocks "github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/boardStorage/mock"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestStorage_CreateBoard(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	//тут создание объектов
	storage := &storage{db: db}

	boardInput := models.BoardChangeInput{
		UserID:    1,
		BoardID:   0,
		BoardName: "test",
		Theme:     "dark",
		Star:      false,
	}

	expectBoardOutside := models.BoardInternal{
		BoardID:  1,
		AdminID:  boardInput.UserID,
		Name:     boardInput.BoardName,
		Theme:    boardInput.Theme,
		Star:     boardInput.Star,
		Cards:    []models.CardOutside{},
		UsersIDs: []int64{},
	}

	mock.
		ExpectQuery("INSERT INTO boards").
		WithArgs(boardInput.UserID, boardInput.BoardName, boardInput.Theme, boardInput.Star).
		WillReturnRows(sqlmock.NewRows([]string{"boardID"}).AddRow(1))

	board, err := storage.CreateBoard(boardInput)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(board, expectBoardOutside) {
		t.Errorf("results not match, want %v, have %v", expectBoardOutside, board)
		return
	}
}

func TestStorage_CreateBoardFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	//тут создание объектов
	storage := &storage{db: db}

	boardInput := models.BoardChangeInput{
		UserID:    1,
		BoardID:   0,
		BoardName: "test",
		Theme:     "dark",
		Star:      false,
	}

	mock.
		ExpectQuery("INSERT INTO boards").
		WithArgs(boardInput.UserID, boardInput.BoardName, boardInput.Theme, boardInput.Star).
		WillReturnError(sql.ErrNoRows)

	_, err = storage.CreateBoard(boardInput)
	if err == nil {
		t.Error("expected error")
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestStorage_ChangeBoard(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	boardInput := models.BoardChangeInput{
		UserID:    1,
		BoardID:   1,
		BoardName: "test changed",
		Theme:     "dark",
		Star:      false,
	}

	expectedCards := make([]models.CardOutside, 0)
	card1 := models.CardOutside{
		CardID: 1,
		Name:   "todo",
		Order:  1,
	}
	expectedCards = append(expectedCards, card1)

	expectedTasks := make([]models.TaskOutside, 0)
	task1 := models.TaskOutside{
		TaskID:      1,
		Name:        "task 1",
		Description: "first task ever",
		Order:       1,
	}

	task2 := models.TaskOutside{
		TaskID:      2,
		Name:        "task 2",
		Description: "second task",
		Order:       2,
	}

	expectedTasks = append(expectedTasks, task1, task2)

	ctrlCards := gomock.NewController(t)
	defer ctrlCards.Finish()

	mockCards := mocks.NewMockCardsStorage(ctrlCards)
	mockCards.EXPECT().GetCardsByBoard(models.BoardInput{BoardID: boardInput.BoardID}).Times(1).Return(expectedCards, nil)

	ctrlTasks := gomock.NewController(t)
	defer ctrlTasks.Finish()

	mockTasks := mocks.NewMockTasksStorage(ctrlTasks)
	mockTasks.EXPECT().GetTasksByCard(models.CardInput{CardID: expectedCards[0].CardID}).Times(1).Return(expectedTasks, nil)

	storage := NewStorage(db, mockCards, mockTasks)

	expectBoardOutside := models.BoardInternal{
		BoardID: boardInput.BoardID,
		AdminID: boardInput.UserID,
		Name:    boardInput.BoardName,
		Theme:   boardInput.Theme,
		Star:    boardInput.Star,
	}

	mock.
		ExpectQuery("UPDATE boards SET").
		WithArgs(boardInput.BoardName, boardInput.Theme, boardInput.Star, boardInput.BoardID).
		WillReturnRows(sqlmock.NewRows([]string{"adminID"}).AddRow(1))

	mock.
		ExpectQuery("SELECT userID from board_members").
		WithArgs(boardInput.BoardID).
		WillReturnRows(sqlmock.NewRows([]string{"userID"}).AddRow(2).AddRow(3))

	expectedCards[0].Tasks = append(expectedCards[0].Tasks, expectedTasks...)
	expectBoardOutside.Cards = append(expectBoardOutside.Cards, expectedCards...)
	expectBoardOutside.UsersIDs = []int64{2, 3}

	board, err := storage.ChangeBoard(boardInput)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(board, expectBoardOutside) {
		t.Errorf("results not match, want \n%v, have \n%v", expectBoardOutside, board)
		return
	}
}

func TestStorage_ChangeBoardGetCardsFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	boardInput := models.BoardChangeInput{
		UserID:    1,
		BoardID:   1,
		BoardName: "test changed",
		Theme:     "dark",
		Star:      false,
	}

	expectedCards := make([]models.CardOutside, 0)

	ctrlCards := gomock.NewController(t)
	defer ctrlCards.Finish()

	mockCards := mocks.NewMockCardsStorage(ctrlCards)
	mockCards.EXPECT().GetCardsByBoard(models.BoardInput{BoardID: boardInput.BoardID}).Times(1).Return(expectedCards, errors.New("error getting cards"))

	storage := storage{db: db, cardsStorage: mockCards}

	mock.
		ExpectQuery("UPDATE boards SET").
		WithArgs(boardInput.BoardName, boardInput.Theme, boardInput.Star, boardInput.BoardID).
		WillReturnRows(sqlmock.NewRows([]string{"adminID"}).AddRow(1))

	mock.
		ExpectQuery("SELECT userID from board_members").
		WithArgs(boardInput.BoardID).
		WillReturnRows(sqlmock.NewRows([]string{"userID"}).AddRow(2).AddRow(3))

	_, err = storage.ChangeBoard(boardInput)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected err")
		return
	}
}

func TestStorage_ChangeBoardGetTasksFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	boardInput := models.BoardChangeInput{
		UserID:    1,
		BoardID:   1,
		BoardName: "test changed",
		Theme:     "dark",
		Star:      false,
	}

	expectedCards := make([]models.CardOutside, 0)
	card1 := models.CardOutside{
		CardID: 1,
		Name:   "todo",
		Order:  1,
	}
	expectedCards = append(expectedCards, card1)

	expectedTasks := make([]models.TaskOutside, 0)

	ctrlCards := gomock.NewController(t)
	defer ctrlCards.Finish()

	mockCards := mocks.NewMockCardsStorage(ctrlCards)
	mockCards.EXPECT().GetCardsByBoard(models.BoardInput{BoardID: boardInput.BoardID}).Times(1).Return(expectedCards, nil)

	ctrlTasks := gomock.NewController(t)
	defer ctrlTasks.Finish()

	mockTasks := mocks.NewMockTasksStorage(ctrlTasks)
	mockTasks.EXPECT().GetTasksByCard(models.CardInput{CardID: expectedCards[0].CardID}).Times(1).Return(expectedTasks, errors.New("error getting tasks"))

	storage := NewStorage(db, mockCards, mockTasks)

	mock.
		ExpectQuery("UPDATE boards SET").
		WithArgs(boardInput.BoardName, boardInput.Theme, boardInput.Star, boardInput.BoardID).
		WillReturnRows(sqlmock.NewRows([]string{"adminID"}).AddRow(1))

	mock.
		ExpectQuery("SELECT userID from board_members").
		WithArgs(boardInput.BoardID).
		WillReturnRows(sqlmock.NewRows([]string{"userID"}).AddRow(2).AddRow(3))

	_, err = storage.ChangeBoard(boardInput)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected err")
		return
	}
}

func TestStorage_ChangeBoardUpdateQueryFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	boardInput := models.BoardChangeInput{
		UserID:    1,
		BoardID:   1,
		BoardName: "test changed",
		Theme:     "dark",
		Star:      false,
	}

	storage := storage{db: db}
	mock.
		ExpectQuery("UPDATE boards SET").
		WithArgs(boardInput.BoardName, boardInput.Theme, boardInput.Star, boardInput.BoardID).
		WillReturnError(sql.ErrNoRows)

	_, err = storage.ChangeBoard(boardInput)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected err")
		return
	}
}

func TestStorage_ChangeBoardGetMembersFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	boardInput := models.BoardChangeInput{
		UserID:    1,
		BoardID:   1,
		BoardName: "test changed",
		Theme:     "dark",
		Star:      false,
	}
	storage := storage{db: db}

	mock.
		ExpectQuery("UPDATE boards SET").
		WithArgs(boardInput.BoardName, boardInput.Theme, boardInput.Star, boardInput.BoardID).
		WillReturnRows(sqlmock.NewRows([]string{"adminID"}).AddRow(1))

	mock.
		ExpectQuery("SELECT userID from board_members").
		WithArgs(boardInput.BoardID).
		WillReturnError(errors.New("internal db error"))

	_, err = storage.ChangeBoard(boardInput)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected err")
		return
	}
}

func TestStorage_ChangeBoardGetMembersScanFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	boardInput := models.BoardChangeInput{
		UserID:    1,
		BoardID:   1,
		BoardName: "test changed",
		Theme:     "dark",
		Star:      false,
	}
	storage := storage{db: db}

	mock.
		ExpectQuery("UPDATE boards SET").
		WithArgs(boardInput.BoardName, boardInput.Theme, boardInput.Star, boardInput.BoardID).
		WillReturnRows(sqlmock.NewRows([]string{"adminID"}).AddRow(1))

	mock.
		ExpectQuery("SELECT userID from board_members").
		WithArgs(boardInput.BoardID).
		WillReturnRows(sqlmock.NewRows([]string{"userID", "username"}).AddRow(1, "pringlz"))

	_, err = storage.ChangeBoard(boardInput)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected err")
		return
	}
}

func TestStorage_DeleteBoard(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	boardInput := models.BoardInput{ BoardID: 1 }

	mock.
		ExpectExec("DELETE FROM boards WHERE boardID").
		WithArgs(boardInput.BoardID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = storage.DeleteBoard(boardInput)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestStorage_DeleteBoardFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	boardInput := models.BoardInput{ BoardID: 1 }

	mock.
		ExpectExec("DELETE FROM boards WHERE boardID").
		WithArgs(boardInput.BoardID).
		WillReturnError(errors.New("error while deleting"))

	err = storage.DeleteBoard(boardInput)
	if err == nil {
		t.Errorf("expected err")
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestStorage_GetBoard(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	boardInput := models.BoardInput{
		BoardID:   1,
	}

	expectedCards := make([]models.CardOutside, 0)
	card1 := models.CardOutside{
		CardID: 1,
		Name:   "todo",
		Order:  1,
	}

	expectedCards = append(expectedCards, card1)

	ctrlCards := gomock.NewController(t)
	defer ctrlCards.Finish()

	mockCards := mocks.NewMockCardsStorage(ctrlCards)

	fmt.Println("expected cards ", expectedCards)
	mockCards.EXPECT().GetCardsByBoard(models.BoardInput{BoardID: boardInput.BoardID}).Times(1).Return(expectedCards, nil)

	expectedTasks := make([]models.TaskOutside, 0)
	task1 := models.TaskOutside{
		TaskID:      1,
		Name:        "task 1",
		Description: "first task ever",
		Order:       1,
	}

	task2 := models.TaskOutside{
		TaskID:      2,
		Name:        "task 2",
		Description: "second task",
		Order:       2,
	}

	expectedTasks = append(expectedTasks, task1, task2)

	ctrlTasks := gomock.NewController(t)
	defer ctrlTasks.Finish()

	mockTasks := mocks.NewMockTasksStorage(ctrlTasks)
	mockTasks.EXPECT().GetTasksByCard(models.CardInput{CardID: expectedCards[0].CardID}).Times(1).Return(expectedTasks, nil)

	storage := &storage{
		db:           db,
		cardsStorage: mockCards,
		tasksStorage: mockTasks,
	}

	expectBoardOutside := models.BoardInternal{
		BoardID: boardInput.BoardID,
		AdminID: 1,
		Name:    "test board",
		Theme:   "dark",
		Star:    false,
	}

	cards := make([]models.CardOutside, 0)
	cards = append(cards, expectedCards...)
	cards[0].Tasks = append(cards[0].Tasks, expectedTasks...)

	expectBoardOutside.Cards = append(expectBoardOutside.Cards, cards...)
	expectBoardOutside.UsersIDs = []int64{2, 3}

	rows := sqlmock.NewRows([]string{"adminID", "boardName", "theme", "star"}).AddRow(1, "test board", "dark", false)

	mock.
		ExpectQuery("SELECT adminID, boardName, theme, star FROM boards WHERE boardID").
		WithArgs(boardInput.BoardID).
		WillReturnRows(rows)

	mock.
		ExpectQuery("SELECT userID from board_members").
		WithArgs(boardInput.BoardID).
		WillReturnRows(sqlmock.NewRows([]string{"userID"}).AddRow(2).AddRow(3))

	board, err := storage.GetBoard(boardInput)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(board, expectBoardOutside) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectBoardOutside, board)
		return
	}
}

func TestStorage_GetBoardQueryFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	boardInput := models.BoardInput{
		BoardID:   1,
	}

	storage := &storage{
		db:           db,
	}

	mock.
		ExpectQuery("SELECT adminID, boardName, theme, star FROM boards WHERE boardID").
		WithArgs(boardInput.BoardID).
		WillReturnError(sql.ErrNoRows)

	_, err = storage.GetBoard(boardInput)
	if err == nil {
		t.Errorf("expected err")
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestStorage_GetBoardMembersFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	boardInput := models.BoardInput{
		BoardID:   1,
	}

	storage := &storage{
		db:           db,
	}

	rows := sqlmock.NewRows([]string{"adminID", "boardName", "theme", "star"}).AddRow(1, "test board", "dark", false)

	mock.
		ExpectQuery("SELECT adminID, boardName, theme, star FROM boards WHERE boardID").
		WithArgs(boardInput.BoardID).
		WillReturnRows(rows)

	mock.
		ExpectQuery("SELECT userID from board_members").
		WithArgs(boardInput.BoardID).
		WillReturnError(sql.ErrNoRows)

	_, err = storage.GetBoard(boardInput)
	if err == nil {
		t.Errorf("expected err")
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestStorage_GetBoardCardsFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	boardInput := models.BoardInput{
		BoardID:   1,
	}

	expectedCards := make([]models.CardOutside, 0)

	ctrlCards := gomock.NewController(t)
	defer ctrlCards.Finish()

	mockCards := mocks.NewMockCardsStorage(ctrlCards)

	mockCards.EXPECT().GetCardsByBoard(models.BoardInput{BoardID: boardInput.BoardID}).Times(1).Return(expectedCards, errors.New(""))

	storage := &storage{
		db:           db,
		cardsStorage: mockCards,
	}

	rows := sqlmock.NewRows([]string{"adminID", "boardName", "theme", "star"}).AddRow(1, "test board", "dark", false)

	mock.
		ExpectQuery("SELECT adminID, boardName, theme, star FROM boards WHERE boardID").
		WithArgs(boardInput.BoardID).
		WillReturnRows(rows)

	mock.
		ExpectQuery("SELECT userID from board_members").
		WithArgs(boardInput.BoardID).
		WillReturnRows(sqlmock.NewRows([]string{"userID"}).AddRow(2).AddRow(3))

	_, err = storage.GetBoard(boardInput)
	if err == nil {
		t.Errorf("expected err")
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}


func TestStorage_GetBoardTasksFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	boardInput := models.BoardInput{
		BoardID:   1,
	}

	expectedCards := make([]models.CardOutside, 0)
	card1 := models.CardOutside{
		CardID: 1,
		Name:   "todo",
		Order:  1,
	}

	expectedCards = append(expectedCards, card1)

	ctrlCards := gomock.NewController(t)
	defer ctrlCards.Finish()

	mockCards := mocks.NewMockCardsStorage(ctrlCards)

	fmt.Println("expected cards ", expectedCards)
	mockCards.EXPECT().GetCardsByBoard(models.BoardInput{BoardID: boardInput.BoardID}).Times(1).Return(expectedCards, nil)

	expectedTasks := make([]models.TaskOutside, 0)

	ctrlTasks := gomock.NewController(t)
	defer ctrlTasks.Finish()

	mockTasks := mocks.NewMockTasksStorage(ctrlTasks)
	mockTasks.EXPECT().GetTasksByCard(models.CardInput{CardID: expectedCards[0].CardID}).Times(1).Return(expectedTasks, errors.New(""))

	storage := &storage{
		db:           db,
		cardsStorage: mockCards,
		tasksStorage: mockTasks,
	}

	rows := sqlmock.NewRows([]string{"adminID", "boardName", "theme", "star"}).AddRow(1, "test board", "dark", false)

	mock.
		ExpectQuery("SELECT adminID, boardName, theme, star FROM boards WHERE boardID").
		WithArgs(boardInput.BoardID).
		WillReturnRows(rows)

	mock.
		ExpectQuery("SELECT userID from board_members").
		WithArgs(boardInput.BoardID).
		WillReturnRows(sqlmock.NewRows([]string{"userID"}).AddRow(2).AddRow(3))

	_, err = storage.GetBoard(boardInput)
	if err == nil {
		t.Errorf("expected err")
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestStorage_GetBoardsList(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{
		db: db,
	}

	userInput := models.UserInput{ID: 1}

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

	rows := sqlmock.NewRows([]string{"boardID", "boardName", "theme", "star"}).
					AddRow(board1.BoardID, board1.Name, board1.Theme, board1.Star).
					AddRow(board2.BoardID, board2.Name, board2.Theme, board2.Star)

	mock.
		ExpectQuery("SELECT DISTINCT B.boardID, B.boardName, B.theme, B.star FROM boards B").
		WithArgs(userInput.ID).
		WillReturnRows(rows)

	boards, err := storage.GetBoardsList(userInput)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(boards, expectedBoards) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedBoards, boards)
		return
	}
}

func TestStorage_GetBoardsListQueryFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{
		db: db,
	}

	userInput := models.UserInput{ID: 1}

	mock.
		ExpectQuery("SELECT DISTINCT B.boardID, B.boardName, B.theme, B.star FROM boards B").
		WithArgs(userInput.ID).
		WillReturnError(sql.ErrNoRows)

	_, err = storage.GetBoardsList(userInput)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected err")
		return
	}
}

func TestStorage_GetBoardsListScanFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{
		db: db,
	}

	userInput := models.UserInput{ID: 1}

	board1 := models.BoardOutsideShort{
		BoardID: 1,
		Name:    "board_1",
		Theme:   "dark",
		Star:    false,
	}

	expectedBoards := make([]models.BoardOutsideShort, 0)
	expectedBoards = append(expectedBoards, board1)

	rows := sqlmock.NewRows([]string{"boardID", "boardName", "theme"}).
		AddRow(board1.BoardID, board1.Name, board1.Theme)

	mock.
		ExpectQuery("SELECT DISTINCT B.boardID, B.boardName, B.theme, B.star FROM boards B").
		WithArgs(userInput.ID).
		WillReturnRows(rows)

	_, err = storage.GetBoardsList(userInput)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected err")
		return
	}
}

func TestStorage_CheckBoardPermissionAdmin(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{
		db: db,
	}

	ifAdmin := true
	boardID := int64(1)
	userID := int64(1)

	mock.
		ExpectQuery("SELECT boardID FROM boards").
		WithArgs(boardID, userID).
		WillReturnRows(sqlmock.NewRows([]string{"adminID"}).AddRow(1))

	err = storage.CheckBoardPermission(userID, boardID, ifAdmin)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
}

func TestStorage_CheckBoardPermissionAdminQueryFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{
		db: db,
	}

	ifAdmin := true
	boardID := int64(1)
	userID := int64(1)

	mock.
		ExpectQuery("SELECT boardID FROM boards").
		WithArgs(boardID, userID).
		WillReturnError(errors.New(""))

	err = storage.CheckBoardPermission(userID, boardID, ifAdmin)
	if err == nil {
		t.Errorf("expected err")
		return
	}
}

func TestStorage_CheckBoardPermissionAdminFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{
		db: db,
	}

	ifAdmin := true
	boardID := int64(1)
	userID := int64(1)

	mock.
		ExpectQuery("SELECT boardID FROM boards").
		WithArgs(boardID, userID).
		WillReturnError(sql.ErrNoRows)

	err = storage.CheckBoardPermission(userID, boardID, ifAdmin)
	if err == nil {
		t.Errorf("expected err")
		return
	}
}

func TestStorage_CheckBoardPermissionUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{
		db: db,
	}

	ifAdmin := false
	boardID := int64(1)
	userID := int64(1)

	mock.
		ExpectQuery("SELECT boardID FROM board_members").
		WithArgs(boardID, userID).
		WillReturnRows(sqlmock.NewRows([]string{"userID"}).AddRow(1))

	err = storage.CheckBoardPermission(userID, boardID, ifAdmin)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
}

func TestStorage_CheckBoardPermissionUserFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{
		db: db,
	}

	ifAdmin := false
	boardID := int64(1)
	userID := int64(1)

	mock.
		ExpectQuery("SELECT boardID FROM board_members").
		WithArgs(boardID, userID).
		WillReturnError(sql.ErrNoRows)

	err = storage.CheckBoardPermission(userID, boardID, ifAdmin)
	if err == nil {
		t.Errorf("expected err")
		return
	}
}

func TestStorage_CheckBoardPermissionUserQueryFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{
		db: db,
	}

	ifAdmin := false
	boardID := int64(1)
	userID := int64(1)

	mock.
		ExpectQuery("SELECT boardID FROM board_members").
		WithArgs(boardID, userID).
		WillReturnError(errors.New(""))

	err = storage.CheckBoardPermission(userID, boardID, ifAdmin)
	if err == nil {
		t.Errorf("expected err")
		return
	}
}

func TestStorage_CreateCard(t *testing.T) {
	cardInput := models.CardInput{
		UserID:  1,
		CardID:  1,
		BoardID: 1,
		Name:    "todo",
		Order:   1,
	}

	cardOutside := models.CardOutside{
		CardID: 1,
		Name:   "todo",
		Order:  1,
	}

	ctrlCards := gomock.NewController(t)
	defer ctrlCards.Finish()

	mockCards := mocks.NewMockCardsStorage(ctrlCards)
	mockCards.EXPECT().CreateCard(cardInput).Times(1).Return(cardOutside, nil)

	storage := &storage{
		cardsStorage: mockCards,
	}

	card, err := storage.CreateCard(cardInput)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(card, cardOutside) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", cardOutside, card)
		return
	}
}

func TestStorage_CreateCardFail(t *testing.T) {
	cardInput := models.CardInput{
		UserID:  1,
		CardID:  1,
		BoardID: 1,
		Name:    "todo",
		Order:   1,
	}

	ctrlCards := gomock.NewController(t)
	defer ctrlCards.Finish()

	mockCards := mocks.NewMockCardsStorage(ctrlCards)
	mockCards.EXPECT().CreateCard(cardInput).Times(1).Return(models.CardOutside{}, errors.New("error creating card"))

	storage := &storage{
		cardsStorage: mockCards,
	}

	_, err := storage.CreateCard(cardInput)
	if err == nil {
		t.Errorf("expected err")
		return
	}
}

func TestStorage_ChangeCard(t *testing.T) {
	cardInput := models.CardInput{
		UserID:  1,
		CardID:  0,
		BoardID: 1,
		Name:    "todo changed",
		Order:   1,
	}

	cardOutside := models.CardOutside{
		CardID: 1,
		Name:   "todo changed",
		Order:  1,
		Tasks:  []models.TaskOutside{},
	}

	ctrlCards := gomock.NewController(t)
	defer ctrlCards.Finish()

	mockCards := mocks.NewMockCardsStorage(ctrlCards)
	mockCards.EXPECT().ChangeCard(cardInput).Times(1).Return(cardOutside, nil)

	ctrlTasks := gomock.NewController(t)
	defer ctrlTasks.Finish()

	mockTasks := mocks.NewMockTasksStorage(ctrlTasks)
	mockTasks.EXPECT().GetTasksByCard(cardInput).Times(1).Return([]models.TaskOutside{}, nil)

	storage := &storage{
		cardsStorage: mockCards,
		tasksStorage: mockTasks,
	}

	card, err := storage.ChangeCard(cardInput)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(card, cardOutside) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", cardOutside, card)
		return
	}
}

func TestStorage_ChangeCardFail(t *testing.T) {
	cardInput := models.CardInput{
		UserID:  1,
		CardID:  0,
		BoardID: 1,
		Name:    "todo changed",
		Order:   1,
	}

	ctrlCards := gomock.NewController(t)
	defer ctrlCards.Finish()

	mockCards := mocks.NewMockCardsStorage(ctrlCards)
	mockCards.EXPECT().ChangeCard(cardInput).Times(1).Return(models.CardOutside{}, errors.New(""))

	storage := &storage{
		cardsStorage: mockCards,
	}

	_, err := storage.ChangeCard(cardInput)
	if err == nil {
		t.Errorf("expected err")
		return
	}
}

func TestStorage_ChangeCardFailGetTasks(t *testing.T) {
	cardInput := models.CardInput{
		UserID:  1,
		CardID:  0,
		BoardID: 1,
		Name:    "todo changed",
		Order:   1,
	}

	cardOutside := models.CardOutside{
		CardID: 1,
		Name:   "todo changed",
		Order:  1,
		Tasks:  []models.TaskOutside{},
	}

	ctrlCards := gomock.NewController(t)
	defer ctrlCards.Finish()

	mockCards := mocks.NewMockCardsStorage(ctrlCards)
	mockCards.EXPECT().ChangeCard(cardInput).Times(1).Return(cardOutside, nil)

	ctrlTasks := gomock.NewController(t)
	defer ctrlTasks.Finish()

	mockTasks := mocks.NewMockTasksStorage(ctrlTasks)
	mockTasks.EXPECT().GetTasksByCard(cardInput).Times(1).Return([]models.TaskOutside{}, errors.New(""))

	storage := &storage{
		cardsStorage: mockCards,
		tasksStorage: mockTasks,
	}

	_, err := storage.ChangeCard(cardInput)
	if err == nil {
		t.Errorf("expected err")
		return
	}
}

func TestStorage_ChangeCardOrder(t *testing.T) {
	card := models.CardOrder{
		CardID: 1,
		Order:  2,
	}
	input := models.CardsOrderInput{}
	input.Cards = append(input.Cards, card)

	ctrlCards := gomock.NewController(t)
	defer ctrlCards.Finish()

	mockCards := mocks.NewMockCardsStorage(ctrlCards)
	mockCards.EXPECT().ChangeCardOrder(input).Times(1).Return(nil)

	storage := &storage{
		cardsStorage: mockCards,
	}

	err := storage.ChangeCardOrder(input)
	if err != nil {
		t.Errorf("unexpected err %s", err)
		return
	}
}

func TestStorage_ChangeCardOrderFail(t *testing.T) {
	card := models.CardOrder{
		CardID: 1,
		Order:  2,
	}
	input := models.CardsOrderInput{}
	input.Cards = append(input.Cards, card)

	ctrlCards := gomock.NewController(t)
	defer ctrlCards.Finish()

	mockCards := mocks.NewMockCardsStorage(ctrlCards)
	mockCards.EXPECT().ChangeCardOrder(input).Times(1).Return(errors.New(""))

	storage := &storage{
		cardsStorage: mockCards,
	}

	err := storage.ChangeCardOrder(input)
	if err == nil {
		t.Errorf("expected err")
		return
	}
}

func TestStorage_ChangeTaskOrder(t *testing.T) {
	tasks := models.TaskOrder{
		TaskID: 1,
		Order:  2,
	}
	card := models.TasksOrder{CardID: 1}
	card.Tasks = append(card.Tasks, tasks)
	input := models.TasksOrderInput{ UserID: 1 }
	input.Tasks = append(input.Tasks, card)

	ctrlTasks := gomock.NewController(t)
	defer ctrlTasks.Finish()

	mockTasks := mocks.NewMockTasksStorage(ctrlTasks)
	mockTasks.EXPECT().ChangeTaskOrder(input).Times(1).Return(nil)

	storage := &storage{
		tasksStorage: mockTasks,
	}

	err := storage.ChangeTaskOrder(input)
	if err != nil {
		t.Errorf("unexpected err %s", err)
		return
	}
}

func TestStorage_ChangeTaskOrderFail(t *testing.T) {
	tasks := models.TaskOrder{
		TaskID: 1,
		Order:  2,
	}
	card := models.TasksOrder{CardID: 1}
	card.Tasks = append(card.Tasks, tasks)
	input := models.TasksOrderInput{ UserID: 1 }
	input.Tasks = append(input.Tasks, card)

	ctrlTasks := gomock.NewController(t)
	defer ctrlTasks.Finish()

	mockTasks := mocks.NewMockTasksStorage(ctrlTasks)
	mockTasks.EXPECT().ChangeTaskOrder(input).Times(1).Return(errors.New(""))

	storage := &storage{
		tasksStorage: mockTasks,
	}

	err := storage.ChangeTaskOrder(input)
	if err == nil {
		t.Errorf("expected err ")
		return
	}
}

func TestStorage_DeleteCard(t *testing.T) {
	cardInput := models.CardInput{
		CardID:  1,
	}

	ctrlCards := gomock.NewController(t)
	defer ctrlCards.Finish()

	mockCards := mocks.NewMockCardsStorage(ctrlCards)
	mockCards.EXPECT().DeleteCard(cardInput).Times(1).Return(nil)

	storage := &storage{
		cardsStorage: mockCards,
	}

	err := storage.DeleteCard(cardInput)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
}

func TestStorage_DeleteCardFail(t *testing.T) {
	cardInput := models.CardInput{
		CardID:  1,
	}

	ctrlCards := gomock.NewController(t)
	defer ctrlCards.Finish()

	mockCards := mocks.NewMockCardsStorage(ctrlCards)
	mockCards.EXPECT().DeleteCard(cardInput).Times(1).Return(errors.New(""))

	storage := &storage{
		cardsStorage: mockCards,
	}

	err := storage.DeleteCard(cardInput)
	if err == nil {
		t.Errorf("expected err")
		return
	}
}

func TestStorage_GetCard(t *testing.T) {
	cardInput := models.CardInput{ CardID: 1 }

	expectedCard := models.CardOutside{
		CardID: 1,
		Name:   "todo",
		Order:  1,
	}

	expectedTasks := make([]models.TaskOutside, 0)
	task1 := models.TaskOutside{
		TaskID:      1,
		Name:        "task 1",
		Description: "first task ever",
		Order:       1,
	}

	task2 := models.TaskOutside{
		TaskID:      2,
		Name:        "task 2",
		Description: "second task",
		Order:       2,
	}

	ctrlCards := gomock.NewController(t)
	defer ctrlCards.Finish()

	mockCards := mocks.NewMockCardsStorage(ctrlCards)
	mockCards.EXPECT().GetCardByID(cardInput).Times(1).Return(expectedCard, nil)

	expectedTasks = append(expectedTasks, task1, task2)
	expectedCard.Tasks = append(expectedCard.Tasks, expectedTasks...)

	ctrlTasks := gomock.NewController(t)
	defer ctrlTasks.Finish()

	mockTasks := mocks.NewMockTasksStorage(ctrlTasks)
	mockTasks.EXPECT().GetTasksByCard(cardInput).Times(1).Return(expectedTasks, nil)

	storage := &storage{
		cardsStorage: mockCards,
		tasksStorage: mockTasks,
	}

	card, err := storage.GetCard(cardInput)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(card, expectedCard) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedCard, card)
		return
	}
}

func TestStorage_GetCardByIDFail(t *testing.T) {
	cardInput := models.CardInput{ CardID: 1 }

	ctrlCards := gomock.NewController(t)
	defer ctrlCards.Finish()

	mockCards := mocks.NewMockCardsStorage(ctrlCards)
	mockCards.EXPECT().GetCardByID(cardInput).Times(1).Return(models.CardOutside{}, errors.New(""))

	storage := &storage{
		cardsStorage: mockCards,
	}

	_, err := storage.GetCard(cardInput)
	if err == nil {
		t.Errorf("expected err")
		return
	}
}

func TestStorage_GetCardTasksFail(t *testing.T) {
	cardInput := models.CardInput{ CardID: 1 }

	expectedCard := models.CardOutside{
		CardID: 1,
		Name:   "todo",
		Order:  1,
	}

	expectedTasks := make([]models.TaskOutside, 0)

	ctrlCards := gomock.NewController(t)
	defer ctrlCards.Finish()

	mockCards := mocks.NewMockCardsStorage(ctrlCards)
	mockCards.EXPECT().GetCardByID(cardInput).Times(1).Return(expectedCard, nil)

	ctrlTasks := gomock.NewController(t)
	defer ctrlTasks.Finish()

	mockTasks := mocks.NewMockTasksStorage(ctrlTasks)
	mockTasks.EXPECT().GetTasksByCard(cardInput).Times(1).Return(expectedTasks, errors.New(""))

	storage := &storage{
		cardsStorage: mockCards,
		tasksStorage: mockTasks,
	}

	_, err := storage.GetCard(cardInput)
	if err == nil {
		t.Errorf("expected err")
		return
	}
}

func TestStorage_CheckCardPermission(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	cardID := int64(1)
	userID := int64(1)

	mock.
		ExpectQuery("SELECT B.boardID").
		WithArgs(userID, cardID).
		WillReturnRows(sqlmock.NewRows([]string{"boardID"}).AddRow(1))

	err = storage.CheckCardPermission(userID, cardID)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestStorage_CheckCardPermissionQueryFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	cardID := int64(1)
	userID := int64(1)

	mock.
		ExpectQuery("SELECT B.boardID").
		WithArgs(userID, cardID).
		WillReturnError(errors.New(""))

	err = storage.CheckCardPermission(userID, cardID)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected err")
		return
	}
}

func TestStorage_CheckCardPermissionFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	cardID := int64(1)
	userID := int64(1)

	mock.
		ExpectQuery("SELECT B.boardID").
		WithArgs(userID, cardID).
		WillReturnError(sql.ErrNoRows)

	err = storage.CheckCardPermission(userID, cardID)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected err")
		return
	}
}

func TestStorage_CreateTask(t *testing.T) {
	taskInput := models.TaskInput{
		UserID:  1,
		CardID:  1,
		Name:    "todo",
		Description: "description",
		Order:   1,
	}

	taskOutside := models.TaskOutside{
		TaskID:      1,
		Name:        taskInput.Name,
		Description: taskInput.Description,
		Order:       taskInput.Order,
	}

	ctrlTasks := gomock.NewController(t)
	defer ctrlTasks.Finish()

	mockTasks := mocks.NewMockTasksStorage(ctrlTasks)
	mockTasks.EXPECT().CreateTask(taskInput).Times(1).Return(taskOutside, nil)

	storage := &storage{
		tasksStorage: mockTasks,
	}

	task, err := storage.CreateTask(taskInput)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(task, taskOutside) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", taskOutside, task)
		return
	}
}

func TestStorage_CreateTaskFail(t *testing.T) {
	taskInput := models.TaskInput{
		UserID:  1,
		CardID:  1,
		Name:    "todo",
		Description: "description",
		Order:   1,
	}

	ctrlTasks := gomock.NewController(t)
	defer ctrlTasks.Finish()

	mockTasks := mocks.NewMockTasksStorage(ctrlTasks)
	mockTasks.EXPECT().CreateTask(taskInput).Times(1).Return(models.TaskOutside{}, errors.New(""))

	storage := &storage{
		tasksStorage: mockTasks,
	}

	_, err := storage.CreateTask(taskInput)
	if err == nil {
		t.Errorf("expected err")
		return
	}
}

func TestStorage_ChangeTask(t *testing.T) {
	taskInput := models.TaskInput{
		UserID:  1,
		CardID:  1,
		TaskID: 1,
		Name:    "todo changed",
		Description: "description changed",
		Order:   1,
	}

	taskOutside := models.TaskOutside{
		TaskID:      1,
		Name:        taskInput.Name,
		Description: taskInput.Description,
		Order:       taskInput.Order,
	}

	ctrlTasks := gomock.NewController(t)
	defer ctrlTasks.Finish()

	mockTasks := mocks.NewMockTasksStorage(ctrlTasks)
	mockTasks.EXPECT().ChangeTask(taskInput).Times(1).Return(taskOutside, nil)

	storage := &storage{
		tasksStorage: mockTasks,
	}

	task, err := storage.ChangeTask(taskInput)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(task, taskOutside) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", taskOutside, task)
		return
	}
}

func TestStorage_ChangeTaskFail(t *testing.T) {
	taskInput := models.TaskInput{
		UserID:  1,
		CardID:  1,
		TaskID: 1,
		Name:    "todo changed",
		Description: "description changed",
		Order:   1,
	}

	ctrlTasks := gomock.NewController(t)
	defer ctrlTasks.Finish()

	mockTasks := mocks.NewMockTasksStorage(ctrlTasks)
	mockTasks.EXPECT().ChangeTask(taskInput).Times(1).Return(models.TaskOutside{}, errors.New(""))

	storage := &storage{
		tasksStorage: mockTasks,
	}

	_, err := storage.ChangeTask(taskInput)
	if err == nil {
		t.Errorf("expected err")
		return
	}
}

func TestStorage_DeleteTask(t *testing.T) {
	taskInput := models.TaskInput{ TaskID: 1 }

	ctrlTasks := gomock.NewController(t)
	defer ctrlTasks.Finish()

	mockTasks := mocks.NewMockTasksStorage(ctrlTasks)
	mockTasks.EXPECT().DeleteTask(taskInput).Times(1).Return(nil)

	storage := &storage{
		tasksStorage: mockTasks,
	}

	err := storage.DeleteTask(taskInput)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
}

func TestStorage_GetTask(t *testing.T) {
	taskInput := models.TaskInput{ TaskID: 1 }

	taskOutside := models.TaskOutside{
		TaskID:      1,
		Name:        "task",
		Description: "description",
		Order:       1,
	}

	ctrlTasks := gomock.NewController(t)
	defer ctrlTasks.Finish()

	mockTasks := mocks.NewMockTasksStorage(ctrlTasks)
	mockTasks.EXPECT().GetTaskByID(taskInput).Times(1).Return(taskOutside, nil)

	storage := &storage{
		tasksStorage: mockTasks,
	}

	task, err := storage.GetTask(taskInput)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(task, taskOutside) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", taskOutside, task)
		return
	}
}

func TestStorage_CheckTaskPermission(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	taskID := int64(1)
	userID := int64(1)

	mock.
		ExpectQuery("SELECT B.boardID").
		WithArgs(userID, taskID).
		WillReturnRows(sqlmock.NewRows([]string{"boardID"}).AddRow(1))

	err = storage.CheckTaskPermission(userID, taskID)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestStorage_CheckTaskPermissionQueryFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	taskID := int64(1)
	userID := int64(1)

	mock.
		ExpectQuery("SELECT B.boardID").
		WithArgs(userID, taskID).
		WillReturnError(errors.New(""))

	err = storage.CheckTaskPermission(userID, taskID)
	if err == nil {
		t.Errorf("expected err")
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestStorage_CheckTaskPermissionFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	taskID := int64(1)
	userID := int64(1)

	mock.
		ExpectQuery("SELECT B.boardID").
		WithArgs(userID, taskID).
		WillReturnError(sql.ErrNoRows)

	err = storage.CheckTaskPermission(userID, taskID)
	if err == nil {
		t.Errorf("expected err")
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}