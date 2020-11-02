package authHandler

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
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
	authService   authService
	authTransport authTransport
	authSessions authSessions
	errorWorker     errorWorker
}

func NewHandler(authService authService, authTransport authTransport, authSessions authSessions, errorWorker errorWorker) *handler {
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
		return h.errorWorker.TransportError(c)
	}

	user, err := h.authService.Auth(userInput)
	if err != nil {
		return h.errorWorker.RespError(c, err)
	}

	response, err := h.authTransport.AuthWrite(user)
	if err != nil {
		return h.errorWorker.TransportError(c)
	}

	return c.JSON(http.StatusOK, response)
}

func (h *handler) Login(c echo.Context) error {
	userInput, err := h.authTransport.LoginRead(c)
	if err != nil {
		return h.errorWorker.TransportError(c)
	}

	userID, _, err := h.authService.Login(userInput)
	if err != nil {
		return h.errorWorker.RespError(c, err)
	}

	response, err := h.authTransport.LoginWrite()
	if err != nil {
		return h.errorWorker.TransportError(c)
	}

	if err := h.authSessions.SetCookie(c, userID); err != nil {
		return c.JSON(http.StatusInternalServerError, response)
	}

	return c.JSON(http.StatusOK, response)
}

func (h *handler) Logout(c echo.Context) error {
	err := h.authSessions.DeleteCookie(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseUser{})
	}
	return c.JSON(http.StatusOK, models.ResponseUser{})
}

func (h *handler) Registration(c echo.Context) error{
	userInput, err := h.authTransport.RegRead(c)
	if err != nil {
		return h.errorWorker.TransportError(c)
	}

	userID, _, err := h.authService.Registration(userInput)
	if err != nil {
		return h.errorWorker.RespError(c, err)
	}

	response, err := h.authTransport.RegWrite()
	if err != nil {
		return h.errorWorker.TransportError(c)
	}

	h.authSessions.SetCookie(c, userID)

	return c.JSON(http.StatusOK, response)
}