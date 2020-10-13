package sources

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

func Router(e *echo.Echo) {
	e.Any("/", root)
	e.POST("/login/", login)
	e.GET("/logout/", logout)
	e.POST("/reg/", registration)
	e.GET("/profile/", profile)
	e.GET("/accounts/", accounts)
	e.Static("/avatar", "")
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
	response, userID, err = cc.CheckUser(*userInput)
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
		return c.JSON(http.StatusUnauthorized, ResponseUser{})
	}
	sessionID := session.Value

	delete(*cc.Sessions, sessionID)
	session.Expires = time.Now().AddDate(0, 0, -1)
	c.SetCookie(session)
	return c.JSON(http.StatusOK, ResponseUser{})
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
	response, userID, err = cc.CreateUser(*userInput)
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
	userID := (*cc.Sessions)[sessionID]

	response := new(ResponseUser)
	for _, user := range *cc.Users {
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

	userID := (*cc.Sessions)[sessionID]

	response := new(ResponseUserLinks)
	for _, user := range *cc.Users {
		if user.ID == userID {
			response.WriteResponse(user.Username, *user.Links, user.Avatar)
		}
	}
	return c.JSON(http.StatusOK, response)
}

func accountsChange(c echo.Context) error {
	cc := c.(*Handlers)

	session, _ := c.Cookie("tabutask_id")
	sessionID := session.Value

	userID := (*cc.Sessions)[sessionID]

	userInput := new(UserLinks)
	if err := c.Bind(userInput); err != nil {
		return err
	}

	var response ResponseUserLinks
	for i, user := range *cc.Users {
		if user.ID == userID {
			response, _ = cc.ChangeUserAccounts(userInput, &(*cc.Users)[i])
		}
	}

	return c.JSON(http.StatusOK, response)
}

func profileChange(c echo.Context) error {
	cc := c.(*Handlers)

	session, _ := c.Cookie("tabutask_id")
	sessionID := session.Value

	userID := (*cc.Sessions)[sessionID]

	formParams, err := c.FormParams()
	if err != nil {
		return err
	}

	var errAvatar, errProfile error
	errorCodes := make([]string, 0)

	file, err := c.FormFile("avatar")
	if err != nil {
		fmt.Println(err)
	} else {
		errAvatar, _ = cc.UploadAvatar(file, userID)
		if errAvatar != nil {
			errorCodes = append(errorCodes, errAvatar.(ResponseError).Codes...)
			fmt.Println(err)
		}
	}

	userInput := getFormParams(formParams)

	var response ResponseUser
	for i, user := range *cc.Users {
		if user.ID == userID {
			response, errProfile = cc.ChangeUserProfile(userInput, &(*cc.Users)[i])
			if errProfile != nil {
				errorCodes = append(errorCodes, errProfile.(ResponseError).Codes...)
			}
			break
		}
	}
	if len(errorCodes) != 0 {
		return c.JSON(http.StatusOK, ResponseError{Codes: errorCodes, Status: 500})
	}
	/*if err != nil {
		return c.JSON(http.StatusOK, err)
	}*/

	return c.JSON(http.StatusOK, response)
}

func passwordChange(c echo.Context) error {
	cc := c.(*Handlers)

	session, _ := c.Cookie("tabutask_id")
	sessionID := session.Value

	userID := (*cc.Sessions)[sessionID]

	userInput := new(UserInputPassword)
	if err := c.Bind(userInput); err != nil {
		return err
	}

	var response ResponseUser
	var err error
	for i, user := range *cc.Users {
		if user.ID == userID {
			response, err = cc.ChangeUserPassword(userInput, &(*cc.Users)[i])
		}
	}

	if err != nil {
		return c.JSON(http.StatusUnauthorized, err)
	}

	return c.JSON(http.StatusOK, response)
}
