package middlewares

import "github.com/labstack/echo"

type sessionsService interface {
	SetCookie(c echo.Context, userID uint64)
	DeleteCookie(c echo.Context) error
	CheckCookie(c echo.Context) (uint64, error)
}

type errorWorker interface {
	RespError(c echo.Context, serveError error) (err error)
	TransportError(c echo.Context) (err error)
}