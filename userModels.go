package main

import (
	"errors"
	"fmt"
	"github.com/labstack/echo"
	"io"
	"mime/multipart"
	"net/url"
	"os"
	"strconv"
)

type Handlers struct {
	echo.Context
	users    *[]User
	sessions *map[string]uint64 //map[sessionID]userID
}

type User struct {
	ID       uint64 `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"fullname"`
	Links    *UserLinks
}

type UserLinks struct {
	Telegram  string `json:"telegram"`
	Instagram string `json:"instagram"`
	Github    string `json:"github"`
	Bitbucket string `json:"bitbucket"`
	Vk        string `json:"vkontakte"`
	Facebook  string `json:"facebook"`
}

type UserInputLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserInputReg struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserInputProfile struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	FullName string `json:"fullname"`
}

type UserInputPassword struct {
	OldPassword string `json:"oldpassword"`
	Password    string `json:"password"`
}

func (h *Handlers) checkUserAuthorized(c echo.Context) (responseUser, error) {
	session, err := c.Cookie("tabutask_id")
	if err != nil {
		fmt.Println(err)
		return responseUser{}, err
	}
	sessionID := session.Value
	userID, authorized := (*h.sessions)[sessionID]

	if authorized {
		for _, user := range *h.users {
			if user.ID == userID {
				response := new(responseUser)
				response.WriteResponse(user)
				return *response, nil
			}
		}
	}
	return responseUser{}, errors.New("No such session ")
}

func (h *Handlers) checkUser(userInput UserInputLogin) (responseUser, uint64, error) {
	response := new(responseUser)
	for _, user := range *h.users {
		if userInput.Email == user.Email && userInput.Password == user.Password {
			response.WriteResponse(user)
			return *response, user.ID, nil
		}
	}

	errorMessage := []Messages{{Message: "Неверная электронная почта или пароль", ErrorName: "password"}}
	return responseUser{}, 0, responseError{Messages: errorMessage, Status: 500}
}

func (h *Handlers) createUser(userInput UserInputReg) (responseUser, uint64, error) {
	errorMessage := make([]Messages, 0)
	for _, user := range *h.users {
		if userInput.Email == user.Email {
			msg := Messages{
				Message: "Такой адрес электронной почты уже зарегистрирован", ErrorName: "email",
			}
			errorMessage = append(errorMessage, msg)
		}

		if userInput.Username == user.Username {
			msg := Messages{
				Message: "Такое имя пользователя уже существует", ErrorName: "username",
			}
			errorMessage = append(errorMessage, msg)
		}
	}

	if len(errorMessage) != 0 {
		return responseUser{}, 0, responseError{Messages: errorMessage, Status: 500}
	}

	var id uint64 = 0
	if len(*h.users) > 0 {
		id = (*h.users)[len(*h.users)-1].ID + 1
	}

	newUser := User{
		ID:       id,
		Username: userInput.Username,
		Email:    userInput.Email,
		Password: userInput.Password,
		Links:    &UserLinks{},
	}

	*h.users = append(*h.users, newUser)

	response := new(responseUser)
	response.WriteResponse(newUser)

	return *response, id, nil
}

func (h *Handlers) changeUserProfile(userInput *UserInputProfile, userExist *User) (responseUser, error) {
	errorMessage := make([]Messages, 0)
	for _, user := range *h.users {
		if (userInput.Email == user.Email) && (user.ID != userExist.ID) {
			msg := Messages{
				Message: "Такой адрес электронной почты уже зарегистрирован", ErrorName: "email",
			}
			errorMessage = append(errorMessage, msg)
		}

		if (userInput.Username == user.Username) && (user.ID != userExist.ID) {
			msg := Messages{
				Message: "Такое имя пользователя уже существует", ErrorName: "username",
			}
			errorMessage = append(errorMessage, msg)
		}
	}

	if len(errorMessage) != 0 {
		return responseUser{}, responseError{Messages: errorMessage, Status: 500}
	}

	response := new(responseUser)

	userExist.Username = userInput.Username
	userExist.Email = userInput.Email
	userExist.FullName = userInput.FullName

	response.WriteResponse(*userExist)
	return *response, nil
}

func (h *Handlers) changeUserAccounts(userInput *UserLinks, userExist *User) (responseUserLinks, error) {
	userExist.Links.Bitbucket = userInput.Bitbucket
	userExist.Links.Github = userInput.Github
	userExist.Links.Instagram = userInput.Instagram
	userExist.Links.Telegram = userInput.Telegram
	userExist.Links.Facebook = userInput.Facebook
	userExist.Links.Vk = userInput.Vk

	response := new(responseUserLinks)
	response.WriteResponse(userExist.Username, *userExist.Links)

	return *response, nil
}

func (h *Handlers) changeUserPassword(userInput *UserInputPassword, userExist *User) (responseUser, error) {
	if userInput.OldPassword != userExist.Password {
		errorMessage := []Messages{{Message: "Неверный пароль", ErrorName: "oldPassword"}}
		return responseUser{}, responseError{Messages: errorMessage, Status: 500}
	}

	userExist.Password = userInput.Password

	response := new(responseUser)
	response.WriteResponse(*userExist)

	return *response, nil
}

func getFormParams(params url.Values) (userInput *UserInputProfile) {
	userInput = new(UserInputProfile)
	userInput.Username = params.Get("username")
	userInput.Email = params.Get("email")
	userInput.FullName = params.Get("fullName")

	return
}

func uploadAvatar(file *multipart.FileHeader, userID uint64) error {
	src, err := file.Open()
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer src.Close()

	filename := strconv.FormatUint(userID, 10)
	dst, err := os.Create("./avatars/" + filename + ".png")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		fmt.Println(err)
		return err
	}

	return err
}
