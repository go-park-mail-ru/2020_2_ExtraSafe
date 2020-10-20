package models

import "github.com/labstack/echo"

type CustomContext struct {
	echo.Context
	UserId    uint64
}
