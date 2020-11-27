package models

type BoardInternal struct {
	BoardID  int64        `json:"boardID"`
	AdminID  int64        `json:"adminID"`
	Name     string        `json:"name"` // название доски
	Theme    string        `json:"theme"`
	Star     bool          `json:"star"`
	Cards    []CardOutside `json:"cards"`
	UsersIDs []int64      `json:"usersIDs"` // массив с пользователями этой доски
	//FIXME added
	Tags 	[]TagOutside
}

// для пользователя, отправившего запрос формируется эта структура
type BoardOutside struct {
	BoardID int64             `json:"boardID"`
	Admin   UserOutsideShort   `json:"admin"` // структура владельца доски
	Name    string             `json:"boardName"`  // название доски
	Theme   string             `json:"theme"`
	Star    bool               `json:"star"`
	Users   []UserOutsideShort `json:"users"` // массив с пользователями этой доски
	Cards   []CardOutside      `json:"cards"`
	//FIXME added
	Tags 	[]TagOutside
}

type BoardOutsideShort struct {
	BoardID int64          `json:"boardID"`
	Name    string          `json:"boardName"`  // название доски
	Theme   string          `json:"theme"`
	Star    bool            `json:"star"`
}

type BoardInput struct {
	UserID  int64 `json:"-"`
	BoardID int64 `json:"boardID"`
}

type BoardChangeInput struct {
	UserID    int64 `json:"-"`
	BoardID   int64 `json:"boardID"`
	BoardName string `json:"boardName"`
	Theme     string `json:"theme"`
	Star      bool   `json:"star"`
}

// для работы в бд
type Board struct {
	BoardID  int64   `json:"boardID"`
	AdminID  int64   `json:"adminID"`
	Name     string   `json:"boardName"` // название доски
	Theme    string   `json:"theme"`
	Star     bool     `json:"star"`
	UsersIDs []int64 `json:"usersIDs"` // массив с пользователями этой доски
}

// add user to board
type BoardMember struct {
	UserID    int64 `json:"-"`
	BoardID   int64 `json:"boardID"`
	MemberID  int64
}
// TODO return UserShortOutside

