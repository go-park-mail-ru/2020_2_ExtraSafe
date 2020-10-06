package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	someUsers := make([]User, 0)
	sessions := make(map[string]uint64, 10)

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://95.163.213.142:80"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
	}))

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &Handlers{c,
				&someUsers,
				&sessions,
			}
			return next(cc)
		}
	})

	router(e)

	e.Logger.Fatal(e.Start(":8080"))
}
