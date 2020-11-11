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
	CheckUser(userInput models.UserInputLogin) (uint64, models.UserOutside, error)
	CreateUser(userInput models.UserInputReg) (uint64, models.UserOutside, error)

	GetUserProfile(userInput models.UserInput) (models.UserOutside, error)
	GetUserAccounts(userInput models.UserInput) (models.UserOutside, error)
	GetUserAvatar(userInput models.UserInput) (models.UserAvatar, error)
	GetBoardMembers(userIDs []uint64) ([] models.UserOutsideShort, error) // 0 структура - админ доски

	ChangeUserProfile(userInput models.UserInputProfile, userAvatar models.UserAvatar) (models.UserOutside, error)
	ChangeUserAccounts(userInput models.UserInputLinks) (models.UserOutside, error)
	ChangeUserPassword(userInput models.UserInputPassword) (models.UserOutside, error)

	CheckExistingUser(email string, username string) (errorCodes []string)
	CheckExistingUserOnUpdate(email string, username string, userID uint64) (errorCodes []string)
	GetInternalUser(userInput models.UserInput) (models.UserOutside, []byte, error)
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

func (s *storage) CheckUser(userInput models.UserInputLogin) (uint64, models.UserOutside, error) {
	user := models.UserOutside{}
	var userID uint64

	pass := make([]byte, 0)
	err := s.db.QueryRow("SELECT userID, password, username, fullname, avatar FROM users WHERE email = $1", userInput.Email).
				Scan(&userID, &pass, &user.Username, &user.FullName, &user.Avatar)

	//TODO error
	switch err {
	case sql.ErrNoRows:
		return 0, models.UserOutside{}, models.ServeError{Codes: []string{"101"}}
	case nil:
		/*passValid := checkPass(pass, userInput.Password)
		fmt.Printf("passValid: %v, pass: %x\n", passValid, pass)
		user.Email = userInput.Email
		user.Links = &models.UserLinks{}
		return userID, user, nil*/
		break
	default:
		fmt.Println(err)
		return 0, models.UserOutside{}, models.ServeError{Codes: []string{"500"}}
	}

	passValid := checkPass(pass, userInput.Password)
	if !passValid {
		return 0, models.UserOutside{}, models.ServeError{Codes: []string{"101"}}
	}
	fmt.Printf("passValid: %v\n", passValid)
	user.Email = userInput.Email
	user.Links = &models.UserLinks{}
	return userID, user, nil
}

