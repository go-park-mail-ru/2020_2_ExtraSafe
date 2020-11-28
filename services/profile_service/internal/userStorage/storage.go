package userStorage

import (
	"bytes"
	"crypto/rand"
	"database/sql"
	"fmt"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"golang.org/x/crypto/argon2"
	"reflect"
)

type Storage interface {
	CheckUser(userInput models.UserInputLogin) (int64, models.UserOutside, error)
	CreateUser(userInput models.UserInputReg) (int64, models.UserOutside, error)

	GetUserProfile(userInput models.UserInput) (models.UserOutside, error)
	GetUserAccounts(userInput models.UserInput) (models.UserOutside, error)
	GetUserAvatar(userInput models.UserInput) (models.UserAvatar, error)

	GetUsersByIDs(userIDs []int64) ([] models.UserOutsideShort, error) // 0 структура - админ доски
	GetUserByUsername(username string) (user models.UserInternal, err error)
	//GetUserByID(userID int64) (models.UserOutsideShort, error)

	ChangeUserProfile(userInput models.UserInputProfile, userAvatar models.UserAvatar) (models.UserOutside, error)
	ChangeUserAccounts(userInput models.UserInputLinks) (models.UserOutside, error)
	ChangeUserPassword(userInput models.UserInputPassword) (models.UserOutside, error)

	CheckExistingUser(email string, username string) (errors models.MultiErrors)
	checkExistingUserOnUpdate(email string, username string, userID int64) (errors models.MultiErrors)
	GetInternalUser(userInput models.UserInput) (models.UserOutside, []byte, error)
}

type storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) Storage {
	return &storage{
		db: db,
	}
}

func (s *storage) CheckUser(userInput models.UserInputLogin) (int64, models.UserOutside, error) {
	user := models.UserOutside{}
	var userID int64

	pass := make([]byte, 0)
	err := s.db.QueryRow("SELECT userID, password, username, fullname, avatar FROM users WHERE email = $1", userInput.Email).
				Scan(&userID, &pass, &user.Username, &user.FullName, &user.Avatar)

	if err == sql.ErrNoRows {
		return 0, models.UserOutside{}, models.ServeError{Codes: []string{"101"}, Descriptions: []string{"Invalid email "}, MethodName: "CheckUser"}
	}

	if err != nil {
		return 0, models.UserOutside{}, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "CheckUser"}
	}

	if checkPass(pass, userInput.Password) {
		user.Email = userInput.Email
		user.Links = &models.UserLinks{}
		return userID, user, nil
	}

	return 0, models.UserOutside{}, models.ServeError{
			Codes: []string{"101"},
			Descriptions: []string{"Invalid password"},
			MethodName: "CheckUser" }
}

func (s *storage) CreateUser(userInput models.UserInputReg) (int64, models.UserOutside, error) {
	multiErrors := s.CheckExistingUser(userInput.Email, userInput.Username)
	if len(multiErrors.Codes) != 0 {
		return 0, models.UserOutside{}, models.ServeError{Codes: multiErrors.Codes,
			Descriptions: multiErrors.Descriptions}
	}

	var ID int64
	salt := make([]byte, 8)

	rand.Read(salt)
	hashedPass := hashPass(salt, userInput.Password)

	err := s.db.QueryRow("INSERT INTO users (email, password, username, fullname, avatar) VALUES ($1, $2, $3, $4, $5) RETURNING userID",
						userInput.Email,
						hashedPass,
						userInput.Username,
						"",
						"default/default_avatar.png").Scan(&ID)
	if err != nil {
		return 0, models.UserOutside{}, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "CreateUser"}
	}

	user := models.UserOutside{
		Email:    userInput.Email,
		Username: userInput.Username,
		FullName: "",
		Links:    &models.UserLinks{},
		Avatar:   "default/default_avatar.png",
	}

	return ID, user, nil
}

func (s *storage) CheckExistingUser(email string, username string) (multiErrors models.MultiErrors) {
	err := s.db.QueryRow("SELECT userID FROM users WHERE email = $1", email).Scan()
	if err != sql.ErrNoRows {
		multiErrors.Codes = append(multiErrors.Codes, "201")
		multiErrors.Descriptions = append(multiErrors.Descriptions, "Email is already exist")
	}

	err = s.db.QueryRow("SELECT userID FROM users WHERE username = $1", username).Scan()
	if err != sql.ErrNoRows {
		multiErrors.Codes = append(multiErrors.Codes, "202")
		multiErrors.Descriptions = append(multiErrors.Descriptions, "Username is already exist")
	}

	fmt.Println(err)
	return multiErrors
}

