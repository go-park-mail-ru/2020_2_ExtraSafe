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
	//errorWorker     errorWorker
}

func NewHandler(authService authService, authTransport authTransport, authSessions authSessions) *handler {
	return &handler{
		authService:   authService,
		authTransport: authTransport,
		authSessions: authSessions,
		//errorWorker:     errorWorker,
	}
}

func (h *handler) Auth(c echo.Context) error {
	userId, err := h.authSessions.CheckCookie(c)
	if err == nil {
		userInput := new(models.UserInput)
		userInput.ID = userId

		user, _ := h.authService.Auth(*userInput)

		response, err := h.authTransport.AuthWrite(user)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, err)
		}
		return c.JSON(http.StatusOK, response)
	}
	return c.JSON(http.StatusUnauthorized, err)
}

func (h *handler) Login(c echo.Context) error {
	cc := c.(*models.CustomContext)

	userId, err := h.authSessions.CheckCookie(c)
	if err == nil {
		userInput := new(models.UserInput)
		userInput.ID = userId
		user, _ := h.authService.Auth(*userInput)
		response, err := h.authTransport.AuthWrite(user)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, err)
		}
		return c.JSON(http.StatusOK, response)
	}

	userInput, err := h.authTransport.LoginRead(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err)
	}

	user, err := h.authService.Login(userInput)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err)
	}

	response, err := h.authTransport.AuthWrite(user)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err)
	}

	h.authSessions.SetCookie(c, cc.UserId)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) Logout(c echo.Context) error {
	h.authSessions.DeleteCookie(c)
	return c.JSON(http.StatusOK, models.ResponseUser{})
}

func (h *handler) Registration(c echo.Context) error{
	cc := c.(*models.CustomContext)

	userId, err := h.authSessions.CheckCookie(c)
	if err == nil {
		userInput := new(models.UserInput)
		userInput.ID = userId
		user, _ := h.authService.Auth(*userInput)
		response, err := h.authTransport.AuthWrite(user)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, err)
		}
		return c.JSON(http.StatusOK, response)
	}

	userInput, err := h.authTransport.RegRead(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err)
	}

	user, err := h.authService.Registration(userInput)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err)
	}

	response, err := h.authTransport.AuthWrite(user)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err)
	}

	h.authSessions.SetCookie(c, cc.UserId)

	return c.JSON(http.StatusOK, response)
}