package models

type Column struct {
	ColumnID    uint64 `json:"column_id"`
	BoardID       uint64 `json:"board_id"`
	Name string `json:"name"`
	Order uint8 `json:"order"`
}


