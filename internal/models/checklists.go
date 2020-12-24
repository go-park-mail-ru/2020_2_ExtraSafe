package models

import "encoding/json"

type Checklist struct {
	ChecklistID int64
	TaskID      int64
	Name        string
	Items       json.RawMessage
}

// ===================================================<-Input
type ChecklistInput struct {
	SessionID   string          `json:"-"`
	UserID      int64           `json:"-"`
	BoardID     int64           `json:"-"`
	ChecklistID int64           `json:"checklistID"`
	TaskID      int64           `json:"taskID"`
	Name        string          `json:"checklistName"`
	Items       json.RawMessage `json:"checklistItems"`
}

// ===================================================<-Internal

// ===================================================<-Outside
type ChecklistOutside struct {
	ChecklistID int64           `json:"checklistID,omitempty"`
	TaskID      int64           `json:"taskID,omitempty"`
	CardID      int64           `json:"cardID,omitempty"`
	Name        string          `json:"checklistName,omitempty"`
	Items       json.RawMessage `json:"checklistItems,omitempty"`
}

// ===================================================<-Other
