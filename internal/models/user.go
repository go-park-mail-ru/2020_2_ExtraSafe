package models

import "mime/multipart"

type User struct {
	ID       uint64 `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"fullName"`
	Links    *UserLinks
	Avatar   string `json:"avatar"`
}

type UserLinks struct {
	Telegram  string `json:"telegram"`
	Instagram string `json:"instagram"`
	Github    string `json:"github"`
	Bitbucket string `json:"bitbucket"`
	Vk        string `json:"vkontakte"`
	Facebook  string `json:"facebook"`
}

type UserInputLinks struct {
	ID        uint64 `json:"-"`
	Telegram  string `json:"telegram"`
	Instagram string `json:"instagram"`
	Github    string `json:"github"`
	Bitbucket string `json:"bitbucket"`
	Vk        string `json:"vkontakte"`
	Facebook  string `json:"facebook"`
}

type UserInput struct {
	ID uint64 `json:"id"`
}

type UserInputLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserInputReg struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserInputProfile struct {
	ID       uint64 `json:"-"`
	Email    string `json:"email"`
	Username string `json:"username"`
	FullName string `json:"fullName"`
	Avatar   *multipart.FileHeader `json:"-"`
}

type UserInputPassword struct {
	ID          uint64 `json:"-"`
	OldPassword string `json:"oldpassword"`
	Password    string `json:"password"`
}
