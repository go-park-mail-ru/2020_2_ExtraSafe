package models

type Task struct {
	TaskID      uint64 `json:"taskID"`
	CardID      uint64 `json:"cardID"`
	Name        string `json:"taskName"`
	Description string `json:"description"`
	Order       uint8  `json:"order"`
}

type TaskOutside struct {
	TaskID      uint64 `json:"taskID"`
	Name        string `json:"taskName"`
	Description string `json:"description"`
	Order       uint8  `json:"order"`
}

type TaskInput struct {
	UserID      uint64 `json:"-"`
	TaskID      uint64 `json:"taskID"`
	CardID      uint64 `json:"cardID"`
	Name        string `json:"taskName"`
	Description string `json:"description"`
	Order       uint8  `json:"order"`
}

type TaskOrder struct {
	TaskID uint64 `json:"taskID"`
	Order  uint8  `json:"order"`
}

type TasksOrder struct {
	CardID uint64      `json:"cardID"`
	Tasks  []TaskOrder `json:"tasks"`
}

type TasksOrderInput struct {
	UserID      uint64 `json:"-"`
	Cards []TasksOrder `json:"cards"`
}