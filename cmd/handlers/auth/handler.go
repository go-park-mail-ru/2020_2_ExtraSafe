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
	return nil
}

func (h *handler) Login(c echo.Context) error {
	cc := c.(*models.CustomContext)

	/*response, err := cc.checkUserAuthorized(c)
	if err == nil {
		return c.JSON(http.StatusOK, response)
	}*/

	userInput, err := h.authTransport.LoginRead(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err)
	}

	response, err := h.authService.Login(userInput)
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

	/*response, err := cc.checkUserAuthorized(c)
	if err == nil {
		return c.JSON(http.StatusOK, response)
	}*/

	userInput, err := h.authTransport.RegRead(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err)
	}

	response, err := h.authService.Registration(userInput)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err)
	}

	h.authSessions.SetCookie(c, cc.UserId)

	return c.JSON(http.StatusOK, response)
}