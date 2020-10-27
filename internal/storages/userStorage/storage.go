package userStorage

import (
	"database/sql"
	"fmt"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
)
/* users table
	userID      BIGSERIAL PRIMARY KEY,
	email       TEXT,
	password    TEXT,
	username    TEXT,
	fullname    TEXT,
	avatar      TEXT
*/

/* social_links table
	userID      BIGSERIAL,
	networkName TEXT,
	link TEXT
*/
type Storage interface {
	LoginUser(userInput models.UserInputLogin) (models.User, error)
	RegisterUser(userInput models.UserInputReg) (models.User, error)

	GetUserProfile(userInput models.UserInput) (models.User, error)
	GetUserAccounts(userInput models.UserInput) (models.User, error)

	ChangeUserProfile(user *models.User, userInput models.UserInputProfile) error

	ChangeUserAccounts(userInput models.UserInputLinks) (models.User, error)
	ChangeUserPassword(userInput models.UserInputPassword) (models.User, error)

	CheckExistingUser(email string, username string) (errorCodes []string)
	CheckExistingUserOnUpdate(email string, username string, userID uint64) (errorCodes []string)
}

type storage struct {
	Users    *[]models.User
	db *sql.DB
}

func NewStorage(db *sql.DB, someUsers *[]models.User) Storage {
	return &storage{
		db: db,
		Users: someUsers,
	}
}

func (s *storage) LoginUser(userInput models.UserInputLogin) (models.User, error) {
	user := models.User{}

	err := s.db.QueryRow("SELECT userID, username, fullname, avatar FROM users WHERE email = $1 AND password = $2", userInput.Email, userInput.Password).
				Scan(&user.ID, &user.Username, &user.FullName, &user.Avatar)

	//TODO сделать корректную обработку ошибок, приходящих из БД
	if err != sql.ErrNoRows {
		user.Password = userInput.Password
		user.Email = userInput.Email
		user.Links = &models.UserLinks{}

		fmt.Println(err)
		fmt.Printf("login info %+v\n", user)
		return user, nil
	}

	return models.User{}, models.ServeError{Codes: []string{"101"}}
}

func (s *storage) RegisterUser(userInput models.UserInputReg) (models.User, error) {
	errors := s.CheckExistingUser(userInput.Email, userInput.Username)
	if len(errors) != 0 {
		return models.User{}, models.ServeError{Codes: errors}
	}

	//TODO посмотреть способ нахождения последнего индекса из лекций по СУБД
	var ID uint64 = 0
	var quantityUsers uint64 = 0
	err := s.db.QueryRow("SELECT COUNT(*) FROM users").Scan(&quantityUsers)
	ID = quantityUsers + 1

	_, err = s.db.Exec("INSERT INTO users (userID, email, password, username, fullname, avatar) VALUES ($1, $2, $3, $4, $5, $6)",
						ID,
						userInput.Email,
						userInput.Password,
						userInput.Username,
						"",
						"default/default_avatar.png")
	if err != nil {
		//TODO разработать код ошибок на стороне БД
		fmt.Println(err)
		return models.User{}, models.ServeError{Codes: []string{"500"}}
	}

	user := models.User{
		ID:       ID,
		Username: userInput.Username,
		Email:    userInput.Email,
		Password: userInput.Password,
		Links:    &models.UserLinks{},
		Avatar:   "default/default_avatar.png",
	}

	fmt.Printf("register info %+v\n", user)
	*s.Users = append(*s.Users, user)

	return user, nil
}

//TODO подумать, как сделать эту проверку компактнее
func (s *storage) CheckExistingUser(email string, username string) (errorCodes []string){
	err := s.db.QueryRow("SELECT userID FROM users WHERE email = $1", email).Scan()
	if err != sql.ErrNoRows {
		errorCodes = append(errorCodes, "201")
	}

	//TODO другие ошибки БД
	err = s.db.QueryRow("SELECT userID FROM users WHERE username = $1", username).Scan()
	if err != sql.ErrNoRows {
		errorCodes = append(errorCodes, "202")
	}

	fmt.Println(err)
	return errorCodes
}

