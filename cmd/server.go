package main

import (
	"database/sql"
	"fmt"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/cmd/handlers"
	authHandler "github.com/go-park-mail-ru/2020_2_ExtraSafe/cmd/handlers/auth"
	boardsHandler "github.com/go-park-mail-ru/2020_2_ExtraSafe/cmd/handlers/boards"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/cmd/handlers/middlewares"
	profileHandler "github.com/go-park-mail-ru/2020_2_ExtraSafe/cmd/handlers/profile"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/errorWorker"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/services/auth"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/services/boards"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/services/profile"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/services/sessions"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/storages/userStorage"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/tools/validation"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/auth_servise/internal/sessionsStorage"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/storage/cardsStorage"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/storage/tasksStorage"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/profile_service/internal/imgStorage"
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	_ "github.com/rs/zerolog/log"
	"github.com/tarantool/go-tarantool"
	"os"
)

func main() {
	var cfg config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return
	}

	//TODO строку подключения в конфиг

	db, err := sql.Open("postgres", "user=tabutask_admin password=1221 dbname=tabutask_db")
	if err != nil {
		fmt.Println(err)
		return
	}

	db.SetMaxIdleConns(3)
	db.SetMaxOpenConns(10)

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}

	tConn, err := tarantool.Connect("127.0.0.1:3301", tarantool.Opts{ User: "guest" })

	if err != nil {
		fmt.Println("Connection refused")
	}
	defer tConn.Close()

	log := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout})

	errWorker := errorWorker.NewErrorWorker()

	usersStorage := userStorage.NewStorage(db)
	sessionStorage := sessionsStorage.NewStorage(tConn)
	avatarStorage := imgStorage.NewStorage()
	cardStorage := cardsStorage.NewStorage(db)
	taskStorage := tasksStorage.NewStorage(db)
	boardsStorage := boardStorage.NewStorage(db, cardStorage, taskStorage)

	validationService := validation.NewService()
	sessionService := sessions.NewService(sessionStorage)
	authService := auth.NewService(usersStorage, boardsStorage, validationService)
	authTransport := auth.NewTransport()
	profileService := profile.NewService(usersStorage, avatarStorage, boardsStorage, validationService)
	profileTransport := profile.NewTransport()
	boardsService := boards.NewService(usersStorage, boardsStorage, validationService)
	boardsTransport := boards.NewTransport()

	middlewaresService := middlewares.NewMiddleware(sessionService, errWorker, authService, authTransport, boardsStorage, &log)

	aHandler := authHandler.NewHandler(authService, authTransport, sessionService, errWorker)
	profHandler := profileHandler.NewHandler(profileService, profileTransport, errWorker)
	boardHandler := boardsHandler.NewHandler(boardsService, boardsTransport, errWorker)

	e := echo.New()

	e.Use(middlewaresService.CORS())

	handlers.Router(e, profHandler, aHandler, boardHandler, middlewaresService)

	e.Logger.Fatal(e.Start(":8080"))
}
