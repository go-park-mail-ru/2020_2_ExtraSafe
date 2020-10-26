package models

type Board struct {
	BoardID       uint64 `json:"board_id"`
	AdminID    uint64 `json:"admin_id"`
	Name string `json:"name"`
	Theme string `json:"theme"`
}

type BoardMembers struct {
	BoardID       uint64 `json:"board_id"`
	UserID    uint64 `json:"user_id"`
}