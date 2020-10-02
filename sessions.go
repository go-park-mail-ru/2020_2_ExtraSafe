package main

import (
	"errors"
	"fmt"
	"github.com/labstack/echo"
	"math/rand"
	"net/http"
	"time"
)

func setCookie(c echo.Context, userID uint64) {
	cookie := new(http.Cookie)
	SID := RandStringRunes(32)
	cc := c.(*Handlers)

	(*cc.sessions)[SID] = userID

	cookie.Name = "tabutask_id"
	cookie.Value = SID
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	fmt.Println(cc.sessions)
	c.SetCookie(cookie)
}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
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
