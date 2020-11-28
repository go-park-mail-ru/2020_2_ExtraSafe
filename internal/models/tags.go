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
	TaskID int64 `json:"taskID"`
	TagID int64 `json:"tagID"`
	BoardID int64 `json:"boardID"`
	Color string `json:"color"`
	Name string `json:"TagName"`
}

type TagOutside struct {
	TagID int64
	Color string
	Name string
}