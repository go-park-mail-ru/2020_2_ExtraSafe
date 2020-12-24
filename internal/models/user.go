package models

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"fullName"`
	Avatar   string `json:"avatar"`
	Boards   []BoardInternal
}

//===================================================<-Input
type UserInput struct {
	SessionID string `json:"-"`
	ID        int64  `json:"id"`
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
	ID       int64  `json:"-"`
	Email    string `json:"email" valid:"required~311, emailValid~311"`
	Username string `json:"username" valid:"required~312, userNameValid~312"`
	FullName string `json:"fullName" valid:"fullNameValid~314"`
	Avatar   []byte `json:"-"`
}

type UserInputPassword struct {
	ID          int64  `json:"-"`
	OldPassword string `json:"oldpassword" valid:"required~511, passwordValid~511"`
	Password    string `json:"password" valid:"required~512, passwordValid~512"`
}

//===================================================<-Internal
type UserInternal struct {
	ID       int64  `json:"-"`
	Email    string `json:"email"`
	Username string `json:"username"`
	FullName string `json:"fullName"`
	Avatar   string `json:"avatar"`
}

//===================================================<-Outside
type UserOutside struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	FullName string `json:"fullName"`
	Avatar   string `json:"avatar"`
}

type UserOutsideShort struct {
	ID       int64  `json:"-"`
	Email    string `json:"email"`
	Username string `json:"username"`
	FullName string `json:"fullName"`
	Avatar   string `json:"avatar"`
}

type UserBoardsOutside struct {
	Email    string              `json:"email"`
	Username string              `json:"username"`
	FullName string              `json:"fullName"`
	Avatar   string              `json:"avatar"`
	Boards   []BoardOutsideShort `json:"boards"`
}

//===================================================<-Other
type UserAvatar struct {
	ID     int64
	Avatar string
}

type UserSession struct {
	SessionID string
	UserID    int64
}
