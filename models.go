package main

import (
	"github.com/labstack/echo"
	"sync"
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
