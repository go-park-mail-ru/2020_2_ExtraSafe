package models

import "mime/multipart"

// для работы в бд
type User struct {
	ID       uint64 `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"fullName"`
	Links    *UserLinks
	Avatar   string `json:"avatar"`
	Boards   []Board
}

// для формирования ответа пользователю
type UserOutside struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	FullName string `json:"fullName"`
	Links    *UserLinks
	Avatar   string `json:"avatar"`
	Boards   []Board `json:"boards"`
}

type UserLinks struct {
	Telegram  string `json:"telegram"`
	Instagram string `json:"instagram"`
	Github    string `json:"github"`
	Bitbucket string `json:"bitbucket"`
	Vk        string `json:"vkontakte"`
	Facebook  string `json:"facebook"`
}

type UserAvatar struct {
	ID uint64
	Avatar string
}

type UserInputLinks struct {
	ID        uint64 `json:"-"`
	Telegram  string `json:"telegram" valid:"telegramValid~611"`
	Instagram string `json:"instagram" valid:"instagramValid~612"`
	Github    string `json:"github" valid:"githubValid~613"`
	Bitbucket string `json:"bitbucket" valid:"bitbucketValid~614"`
	Vk        string `json:"vkontakte" valid:"vkValid~615"`
	Facebook  string `json:"facebook" valid:"facebookValid~616"`
}

type UserInput struct {
	ID uint64 `json:"id"`
}

type UserInputLogin struct {
	Email    string `json:"email" valid:"required~100, emailValid~110"`
	Password string `json:"password" valid:"required~100, passwordValid~110"`
}

type UserInputReg struct {
	Email    string `json:"email" valid:"required~211, emailValid~211"`
	Username string `json:"username" valid:"required~212, userNameValid~212"`
	Password string `json:"password" valid:"required~213, passwordValid~213"`
}

type UserInputProfile struct {
	ID       uint64 `json:"-"`
	Email    string `json:"email" valid:"required~311, emailValid~311"`
	Username string `json:"username" valid:"required~312, userNameValid~312"`
	FullName string `json:"fullName" valid:"fullNameValid~314"`
	Avatar   *multipart.FileHeader `json:"-"`
}

type UserInputPassword struct {
	ID          uint64 `json:"-"`
	OldPassword string `json:"oldpassword" valid:"required~511, passwordValid~511"`
	Password    string `json:"password" valid:"required~512, passwordValid~512"`
}
