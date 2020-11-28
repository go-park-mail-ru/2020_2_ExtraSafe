package models

import "encoding/json"

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
	Token    string              `json:"token"`
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

type ResponseBoardShort struct {
	Status  int                `json:"status"`
	BoardID int64             `json:"boardID"`
	Name    string             `json:"boardName"`  // название доски
	Theme   string             `json:"theme"`
	Star    bool               `json:"star"`
}

type ResponseBoard struct {
	Status  int                `json:"status"`
	BoardID int64              `json:"boardID"`
	Admin   UserOutsideShort   `json:"admin"`     // структура владельца доски
	Name    string             `json:"boardName"` // название доски
	Theme   string             `json:"theme"`
	Star    bool               `json:"star"`
	Users   []UserOutsideShort `json:"users"` // массив с пользователями этой доски
	Cards   []CardOutside      `json:"cards"`
	Tags    []TagOutside       `json:"tags"`
}

type ResponseBoards struct {
	Status int                 `json:"status"`
	Boards []BoardOutsideShort `json:"boards"`
}

type ResponseCard struct {
	Status int           `json:"status"`
	CardID int64         `json:"cardID"`
	Name   string        `json:"cardName"`
	Order  int64         `json:"order"`
	Tasks  []TaskOutsideShort `json:"tasks"`
}

type ResponseTask struct {
	Status      int                `json:"status"`
	TaskID      int64              `json:"taskID"`
	Name        string             `json:"taskName"`
	Description string             `json:"description"`
	Order       int64              `json:"order"`
	Tags        []TagOutside       `json:"tags"`
	Users       []UserOutsideShort `json:"users"`
	Checklists  []ChecklistOutside `json:"checklists"`
	Comments    []CommentOutside   `json:"comments"`
}

type ResponseTag struct {
	TagID   int64  `json:"tagID"`
	Color   string `json:"color"`
	TagName string `json:"TagName"`
}

type ResponseComment struct {
	CommentID int64            `json:"commentID"`
	Message   string           `json:"message"`
	Order     int64            `json:"order"`
	User      UserOutsideShort `json:"user"`
}

type ResponseChecklist struct {
	ChecklistID int64           `json:"checklistID"`
	Name        string          `json:"checklistName"`
	Items       json.RawMessage `json:"items"`
}