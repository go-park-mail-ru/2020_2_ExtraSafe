package userStorage

import (
	"../../../internal/errorWorker"
	"../../../internal/models"
	"fmt"
	"github.com/labstack/echo"
)

type Storage interface {
	checkUserAuthorized(c echo.Context) (models.User, error)
	CheckUser(userInput models.UserInputLogin) (models.User, error)
	CreateUser(userInput models.UserInputReg) (models.User, error)

	GetUserProfile(userInput models.UserInput) (models.User, error)
	GetUserAccounts(userInput models.UserInput) (models.User, error)
	ChangeUserProfile(userInput models.UserInputProfile) (models.User, error)
	ChangeUserAccounts(userInput models.UserInputLinks) (models.User, error)
	ChangeUserPassword(userInput models.UserInputPassword) (models.User, error)
}

type storage struct {
	Sessions *map[string]uint64
	Users    *[]models.User
}

func NewStorage(someUsers []models.User, sessions map[string]uint64) Storage {
	return &storage{
		Sessions: &sessions,
		Users: &someUsers,
	}
}

func (s *storage) checkUserAuthorized(c echo.Context) (models.User, error) {
	session, err := c.Cookie("tabutask_id")
	if err != nil {
		fmt.Println(err)
		return models.User{}, errorWorker.ResponseError{Status: 500, Codes: nil}
	}
	sessionID := session.Value
	userID, authorized := (*s.Sessions)[sessionID]

	if authorized {
		for _, user := range *s.Users {
			if user.ID == userID {
				return user, nil
			}
		}
	}
	return models.User{}, errorWorker.ResponseError{Status: 500, Codes: nil}
}

func (s *storage) CheckUser(userInput models.UserInputLogin) (models.User, error) {
	for _, user := range *s.Users {
		if userInput.Email == user.Email && userInput.Password == user.Password {
			return user, nil
		}
	}

	errorCodes := []string{"101"}
	return models.User{}, errorWorker.ResponseError{Codes: errorCodes, Status: 500}
}

func (s *storage) CreateUser(userInput models.UserInputReg) (models.User, error) {
	errorCodes := make([]string, 0)
	for _, user := range *s.Users {
		if userInput.Email == user.Email {
			errorCodes = append(errorCodes, "201")
		}

		if userInput.Username == user.Username {
			errorCodes = append(errorCodes, "202")
		}
	}

	if len(errorCodes) != 0 {
		return models.User{}, errorWorker.ResponseError{Codes: errorCodes, Status: 500}
	}

	var id uint64 = 0
	if len(*s.Users) > 0 {
		id = (*s.Users)[len(*s.Users)-1].ID + 1
	}

	newUser := models.User{
		ID:       id,
		Username: userInput.Username,
		Email:    userInput.Email,
		Password: userInput.Password,
		Links:    &models.UserLinks{},
		Avatar:   "default/default_avatar.png",
	}

	*s.Users = append(*s.Users, newUser)

	return newUser, nil
}


func (s *storage) GetUserProfile(userInput models.UserInput) (models.User, error) {
	someUser := new(models.User)
	for i, user := range *s.Users {
		if user.ID == userInput.ID {
			someUser = &(*s.Users)[i]
		}
	}
	return *someUser, nil
}

func (s *storage) GetUserAccounts(userInput models.UserInput) (models.User, error) {
	someUser := new(models.User)
	for i, user := range *s.Users {
		if user.ID == userInput.ID {
			someUser = &(*s.Users)[i]
		}
	}
	return *someUser, nil
}


func (s *storage) ChangeUserProfile(userInput models.UserInputProfile) (models.User, error) {
	errorCodes := make([]string, 0)

	userExist := new(models.User)

	for i, user := range *s.Users {
		if user.ID == userInput.ID {
			userExist = &(*s.Users)[i]
		}
	}

	for _, user := range *s.Users {
		if (userInput.Email == user.Email) && (user.ID != userExist.ID) {
			errorCodes = append(errorCodes, "301")
		}

		if (userInput.Username == user.Username) && (user.ID != userExist.ID) {
			errorCodes = append(errorCodes, "302")
		}
	}

	if len(errorCodes) != 0 {
		return models.User{}, errorWorker.ResponseError{Codes: errorCodes, Status: 500}
	}

	userExist.Username = userInput.Username
	userExist.Email = userInput.Email
	userExist.FullName = userInput.FullName

	return *userExist, nil
}

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
	userExist := new(models.User)

	for i, user := range *s.Users {
		if user.ID == userInput.ID {
			userExist = &(*s.Users)[i]
		}
	}

	if userInput.OldPassword != userExist.Password {
		errorCodes := []string{"501"}
		return models.User{}, errorWorker.ResponseError{Codes: errorCodes, Status: 500}
	}

	userExist.Password = userInput.Password

	return *userExist, nil
}

