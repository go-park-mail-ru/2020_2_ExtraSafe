package models

import "encoding/json"

type Checklist struct {
	ChecklistID int64
	TaskID int64
	Name string
	Items json.RawMessage
	//Items []Item
}

/*type Item struct {
	Description string `json:"description"`
	State bool `json:"state"`
}
*/
type ChecklistInput struct {
	UserID int64
	ChecklistID int64
	TaskID int64
	Name string
	Items json.RawMessage
}

type ChecklistOutside struct {
	ChecklistID int64
	Name string
	Items json.RawMessage
}
