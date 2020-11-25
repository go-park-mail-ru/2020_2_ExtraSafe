package middlewares

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"github.com/labstack/echo"
)

type sessionsService interface {
	SetCookie(c echo.Context, userID int64) error
	DeleteCookie(c echo.Context) error
	CheckCookie(c echo.Context) (int64, error)
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
	AuthWrite(user models.UserBoardsOutside, token string) (response models.ResponseUserAuth, err error)
}

type boardStorage interface {
	CheckBoardPermission(userID int64, boardID int64, ifAdmin bool) (err error)
	CheckCardPermission(userID int64, cardID int64) (err error)
	CheckTaskPermission(userID int64, taskID int64) (err error)
}