package models

type Comment struct {
	CommentID int64
	TaskID int64
	Message string
	Order int64
	UserID int64
}

type CommentInput struct {
	CommentID int64
	TaskID int64
	Message string
	Order int64
	UserID int64
}

type CommentOutside struct {
	CommentID int64
	Message string
	Order int64
	User UserOutsideShort
}