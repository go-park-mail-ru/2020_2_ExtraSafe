package models

type Column struct {
	ColumnID uint64 `json:"columnID"`
	BoardID  uint64 `json:"boardID"`
	Name     string `json:"name"`
	Order    uint8  `json:"order"`
}

type ColumnOutside struct {
	ColumnID uint64        `json:"columnID"`
	Name     string        `json:"name"`
	Order    uint8         `json:"order"`
	Tasks    []TaskOutside `json:"tasks"`
}

type ColumnInput struct {
	ColumnID uint64 `json:"columnID"`
	BoardID  uint64 `json:"boardID"`
	Name     string `json:"name"`
}


