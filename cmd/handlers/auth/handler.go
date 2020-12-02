package authHandler

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/services/auth"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/tools/csrf"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/tools/errorWorker"
	"github.com/labstack/echo"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"github.com/mailru/easyjson"
	"net/http"
	"time"
)

type Handler interface {
	Auth(c echo.Context) error
	Login(c echo.Context) error
	Logout(c echo.Context) error
	Registration(c echo.Context) error
}

type handler struct {
	authService   auth.ServiceAuth
	authTransport auth.TransportAuth
	errorWorker   errorWorker.ErrorWorker
}

func NewHandler(authService auth.ServiceAuth, authTransport auth.TransportAuth, errorWorker errorWorker.ErrorWorker) *handler {
	return &handler{
		authService:   authService,
		authTransport: authTransport,
		errorWorker:   errorWorker,
	}
}

func (h *handler) Auth(c echo.Context) error {
	userInput, err := h.authTransport.AuthRead(c)
	if err != nil {
		return h.errorWorker.TransportError(c)
	}

	user, err := h.authService.Auth(userInput)
	if err != nil {
		return h.errorWorker.RespError(c, err)
	}

	token, _ := csrf.GenerateToken(userInput.ID)

	response, _ := h.authTransport.AuthWrite(user, token)

	response.MarshalJSON

	return c.JSON(http.StatusOK, response)
}

func (h *handler) Login(c echo.Context) error {
	userInput, err := h.authTransport.LoginRead(c)
	if err != nil {
		return h.errorWorker.TransportError(c)
	}

	user, err := h.authService.Login(userInput)
	if err != nil {
		return h.errorWorker.RespError(c, err)
	}

	token, _ := csrf.GenerateToken(user.UserID)

	response, _ := h.authTransport.LoginWrite(token)

	cookie := new(http.Cookie)
	cookie.Name = "tabutask_id"
	cookie.Value = user.SessionID
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.HttpOnly = true
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) Logout(c echo.Context) error {
	err := h.authService.Logout(c)
	if err != nil {
		return h.errorWorker.RespError(c, err)
	}

	cookie := new(http.Cookie)
	cookie.Name = "tabutask_id"
	cookie.Expires = time.Now().AddDate(0, 0, -1)
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}

func (h *handler) Registration(c echo.Context) error{
	userInput, err := h.authTransport.RegRead(c)
	if err != nil {
		return h.errorWorker.TransportError(c)
	}

	user, err := h.authService.Registration(userInput)
	if err != nil {
		return h.errorWorker.RespError(c, err)
	}

	response, _ := h.authTransport.RegWrite()

	cookie := new(http.Cookie)
	cookie.Name = "tabutask_id"
	cookie.Value = user.SessionID
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.HttpOnly = true
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, response)
}