package userStorage

import (
	"bytes"
	"crypto/rand"
	"database/sql"
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"reflect"
	"testing"
)

func TestStorage_CreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	userInput := models.UserInputReg{
		Email:    "epridius@yandex.ru",
		Username: "pringleskate",
		Password: "1234567",
	}

	expect := models.UserOutside{
		Email:    userInput.Email,
		Username: userInput.Username,
		FullName: "",
		Links: &models.UserLinks{},
		Avatar: "default/default_avatar.png",
	}

	mock.ExpectQuery("SELECT userID FROM users WHERE email").
		WithArgs(userInput.Email).
		WillReturnError(sql.ErrNoRows)

	mock.ExpectQuery("SELECT userID FROM users WHERE username").
		WithArgs(userInput.Username).
		WillReturnError(sql.ErrNoRows)

	mock.ExpectQuery("INSERT").
		WithArgs(userInput.Email, sqlmock.AnyArg(), userInput.Username, "", "default/default_avatar.png").
		WillReturnRows(sqlmock.NewRows([]string{"userID"}).AddRow(1))

	_, user, err := storage.CreateUser(userInput)
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

func TestStorage_CreateUserFailOnCheck(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	userInput := models.UserInputReg{
		Email:    "epridius@yandex.ru",
		Username: "pringleskate",
		Password: "1234567",
	}

	mock.ExpectQuery("SELECT userID FROM users WHERE email").
		WithArgs(userInput.Email).
		WillReturnRows(sqlmock.NewRows([]string{"userID"}).AddRow(1))

	_, _, err = storage.CreateUser(userInput)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Error("expected err, got nil")
		return
	}
}

func TestStorage_CreateUserFailOnInsert(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	userInput := models.UserInputReg{
		Email:    "epridius@yandex.ru",
		Username: "pringleskate",
		Password: "1234567",
	}

	mock.ExpectQuery("SELECT userID FROM users WHERE email").
		WithArgs(userInput.Email).
		WillReturnError(sql.ErrNoRows)

	mock.ExpectQuery("SELECT userID FROM users WHERE username").
		WithArgs(userInput.Username).
		WillReturnError(sql.ErrNoRows)

	mock.ExpectQuery("INSERT").
		WithArgs(userInput.Email, sqlmock.AnyArg(), userInput.Username, "", "default/default_avatar.png").
		WillReturnError(sql.ErrNoRows)

	_, _, err = storage.CreateUser(userInput)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Error("expected err, got nil")
		return
	}
}

func TestStorage_CheckUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	userInput := models.UserInputLogin{
		Email:   "epridius@yandex.ru",
		Password: "1234567",
	}

	salt := make([]byte, 8)
	rand.Read(salt)
	hashedPass := hashPass(salt, userInput.Password)

	expect := models.UserOutside{
		Email:    userInput.Email,
		Username: "lala",
		FullName: "",
		Links: &models.UserLinks{},
		Avatar: "default/default_avatar.png",
	}

	rows := sqlmock.NewRows([]string{"userID", "password", "username", "fullname", "avatar"})
	rows.AddRow(1, hashedPass, expect.Username, expect.FullName, expect.Avatar)
	mock.ExpectQuery("SELECT").
		WithArgs(userInput.Email).
		WillReturnRows(rows)

	_, user, err := storage.CheckUser(userInput)
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

func TestStorage_CheckUserFailOnSelect(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	userInput := models.UserInputLogin{
		Email:   "epridius@yandex.ru",
		Password: "1234567",
	}

	mock.ExpectQuery("SELECT").
		WithArgs(userInput.Email).
		WillReturnError(sql.ErrNoRows)

	_, _, err = storage.CheckUser(userInput)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Error("expected err, got nil")
		return
	}
}


func TestStorage_CheckUserFailOnInternalError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	userInput := models.UserInputLogin{
		Email:   "epridius@yandex.ru",
		Password: "1234567",
	}

	mock.ExpectQuery("SELECT").
		WithArgs(userInput.Email).
		WillReturnError(errors.New("internal db error"))

	_, _, err = storage.CheckUser(userInput)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Error("expected err, got nil")
		return
	}
}

func TestStorage_CheckUserFailOnCheckPassword(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	userInput := models.UserInputLogin{
		Email:   "epridius@yandex.ru",
		Password: "1234",
	}

	salt := make([]byte, 8)
	rand.Read(salt)
	hashedPass := hashPass(salt, "1234567")

	expect := models.UserOutside{
		Email:    userInput.Email,
		Username: "lala",
		FullName: "",
		Links: &models.UserLinks{},
		Avatar: "default/default_avatar.png",
	}

	rows := sqlmock.NewRows([]string{"userID", "password", "username", "fullname", "avatar"})
	rows.AddRow(1, hashedPass, expect.Username, expect.FullName, expect.Avatar)
	mock.ExpectQuery("SELECT").
		WithArgs(userInput.Email).
		WillReturnRows(rows)

	_, _, err = storage.CheckUser(userInput)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Error("expected err, got nil")
		return
	}
}

