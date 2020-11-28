package models

import "encoding/json"

type Checklist struct {
	ChecklistID int64
	TaskID      int64
	Name        string
	Items       json.RawMessage
	//Items []Item
}

/*type Item struct {
	Description string `json:"description"`
	State bool `json:"state"`
}
*/
type ChecklistInput struct {
	UserID      int64           `json:"-"`
	ChecklistID int64           `json:"checklistID"`
	TaskID      int64           `json:"taskID"`
	Name        string          `json:"checklistName"`
	Items       json.RawMessage `json:"items"`
}

type ChecklistOutside struct {
	ChecklistID int64
	Name        string
	Items       json.RawMessage
}
