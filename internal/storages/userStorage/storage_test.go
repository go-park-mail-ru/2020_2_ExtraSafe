package userStorage

import (
	"bytes"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"reflect"
	"testing"
)

//TODO TESTS сделать тесты на ошибки

func TestStorage_GetUserProfile(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	userInput := models.UserInput{ ID: 1 }

	rows := sqlmock.NewRows([]string{"email", "username", "fullname", "avatar"})
	rows.AddRow("epridius@yandex.ru", "pkaterinaa", "", "default/default_avatar.png")
	expect := models.UserOutside{Email: "epridius@yandex.ru", Username: "pkaterinaa", Links: &models.UserLinks{}, Avatar: "default/default_avatar.png"}

	//query success
	mock.
		ExpectQuery("SELECT").
		WithArgs(userInput.ID).
		WillReturnRows(rows)

	user, err := storage.GetUserProfile(userInput)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(user, expect) {
		t.Errorf("results not match, want %v, have %v", expect, user)
		return
	}

	//no rows in db
	mock.
		ExpectQuery("SELECT").
		WithArgs(userInput.ID).
		WillReturnError(sql.ErrNoRows)

	_, err = storage.GetUserProfile(userInput)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}

	rows = sqlmock.NewRows([]string{"email", "username"}).
		AddRow("epridius@yandex.ru", "pkaterinaa")

	//bad scan result
	mock.
		ExpectQuery("SELECT").
		WithArgs(userInput.ID).
		WillReturnRows(rows)

	_, err = storage.GetUserProfile(userInput)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestStorage_GetUserAccounts(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	userInput := models.UserInput{ ID: 1 }

	profileRows := sqlmock.NewRows([]string{"email", "username", "fullname", "avatar"})
	profileRows.AddRow("epridius@yandex.ru", "pkaterinaa", "", "default/default_avatar.png")

	accountsRows := sqlmock.NewRows([]string{"networkName", "link"})
	accountsRows.AddRow("Instagram", "pkaterinaa").AddRow("Github", "pringleskate")
	
	accountsExpect := models.UserOutside{
		Email: "epridius@yandex.ru",
		Username: "pkaterinaa",
		Links: &models.UserLinks{
			Instagram: "pkaterinaa",
			Github: "pringleskate",
		},
		Avatar: "default/default_avatar.png",
	}

	mock.
		ExpectQuery("SELECT").
		WithArgs(userInput.ID).
		WillReturnRows(profileRows)

	mock.
		ExpectQuery("SELECT networkName, link FROM").
		WithArgs(userInput.ID).
		WillReturnRows(accountsRows)

	user, err := storage.GetUserAccounts(userInput)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(user, accountsExpect) {
		t.Errorf("results not match, want %v, have %v", accountsExpect, user)
		return
	}
}

