package models

type Task struct {
	TaskID      uint64 `json:"task_id"`
	ColumnID    uint64 `json:"column_id"`
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
	TaskID      uint64 `json:"task_id"`
	ColumnID    uint64 `json:"column_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Order       uint8  `json:"order"`
}
