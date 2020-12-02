package models


type Comment struct {
	CommentID int64
	TaskID    int64
	Message   string
	Order     int64
	UserID    int64
}

//===================================================<-Input
type CommentInput struct {
	CommentID int64  `json:"commentID"`
	TaskID    int64  `json:"taskID"`
	Message   string `json:"commentMessage"`
	Order     int64  `json:"commentOrder"`
	UserID    int64  `json:"-"`
}

//===================================================<-Internal
type CommentInternal struct {
	CommentID int64
	Message   string
	Order     int64
	UserID    int64
}

//===================================================<-Outside
type CommentOutside struct {
	CommentID int64            `json:"commentID"`
	Message   string           `json:"commentMessage"`
	Order     int64            `json:"commentOrder"`
	User      UserOutsideShort `json:"commentAuthor"`
}

//===================================================<-Other
