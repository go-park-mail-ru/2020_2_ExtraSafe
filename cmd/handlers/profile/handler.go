package profileHandler

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/services/profile"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/tools/errorWorker"
	"github.com/labstack/echo"
	"net/http"
)

type Handler interface {
	Profile(c echo.Context) error
	Boards(c echo.Context) error
	ProfileChange(c echo.Context) error
	PasswordChange(c echo.Context) error
}

type handler struct {
	profileService   profile.ServiceProfile
	profileTransport profile.Transport
	errorWorker     errorWorker.ErrorWorker
}

func NewHandler(profileService profile.ServiceProfile, profileTransport profile.Transport, errorWorker errorWorker.ErrorWorker) *handler {
	return &handler{
		profileService:   profileService,
		profileTransport: profileTransport,
		errorWorker:     errorWorker,
	}
}

func (h *handler) Profile(c echo.Context) error {
	userInput, err := h.profileTransport.ProfileRead(c)
	if err != nil {
		return h.errorWorker.TransportError(c)
	}

	user, err := h.profileService.Profile(userInput)
	if err != nil {
		return h.errorWorker.RespError(c, err)
	}

	response, _ := h.profileTransport.ProfileWrite(user)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) Boards(c echo.Context) error {
	userInput, err := h.profileTransport.ProfileRead(c)
	if err != nil {
		return h.errorWorker.TransportError(c)
	}

	boards, err := h.profileService.Boards(userInput)
	if err != nil {
		return h.errorWorker.RespError(c, err)
	}

	response, _ := h.profileTransport.BoardsWrite(boards)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) ProfileChange(c echo.Context) error {
	userInput, err := h.profileTransport.ProfileChangeRead(c)
	if err != nil {
		return h.errorWorker.TransportError(c)
	}

	user, err := h.profileService.ProfileChange(userInput)
	if err != nil {
		return h.errorWorker.RespError(c, err)
	}

	response, _ := h.profileTransport.ProfileWrite(user)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) PasswordChange(c echo.Context) error {
	userInput, err := h.profileTransport.PasswordChangeRead(c)
	if err != nil {
		return h.errorWorker.TransportError(c)
	}

	user, err := h.profileService.PasswordChange(userInput)
	if err != nil {
		return h.errorWorker.RespError(c, err)
	}

	response, _ := h.profileTransport.ProfileWrite(user)

	return c.JSON(http.StatusOK, response)
}