func (s *storage) CreateUser(userInput models.UserInputReg) (uint64, models.UserOutside, error) {
	errors := s.CheckExistingUser(userInput.Email, userInput.Username)
	if len(errors) != 0 {
		return 0, models.UserOutside{}, models.ServeError{Codes: errors}
	}

	var ID uint64


	salt := make([]byte, 8)
	rand.Read(salt)
	fmt.Printf("salt: %x\n", salt)
	hashedPass := hashPass(salt, userInput.Password)

	err := s.db.QueryRow("INSERT INTO users (email, password, username, fullname, avatar) VALUES ($1, $2, $3, $4, $5) RETURNING userID",
						userInput.Email,
						hashedPass,
						userInput.Username,
						"",
						"default/default_avatar.png").Scan(&ID)
	if err != nil {
		//TODO error
		fmt.Println(err)
		return 0, models.UserOutside{}, models.ServeError{Codes: []string{"500"}}
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

//FIXME подумать, как сделать эту проверку компактнее
//TODO разобраться с pq.Error (какая информация, как логировать)
func (s *storage) CheckExistingUser(email string, username string) (errorCodes []string) {
	err := s.db.QueryRow("SELECT userID FROM users WHERE email = $1", email).Scan()
	if err != sql.ErrNoRows {
		errorCodes = append(errorCodes, "201")
	}

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

func (s *storage) GetUserProfile(userInput models.UserInput) (models.UserOutside, error) {
	user := models.UserOutside{Links: &models.UserLinks{}}

	err := s.db.QueryRow("SELECT email, username, fullname, avatar FROM users WHERE userID = $1", userInput.ID).
		Scan(&user.Email, &user.Username, &user.FullName, &user.Avatar)

	if err != sql.ErrNoRows {
		if err != nil {
			fmt.Println(err)
			return models.UserOutside{}, models.ServeError{Codes: []string{"500"}}
		}
		return user, nil
	}

	//TODO error
	return user, models.ServeError{Codes: []string{"101"}}
}

func (s *storage) GetUserAccounts(userInput models.UserInput) (models.UserOutside, error) {
	user, err := s.GetUserProfile(userInput)
	if err != nil {
		return models.UserOutside{}, err
	}

	rows, err := s.db.Query("SELECT networkName, link FROM social_links WHERE userID = $1", userInput.ID)
	if err != nil {
		return models.UserOutside{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var networkName, link string

		err = rows.Scan(&networkName, &link)
		if err != nil {
			return models.UserOutside{}, err
		}

		//FIXME поиграться с рефлектами
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

	//TODO error
	fmt.Println(err)
	return user, models.ServeError{Codes: []string{"500"}}

}

func (s *storage) ChangeUserProfile(userInput models.UserInputProfile, userAvatar models.UserAvatar) (models.UserOutside, error) {
	errorCodes := s.CheckExistingUserOnUpdate(userInput.Email, userInput.Username, userInput.ID)

	if len(errorCodes) != 0 {
		return models.UserOutside{}, models.ServeError{Codes: errorCodes}
	}

	_, err := s.db.Exec("UPDATE users SET email = $1, username = $2, fullname = $3, avatar = $4 WHERE userID = $5",
		userInput.Email, userInput.Username, userInput.FullName, userAvatar.Avatar, userInput.ID)
	if err != nil {
		//TODO error
		fmt.Println(err)
		return models.UserOutside{}, models.ServeError{Codes: []string{"500"}}
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

	//FIXME поиграться с рефлектами
	input := reflect.ValueOf(userInput)
	for i := 1; i < input.NumField(); i++ {
		inputLink := input.Field(i).Interface().(string)
		if  inputLink == reflect.Indirect(reflect.ValueOf(user.Links)).FieldByName(networkNames[i]).String() {
			continue
		}

		//curNetwork :=
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
		errorCodes := []string{"501"}
		return models.UserOutside{}, models.ServeError{Codes: errorCodes}
	}
	/*if userInput.OldPassword != password {
		errorCodes := []string{"501"}
		return models.UserOutside{}, models.ServeError{Codes: errorCodes}
	}
*/
	salt := make([]byte, 8)
	rand.Read(salt)
	userPassHash := hashPass(salt, userInput.Password)
	_, err = s.db.Exec("UPDATE users SET password = $1 WHERE userID = $2", userPassHash, userInput.ID)
	if err != nil {
		//TODO error

		fmt.Println(err)
		return  models.UserOutside{}, models.ServeError{Codes: []string{"500"}}
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
	//TODO error
	/*if err == sql.ErrNoRows {
		return user, "", models.ServeError{Codes: []string{"101"}}
	}*/

	//могут ли тут быть 2 рзных вида ошибок - обращение к БД и само отсутствие такой записи о пользователе?
	return user, nil, models.ServeError{Codes: []string{"500"}}
}

func (s *storage) GetBoardMembers(userIDs []uint64) ([] models.UserOutsideShort, error) {
	members := make([]models.UserOutsideShort, 0)

	for _, userID := range userIDs {
		user := models.UserOutsideShort{}
		err := s.db.QueryRow("SELECT email, username, fullname, avatar FROM users WHERE userID = $1", userID).
			Scan(&user.Email, &user.Username, &user.FullName, &user.Avatar)

		if err != nil {
			fmt.Println(err)
			return []models.UserOutsideShort{}, models.ServeError{Codes: []string{"500"}}
		}

		members = append(members, user)
		}

	return members, nil
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
