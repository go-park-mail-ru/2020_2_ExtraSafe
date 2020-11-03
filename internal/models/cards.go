package models

type Card struct {
	ColumnID uint64 `json:"columnID"`
	BoardID  uint64 `json:"boardID"`
	Name     string `json:"name"`
	Order    uint8  `json:"order"`
}

type CardOutside struct {
	ColumnID uint64 `json:"columnID"`
	Name     string `json:"name"`
	Order    uint8  `json:"order"`
	Tasks    []Task `json:"tasks"`
}

type CardInput struct {
	ColumnID uint64 `json:"columnID"`
	BoardID  uint64 `json:"boardID"`
	Name     string `json:"name"`
}


