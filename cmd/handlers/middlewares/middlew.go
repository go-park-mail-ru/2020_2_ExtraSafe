package middlewares

import (
	"../../../internal/models"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Middleware interface {
	CORS() echo.MiddlewareFunc
	CookieSession(next echo.HandlerFunc) echo.HandlerFunc
}

type middlew struct {
	sessionsService sessionsService
	errorWorker errorWorker
}

func NewMiddleware(sessionsService sessionsService, errorWorker errorWorker) Middleware {
	return middlew{
		sessionsService: sessionsService,
		errorWorker: errorWorker,
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

/*func (m middlew) CookieSession() echo.MiddlewareFunc {

}*/

func (m middlew) CookieSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := c.(*models.CustomContext)

		userId, err := m.sessionsService.CheckCookie(c)
		if err != nil {
			return m.errorWorker.TransportError(c)
		}
		cc.UserId = userId
		return next(c)
	}
}