package middlewares

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

type Middleware interface {
	CORS() echo.MiddlewareFunc
	CookieSession(next echo.HandlerFunc) echo.HandlerFunc
	AuthCookieSession(next echo.HandlerFunc) echo.HandlerFunc
}

type middlew struct {
	sessionsService sessionsService
	errorWorker errorWorker
	authService authService
	authTransport authTransport
}

func NewMiddleware(sessionsService sessionsService, errorWorker errorWorker, authService authService, authTransport authTransport) Middleware {
	return middlew{
		sessionsService: sessionsService,
		errorWorker: errorWorker,
		authService: authService,
		authTransport: authTransport,
	}
}

func (m middlew) CORS() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://95.163.213.142:80",
			"http://95.163.213.142",
			"http://127.0.0.1:3033",
			"http://tabutask.ru"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
	})
}

func (m middlew) CookieSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId, err := m.sessionsService.CheckCookie(c)
		if err != nil {
			return m.errorWorker.TransportError(c)
		}
		c.Set("userId", userId)
		return next(c)
	}
}

func (m middlew) AuthCookieSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId, err := m.sessionsService.CheckCookie(c)
		if err == nil {
			userInput := new(models.UserInput)
			userInput.ID = userId
			user, _ := m.authService.Auth(*userInput)
			response, err := m.authTransport.AuthWrite(user)
			if err != nil {
				return m.errorWorker.TransportError(c)
			}
			return c.JSON(http.StatusOK, response)
		}
		return next(c)
	}
}