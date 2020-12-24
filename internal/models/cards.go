package models

type Card struct {
	CardID  int64  `json:"columnID"`
	BoardID int64  `json:"boardID"`
	Name    string `json:"cardName"`
	Order   int64  `json:"order"`
}

// ===================================================<-Input
type CardInput struct {
	SessionID string `json:"-"`
	UserID    int64  `json:"-"`
	CardID    int64  `json:"cardID"`
	BoardID   int64  `json:"boardID"`
	Name      string `json:"cardName"`
	Order     int64  `json:"cardOrder"`
}

type CardOrder struct {
	CardID int64 `json:"cardID"`
	Order  int64 `json:"cardOrder"`
}

type CardsOrderInput struct {
	SessionID string      `json:"-"`
	UserID    int64       `json:"-"`
	BoardID   int64       `json:"-"`
	Cards     []CardOrder `json:"cards"`
}

// ===================================================<-Internal
type CardInternal struct {
	CardID int64
	Name   string
	Order  int64
	Tasks  []TaskInternalShort
}

// ===================================================<-Outside
type CardOutside struct {
	CardID int64              `json:"cardID"`
	Name   string             `json:"cardName"`
	Order  int64              `json:"cardOrder"`
	Tasks  []TaskOutsideShort `json:"cardTasks"`
}

type CardOutsideShort struct {
	CardID int64  `json:"cardID"`
	Name   string `json:"cardName"`
}

// ===================================================<-Other
