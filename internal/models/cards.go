package models

type Card struct {
	CardID  int64 `json:"columnID"`
	BoardID int64 `json:"boardID"`
	Name    string `json:"cardName"`
	Order   int64  `json:"order"`
}

type CardOutside struct {
	CardID int64               `json:"cardID"`
	Name   string              `json:"cardName"`
	Order  int64               `json:"order"`
	Tasks  []TaskInternalShort `json:"tasks"`
}

type CardInput struct {
	UserID  int64 `json:"-"`
	CardID  int64 `json:"cardID"`
	BoardID int64 `json:"boardID"`
	Name    string `json:"cardName"`
	Order   int64  `json:"order"`
}

type CardOrder struct {
	CardID int64 `json:"cardID"`
	Order  int64  `json:"order"`
}

type CardsOrderInput struct {
	UserID int64      `json:"-"`
	Cards  []CardOrder `json:"cards"`
}
