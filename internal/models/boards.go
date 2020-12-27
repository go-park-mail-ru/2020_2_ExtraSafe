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
	UserID    int64  `json:"-"`
	SessionID string `json:"-"`
	BoardID   int64  `json:"boardID"`
}

type BoardChangeInput struct {
	UserID    int64  `json:"-"`
	SessionID string `json:"-"`
	BoardID   int64  `json:"boardID"`
	BoardName string `json:"boardName"`
	Theme     string `json:"boardTheme"`
	Star      bool   `json:"boardStar"`
}

type BoardMemberInput struct {
	UserID     int64  `json:"-"`
	SessionID  string `json:"-"`
	BoardID    int64  `json:"boardID"`
	MemberName string `json:"memberUsername"`
}

type BoardInviteInput struct {
	UserID  int64
	BoardID int64
	UrlHash string
}

type BoardInputTemplate struct {
	UserID       int64  `json:"-"`
	SessionID    string `json:"-"`
	TemplateSlug string `json:"templateSlug"`
	BoardName    string `json:"boardName"`
}

//===================================================<-Internal
type BoardInternalTemplate struct {
	AdminID  int64          `json:"adminID"`
	BoardName     string         `json:"boardName"`
	Cards    []CardInternal `json:"cards"`
	Tags     []TagOutside	`json:"tags"`
}

type BoardInternal struct {
	BoardID  int64          `json:"boardID"`
	AdminID  int64          `json:"adminID"`
	Name     string         `json:"boardName"`
	Theme    string         `json:"boardTheme"`
	Star     bool           `json:"boardStar"`
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

type BoardTemplateOutsideShort struct {
	TemplateSlug string `json:"templateSlug"`
	TemplateName string `json:"templateName"`
	Description  string `json:"templateDescription"`
}

type BoardMemberOutside struct {
	BoardName string `json:"boardName"`
	UserOutsideShort
	Initiator string `json:"initiator"`
}

//===================================================<-Other
type BoardMember struct {
	UserID   int64
	BoardID  int64
	MemberID int64
}
