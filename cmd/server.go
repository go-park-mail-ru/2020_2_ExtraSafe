package main

import (
	"../internal/errorWorker"
	"../internal/models"
	"../internal/services/auth"
	"../internal/services/profile"
	"../internal/services/sessions"
	"../internal/storages/imgStorage"
	"../internal/storages/sessionsStorage"
	"../internal/storages/userStorage"
	"./handlers"
	authHandler "./handlers/auth"
	"./handlers/middlewares"
	profileHandler "./handlers/profile"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo"
	"os"
	"path/filepath"
)

func main() {
	clearDataStore()

	var cfg config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return
	}

	someUsers := make([]models.User, 0)
	userSessions := make(map[string]uint64, 10)

	errWorker := errorWorker.NewErrorWorker()

	usersStorage := userStorage.NewStorage(&someUsers, &userSessions)
	sessionStorage := sessionsStorage.NewStorage(&userSessions)
	avatarStorage := imgStorage.NewStorage(&someUsers, &userSessions)

	sessionService := sessions.NewService(sessionStorage)
	authService := auth.NewService(usersStorage)
	authTransport := auth.NewTransport()
	profileService := profile.NewService(usersStorage, avatarStorage)
	profileTransport := profile.NewTransport()

	middlewaresService := middlewares.NewMiddleware(sessionService, errWorker)

	aHandler := authHandler.NewHandler(authService, authTransport, sessionService, errWorker)
	profHandler := profileHandler.NewHandler(profileService, profileTransport, errWorker)

	e := echo.New()

	e.Use(middlewaresService.CORS())

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &models.CustomContext{
				Context: c,
			}
			return next(cc)
		}
	})

	handlers.Router(e, profHandler, aHandler, middlewaresService)

	e.Logger.Fatal(e.Start(":8080"))
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

