package main

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/cmd/handlers"
	authHandler "github.com/go-park-mail-ru/2020_2_ExtraSafe/cmd/handlers/auth"
	boardsHandler "github.com/go-park-mail-ru/2020_2_ExtraSafe/cmd/handlers/boards"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/cmd/handlers/middlewares"
	profileHandler "github.com/go-park-mail-ru/2020_2_ExtraSafe/cmd/handlers/profile"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/errorWorker"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/services/auth"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/services/boards"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/services/profile"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/tools/validation"
<<<<<<< HEAD
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/auth_servise/internal/sessionsStorage"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/boardStorage/cardsStorage"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/boardStorage/tasksStorage"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/profile_service/internal/imgStorage"
	"github.com/kelseyhightower/envconfig"
=======
	protoAuth "github.com/go-park-mail-ru/2020_2_ExtraSafe/services/proto/auth"
	protoBoard "github.com/go-park-mail-ru/2020_2_ExtraSafe/services/proto/board"
	protoProfile "github.com/go-park-mail-ru/2020_2_ExtraSafe/services/proto/profile"
	_"github.com/kelseyhightower/envconfig"
>>>>>>> 459425ed2a488aa32ace632ce5a864946511248a
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	_ "github.com/rs/zerolog/log"
	_"github.com/tarantool/go-tarantool"
	"google.golang.org/grpc"
	"log"
	"os"
)

func main() {
	/*var cfg config
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
	defer tConn.Close()*/

	// =============================

	grpcConn3, err := grpc.Dial(
		"127.0.0.1:9083",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("cant connect to grpc")
	}
	defer grpcConn3.Close()

	// =============================
	// =============================

	grpcConn2, err := grpc.Dial(
		"127.0.0.1:9082",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("cant connect to grpc")
	}
	defer grpcConn2.Close()

	// =============================

	// =============================

	grpcConn1, err := grpc.Dial(
		"127.0.0.1:9081",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("cant connect to grpc")
	}
	defer grpcConn1.Close()

	// =============================

	loggg := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout})

	errWorker := errorWorker.NewErrorWorker()

	/*usersStorage := userStorage.NewStorage(db)
	sessionStorage := sessionsStorage.NewStorage(tConn)
	avatarStorage := imgStorage.NewStorage()
	cardStorage := cardsStorage.NewStorage(db)
	taskStorage := tasksStorage.NewStorage(db)
	boardsStorage := boardStorage.NewStorage(db, cardStorage, taskStorage)*/

	boardClient := protoBoard.NewBoardClient(grpcConn3)
	profileClient := protoProfile.NewProfileClient(grpcConn2)
	authClient := protoAuth.NewAuthClient(grpcConn1)

	validationService := validation.NewService()
	//sessionService := sessions.NewService(sessionStorage)
	authService := auth.NewService(authClient, validationService)
	authTransport := auth.NewTransport()
	profileService := profile.NewService(profileClient, validationService)
	profileTransport := profile.NewTransport()
	boardsService := boards.NewService(boardClient, validationService)
	boardsTransport := boards.NewTransport()

	middlewaresService := middlewares.NewMiddleware(errWorker, authService, authTransport, boardsService, &loggg)

	aHandler := authHandler.NewHandler(authService, authTransport, errWorker)
	profHandler := profileHandler.NewHandler(profileService, profileTransport, errWorker)
	boardHandler := boardsHandler.NewHandler(boardsService, boardsTransport, errWorker)

	e := echo.New()

	e.Use(middlewaresService.CORS())

	handlers.Router(e, profHandler, aHandler, boardHandler, middlewaresService)

	e.Logger.Fatal(e.Start(":8080"))
}
