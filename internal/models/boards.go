package models

type BoardInternal struct {
	BoardID  uint64        `json:"boardID"`
	AdminID  uint64        `json:"adminID"`
	Name     string        `json:"name"` // название доски
	Theme    string        `json:"theme"`
	Star     bool          `json:"star"`
	Cards    []CardOutside `json:"cards"`
	UsersIDs []uint64      `json:"usersIDs"` // массив с пользователями этой доски
}

// для пользователя, отправившего запрос формируется эта структура
type BoardOutside struct {
	BoardID uint64             `json:"boardID"`
	Admin   UserOutsideShort   `json:"admin"` // структура владельца доски
	Name    string             `json:"name"`  // название доски
	Theme   string             `json:"theme"`
	Star    bool               `json:"star"`
	Users   []UserOutsideShort `json:"users"` // массив с пользователями этой доски
	Cards   []CardOutside      `json:"cards"`
}

type BoardOutsideShort struct {
	BoardID uint64          `json:"boardID"`
	Name    string          `json:"name"`  // название доски
	Theme   string          `json:"theme"`
	Star    bool            `json:"star"`
}

type BoardInput struct {
	UserID   uint64 `json:"-"`
	BoardID uint64 `json:"boardID"`
}

type BoardChangeInput struct {
	UserID    uint64 `json:"-"`
	BoardID   uint64 `json:"boardID"`
	BoardName string `json:"boardName"`
	Theme     string `json:"theme"`
	Star      bool   `json:"star"`
}

// для работы в бд
type Board struct {
	BoardID  uint64   `json:"boardID"`
	AdminID  uint64   `json:"adminID"`
	Name     string   `json:"name"` // название доски
	Theme    string   `json:"theme"`
	Star     bool     `json:"star"`
	UsersIDs []uint64 `json:"usersIDs"` // массив с пользователями этой доски
}
