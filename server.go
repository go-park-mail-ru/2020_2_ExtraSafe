package main

import (
	"errors"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

type Handlers struct {
	echo.Context
	users    *[]User
	mu       *sync.Mutex
	sessions *map[string]uint64 //map[sessionID]userID
}

type UserInputLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserInputReg struct {
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

type User struct {
	ID       uint64 `json:"id"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
	Password string `json:"-"`
}

type responseUser struct {
	Status   int    `json:"status"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
}

var (
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func (h *Handlers) createUser(userInput UserInputReg) (responseUser, uint64, error) {
	for _, user := range *h.users {
		if userInput.Email == user.Email {
			fmt.Println("Email already exist ")
			return responseUser{}, 0, errors.New("Email already exist ")
		}
	}

	for _, user := range *h.users {
		if userInput.Nickname == user.Nickname {
			fmt.Println("Nickname already exist ")
			return responseUser{}, 0, errors.New("Nickname already exist ")
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
	}
	*h.users = append(*h.users, newUser)
	h.mu.Unlock()

	response := new(responseUser)
	response.WriteResponse(newUser)

	return *response, id, nil
}

func (response *responseUser) WriteResponse(user User) {
	response.Status = 200
	response.Email = user.Email
	response.Nickname = user.Nickname
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

func (h *Handlers) checkUserAuthorized(c echo.Context) (responseUser, error) {
	session, err := c.Cookie("session_id")
	if err != nil {
		fmt.Println(err)
		return responseUser{}, err
	}
	sessionID := session.Value
	fmt.Println("Got cookie:", sessionID)

	userID, authorized := (*h.sessions)[sessionID]
	fmt.Println(userID, authorized)

	if authorized {
		for _, user := range *h.users {
			if user.ID == userID {
				response := new(responseUser)
				response.WriteResponse(user)
				fmt.Println("Cookie exists")
				return *response, nil
			}
		}
	}
	fmt.Println("Cookie not exists")
	return responseUser{}, errors.New("No such session ")
}

func setCookie(c echo.Context, userID uint64) {
	cookie := new(http.Cookie)
	SID := RandStringRunes(32)
	cc := c.(*Handlers)

	(*cc.sessions)[SID] = userID

	fmt.Println("Setting: ", SID)

	cookie.Name = "session_id"
	cookie.Value = SID
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	fmt.Println(cc.sessions)
	c.SetCookie(cookie)
}

func login(c echo.Context) error {
	cc := c.(*Handlers)

	//проверка на авторизованность
	response, err := cc.checkUserAuthorized(c)
	if err == nil {
		fmt.Println("Authorized")
		return c.JSON(http.StatusOK, response)
	}

	userInput := new(UserInputLogin)
	if err := c.Bind(userInput); err != nil {
		return err
	}

	var userID uint64
	response, userID, err = cc.checkUser(*userInput)
	if err != nil {
		fmt.Println("Not authorized")
		return c.JSON(http.StatusUnauthorized, response)
	}

	setCookie(c, userID)
	return c.JSON(http.StatusOK, response)
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

func root(c echo.Context) error {
	cc := c.(*Handlers)

	response, err := cc.checkUserAuthorized(c)
	if err == nil {
		return c.JSON(http.StatusOK, response)
	}
	return c.JSON(http.StatusTeapot, response)
}

func urls(e *echo.Echo) {
	e.Any("/", root)
	e.POST("/login/", login)
	e.POST("/reg/", registration)
}

func main() {
	someUsers := make([]User, 0)
	sessions := make(map[string]uint64, 10)

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://127.0.0.1:3033"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
	}))

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &Handlers{c,
				&someUsers,
				&sync.Mutex{},
				&sessions,
			}
			return next(cc)
		}
	})

	urls(e)

	e.Logger.Fatal(e.Start(":8080"))
}
