package models

type Task struct {
	TaskID      int64 `json:"taskID"`
	CardID      int64 `json:"cardID"`
	Name        string `json:"taskName"`
	Description string `json:"description"`
	Order       int64  `json:"order"`
}

type TaskOutside struct {
	TaskID      int64 `json:"taskID"`
	Name        string `json:"taskName"`
	Description string `json:"description"`
	Order       int64  `json:"order"`
}

type TaskInput struct {
	UserID      int64 `json:"-"`
	TaskID      int64 `json:"taskID"`
	CardID      int64 `json:"cardID"`
	Name        string `json:"taskName"`
	Description string `json:"description"`
	Order       int64  `json:"order"`
}

type TaskOrder struct {
	TaskID int64 `json:"taskID"`
	Order  int64  `json:"order"`
}

type TasksOrder struct {
	CardID int64      `json:"cardID"`
	Tasks  []TaskOrder `json:"tasks"`
}

type TasksOrderInput struct {
	UserID      int64 `json:"-"`
	Tasks []TasksOrder `json:"cards"`
}

// add, delete
type TaskTagInput struct {
	UserID      int64 `json:"-"`
	TaskID int64 `json:"taskID"`
	TagID  int64 `json:"tagID"`
}

type TaskAssigner struct {
	UserID    int64 `json:"-"`
	TaskID   int64 `json:"taskID"`
	AssignerID  int64
}