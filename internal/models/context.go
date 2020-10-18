package models

import "github.com/labstack/echo"

type CustomContext struct {
	echo.Context
	UserId    uint64
	//Sessions *map[string]uint64 //map[sessionID]userID
}
