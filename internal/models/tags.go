package models

type Tag struct {
	TagID int64
	BoardID int64
	Color string
	Name string
}

// create, update, delete
type TagInput struct {
	UserID int64 `json:"-"`
	TaskID int64
	TagID int64
	BoardID int64
	Color string
	Name string
}

type TagOutside struct {
	TagID int64
	Color string
	Name string
}