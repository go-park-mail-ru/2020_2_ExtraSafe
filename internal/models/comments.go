package models

type Comment struct {
	CommentID int64
	TaskID    int64
	Message   string
	Order     int64
	UserID    int64
}

type CommentInput struct {
	CommentID int64  `json:"commentID"`
	TaskID    int64  `json:"taskID"`
	Message   string `json:"message"`
	Order     int64  `json:"order"`
	UserID    int64  `json:"-"`
}

type CommentOutside struct {
	CommentID int64
	Message   string
	Order     int64
	User      UserOutsideShort
}

type CommentInternal struct {
	CommentID int64
	Message   string
	Order     int64
	UserID    int64
}