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
	Telegram  string `json:"telegram" valid:"alphanum"`
	Instagram string `json:"instagram" valid:"alphanum"`
	Github    string `json:"github" valid:"alphanum"`
	Bitbucket string `json:"bitbucket" valid:"alphanum"`
	Vk        string `json:"vkontakte" valid:"alphanum"`
	Facebook  string `json:"facebook" valid:"alphanum"`
}

//TODO links validation
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
	Email    string `json:"email" valid:"email"`
	Password string `json:"password" valid:"passwordValid"`
}

type UserInputReg struct {
	Email    string `json:"email" valid:"email"`
	Username string `json:"username" valid:"userNameValid"`
	Password string `json:"password" valid:"passwordValid"`
}

type UserInputProfile struct {
	ID       uint64 `json:"-"`
	Email    string `json:"email" valid:"email"`
	Username string `json:"username" valid:"userNameValid"`
	FullName string `json:"fullName" valid:"fullNameValid"`
	Avatar   *multipart.FileHeader `json:"-"`
}

type UserInputPassword struct {
	ID          uint64 `json:"-"`
	OldPassword string `json:"oldpassword" valid:"passwordValid"`
	Password    string `json:"password" valid:"passwordValid"`
}
