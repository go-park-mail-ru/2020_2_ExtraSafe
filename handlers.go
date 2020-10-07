package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

func router(e *echo.Echo) {
	e.Any("/", root)
	e.POST("/login/", login)
	e.GET("/logout/", logout)
	e.POST("/reg/", registration)
	e.GET("/profile/", profile)
	e.GET("/accounts/", accounts)
	e.GET("/avatar/:name", avatar)
	e.POST("/profile/", profileChange)
	e.POST("/accounts/", accountsChange)
	e.POST("/password/", passwordChange)
}

func root(c echo.Context) error {
	cc := c.(*Handlers)

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
		return c.JSON(http.StatusUnauthorized, err)
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
		return c.JSON(http.StatusUnauthorized, err)
	}

	setCookie(c, userID)
	return c.JSON(http.StatusOK, response)
}

func profile(c echo.Context) error {
	cc := c.(*Handlers)

	session, _ := c.Cookie("tabutask_id")
	sessionID := session.Value
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
			response.WriteResponse(user.Username, *user.Links, user.Avatar)
		}
	}
	return c.JSON(http.StatusOK, response)
}

func avatar(c echo.Context) error {
	filename := c.Param("name")

	if filename == "default_avatar.png" {
		return c.File("./default/default_avatar.png")
	}
	return c.File("./avatars/" + filename)
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
		err, _ := cc.uploadAvatar(file, userID)
		if err != nil {
			fmt.Println(err)
		}
	}

	userInput := getFormParams(formParams)

	var response responseUser
	for i, user := range *cc.users {
		if user.ID == userID {
			response, err = cc.changeUserProfile(userInput, &(*cc.users)[i])
			break
		}
	}
	if err != nil {
		return c.JSON(http.StatusOK, err)
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
		return c.JSON(http.StatusUnauthorized, err)
	}

	return c.JSON(http.StatusOK, response)
}
