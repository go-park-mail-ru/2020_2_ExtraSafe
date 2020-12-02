package main

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/cmd/handlers"
	authHandler "github.com/go-park-mail-ru/2020_2_ExtraSafe/cmd/handlers/auth"
	boardsHandler "github.com/go-park-mail-ru/2020_2_ExtraSafe/cmd/handlers/boards"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/cmd/handlers/middlewares"
	profileHandler "github.com/go-park-mail-ru/2020_2_ExtraSafe/cmd/handlers/profile"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/services/auth"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/services/boards"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/services/profile"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/tools/errorWorker"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/tools/logger"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/tools/validation"
	protoAuth "github.com/go-park-mail-ru/2020_2_ExtraSafe/services/proto/auth"
	protoBoard "github.com/go-park-mail-ru/2020_2_ExtraSafe/services/proto/board"
	protoProfile "github.com/go-park-mail-ru/2020_2_ExtraSafe/services/proto/profile"
	_ "github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	_ "github.com/rs/zerolog/log"
	_ "github.com/tarantool/go-tarantool"
	"google.golang.org/grpc"
	"log"
	"os"
)

func main() {

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

	zeroLogger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout})

	internalLogger := logger.NewLogger(&zeroLogger)
	errWorker := errorWorker.NewErrorWorker()

	boardClient := protoBoard.NewBoardClient(grpcConn3)
	profileClient := protoProfile.NewProfileClient(grpcConn2)
	authClient := protoAuth.NewAuthClient(grpcConn1)

	validationService := validation.NewService()
	authService := auth.NewService(authClient, validationService)
	authTransport := auth.NewTransport()
	profileService := profile.NewService(profileClient, validationService)
	profileTransport := profile.NewTransport()
	boardsService := boards.NewService(boardClient, validationService)
	boardsTransport := boards.NewTransport()

	middlewaresService := middlewares.NewMiddleware(errWorker, authService, authTransport, boardsService, internalLogger)

	aHandler := authHandler.NewHandler(authService, authTransport, errWorker)
	profHandler := profileHandler.NewHandler(profileService, profileTransport, errWorker)
	boardHandler := boardsHandler.NewHandler(boardsService, boardsTransport, errWorker)

	e := echo.New()

	e.Use(middlewaresService.CORS())

	handlers.Router(e, profHandler, aHandler, boardHandler, middlewaresService)

	e.Logger.Fatal(e.Start(":8080"))
}
