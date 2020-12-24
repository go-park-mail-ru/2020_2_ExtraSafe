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
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/tools/errorworker"
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
	boardServiceAddr := os.Getenv("BOARDS_SERVICE_ADDR")
	profileServiceAddr := os.Getenv("PROFILE_SERVICE_ADDR")
	authServiceAddr := os.Getenv("AUTH_SERVICE_ADDR")
	mainServiceAddr := os.Getenv("TABUTASK_SERVER_ADDR")

	// =============================

	grpcConnBoard, err := grpc.Dial(
		boardServiceAddr,
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("cant connect to grpc")
	}
	defer grpcConnBoard.Close()

	// =============================

	grpcConnProfile, err := grpc.Dial(
		profileServiceAddr,
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("cant connect to grpc")
	}
	defer grpcConnProfile.Close()

	// =============================

	grpcConnAuth, err := grpc.Dial(
		authServiceAddr,
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("cant connect to grpc")
	}
	defer grpcConnAuth.Close()

	// =============================

	zeroLogger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout})

	internalLogger := logger.NewLogger(&zeroLogger)
	errWorker := errorworker.NewErrorWorker()

	boardClient := protoBoard.NewBoardClient(grpcConnBoard)
	profileClient := protoProfile.NewProfileClient(grpcConnProfile)
	authClient := protoAuth.NewAuthClient(grpcConnAuth)

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

	e.Logger.Fatal(e.Start(mainServiceAddr))
}
