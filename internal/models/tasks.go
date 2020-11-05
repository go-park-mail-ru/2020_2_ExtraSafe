package models

type Task struct {
	TaskID      uint64 `json:"task_id"`
	CardID      uint64 `json:"card_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Order       uint8  `json:"order"`
	//Deadline timestamp
}

type TaskOutside struct {
	TaskID      uint64 `json:"task_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Order       uint8  `json:"order"`
	//Deadline timestamp
}

type TaskInput struct {
	UserID      uint64 `json:"-"`
	TaskID      uint64 `json:"taskID"`
	CardID      uint64 `json:"cardID"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Order       uint8  `json:"order"`
}
