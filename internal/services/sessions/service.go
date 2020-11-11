package sessions

import (
	"errors"
	"fmt"
	"github.com/labstack/echo"
	"math/rand"
	"net/http"
	"time"
)

type Service interface {
	SetCookie(c echo.Context, userID uint64) error
	DeleteCookie(c echo.Context) error
	CheckCookie(c echo.Context) (uint64, error)
}

type service struct {
	sessionsStorage sessionsStorage
}

func NewService(sessionsStorage sessionsStorage) Service {
	return &service{
		sessionsStorage: sessionsStorage,
	}
}

var (
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func (s *service)SetCookie(c echo.Context, userID uint64) error {
	cookie := new(http.Cookie)
	SID := RandStringRunes(32)

	if err := s.sessionsStorage.CreateUserSession(userID, SID); err != nil {
		fmt.Println(err)
		return err
	}

	cookie.Name = "tabutask_id"
	cookie.Value = SID
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.HttpOnly = true
	c.SetCookie(cookie)

	return nil
}

func (s *service)DeleteCookie(c echo.Context) error {
	session, err := c.Cookie("tabutask_id")
	if err != nil {
		return err
	}
	sessionID := session.Value

	if err := s.sessionsStorage.DeleteUserSession(sessionID); err != nil {
		fmt.Println(err)
		return err
	}

	session.Expires = time.Now().AddDate(0, 0, -1)
	c.SetCookie(session)
	return nil
}

func (s *service)CheckCookie(c echo.Context) (uint64, error) {
	session, err := c.Cookie("tabutask_id")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	sessionID := session.Value

	userId, err := s.sessionsStorage.CheckUserSession(sessionID)
	if err != nil {
		return 0, errors.New("Not auth ")

	}
	return userId, nil
}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
