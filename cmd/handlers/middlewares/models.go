package middlewares

import "github.com/labstack/echo"

type sessionsService interface {
	SetCookie(c echo.Context, userID uint64)
	DeleteCookie(c echo.Context) error
}