func (s *storage) checkExistingUserOnUpdate(email string, username string, userID int64) (multiErrors models.MultiErrors){
	var existingUserID int64 = 0

	err := s.db.QueryRow("SELECT userID FROM users WHERE email = $1", email).Scan(&existingUserID)
	if err != sql.ErrNoRows && existingUserID != userID {
		multiErrors.Codes = append(multiErrors.Codes, "301")
		multiErrors.Descriptions = append(multiErrors.Descriptions, "Email is already exist")
	}

	err = s.db.QueryRow("SELECT userID FROM users WHERE username = $1", username).Scan(&existingUserID)
	if err != sql.ErrNoRows && existingUserID != userID  {
		multiErrors.Codes = append(multiErrors.Codes, "302")
		multiErrors.Descriptions = append(multiErrors.Descriptions, "Username is already exist")
	}

	fmt.Println(err)
	return multiErrors
}

func (s *storage) GetUserProfile(userInput models.UserInput) (models.UserOutside, error) {
	user := models.UserOutside{Links: &models.UserLinks{}}

	err := s.db.QueryRow("SELECT email, username, fullname, avatar FROM users WHERE userID = $1", userInput.ID).
		Scan(&user.Email, &user.Username, &user.FullName, &user.Avatar)

	if err != sql.ErrNoRows {
		if err != nil {
			fmt.Println(err)
			return models.UserOutside{}, models.ServeError{Codes: []string{"500"}, OriginalError: err,
				MethodName: "GetUserProfile"}
		}
		return user, nil
	}

	return user, models.ServeError{Codes: []string{"500"}, OriginalError: err, MethodName: "GetUserProfile"}
}

func (s *storage) GetUserAccounts(userInput models.UserInput) (models.UserOutside, error) {
	user, err := s.GetUserProfile(userInput)
	if err != nil {
		return models.UserOutside{}, err
	}

	rows, err := s.db.Query("SELECT networkName, link FROM social_links WHERE userID = $1", userInput.ID)
	if err != nil {
		return models.UserOutside{}, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "GetUserAccounts"}
	}
	defer rows.Close()

	for rows.Next() {
		var networkName, link string

		err = rows.Scan(&networkName, &link)
		if err != nil {
			return models.UserOutside{}, models.ServeError{Codes: []string{"500"}, OriginalError: err,
				MethodName: "GetUserAccounts"}
		}

		reflect.Indirect(reflect.ValueOf(user.Links)).FieldByName(networkName).SetString(link)
	}

	return user, nil
}


func (s *storage) GetUserAvatar(userInput models.UserInput) (models.UserAvatar, error) {
	user := models.UserAvatar{}

	err := s.db.QueryRow("SELECT avatar FROM users WHERE userID = $1", userInput.ID).
		Scan(&user.Avatar)

	if err == nil {
		user.ID = userInput.ID
		return user, nil
	}

	fmt.Println(err)
	return user, models.ServeError{Codes: []string{"500"}, OriginalError: err, MethodName: "GetUserAvatar"}
}

func (s *storage) ChangeUserProfile(userInput models.UserInputProfile, userAvatar models.UserAvatar) (models.UserOutside, error) {
	multiErrors := s.checkExistingUserOnUpdate(userInput.Email, userInput.Username, userInput.ID)

	if len(multiErrors.Codes) != 0 {
		return models.UserOutside{}, models.ServeError{Codes: multiErrors.Codes, Descriptions: multiErrors.Descriptions,
			MethodName: "ChangeUserProfile"}
	}

	_, err := s.db.Exec("UPDATE users SET email = $1, username = $2, fullname = $3, avatar = $4 WHERE userID = $5",
		userInput.Email, userInput.Username, userInput.FullName, userAvatar.Avatar, userInput.ID)
	if err != nil {
		fmt.Println(err)
		return models.UserOutside{}, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "ChangeUserProfile"}
	}

	user := models.UserOutside {
		Username : userInput.Username,
		Email : userInput.Email,
		FullName : userInput.FullName,
		Avatar : userAvatar.Avatar,
	}

	return user, nil
}

