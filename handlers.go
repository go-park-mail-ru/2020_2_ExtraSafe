package main

import (
	"errors"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func urls(e *echo.Echo) {
	e.Any("/", root)
	e.POST("/login/", login)
	e.POST("/reg/", registration)
	e.GET("/profile/", profile)
	e.GET("/accounts/", accounts)
	e.POST("/profile/", profileChange)
	e.POST("/accounts/", accountsChange)
	e.POST("/password/", passwordChange)
}

func passwordChange(c echo.Context) error {
	cc := c.(*Handlers)

	session, _ := c.Cookie("tabutask_id")
	sessionID := session.Value

	userID := (*cc.sessions)[sessionID]

	userInput := new(UserInputPassword)
	if err := c.Bind(userInput); err != nil {
		return err
	}

	var response responseUser
	var err error
	for _, user := range *cc.users {
		if user.ID == userID {
			response, err = cc.changeUserPassword(userInput, user)
		}
	}

	if err != nil {
		return c.JSON(http.StatusUnauthorized, response)
	}

	return c.JSON(http.StatusOK, response)
}

func accountsChange(c echo.Context) error {
	cc := c.(*Handlers)

	session, _ := c.Cookie("tabutask_id")
	sessionID := session.Value

	userID := (*cc.sessions)[sessionID]

	userInput := new(UserLinks)
	if err := c.Bind(userInput); err != nil {
		return err
	}

	var response responseUserLinks
	for _, user := range *cc.users {
		if user.ID == userID {
			response, _ = cc.changeUserAccounts(userInput, user)
		}
	}

	return c.JSON(http.StatusOK, response)
}

func profileChange(c echo.Context) error {
	cc := c.(*Handlers)

	session, _ := c.Cookie("tabutask_id")
	sessionID := session.Value

	userID := (*cc.sessions)[sessionID]

	userInput := new(UserInputProfile)
	if err := c.Bind(userInput); err != nil {
		return err
	}

	var response responseUser
	for _, user := range *cc.users {
		if user.ID == userID {
			response, _ = cc.changeUserProfile(userInput, user)
		}
	}

	return c.JSON(http.StatusOK, response)
}

func profile(c echo.Context) error {
	cc := c.(*Handlers)

	session, _ := c.Cookie("tabutask_id")
	sessionID := session.Value // не может быть nil, тк мы на руте проверяем авторизованность,
	// а на авторизации/регистрации выдаем куки
	userID := (*cc.sessions)[sessionID]

	response := new(responseUser)
	for _, user := range *cc.users {
		if user.ID == userID {
			response.WriteResponse(user)
		}
	}
	return c.JSON(http.StatusOK, response)
}

func accounts(c echo.Context) error {
	cc := c.(*Handlers)

	session, _ := c.Cookie("tabutask_id")
	sessionID := session.Value

	userID := (*cc.sessions)[sessionID]

	response := new(responseUserLinks)
	for _, user := range *cc.users {
		if user.ID == userID {
			response.WriteResponse(user.Nickname, *user.Links)
		}
	}
	return c.JSON(http.StatusOK, response)
}

func login(c echo.Context) error {
	cc := c.(*Handlers)

	//TODO убрать отсюда проверку куки
	response, err := cc.checkUserAuthorized(c)
	if err == nil {
		return c.JSON(http.StatusOK, response)
	}

	userInput := new(UserInputLogin)
	if err := c.Bind(userInput); err != nil {
		return err
	}

	var userID uint64
	response, userID, err = cc.checkUser(*userInput)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, response)
	}

	setCookie(c, userID)
	return c.JSON(http.StatusOK, response)
}

func registration(c echo.Context) error {
	cc := c.(*Handlers)

	//TODO убрать отсюда проверку куки
	response, err := cc.checkUserAuthorized(c)
	if err == nil {
		return c.JSON(http.StatusOK, response)
	}

	userInput := new(UserInputReg)
	if err := c.Bind(userInput); err != nil {
		return err
	}

	var userID uint64
	response, userID, err = cc.createUser(*userInput)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, response)
	}

	setCookie(c, userID)
	return c.JSON(http.StatusOK, response)
}

