package models


type Task struct {
	TaskID      int64  `json:"taskID"`
	CardID      int64  `json:"cardID"`
	Name        string `json:"taskName"`
	Description string `json:"description"`
	Order       int64  `json:"order"`
}

//===================================================<-Input
type TaskInput struct {
	UserID      int64  `json:"-"`
	TaskID      int64  `json:"taskID"`
	CardID      int64  `json:"cardID"`
	Name        string `json:"taskName"`
	Description string `json:"description"`
	Order       int64  `json:"order"`
}

type TaskOrder struct {
	TaskID int64 `json:"taskID"`
	Order  int64 `json:"order"`
}

type TasksOrder struct {
	CardID int64       `json:"cardID"`
	Tasks  []TaskOrder `json:"tasks"`
}

type TasksOrderInput struct {
	UserID int64        `json:"-"`
	Tasks  []TasksOrder `json:"cards"`
}

type TaskTagInput struct {
	UserID int64 `json:"-"`
	TaskID int64 `json:"taskID"`
	TagID  int64 `json:"tagID"`
}

type TaskAssignerInput struct {
	UserID       int64  `json:"-"`
	TaskID       int64  `json:"taskID"`
	AssignerName string `json:"assignerName"`
}

//===================================================<-Internal
type TaskInternalShort struct {
	TaskID      int64
	Name        string
	Description string
	Order       int64
	Tags        []TagOutside
	Users       []int64
	Checklists  []ChecklistOutside
}

type TaskInternal struct {
	TaskID      int64
	Name        string
	Description string
	Order       int64
	Tags        []TagOutside
	Users       []int64
	Checklists  []ChecklistOutside
	Comments    []CommentOutside
	Attachments []AttachmentOutside
}

//===================================================<-Outside
type TaskOutside struct {
	TaskID      int64               `json:"taskID"`
	Name        string              `json:"taskName"`
	Description string              `json:"description"`
	Order       int64               `json:"order"`
	Tags        []TagOutside        `json:"tags"`
	Users       []UserOutsideShort  `json:"users"`
	Checklists  []ChecklistOutside  `json:"checklists"`
	Comments    []CommentOutside    `json:"comments"`
	Attachments []AttachmentOutside `json:"attachments"`
}

type TaskOutsideShort struct {
	TaskID      int64              `json:"taskID"`
	Name        string             `json:"taskName"`
	Description string             `json:"description"`
	Order       int64              `json:"order"`
	Tags        []TagOutside       `json:"tags"`
	Users       []UserOutsideShort `json:"users"`
	Checklists  []ChecklistOutside `json:"checklists"`
}

type TaskOutsideSuperShort struct {
	TaskID      int64  `json:"taskID"`
	Name        string `json:"taskName"`
	Description string `json:"description"`
}

//===================================================<-Other
type TaskAssigner struct {
	UserID      int64
	TaskID      int64
	AssignerID  int64
}