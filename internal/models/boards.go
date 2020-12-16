package models


type Board struct {
	BoardID  int64   `json:"boardID"`
	AdminID  int64   `json:"adminID"`
	Name     string  `json:"boardName"`
	Theme    string  `json:"theme"`
	Star     bool    `json:"star"`
	UsersIDs []int64 `json:"usersIDs"`
}

//===================================================<-Input
type BoardInput struct {
	UserID  int64 `json:"-"`
	BoardID int64 `json:"boardID"`
}

type BoardChangeInput struct {
	UserID    int64  `json:"-"`
	BoardID   int64  `json:"boardID"`
	BoardName string `json:"boardName"`
	Theme     string `json:"boardTheme"`
	Star      bool   `json:"boardStar"`
}

type BoardMemberInput struct {
	UserID     int64  `json:"-"`
	BoardID    int64  `json:"boardID"`
	MemberName string `json:"memberUsername"`
}

type BoardInviteInput struct {
	UserID  int64
	BoardID int64
	UrlHash string
}

//===================================================<-Internal
type BoardInternal struct {
	BoardID  int64          `json:"boardID"`
	AdminID  int64          `json:"adminID"`
	Name     string         `json:"name"`
	Theme    string         `json:"theme"`
	Star     bool           `json:"star"`
	Cards    []CardInternal `json:"cards"`
	UsersIDs []int64        `json:"usersIDs"`
	Tags     []TagOutside
}

//===================================================<-Outside
type BoardOutside struct {
	BoardID int64              `json:"boardID"`
	Admin   UserOutsideShort   `json:"boardAdmin"`
	Name    string             `json:"boardName"`
	Theme   string             `json:"boardTheme"`
	Star    bool               `json:"boardStar"`
	Users   []UserOutsideShort `json:"boardMembers"`
	Cards   []CardOutside      `json:"boardCards"`
	Tags    []TagOutside       `json:"boardTags"`
}

type BoardOutsideShort struct {
	BoardID int64  `json:"boardID"`
	Name    string `json:"boardName"`
	Theme   string `json:"boardTheme"`
	Star    bool   `json:"boardStar"`
}

//===================================================<-Other
type BoardMember struct {
	UserID   int64
	BoardID  int64
	MemberID int64
}