func root(c echo.Context) error {
	cc := c.(*Handlers)
	fmt.Println("root")

	response, err := cc.checkUserAuthorized(c)
	if err == nil {
		return c.JSON(http.StatusOK, response)
	}
	return c.JSON(http.StatusUnauthorized, response)
}

func (h *Handlers) createUser(userInput UserInputReg) (responseUser, uint64, error) {
	for _, user := range *h.users {
		if userInput.Email == user.Email || userInput.Nickname == user.Nickname {
			fmt.Println("Email or nickname already exist ")
			return responseUser{}, 0, errors.New("Email already exist ")
		}
	}

	h.mu.Lock()

	var id uint64 = 0
	if len(*h.users) > 0 {
		id = (*h.users)[len(*h.users)-1].ID + 1
	}

	newUser := User{
		ID:       id,
		Nickname: userInput.Nickname,
		Email:    userInput.Email,
		Password: userInput.Password,
		Links:    &UserLinks{},
	}

	*h.users = append(*h.users, newUser)
	h.mu.Unlock()

	response := new(responseUser)
	response.WriteResponse(newUser)

	return *response, id, nil
}

func (h *Handlers) changeUserProfile(userInput *UserInputProfile, userExist User) (responseUser, error) {
	for _, user := range *h.users {
		if userInput.Email == user.Email || userInput.Nickname == user.Nickname {
			fmt.Println("Email or nickname already exist ")
			return responseUser{}, errors.New("Email already exist ")
		}
	}

	h.mu.Lock()

	userExist.Nickname = userInput.Nickname
	userExist.Email = userInput.Email
	userExist.FullName = userInput.FullName

	h.mu.Unlock()

	response := new(responseUser)
	response.WriteResponse(userExist)

	return *response, nil
}

func (h *Handlers) changeUserAccounts(userInput *UserLinks, userExist User) (responseUserLinks, error) {
	h.mu.Lock()

	userExist.Links.Bitbucket = userInput.Bitbucket
	userExist.Links.Github = userInput.Github
	userExist.Links.Instagram = userInput.Instagram
	userExist.Links.Telegram = userInput.Telegram
	userExist.Links.Facebook = userInput.Facebook

	h.mu.Unlock()

	response := new(responseUserLinks)
	response.WriteResponse(userExist.Nickname, *userExist.Links)

	return *response, nil
}

func (h *Handlers) changeUserPassword(userInput *UserInputPassword, userExist User) (responseUser, error) {
	if userInput.OldPassword != userExist.Password {
		return responseUser{}, errors.New("Invalid password ")
	}

	h.mu.Lock()

	userExist.Password = userInput.Password

	h.mu.Unlock()

	response := new(responseUser)
	response.WriteResponse(userExist)

	return *response, nil
}

func (response *responseUser) WriteResponse(user User) {
	response.Status = 200
	response.Email = user.Email
	response.Nickname = user.Nickname
	response.FullName = user.FullName
}

func (response *responseUserLinks) WriteResponse(nickname string, links UserLinks) {
	response.Status = 200
	response.Nickname = nickname
	response.Telegram = links.Telegram
	response.Instagram = links.Instagram
	response.Github = links.Github
	response.Bitbucket = links.Bitbucket
	response.Vk = links.Bitbucket
	response.Facebook = links.Facebook
}

func (h *Handlers) checkUser(userInput UserInputLogin) (responseUser, uint64, error) {
	response := new(responseUser)
	for _, user := range *h.users {
		if userInput.Email == user.Email && userInput.Password == user.Password {
			response.WriteResponse(user)
			return *response, user.ID, nil
		}
	}
	return responseUser{}, 0, errors.New("No such user ")
}