func (s *storage) ChangeUserAccounts(userInput models.UserInputLinks) (models.UserOutside, error) {
	user, err := s.GetUserAccounts(models.UserInput{ ID : userInput.ID })
	if err != nil {
		fmt.Println(err)
		return models.UserOutside{}, err
	}

	networkNames := []string{"", "Telegram", "Instagram", "Github", "Bitbucket", "Vk", "Facebook"}

	input := reflect.ValueOf(userInput)
	for i := 1; i < input.NumField(); i++ {
		inputLink := input.Field(i).Interface().(string)
		if  inputLink == reflect.Indirect(reflect.ValueOf(user.Links)).FieldByName(networkNames[i]).String() {
			continue
		}

		if reflect.Indirect(reflect.ValueOf(user.Links)).FieldByName(networkNames[i]).String() == "" {
			_, err = s.db.Exec("INSERT INTO social_links (userID, networkName, link) VALUES ($1, $2, $3)", userInput.ID, networkNames[i], inputLink)
		} else if inputLink == "" {
			_, err = s.db.Exec("DELETE FROM social_links WHERE userID = $1 AND networkName = $2", userInput.ID, networkNames[i])
		} else {
			_, err = s.db.Exec("UPDATE social_links SET link = $1 WHERE userID = $2 AND networkName = $3", inputLink, userInput.ID, networkNames[i])
		}

		reflect.Indirect(reflect.ValueOf(user.Links)).FieldByName(networkNames[i]).SetString(inputLink)
	}

	return user, nil
}

func (s *storage) ChangeUserPassword(userInput models.UserInputPassword) (models.UserOutside, error) {
	user, password, err := s.GetInternalUser(models.UserInput{ ID : userInput.ID})
	if err != nil {
		fmt.Println(err)
		return models.UserOutside{}, err
	}

	if !checkPass(password, userInput.OldPassword) {
		return models.UserOutside{}, models.ServeError{Codes: []string{"501"},
			Descriptions: []string{"Passwords is not equal"}, MethodName: "ChangeUserPassword"}
	}

	salt := make([]byte, 8)

	rand.Read(salt)
	userPassHash := hashPass(salt, userInput.Password)

	_, err = s.db.Exec("UPDATE users SET password = $1 WHERE userID = $2", userPassHash, userInput.ID)
	if err != nil {
		fmt.Println(err)
		return  models.UserOutside{}, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "ChangeUserPassword"}
	}

	return user, nil
}

func (s *storage) GetInternalUser(userInput models.UserInput) (models.UserOutside, []byte, error) {
	user := models.UserOutside{Links: &models.UserLinks{}}
	var password []byte

	err := s.db.QueryRow("SELECT email, password, username, fullname, avatar FROM users WHERE userID = $1", userInput.ID).
		Scan(&user.Email, &password, &user.Username, &user.FullName, &user.Avatar)

	if err == nil {
		return user, password , nil
	}

	return user, nil, models.ServeError{Codes: []string{"500"}, OriginalError: err, MethodName: "getInternalUser"}
}

func (s *storage) GetUsersByIDs(userIDs []int64) ([] models.UserOutsideShort, error) {
	members := make([]models.UserOutsideShort, 0)

	for _, userID := range userIDs {
		user := models.UserOutsideShort{}
		err := s.db.QueryRow("SELECT email, username, fullname, avatar FROM users WHERE userID = $1", userID).
			Scan(&user.Email, &user.Username, &user.FullName, &user.Avatar)

		if err != nil {
			fmt.Println(err)
			return []models.UserOutsideShort{}, models.ServeError{Codes: []string{"500"}, OriginalError: err,
				MethodName: "GetUsersByIDs"}
		}

		members = append(members, user)
	}

	return members, nil
}

func (s *storage) GetUserByUsername(username string) (user models.UserInternal, err error) {
	err = s.db.QueryRow("SELECT userID, email, username,  fullname, avatar FROM users WHERE username = $1", username).
		Scan(&user.ID, &user.Email, &user.Username, &user.FullName, &user.Avatar)

	if err != sql.ErrNoRows {
		if err != nil {
			return user, models.ServeError{Codes: []string{"500"}, OriginalError: err,
				MethodName: "GetUserByUsername"}
		}
		return user, nil
	}

	return user, models.ServeError{Codes: []string{"500"}, OriginalError: err, MethodName: "GetUserByUsername"}
}

func hashPass(salt []byte, plainPassword string) []byte {
	hashedPass := argon2.IDKey([]byte(plainPassword), salt, 1, 64*1024, 4, 32)
	return append(salt, hashedPass...)
}

func checkPass(passHash []byte, plainPassword string) bool {
	salt := make([]byte, 8)
	copy(salt, passHash)

	userPassHash := hashPass(salt, plainPassword)
	return bytes.Equal(userPassHash, passHash)
}
