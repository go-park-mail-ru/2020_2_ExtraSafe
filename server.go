package main

import (
	"errors"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"sync"
)

type Handlers struct {
	echo.Context
	users *[]User
	mu    *sync.Mutex
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

func (h *Handlers) createUser(userInput UserInputReg) (responseUser, error) {
	for _, user := range *h.users {
		if userInput.Email == user.Email {
			fmt.Println("Email already exist ")
			return responseUser{}, errors.New("Email already exist ")
		}
	}

	for _, user := range *h.users {
		if userInput.Nickname == user.Nickname {
			fmt.Println("Nickname already exist ")
			return responseUser{}, errors.New("Nickname already exist ")
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

	return *response, nil
}

func (response *responseUser) WriteResponse(user User) {
	response.Status = 200
	response.Email = user.Email
	response.Nickname = user.Nickname
}

func (h *Handlers) checkUser(userInput UserInputLogin) (responseUser, error) {
	response := new(responseUser)
	for _, user := range *h.users {
		if userInput.Email == user.Email && userInput.Password == user.Password {
			response.WriteResponse(user)
			return *response, nil
		}
	}
	return responseUser{}, errors.New("No such user ")
}

func login(c echo.Context) error {
	cc := c.(*Handlers)
	userInput := new(UserInputLogin)
	if err := c.Bind(userInput); err != nil {
		return err
	}

	response, err := cc.checkUser(*userInput)
	if err != nil {
		return c.JSON(http.StatusForbidden, response)
	}

	return c.JSON(http.StatusOK, response)
}

func registration(c echo.Context) error {
	cc := c.(*Handlers)
	userInput := new(UserInputReg)
	if err := c.Bind(userInput); err != nil {
		return err
	}

	response, err := cc.createUser(*userInput)
	if err != nil {
		return c.JSON(http.StatusForbidden, response)
	}

	return c.JSON(http.StatusOK, response)
}

func urls(e *echo.Echo) {
	e.POST("/login/", login)
	e.POST("/reg/", registration)
}

func main() {
	someUsers := make([]User, 0)

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
			}
			return next(cc)
		}
	})

	urls(e)

	e.Logger.Fatal(e.Start(":8080"))
}
