package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"os"
	"path/filepath"
)

func main() {
	clearDataStore()

	someUsers := make([]User, 0)
	sessions := make(map[string]uint64, 10)

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://95.163.213.142:80", "http://95.163.213.142", "http://127.0.0.1:3033"},
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

	//defer clearDataStore()
}

func clearDataStore() {
	dir := "avatars"
	d, err := os.Open(dir)
	if err != nil {
		fmt.Println("Cannot clear avatars datatore")
		return
	}
	names, err := d.Readdirnames(-1)
	if err != nil {
		fmt.Println("Cannot clear avatars datatore")
		return
	}

	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			fmt.Println("Cannot clear avatars datatore")
			return
		}
	}
}