func TestStorage_ChangeUserPassword(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	userInput := models.UserInputPassword{
		ID:          1,
		OldPassword: "lalala",
		Password:    "newPswd",
	}

	salt := make([]byte, 8)
	rand.Read(salt)
	hashedPass := hashPass(salt, userInput.OldPassword)

	rows := sqlmock.NewRows([]string{"email", "password", "username", "fullname", "avatar"})
	rows.AddRow("epridius@yandex.ru", hashedPass, "pkaterinaa", "", "default/default_avatar.png")
	expect := models.UserOutside{Email: "epridius@yandex.ru", Username: "pkaterinaa", Links: &models.UserLinks{}, Avatar: "default/default_avatar.png"}

	mock.
		ExpectQuery("SELECT").
		WithArgs(userInput.ID).
		WillReturnRows(rows)

	mock.
		ExpectExec("UPDATE").
		WithArgs(sqlmock.AnyArg(), userInput.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	user, err := storage.ChangeUserPassword(userInput)

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
}

func TestStorage_ChangeUserPasswordFailOnGetUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	userInput := models.UserInputPassword{
		ID:          1,
		OldPassword: "lalala",
		Password:    "newPswd",
	}

	mock.
		ExpectQuery("SELECT").
		WithArgs(userInput.ID).WillReturnError(sql.ErrNoRows)

	_, err = storage.ChangeUserPassword(userInput)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Error("expected err, got nil")
		return
	}
}

func TestStorage_ChangeUserPasswordFailOnCheckPass(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	userInput := models.UserInputPassword{
		ID:          1,
		OldPassword: "lalala",
		Password:    "newPswd",
	}

	salt := make([]byte, 8)
	rand.Read(salt)
	hashedPass := hashPass(salt, userInput.Password)

	rows := sqlmock.NewRows([]string{"email", "password", "username", "fullname", "avatar"})
	rows.AddRow("epridius@yandex.ru", hashedPass, "pkaterinaa", "", "default/default_avatar.png")

	mock.
		ExpectQuery("SELECT").
		WithArgs(userInput.ID).
		WillReturnRows(rows)

	_, err = storage.ChangeUserPassword(userInput)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Error("expected err, got nil")
		return
	}
}

func TestStorage_ChangeUserPasswordFailOnUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	userInput := models.UserInputPassword{
		ID:          1,
		OldPassword: "lalala",
		Password:    "newPswd",
	}

	salt := make([]byte, 8)
	rand.Read(salt)
	hashedPass := hashPass(salt, userInput.OldPassword)

	rows := sqlmock.NewRows([]string{"email", "password", "username", "fullname", "avatar"})
	rows.AddRow("epridius@yandex.ru", hashedPass, "pkaterinaa", "", "default/default_avatar.png")

	mock.
		ExpectQuery("SELECT").
		WithArgs(userInput.ID).
		WillReturnRows(rows)

	mock.
		ExpectExec("UPDATE").
		WithArgs(sqlmock.AnyArg(), userInput.ID).
		WillReturnError(errors.New("err on update"))

	_, err = storage.ChangeUserPassword(userInput)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Error("expected err, got nil")
		return
	}
}

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


func TestStorage_GetUserAccountsFirstFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := NewStorage(db)

	userInput := models.UserInput{ ID: 1 }

	mock.
		ExpectQuery("SELECT").
		WithArgs(userInput.ID).
		WillReturnError(sql.ErrNoRows)

	_, err = storage.GetUserAccounts(userInput)
	if err == nil {
		t.Errorf("expected err, got nil: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestStorage_GetUserAccountsSecondFail(t *testing.T) {
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

	mock.
		ExpectQuery("SELECT").
		WithArgs(userInput.ID).
		WillReturnRows(profileRows)

	mock.
		ExpectQuery("SELECT networkName, link FROM").
		WithArgs(userInput.ID).
		WillReturnError(errors.New("some error in selecting accounts data"))

	_, err = storage.GetUserAccounts(userInput)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		fmt.Println("expected error, got nil")
		return
	}
}

func TestStorage_GetUserAccountsThirdFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	userInput := models.UserInput{ ID: 1 }

	profileRows := sqlmock.NewRows([]string{"email", "username", "fullname", "avatar"})
	profileRows.AddRow("epridius@yandex.ru", "pkaterinaa", "", "default/default_avatar.png")

	accountsRows := sqlmock.NewRows([]string{"networkName", "link", "anotherColumn"})
	accountsRows.AddRow("Instagram", "pkaterinaa", "lala")

	mock.
		ExpectQuery("SELECT").
		WithArgs(userInput.ID).
		WillReturnRows(profileRows)

	mock.
		ExpectQuery("SELECT networkName, link FROM").
		WithArgs(userInput.ID).
		WillReturnRows(accountsRows)

	_, err = storage.GetUserAccounts(userInput)
	fmt.Println(err)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		fmt.Println("expected error, got nil")
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

func TestStorage_ChangeUserProfileFirstFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	userInput := models.UserInputProfile{ ID: 1, Email: "epridius@yandex.ru", Username: "pringleskate", FullName : "Pridius Kate"}
	userAvatar := models.UserAvatar{Avatar: "pringlz_avatar"}

	mock.ExpectQuery("SELECT userID FROM users WHERE email").
		WithArgs(userInput.Email).
		WillReturnRows(sqlmock.NewRows([]string{"userID"}).AddRow(1))

	mock.ExpectQuery("SELECT userID FROM users WHERE username").
		WithArgs(userInput.Username).
		WillReturnRows(sqlmock.NewRows([]string{"userID"}).AddRow(3))

	_, err = storage.ChangeUserProfile(userInput, userAvatar)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
	if err == nil {
		t.Errorf("expected err, got nil: %s", err)
		return
	}
}

func TestStorage_ChangeUserProfileSecondFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	userInput := models.UserInputProfile{ ID: 1, Email: "epridius@yandex.ru", Username: "pringleskate", FullName : "Pridius Kate"}
	userAvatar := models.UserAvatar{Avatar: "pringlz_avatar"}

	mock.ExpectQuery("SELECT userID FROM users WHERE email").
		WithArgs(userInput.Email).
		WillReturnRows(sqlmock.NewRows([]string{"userID"}).AddRow(1))

	mock.ExpectQuery("SELECT userID FROM users WHERE username").
		WithArgs(userInput.Username).
		WillReturnError(sql.ErrNoRows)

	mock.ExpectExec("UPDATE").
		WithArgs(userInput.Email, userInput.Username, userInput.FullName, userAvatar.Avatar, userInput.ID).
		WillReturnError(errors.New("error on update"))

	_, err = storage.ChangeUserProfile(userInput, userAvatar)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected err, got nil: %s", err)
		return
	}
}

func TestStorage_ChangeUserProfileThirdFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	userInput := models.UserInputProfile{ ID: 1, Email: "epridius@yandex.ru", Username: "pringleskate", FullName : "Pridius Kate"}
	userAvatar := models.UserAvatar{Avatar: "pringlz_avatar"}

	mock.ExpectQuery("SELECT userID FROM users WHERE email").
		WithArgs(userInput.Email).
		WillReturnRows(sqlmock.NewRows([]string{"userID"}).AddRow(2))

	_, err = storage.ChangeUserProfile(userInput, userAvatar)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected err, got nil: %s", err)
		return
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

func TestStorage_ChangeUserAccountsFail(t *testing.T) {
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

	mock.
		ExpectQuery("SELECT").
		WithArgs(userInput.ID).
		WillReturnError(sql.ErrNoRows)

	_, err = storage.ChangeUserAccounts(userInput)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Error("expected err, got nil")
		return
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

func TestStorage_GetInternalUserFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	userInput := models.UserInput{ ID: 1 }

	mock.
		ExpectQuery("SELECT").
		WithArgs(userInput.ID).
		WillReturnError(sql.ErrNoRows)

	_, _, err = storage.GetInternalUser(userInput)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

	if err == nil {
		t.Error("expected error, got nil")
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

	if len(errorCodes.Codes) != 0 {
		t.Errorf("results not match, want %v, have %v", expect, errorCodes)
		return
	}
}

func TestStorage_CheckExistingUserFirstFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	username := "pkaterinaa"
	email := "epridius@yandex.ru"

	expectedErrors := models.MultiErrors{Codes: []string{"201"}, Descriptions: []string{"Email is already exist"}}

	rows := sqlmock.NewRows([]string{"userID"}).AddRow(1)

	mock.ExpectQuery("SELECT userID FROM users WHERE email").
		WithArgs(email).
		WillReturnRows(rows)

	mock.ExpectQuery("SELECT userID FROM users WHERE username").
		WithArgs(username).
		WillReturnError(sql.ErrNoRows)

	errorCodes := storage.CheckExistingUser(email, username)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

	if !reflect.DeepEqual(errorCodes, expectedErrors) {
		t.Errorf("results not match, want %v, have %v", expectedErrors, errorCodes)
		return
	}
}

func TestStorage_CheckExistingUserSecondFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}

	username := "pkaterinaa"
	email := "epridius@yandex.ru"

	expectedErrors := models.MultiErrors{Codes: []string{"201", "202"}, Descriptions: []string{"Email is already exist", "Username is already exist"}}

	rows := sqlmock.NewRows([]string{"userID"}).AddRow(1)
	secRows := sqlmock.NewRows([]string{"userID"}).AddRow(1)

	mock.ExpectQuery("SELECT userID FROM users WHERE email").
		WithArgs(email).
		WillReturnRows(rows)

	mock.ExpectQuery("SELECT userID FROM users WHERE username").
		WithArgs(username).
		WillReturnRows(secRows)

	errorCodes := storage.CheckExistingUser(email, username)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

	if !reflect.DeepEqual(errorCodes, expectedErrors) {
		t.Errorf("results not match, want %v, have %v", expectedErrors, errorCodes)
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

func TestStorage_GetBoardMembersFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := &storage{db: db}
	userIDs := []uint64{1, 2}

	mock.ExpectQuery("SELECT").
		WithArgs(userIDs[0]).
		WillReturnError(sql.ErrNoRows)

	_, err = storage.GetBoardMembers(userIDs)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Error("expected err, got nil")
		return
	}
}