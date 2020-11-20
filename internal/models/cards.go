package models

type Card struct {
	CardID  uint64 `json:"columnID"`
	BoardID uint64 `json:"boardID"`
	Name    string `json:"cardName"`
	Order   uint8  `json:"order"`
}

type CardOutside struct {
	CardID uint64        `json:"cardID"`
	Name   string        `json:"cardName"`
	Order  uint8         `json:"order"`
	Tasks  []TaskOutside `json:"tasks"`
}

type CardInput struct {
	UserID  uint64 `json:"-"`
	CardID  uint64 `json:"cardID"`
	BoardID uint64 `json:"boardID"`
	Name    string `json:"cardName"`
	Order   uint8  `json:"order"`
}

type CardOrder struct {
	CardID uint64 `json:"cardID"`
	Order  uint8  `json:"order"`
}

type CardsOrderInput struct {
	UserID uint64      `json:"-"`
	Cards  []CardOrder `json:"cards"`
}
