package models

type ResponseStatus struct {
	Status   int    `json:"status"`
}

type ResponseToken struct {
	Status int    `json:"status"`
	Token  string `json:"token"`
}

type ResponseUser struct {
	Status   int    `json:"status"`
	Email    string `json:"email"`
	Username string `json:"username"`
	FullName string `json:"fullName"`
	Avatar   string `json:"avatar"`
}

type ResponseUserAuth struct {
	Status   int                 `json:"status"`
	Email    string              `json:"email"`
	Username string              `json:"username"`
	FullName string              `json:"fullName"`
	Avatar   string              `json:"avatar"`
	Links    UserLinks           `json:"links"`
	Boards   []BoardOutsideShort `json:"boards"`
}

type ResponseUserLinks struct {
	Status    int    `json:"status"`
	Username  string `json:"username"`
	Telegram  string `json:"telegram"`
	Instagram string `json:"instagram"`
	Github    string `json:"github"`
	Bitbucket string `json:"bitbucket"`
	Vk        string `json:"vkontakte"`
	Facebook  string `json:"facebook"`
	Avatar    string `json:"avatar"`
}

type ResponseBoard struct {
	Status  int                `json:"status"`
	BoardID uint64             `json:"boardID"`
	Admin   UserOutsideShort   `json:"admin"` // структура владельца доски
	Name    string             `json:"name"`  // название доски
	Theme   string             `json:"theme"`
	Star    bool               `json:"star"`
	Users   []UserOutsideShort `json:"users"` // массив с пользователями этой доски
	Cards   []CardOutside      `json:"cards"`
}

type ResponseBoards struct {
	Status int                 `json:"status"`
	Boards []BoardOutsideShort `json:"boards"`
}

type ResponseCard struct {
	Status  int          `json:"status"`
	CardID uint64        `json:"cardID"`
	Name   string        `json:"name"`
	Order  uint8         `json:"order"`
	Tasks  []TaskOutside `json:"tasks"`
}

type ResponseTask struct {
	Status int `json:"status"`
	TaskID      uint64 `json:"taskID"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Order       uint8  `json:"order"`
}