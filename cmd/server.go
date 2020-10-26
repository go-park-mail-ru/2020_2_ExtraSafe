package main

import (
	"fmt"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/cmd/handlers"
	authHandler "github.com/go-park-mail-ru/2020_2_ExtraSafe/cmd/handlers/auth"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/cmd/handlers/middlewares"
	profileHandler "github.com/go-park-mail-ru/2020_2_ExtraSafe/cmd/handlers/profile"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/errorWorker"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/services/auth"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/services/profile"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/services/sessions"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/storages/imgStorage"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/storages/sessionsStorage"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/storages/userStorage"
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

	usersStorage := userStorage.NewStorage(&someUsers)
	sessionStorage := sessionsStorage.NewStorage(&userSessions)
	avatarStorage := imgStorage.NewStorage(&someUsers)

	sessionService := sessions.NewService(sessionStorage)
	authService := auth.NewService(usersStorage)
	authTransport := auth.NewTransport()
	profileService := profile.NewService(usersStorage, avatarStorage)
	profileTransport := profile.NewTransport()

	middlewaresService := middlewares.NewMiddleware(sessionService, errWorker, authService, authTransport)

	aHandler := authHandler.NewHandler(authService, authTransport, sessionService, errWorker)
	profHandler := profileHandler.NewHandler(profileService, profileTransport, errWorker)

	e := echo.New()

	e.Use(middlewaresService.CORS())

	handlers.Router(e, profHandler, aHandler, middlewaresService)

	e.Logger.Fatal(e.Start(":8080"))
}

func clearDataStore() {
	dir := "../avatars"
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

