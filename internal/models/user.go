package models

// для формирования ответа пользователю
type UserOutside struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	FullName string `json:"fullName"`
	Links    *UserLinks
	Avatar   string `json:"avatar"`
}

//TODO добавить заполнение ID
type UserOutsideShort struct {
	ID int64
	Email    string `json:"email"`
	Username string `json:"username"`
	FullName string `json:"fullName"`
	Avatar   string `json:"avatar"`
}

type UserInternal struct {
	ID        int64 `json:"-"`
	Email    string `json:"email"`
	Username string `json:"username"`
	FullName string `json:"fullName"`
	Avatar   string `json:"avatar"`
}

type UserBoardsOutside struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	FullName string `json:"fullName"`
	Links    *UserLinks
	Avatar   string `json:"avatar"`
	Boards   []BoardOutsideShort `json:"boards"`
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
	ID int64
	Avatar string
}

type UserInputLinks struct {
	ID        int64 `json:"-"`
	Telegram  string `json:"telegram" valid:"telegramValid~611"`
	Instagram string `json:"instagram" valid:"instagramValid~612"`
	Github    string `json:"github" valid:"githubValid~613"`
	Bitbucket string `json:"bitbucket" valid:"bitbucketValid~614"`
	Vk        string `json:"vkontakte" valid:"vkValid~615"`
	Facebook  string `json:"facebook" valid:"facebookValid~616"`
}

type UserInput struct {
	ID int64 `json:"id"`
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
	ID       int64                `json:"-"`
	Email    string                `json:"email" valid:"required~311, emailValid~311"`
	Username string                `json:"username" valid:"required~312, userNameValid~312"`
	FullName string                `json:"fullName" valid:"fullNameValid~314"`
	Avatar   []byte `json:"-"`
}

type UserInputPassword struct {
	ID          int64 `json:"-"`
	OldPassword string `json:"oldpassword" valid:"required~511, passwordValid~511"`
	Password    string `json:"password" valid:"required~512, passwordValid~512"`
}

type UserSession struct {
	SessionID string
	UserID int64
}

// для работы в бд
type User struct {
	ID       int64 `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"fullName"`
	Links    *UserLinks
	Avatar   string `json:"avatar"`
	Boards   []BoardInternal
}
