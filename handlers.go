package main

import (
	"errors"
	"fmt"
	"github.com/labstack/echo"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strconv"
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
	e.GET("/avatars/", getAvatar)
}

func getAvatar(c echo.Context) error {
	cc := c.(*Handlers)

	session, _ := c.Cookie("tabutask_id")
	sessionID := session.Value

	userID := (*cc.sessions)[sessionID]
	filename := "./avatars/" + strconv.FormatUint(userID, 10)

	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return c.File("./avatars/default_avatar.png")
	}
	return c.File(filename)
}

func getFormParams(params url.Values) (userInput *UserInputProfile) {
	userInput = new(UserInputProfile)
	userInput.Nickname = params.Get("username")
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
	dst, err := os.Create("./avatars/" + filename)
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
	for i, user := range *cc.users {
		if user.ID == userID {
			response, err = cc.changeUserPassword(userInput, &(*cc.users)[i])
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
	for i, user := range *cc.users {
		if user.ID == userID {
			response, _ = cc.changeUserAccounts(userInput, &(*cc.users)[i])
		}
	}

	return c.JSON(http.StatusOK, response)
}

func profileChange(c echo.Context) error {
	cc := c.(*Handlers)

	session, _ := c.Cookie("tabutask_id")
	sessionID := session.Value

	userID := (*cc.sessions)[sessionID]

	formParams, err := c.FormParams()
	if err != nil {
		return err
	}

	file, err := c.FormFile("avatar")
	if err != nil {
		fmt.Println(err)
	} else {
		err = uploadAvatar(file, userID)
		if err != nil {
			fmt.Println(err)
		}
	}

	userInput := getFormParams(formParams)

	var response responseUser
	for i, user := range *cc.users {
		if user.ID == userID {
			response, _ = cc.changeUserProfile(userInput, &(*cc.users)[i])
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

	response := new(responseUser)
	response.WriteResponse(newUser)

	return *response, id, nil
}

func (h *Handlers) changeUserProfile(userInput *UserInputProfile, userExist *User) (responseUser, error) {
	response := new(responseUser)
	for _, user := range *h.users {
		if (userInput.Email == user.Email || userInput.Nickname == user.Nickname) && (user.ID != userExist.ID) {
			fmt.Println("Email or nickname already exist ")
			response.WriteResponse(*userExist)
			return *response, errors.New("Email already exist ")
		}
	}

	userExist.Nickname = userInput.Nickname
	userExist.Email = userInput.Email
	userExist.FullName = userInput.FullName

	response.WriteResponse(*userExist)
	fmt.Println(h.users)
	return *response, nil
}

func (h *Handlers) changeUserAccounts(userInput *UserLinks, userExist *User) (responseUserLinks, error) {
	userExist.Links.Bitbucket = userInput.Bitbucket
	userExist.Links.Github = userInput.Github
	userExist.Links.Instagram = userInput.Instagram
	userExist.Links.Telegram = userInput.Telegram
	userExist.Links.Facebook = userInput.Facebook

	response := new(responseUserLinks)
	response.WriteResponse(userExist.Nickname, *userExist.Links)

	return *response, nil
}

func (h *Handlers) changeUserPassword(userInput *UserInputPassword, userExist *User) (responseUser, error) {
	if userInput.OldPassword != userExist.Password {
		return responseUser{}, errors.New("Invalid password ")
	}

	userExist.Password = userInput.Password

	response := new(responseUser)
	response.WriteResponse(*userExist)

	return *response, nil
}

func (response *responseError) WriteResponse(message string) {
	response.Status = 500
	response.Message = message
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