func (s *storage) CheckExistingUserOnUpdate(email string, username string, userID uint64) (errorCodes []string){
	var existingUserID uint64 = 0

	err := s.db.QueryRow("SELECT userID FROM users WHERE email = $1", email).Scan(&existingUserID)
	if err != sql.ErrNoRows && existingUserID != userID {
		errorCodes = append(errorCodes, "301")
	}

	err = s.db.QueryRow("SELECT userID FROM users WHERE username = $1", username).Scan(&existingUserID)
	if err != sql.ErrNoRows && existingUserID != userID  {
		errorCodes = append(errorCodes, "302")
	}

	fmt.Println(err)
	return errorCodes
}


func (s *storage) GetUserProfile(userInput models.UserInput) (models.User, error) {
	user := models.User{Links: &models.UserLinks{}}

	err := s.db.QueryRow("SELECT email, password, username, fullname, avatar FROM users WHERE userID = $1", userInput.ID).
		Scan(&user.Email, &user.Password, &user.Username, &user.FullName, &user.Avatar)

	//TODO сделать корректную обработку ошибок, приходящих из БД
	if err != sql.ErrNoRows {
		user.ID = userInput.ID
		fmt.Println(err)
		return user, nil
	}

	//TODO сделать правильную ошибку
	return user, models.ServeError{Codes: []string{"101"}}
}

//FIXME
func (s *storage) GetUserAccounts(userInput models.UserInput) (models.User, error) {
	someUser := new(models.User)
	for i, user := range *s.Users {
		if user.ID == userInput.ID {
			someUser = &(*s.Users)[i]
		}
	}
	return *someUser, nil
}


func (s *storage) ChangeUserProfile(user *models.User, userInput models.UserInputProfile) error {
	errorCodes := s.CheckExistingUserOnUpdate(userInput.Email, userInput.Username, user.ID)

	if len(errorCodes) != 0 {
		return models.ServeError{Codes: errorCodes}
	}

	_, err := s.db.Exec("UPDATE users SET email = $1, username = $2, fullname = $3, avatar = $4 WHERE userID = $5",
						userInput.Email, userInput.Username, userInput.FullName, user.Avatar, user.ID)
	if err != nil {
		//TODO разработать код ошибок на стороне БД
		fmt.Println(err)
		return models.ServeError{Codes: []string{"500"}}
	}

	user.Username = userInput.Username
	user.Email = userInput.Email
	user.FullName = userInput.FullName

	fmt.Printf("info after change %+v\n", user)
	return nil
}


//FIXME
func (s *storage) ChangeUserAccounts(userInput models.UserInputLinks) (models.User, error) {
	userExist := new(models.User)

	for i, user := range *s.Users {
		if user.ID == userInput.ID {
			userExist = &(*s.Users)[i]
		}
	}

	userExist.Links.Bitbucket = userInput.Bitbucket
	userExist.Links.Github = userInput.Github
	userExist.Links.Instagram = userInput.Instagram
	userExist.Links.Telegram = userInput.Telegram
	userExist.Links.Facebook = userInput.Facebook
	userExist.Links.Vk = userInput.Vk

	return *userExist, nil
}


func (s *storage) ChangeUserPassword(userInput models.UserInputPassword) (models.User, error) {
	user, err := s.GetUserProfile(models.UserInput{ ID : userInput.ID})
	if err != nil {
		fmt.Println(err)
		return models.User{}, err
	}

	if userInput.OldPassword != user.Password {
		errorCodes := []string{"501"}
		return models.User{}, models.ServeError{Codes: errorCodes}
	}

	_, err = s.db.Exec("UPDATE users SET password = $1 WHERE userID = $2", userInput.Password, user.ID)
	if err != nil {
		//TODO разработать код ошибок на стороне БД
		fmt.Println(err)
		return  models.User{}, models.ServeError{Codes: []string{"500"}}
	}

	return user, nil
}

