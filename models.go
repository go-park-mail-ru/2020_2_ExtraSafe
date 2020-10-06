package main

import (
	"github.com/labstack/echo"
)

type Handlers struct {
	echo.Context
	users    *[]User
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

type UserInputProfile struct {
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
	FullName string `json:"fullname"`
}

type UserInputPassword struct {
	OldPassword string `json:"oldpassword"`
	Password    string `json:"password"`
}

type User struct {
	ID       uint64 `json:"id"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	FullName string `json:"fullname"`
	Links    *UserLinks
}

type responseUser struct {
	Status   int    `json:"status"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
	FullName string `json:"fullname"`
}

type responseUserLinks struct {
	Status    int    `json:"status"`
	Nickname  string `json:"nickname"`
	Telegram  string `json:"telegram"`
	Instagram string `json:"instagram"`
	Github    string `json:"github"`
	Bitbucket string `json:"bitbucket"`
	Vk        string `json:"vk"`
	Facebook  string `json:"facebook"`
}

type responseError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type UserLinks struct {
	Telegram  string `json:"telegram"`
	Instagram string `json:"instagram"`
	Github    string `json:"github"`
	Bitbucket string `json:"bitbucket"`
	Vk        string `json:"vk"`
	Facebook  string `json:"facebook"`
}

var (
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)
