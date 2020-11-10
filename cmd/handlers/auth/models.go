package authHandler

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"github.com/labstack/echo"
)

type authService interface {
	Auth(request models.UserInput) (response models.UserBoardsOutside, err error)
	Login(request models.UserInputLogin) (userID uint64, response models.UserOutside, err error)
	Registration(request models.UserInputReg) (userID uint64, response models.UserOutside, err error)
}

type authTransport interface {
	AuthRead(c echo.Context) (request models.UserInput, err error)
	LoginRead(c echo.Context) (request models.UserInputLogin, err error)
	RegRead(c echo.Context) (request models.UserInputReg, err error)

	AuthWrite(user models.UserBoardsOutside) (response models.ResponseUserAuth, err error)
	LoginWrite(token string) (response models.ResponseToken, err error)
	RegWrite() (response models.ResponseStatus, err error)
}

type authSessions interface {
	SetCookie(c echo.Context, userID uint64) error
	DeleteCookie(c echo.Context) error
	CheckCookie(c echo.Context) (uint64, error)
}

type errorWorker interface {
	RespError(c echo.Context, serveError error) (err error)
	TransportError(c echo.Context) (err error)
}
