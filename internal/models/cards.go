package models

type Card struct {
	CardID  uint64 `json:"columnID"`
	BoardID uint64 `json:"boardID"`
	Name    string `json:"name"`
	Order   uint8  `json:"order"`
}

type CardOutside struct {
	CardID uint64        `json:"cardID"`
	Name   string        `json:"name"`
	Order  uint8         `json:"order"`
	Tasks  []TaskOutside `json:"tasks"`
}

type CardInput struct {
	CardID  uint64 `json:"columnID"`
	BoardID uint64 `json:"boardID"`
	Name    string `json:"name"`
	Order   uint8  `json:"order"`
}


