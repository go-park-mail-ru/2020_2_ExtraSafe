package src

import (
	"github.com/labstack/echo"
	"math/rand"
	"net/http"
	"time"
)

var (
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func setCookie(c echo.Context, userID uint64) {
	cookie := new(http.Cookie)
	SID := RandStringRunes(32)
	cc := c.(*Handlers)

	(*cc.Sessions)[SID] = userID

	cookie.Name = "tabutask_id"
	cookie.Value = SID
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.HttpOnly = true
	c.SetCookie(cookie)
}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
