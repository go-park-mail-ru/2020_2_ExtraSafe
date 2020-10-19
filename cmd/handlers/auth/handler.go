package authHandler

import (
	"../../../internal/models"
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
	userId, err := h.authSessions.CheckCookie(c)
	if err == nil {
		userInput := new(models.UserInput)
		userInput.ID = userId
		user, _ := h.authService.Auth(*userInput)
		response, err := h.authTransport.AuthWrite(user)
		if err != nil {
			return h.errorWorker.TransportError(c)
		}
		return c.JSON(http.StatusOK, response)
	}

	userInput, err := h.authTransport.LoginRead(c)
	if err != nil {
		return h.errorWorker.TransportError(c)
	}

	user, err := h.authService.Login(userInput)
	if err != nil {
		return h.errorWorker.RespError(c, err)
	}

	response, err := h.authTransport.AuthWrite(user)
	if err != nil {
		return h.errorWorker.TransportError(c)
	}

	h.authSessions.SetCookie(c, user.ID)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) Logout(c echo.Context) error {
	h.authSessions.DeleteCookie(c)
	return c.JSON(http.StatusOK, models.ResponseUser{})
}

func (h *handler) Registration(c echo.Context) error{
	userId, err := h.authSessions.CheckCookie(c)
	if err == nil {
		userInput := new(models.UserInput)
		userInput.ID = userId
		user, _ := h.authService.Auth(*userInput)
		response, err := h.authTransport.AuthWrite(user)
		if err != nil {
			return h.errorWorker.TransportError(c)
		}
		return c.JSON(http.StatusOK, response)
	}

	userInput, err := h.authTransport.RegRead(c)
	if err != nil {
		return h.errorWorker.TransportError(c)
	}

	user, err := h.authService.Registration(userInput)
	if err != nil {
		return h.errorWorker.RespError(c, err)
	}

	response, err := h.authTransport.AuthWrite(user)
	if err != nil {
		return h.errorWorker.TransportError(c)
	}

	h.authSessions.SetCookie(c, user.ID)

	return c.JSON(http.StatusOK, response)
}