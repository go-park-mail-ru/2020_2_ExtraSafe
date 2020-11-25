package authHandler

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"github.com/labstack/echo"
)


//go:generate mockgen -destination=./mock/mock_authService.go -package=mock github.com/go-park-mail-ru/2020_2_ExtraSafe/cmd/handlers/auth AuthService
//go:generate mockgen -destination=./mock/mock_authTransport.go -package=mock github.com/go-park-mail-ru/2020_2_ExtraSafe/cmd/handlers/auth AuthTransport
//go:generate mockgen -destination=./mock/mock_authSessions.go -package=mock github.com/go-park-mail-ru/2020_2_ExtraSafe/cmd/handlers/auth AuthSessions


type AuthService interface {
	Auth(request models.UserInput) (response models.UserBoardsOutside, err error)
	Login(request models.UserInputLogin) (userID int64, response models.UserOutside, err error)
	Registration(request models.UserInputReg) (userID int64, response models.UserOutside, err error)
}

type AuthTransport interface {
	AuthRead(c echo.Context) (request models.UserInput, err error)
	LoginRead(c echo.Context) (request models.UserInputLogin, err error)
	RegRead(c echo.Context) (request models.UserInputReg, err error)

	AuthWrite(user models.UserBoardsOutside, token string) (response models.ResponseUserAuth, err error)
	LoginWrite(token string) (response models.ResponseToken, err error)
	RegWrite() (response models.ResponseStatus, err error)
}

type AuthSessions interface {
	SetCookie(c echo.Context, userID int64) error
	DeleteCookie(c echo.Context) error
	CheckCookie(c echo.Context) (int64, error)
}

type ErrorWorker interface {
	RespError(c echo.Context, serveError error) (err error)
	TransportError(c echo.Context) (err error)
}
