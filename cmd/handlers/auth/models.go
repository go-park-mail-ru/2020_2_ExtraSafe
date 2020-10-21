package authHandler

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"github.com/labstack/echo"
)

type authService interface {
	Auth(request models.UserInput) (response models.User, err error)
	Login(request models.UserInputLogin) (response models.User, err error)
	Registration(request models.UserInputReg) (response models.User, err error)
}

type authTransport interface {
	AuthRead(c echo.Context) (request models.UserInput, err error)
	LoginRead(c echo.Context) (request models.UserInputLogin, err error)
	RegRead(c echo.Context) (request models.UserInputReg, err error)

	AuthWrite(user models.User) (response models.ResponseUserAuth, err error)
}

type authSessions interface {
	SetCookie(c echo.Context, userID uint64)
	DeleteCookie(c echo.Context) error
	CheckCookie(c echo.Context) (uint64, error)
}

type errorWorker interface {
	RespError(c echo.Context, serveError error) (err error)
	TransportError(c echo.Context) (err error)
}
