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
	Theme     string `json:"theme"`
	Star      bool   `json:"star"`
}

type BoardMemberInput struct {
	UserID     int64  `json:"-"`
	BoardID    int64  `json:"boardID"`
	MemberName string `json:"memberName"`
}
// TODO return UserShortOutside

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
	//FIXME added
}

//===================================================<-Outside
type BoardOutside struct {
	BoardID int64              `json:"boardID"`
	Admin   UserOutsideShort   `json:"admin"`
	Name    string             `json:"boardName"`
	Theme   string             `json:"theme"`
	Star    bool               `json:"star"`
	Users   []UserOutsideShort `json:"users"`
	Cards   []CardOutside      `json:"cards"`
	Tags    []TagOutside       `json:"tags"`
	//FIXME added
}

type BoardOutsideShort struct {
	BoardID int64  `json:"boardID"`
	Name    string `json:"boardName"`
	Theme   string `json:"theme"`
	Star    bool   `json:"star"`
}

//===================================================<-Other
type BoardMember struct {
	UserID   int64
	BoardID  int64
	MemberID int64
}
