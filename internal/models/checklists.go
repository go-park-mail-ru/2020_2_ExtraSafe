package models

type Checklist struct {
	ChecklistID int64
	TaskID int64
	Name string
	Items []Item
}

type Item struct {
	Description string
	State bool
}

type ChecklistInput struct {
	UserID int64
	ChecklistID int64
	TaskID int64
	Name string
	Items []Item
}

type ChecklistOutside struct {
	ChecklistID int64
	Name string
	Items []Item
}