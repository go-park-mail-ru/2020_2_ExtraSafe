package models

import "encoding/json"

type ResponseStatus struct {
	Status   int    `json:"status"`
}

type ResponseToken struct {
	Status int    `json:"status"`
	Token  string `json:"token"`
}

//===================================================<-User
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
	Boards   []BoardOutsideShort `json:"boards"`
}

//===================================================<-Board
type ResponseBoardShort struct {
	Status  int                `json:"status"`
	BoardID int64             `json:"boardID"`
	Name    string             `json:"boardName"`
	Theme   string             `json:"theme"`
	Star    bool               `json:"star"`
}

type ResponseBoard struct {
	Status  int                `json:"status"`
	BoardID int64              `json:"boardID"`
	Admin   UserOutsideShort   `json:"boardAdmin"`
	Name    string             `json:"boardName"`
	Theme   string             `json:"boardTheme"`
	Star    bool               `json:"boardStar"`
	Users   []UserOutsideShort `json:"boardMembers"`
	Cards   []CardOutside      `json:"boardCards"`
	Tags    []TagOutside       `json:"boardTags"`
}

type ResponseBoards struct {
	Status int                 `json:"status"`
	Boards []BoardOutsideShort `json:"boards"`
}

type ResponseURL struct {
	Status int    `json:"status"`
	URL    string `json:"sharedURL"`
}

//===================================================<-Card
type ResponseCardShort struct {
	Status int           `json:"status"`
	CardID int64         `json:"cardID"`
	Name   string        `json:"cardName"`
}

type ResponseCard struct {
	Status int                `json:"status"`
	CardID int64              `json:"cardID"`
	Name   string             `json:"cardName"`
	Order  int64              `json:"cardOrder"`
	Tasks  []TaskOutsideShort `json:"cardTasks"`
}

//===================================================<-Task
type ResponseTaskSuperShort struct {
	Status      int                 `json:"status"`
	TaskID      int64               `json:"taskID"`
	Name        string              `json:"taskName"`
	Description string              `json:"taskDescription"`
}

type ResponseTask struct {
	Status      int                 `json:"status"`
	TaskID      int64               `json:"taskID"`
	Name        string              `json:"taskName"`
	Description string              `json:"taskDescription"`
	Order       int64               `json:"taskOrder"`
	Tags        []TagOutside        `json:"taskTags"`
	Users       []UserOutsideShort  `json:"taskAssigners"`
	Checklists  []ChecklistOutside  `json:"taskChecklists"`
	Comments    []CommentOutside    `json:"taskComments"`
	Attachments []AttachmentOutside `json:"taskAttachments"`
}

//===================================================<-Tag
type ResponseTag struct {
	Status  int    `json:"status"`
	TagID   int64  `json:"tagID"`
	Color   string `json:"tagColor"`
	TagName string `json:"tagName"`
}

//===================================================<-Comment
type ResponseComment struct {
	Status    int              `json:"status"`
	CommentID int64            `json:"commentID"`
	Message   string           `json:"commentMessage"`
	Order     int64            `json:"commentOrder"`
	User      UserOutsideShort `json:"commentAuthor"`
}

//===================================================<-Checklist
type ResponseChecklist struct {
	Status      int             `json:"status"`
	ChecklistID int64           `json:"checklistID"`
	Name        string          `json:"checklistName"`
	Items       json.RawMessage `json:"checklistItems"`
}

//===================================================<-Attachment
type ResponseAttachment struct {
	Status       int    `json:"status"`
	AttachmentID int64  `json:"attachmentID"`
	Filename     string `json:"attachmentFileName"`
	Filepath     string `json:"attachmentFilePath"`
}

//===================================================<-WebSocket
type WS struct {
	SessionID string      `json:"-"`
	Method    string      `json:"method"`
	Body      interface{} `json:"body"`
}

type NotificationMessage struct {
	UserID int64       `json:"-"`
	Body   interface{} `json:"body"`
}