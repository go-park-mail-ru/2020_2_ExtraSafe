package models

type Comment struct {
	CommentID int64
	TaskID    int64
	Message   string
	Order     int64
	UserID    int64
}

// ===================================================<-Input
type CommentInput struct {
	CommentID int64  `json:"commentID"`
	TaskID    int64  `json:"taskID"`
	Message   string `json:"commentMessage"`
	Order     int64  `json:"commentOrder"`
	BoardID   int64  `json:"-"`
	UserID    int64  `json:"-"`
	SessionID string `json:"-"`
}

// ===================================================<-Internal
type CommentInternal struct {
	CommentID int64
	CardID    int64
	TaskID    int64
	Message   string
	Order     int64
	UserID    int64
}

// ===================================================<-Outside
type CommentOutside struct {
	CommentID int64            `json:"commentID,omitempty"`
	TaskID    int64            `json:"taskID,omitempty"`
	CardID    int64            `json:"cardID,omitempty"`
	Message   string           `json:"commentMessage,omitempty"`
	Order     int64            `json:"commentOrder,omitempty"`
	User      UserOutsideShort `json:"commentAuthor,omitempty"`
}

// ===================================================<-Other
