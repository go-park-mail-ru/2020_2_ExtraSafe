package middlewares

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"github.com/labstack/echo"
)

type sessionsService interface {
	SetCookie(c echo.Context, userID uint64)
	DeleteCookie(c echo.Context) error
	CheckCookie(c echo.Context) (uint64, error)
}

type errorWorker interface {
	RespError(c echo.Context, serveError error) (err error)
	TransportError(c echo.Context) (err error)
}

type authService interface {
	Auth(request models.UserInput) (response models.User, err error)
}

type authTransport interface {
	AuthWrite(user models.User) (response models.ResponseUserAuth, err error)
}