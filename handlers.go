package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"os"
	"strconv"
	"time"
)

func router(e *echo.Echo) {
	e.Any("/", root)
	e.POST("/login/", login)
	e.GET("/logout/", logout)
	e.POST("/reg/", registration)
	e.GET("/profile/", profile)
	e.GET("/accounts/", accounts)
	e.GET("/avatar/", avatar)
	e.POST("/profile/", profileChange)
	e.POST("/accounts/", accountsChange)
	e.POST("/password/", passwordChange)
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

func login(c echo.Context) error {
	cc := c.(*Handlers)

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

func logout(c echo.Context) error {
	cc := c.(*Handlers)

	session, err := c.Cookie("tabutask_id")
	if err != nil {
		return c.JSON(http.StatusUnauthorized, responseUser{})
	}
	sessionID := session.Value

	delete(*cc.sessions, sessionID)
	session.Expires = time.Now().AddDate(0, 0, -1)
	c.SetCookie(session)
	return c.JSON(http.StatusOK, responseUser{})
}

func registration(c echo.Context) error {
	cc := c.(*Handlers)

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

func avatar(c echo.Context) error {
	cc := c.(*Handlers)

	session, _ := c.Cookie("tabutask_id")
	sessionID := session.Value

	userID := (*cc.sessions)[sessionID]
	filename := "./avatars/" + strconv.FormatUint(userID, 10) + ".png"

	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return c.File("./avatars/default_avatar.png")
	}
	return c.File(filename)
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
