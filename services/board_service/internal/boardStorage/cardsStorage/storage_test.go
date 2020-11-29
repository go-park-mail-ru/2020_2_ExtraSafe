package cardsStorage

import (
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"reflect"
	"testing"
)

func TestStorage_CreateCard(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := NewStorage(db)

	cardInput := models.CardInput{
		UserID:  1,
		CardID:  0,
		BoardID: 1,
		Name:    "todo",
		Order:   1,
	}

	expectCardOutside := models.CardOutside{
		CardID: 1,
		Name:   cardInput.Name,
		Order:  cardInput.Order,
		Tasks:  []models.TaskOutsideShort{},
	}

	mock.
		ExpectQuery("INSERT INTO cards").
		WithArgs(cardInput.BoardID, cardInput.Name, cardInput.Order).
		WillReturnRows(sqlmock.NewRows([]string{"cardID"}).AddRow(1))

	card, err := storage.CreateCard(cardInput)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(card, expectCardOutside) {
		t.Errorf("results not match, want %v, have %v", expectCardOutside, card)
		return
	}
}

func TestStorage_CreateCardFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	cardInput := models.CardInput{
		UserID:  1,
		CardID:  0,
		BoardID: 1,
		Name:    "todo",
		Order:   1,
	}

	mock.
		ExpectQuery("INSERT INTO cards").
		WithArgs(cardInput.BoardID, cardInput.Name, cardInput.Order).
		WillReturnError(sql.ErrNoRows)

	_, err = storage.CreateCard(cardInput)
	if err == nil {
		t.Error("expected err")
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestStorage_ChangeCard(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	cardInput := models.CardInput{
		UserID:  1,
		CardID:  1,
		BoardID: 1,
		Name:    "Changed todo",
		Order:   2,
	}

	expectCardOutside := models.CardInternal{
		CardID: 1,
		Name:   cardInput.Name,
		Order:  cardInput.Order,
	}

	mock.
		ExpectExec("UPDATE cards SET").
		WithArgs(cardInput.Name, cardInput.Order, cardInput.CardID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	card, err := storage.ChangeCard(cardInput)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(card, expectCardOutside) {
		t.Errorf("results not match, want %v, have %v", expectCardOutside, card)
		return
	}
}

func TestStorage_ChangeCardFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	cardInput := models.CardInput{
		UserID:  1,
		CardID:  1,
		BoardID: 1,
		Name:    "Changed todo",
		Order:   2,
	}

	mock.
		ExpectExec("UPDATE cards SET").
		WithArgs(cardInput.Name, cardInput.Order, cardInput.CardID).
		WillReturnError(errors.New("update exec error"))

	_, err = storage.ChangeCard(cardInput)
	if err == nil {
		t.Errorf("expected err")
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestStorage_DeleteCard(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	cardInput := models.CardInput{ CardID:  1 }

	mock.
		ExpectExec("DELETE FROM cards WHERE cardID").
		WithArgs(cardInput.CardID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = storage.DeleteCard(cardInput)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestStorage_DeleteCardFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	cardInput := models.CardInput{CardID: 1}

	mock.
		ExpectExec("DELETE FROM cards WHERE cardID").
		WithArgs(cardInput.CardID).
		WillReturnError(errors.New("delete exec error"))

	err = storage.DeleteCard(cardInput)
	if err == nil {
		t.Errorf("expected err")
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestStorage_GetCardByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	cardInput := models.CardInput{ CardID:  1}

	expectCardOutside := models.CardInternal{
		CardID: 1,
		Name:   "todo",
		Order:  1,
	}

	rows := sqlmock.NewRows([]string{"cardName", "cardOrder"}).AddRow("todo", 1)

	mock.
		ExpectQuery("SELECT cardName, cardOrder FROM cards WHERE cardID").
		WithArgs(cardInput.CardID).
		WillReturnRows(rows)

	card, err := storage.GetCardByID(cardInput)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(card, expectCardOutside) {
		t.Errorf("results not match, want %v, have %v", expectCardOutside, card)
		return
	}
}

func TestStorage_GetCardByIDFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	cardInput := models.CardInput{ CardID:  1}

	mock.
		ExpectQuery("SELECT cardName, cardOrder FROM cards WHERE cardID").
		WithArgs(cardInput.CardID).
		WillReturnError(sql.ErrNoRows)

	_, err = storage.GetCardByID(cardInput)
	if err == nil {
		t.Errorf("unexpected not err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestStorage_GetCardsByBoard(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	boardInput := models.BoardInput{ BoardID:  1}

	expectedCards := make([]models.CardInternal, 0)
	card1 := models.CardInternal{
		CardID: 1,
		Name:   "todo",
		Order:  1,
	}

	card2 := models.CardInternal{
		CardID: 2,
		Name:   "todo",
		Order:  2,
	}

	expectedCards = append(expectedCards, card1, card2)

	rows := sqlmock.NewRows([]string{"cardID", "cardName", "cardOrder"})
	rows.AddRow(card1.CardID, card1.Name, card1.Order).AddRow(card2.CardID, card2.Name, card2.Order)

	mock.
		ExpectQuery("SELECT cardID, cardName, cardOrder FROM cards WHERE boardID").
		WithArgs(boardInput.BoardID).
		WillReturnRows(rows)

	cards, err := storage.GetCardsByBoard(boardInput)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(cards, expectedCards) {
		t.Errorf("results not match, want %v, have %v", expectedCards, cards)
		return
	}
}

func TestStorage_GetCardsByBoardScanFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	boardInput := models.BoardInput{ BoardID:  1}

	expectedCards := make([]models.CardInternal, 0)
	card1 := models.CardInternal{
		CardID: 1,
		Name:   "todo",
		Order:  1,
	}

	expectedCards = append(expectedCards, card1)

	rows := sqlmock.NewRows([]string{"cardID", "cardName"})
	rows.AddRow(card1.CardID, card1.Name)

	mock.
		ExpectQuery("SELECT cardID, cardName, cardOrder FROM cards WHERE boardID").
		WithArgs(boardInput.BoardID).
		WillReturnRows(rows)

	_, err = storage.GetCardsByBoard(boardInput)
	if err == nil {
		t.Error("unexpected not err", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestStorage_GetCardsByBoardFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	boardInput := models.BoardInput{ BoardID:  1}

	mock.
		ExpectQuery("SELECT cardID, cardName, cardOrder FROM cards WHERE boardID").
		WithArgs(boardInput.BoardID).
		WillReturnError(sql.ErrNoRows)

	_, err = storage.GetCardsByBoard(boardInput)
	if err == nil {
		t.Error("unexpected not err", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestStorage_ChangeCardOrder(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	card := models.CardOrder{
		CardID: 1,
		Order:  2,
	}
	input := models.CardsOrderInput{}
	input.Cards = append(input.Cards, card)

	mock.
		ExpectExec("UPDATE cards SET").
		WithArgs(card.Order, card.CardID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = storage.ChangeCardOrder(input)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestStorage_ChangeCardOrderFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	card := models.CardOrder{
		CardID: 1,
		Order:  2,
	}
	input := models.CardsOrderInput{}
	input.Cards = append(input.Cards, card)

	mock.
		ExpectExec("UPDATE cards SET").
		WithArgs(card.Order, card.CardID).
		WillReturnError(errors.New("update exec error"))

	err = storage.ChangeCardOrder(input)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Error("expected err")
		return
	}
}