package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Middleware interface {
	CORS() echo.MiddlewareFunc
}

type middlew struct {
	sessionsService sessionsService
}

func NewMiddleware(sessionsService sessionsService) Middleware {
	return middlew{
		sessionsService: sessionsService,
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