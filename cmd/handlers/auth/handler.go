package authHandler

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/tools/csrf"
	"github.com/labstack/echo"
	"net/http"
)

type Handler interface {
	Auth(c echo.Context) error
	Login(c echo.Context) error
	Logout(c echo.Context) error
	Registration(c echo.Context) error
}

type handler struct {
	authService   AuthService
	authTransport AuthTransport
	authSessions  AuthSessions
	errorWorker   ErrorWorker
}

func NewHandler(authService AuthService, authTransport AuthTransport, authSessions AuthSessions, errorWorker ErrorWorker) *handler {
	return &handler{
		authService:   authService,
		authTransport: authTransport,
		authSessions: authSessions,
		errorWorker:     errorWorker,
	}
}

func (h *handler) Auth(c echo.Context) error {
	userInput, err := h.authTransport.AuthRead(c)
	if err != nil {
		return err
	}

	user, err := h.authService.Auth(userInput)
	if err != nil {
		return err
	}

	response, err := h.authTransport.AuthWrite(user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response)
}

func (h *handler) Login(c echo.Context) error {
	userInput, err := h.authTransport.LoginRead(c)
	if err != nil {
		return err
	}

	userID, _, err := h.authService.Login(userInput)
	if err != nil {
		return err
	}

	token, _ := csrf.GenerateToken(userID)

	response, err := h.authTransport.LoginWrite(token)
	if err != nil {
		return err
	}

	if err := h.authSessions.SetCookie(c, userID); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response)
}

func (h *handler) Logout(c echo.Context) error {
	err := h.authSessions.DeleteCookie(c)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}

func (h *handler) Registration(c echo.Context) error{
	userInput, err := h.authTransport.RegRead(c)
	if err != nil {
		return err
	}

	userID, _, err := h.authService.Registration(userInput)
	if err != nil {
		return err
	}

	response, err := h.authTransport.RegWrite()
	if err != nil {
		return err
	}

	h.authSessions.SetCookie(c, userID)

	return c.JSON(http.StatusOK, response)
}