package models

type Tasks struct {
	TaskID    uint64 `json:"task_id"`
	ColumnID       uint64 `json:"column_id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Order uint8 `json:"order"`
	//Deadline timestamp
}
