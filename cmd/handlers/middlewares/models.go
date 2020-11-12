package middlewares

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"github.com/labstack/echo"
)

type sessionsService interface {
	SetCookie(c echo.Context, userID uint64) error
	DeleteCookie(c echo.Context) error
	CheckCookie(c echo.Context) (uint64, error)
}

type errorWorker interface {
	RespError(c echo.Context, serveError error) (err error)
	TransportError(c echo.Context) (err error)
	TokenError(c echo.Context, token string) (err error)
}

type authService interface {
	Auth(request models.UserInput) (response models.UserBoardsOutside, err error)
}

type authTransport interface {
	AuthWrite(user models.UserBoardsOutside) (response models.ResponseUserAuth, err error)
}

type boardStorage interface {
	CheckBoardPermission(userID uint64, boardID uint64, ifAdmin bool) (err error)
	CheckCardPermission(userID uint64, cardID uint64) (err error)
	CheckTaskPermission(userID uint64, taskID uint64) (err error)
}