func TestStorage_GetUserAvatar(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	userInput := models.UserInput{ ID: 1 }
	rows := sqlmock.NewRows([]string{"avatar"}).AddRow("default/default_avatar.png")

	expect := models.UserAvatar{
		ID:     userInput.ID,
		Avatar: "default/default_avatar.png",
	}

	mock.
		ExpectQuery("SELECT").
		WithArgs(userInput.ID).
		WillReturnRows(rows)

	avatar, err := storage.GetUserAvatar(userInput)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(avatar, expect) {
		t.Errorf("results not match, want %v, have %v", expect, avatar)
		return
	}

	mock.
		ExpectQuery("SELECT").
		WithArgs(userInput.ID).
		WillReturnError(sql.ErrNoRows)

	_, err = storage.GetUserAvatar(userInput)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestStorage_ChangeUserProfile(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	userInput := models.UserInputProfile{ ID: 1, Email: "epridius@yandex.ru", Username: "pringleskate", FullName : "Pridius Kate"}
	userAvatar := models.UserAvatar{Avatar: "pringlz_avatar"}
	expect := models.UserOutside{
		Email:    userInput.Email,
		Username: userInput.Username,
		FullName: userInput.FullName,
		Avatar:   userAvatar.Avatar,
	}

	mock.ExpectQuery("SELECT userID FROM users WHERE email").
		WithArgs(userInput.Email).
		WillReturnRows(sqlmock.NewRows([]string{"userID"}).AddRow(1))

	mock.ExpectQuery("SELECT userID FROM users WHERE username").
		WithArgs(userInput.Username).
		WillReturnError(sql.ErrNoRows)

	mock.ExpectExec("UPDATE").
		WithArgs(userInput.Email, userInput.Username, userInput.FullName, userAvatar.Avatar, userInput.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	user, err := storage.ChangeUserProfile(userInput, userAvatar)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(user, expect) {
		t.Errorf("results not match, want %v, have %v", expect, user)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestStorage_ChangeUserAccounts(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}
	userInput := models.UserInputLinks{
		ID : 1,
		Telegram: "pkaterinaa",
		Instagram : "",
		Github: "newPringleskate",
	}

	profileRows := sqlmock.NewRows([]string{"email", "username", "fullname", "avatar"})
	profileRows.AddRow("epridius@yandex.ru", "pkaterinaa", "", "default/default_avatar.png")

	accountsRows := sqlmock.NewRows([]string{"networkName", "link"})
	accountsRows.AddRow("Instagram", "pkaterinaa").AddRow("Github", "pringleskate")

	accountsExpect := models.UserOutside{
		Email: "epridius@yandex.ru",
		Username: "pkaterinaa",
		Links: &models.UserLinks{
			Telegram: "pkaterinaa",
			Github: "newPringleskate",
		},
		Avatar: "default/default_avatar.png",
	}

	mock.
		ExpectQuery("SELECT").
		WithArgs(userInput.ID).
		WillReturnRows(profileRows)

	mock.
		ExpectQuery("SELECT networkName, link FROM").
		WithArgs(userInput.ID).
		WillReturnRows(accountsRows)

	mock.ExpectExec("INSERT").
		WithArgs(userInput.ID, "Telegram", userInput.Telegram).
		WillReturnResult(sqlmock.NewResult(3, 1))

	mock.ExpectExec("DELETE").
		WithArgs(userInput.ID, "Instagram").
		WillReturnResult(sqlmock.NewResult(3, 1))

	mock.ExpectExec("UPDATE").
		WithArgs( userInput.Github, userInput.ID, "Github").
		WillReturnResult(sqlmock.NewResult(3, 1))

	user, err := storage.ChangeUserAccounts(userInput)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(user, accountsExpect) {
		t.Errorf("results not match, want %v, have %v", accountsExpect, user)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestStorage_GetInternalUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	userInput := models.UserInput{ ID: 1 }

	rows := sqlmock.NewRows([]string{"email", "password", "username", "fullname", "avatar"})
	rows.AddRow("epridius@yandex.ru", []byte("lalala"), "pkaterinaa", "", "default/default_avatar.png")
	expect := models.UserOutside{Email: "epridius@yandex.ru", Username: "pkaterinaa", Links: &models.UserLinks{}, Avatar: "default/default_avatar.png"}
	expectPass := []byte("lalala")

	//query success
	mock.
		ExpectQuery("SELECT").
		WithArgs(userInput.ID).
		WillReturnRows(rows)

	user, password, err := storage.GetInternalUser(userInput)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(user, expect) {
		t.Errorf("results not match, want %v, have %v", expect, user)
		return
	}
	if !bytes.Equal(expectPass, password) {
		t.Errorf("results not match, want %v, have %v", expectPass, password)
		return
	}
}

func TestStorage_CheckExistingUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	username := "pkaterinaa"
	email := "epridius@yandex.ru"
	expect := make([]string, 0)

	mock.ExpectQuery("SELECT userID FROM users WHERE email").
		WithArgs(email).
		WillReturnError(sql.ErrNoRows)

	mock.ExpectQuery("SELECT userID FROM users WHERE username").
		WithArgs(username).
		WillReturnError(sql.ErrNoRows)

	errorCodes := storage.CheckExistingUser(email, username)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	//FIXME
	if len(errorCodes.Codes) != 0 {
		t.Errorf("results not match, want %v, have %v", expect, errorCodes)
		return
	}
}

func TestStorage_GetBoardMembers(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	userIDs := []uint64{1, 2}

	expect := make([]models.UserOutsideShort, 0)
	first := models.UserOutsideShort{
		Email:    "epridius@yandex.ru",
		Username: "pkaterinaa",
		FullName: "",
		Avatar:   "default/default_avatar.png",
	}

	second := models.UserOutsideShort{
		Email:    "egoraist@gmail.com",
		Username: "egoraist",
		FullName: "",
		Avatar:   "default/default_avatar.png",
	}
	rowsFirst := sqlmock.NewRows([]string{"email", "username", "fullname", "avatar"})
	rowsFirst.AddRow(first.Email, first.Username, first.FullName, first.Avatar)

	rowsSecond := sqlmock.NewRows([]string{"email", "username", "fullname", "avatar"})
	rowsSecond.AddRow(second.Email, second.Username, second.FullName, second.Avatar)

	expect = append(expect, first, second)

	mock.ExpectQuery("SELECT").
		WithArgs(userIDs[0]).
		WillReturnRows(rowsFirst)

	mock.ExpectQuery("SELECT").
		WithArgs(userIDs[1]).
		WillReturnRows(rowsSecond)

	users, err := storage.GetBoardMembers(userIDs)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(expect, users) {
		t.Errorf("results not match, want %v, have %v", expect, users)
		return
	}
}

/*func TestCreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	userInput := models.UserInputReg{
		Email:    "epridius@yandex.ru",
		Password: "lalala",
		Username: "pkaterinaa",
	}

	//rows := sqlmock.NewRows([]string{"userID"})
	//expect := models.UserOutside{Email: "epridius@yandex.ru", Username: "pkaterinaa", Links: &models.UserLinks{}, Avatar: "default/default_avatar.png"}
	//rows = rows.AddRow(1)

	mock.ExpectQuery("SELECT").
		WithArgs(userInput.Email).
		WillReturnError(sql.ErrNoRows)


	mock.ExpectQuery("SELECT").
		WithArgs(userInput.Username).
		WillReturnError(sql.ErrNoRows)

	/////НЕ ПОЛУЧИТСЯ, ПОТОМУ ЧТО МЫ НЕ ЗНАЕМ СОЛЬ И НЕ МОЖЕМ ПРЕДОСТАВИТЬ ТОТ ЗАПРОС, КОТОРЫЙ УЙДЕТ В БД
	////FUCK
	mock.
		ExpectQuery("INSERT").
		WithArgs(userInput.Email, []byte(userInput.Password), userInput.Username, "", "default/default_avatar.png").
		//WillReturnRows(rows)
		WillReturnError(fmt.Errorf("db_error"))

	_, _, err = storage.CreateUser(userInput)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}

}

*/

/*
func TestCheckUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	userInput := models.UserInputLogin{
		Email:    "epridius@yandex.ru",
		Password: "lalala",
	}

	rows := sqlmock.NewRows([]string{"userID", "password", "username", "fullname", "avatar"})
	expect := models.UserOutside{Email: "epridius@yandex.ru", Username: "pkaterinaa", Links: &models.UserLinks{}, Avatar: "default/default_avatar.png"}

	rows = rows.AddRow(1, []byte{}, expect.Username, expect.FullName, expect.Avatar)
	mock.
		ExpectQuery("SELECT").
		WithArgs(userInput.Email).WillReturnError(models.ServeError{Codes: []string{"101"}})

	_, _, err2 := storage.CheckUser(userInput)
	if err2 != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err2 := mock.ExpectationsWereMet(); err2 != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err2 == nil {
		t.Errorf("expected error, got nil")
		return
	}
}*/
	// query error
	/*mock.
		ExpectQuery("SELECT userID, password, username, fullname, avatar FROM users WHERE email").
		WithArgs(userInput.Email).
		WillReturnError(fmt.Errorf("db_error"))

	_, user, err3 := storage.CheckUser(userInput)
	if err3 := mock.ExpectationsWereMet(); err3 != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err3 == nil {
		t.Errorf("expected error, got nil")
		return
	}*/

	// row scan error
/*	rows = sqlmock.NewRows([]string{"id", "title"}).
		AddRow(1, "title")

	mock.
		ExpectQuery("SELECT id, title, updated, description FROM items WHERE").
		WithArgs(elemID).
		WillReturnRows(rows)

	_, err = repo.GetByID(elemID)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}*/

