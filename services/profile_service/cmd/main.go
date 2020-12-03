package main

import (
	"database/sql"
	"fmt"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/profile_service/internal/imgStorage"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/profile_service/internal/service"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/profile_service/internal/userStorage"
	protoBoard "github.com/go-park-mail-ru/2020_2_ExtraSafe/services/proto/board"
	protoProfile "github.com/go-park-mail-ru/2020_2_ExtraSafe/services/proto/profile"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"strings"
)

func init() {
	if err := godotenv.Load("../../../config.env"); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	driverName, ok := os.LookupEnv("TABUTASK_USERS_DRIVER")
	if !ok {
		log.Fatalf("Cannot get env parameter")
	}

	userName, ok := os.LookupEnv("TABUTASK_USERS_USER")
	if !ok {
		log.Fatalf("Cannot get env parameter")
	}

	password, ok := os.LookupEnv("TABUTASK_USERS_PASSWORD")
	if !ok {
		log.Fatalf("Cannot get env parameter")
	}

	dbName, ok := os.LookupEnv("TABUTASK_USERS_NAME")
	if !ok {
		log.Fatalf("")
	}

	connections := strings.Join([]string{"user=", userName, "password=", password, "dbname=", dbName}, " ")
	db, err := sql.Open(driverName, connections)
	if err != nil {
		log.Fatalf("Cannot connect to database")
	}

	db.SetMaxIdleConns(3)
	db.SetMaxOpenConns(10)

	err = db.Ping()
	if err != nil {
		return
	}

	profileStorage := userStorage.NewStorage(db)
	avatarStorage := imgStorage.NewStorage()

	// =============================

	boardServiceAddr, ok := os.LookupEnv("BOARDS_SERVICE_ADDR")
	if !ok {
		log.Fatalf("")
	}

	grpcConn, err := grpc.Dial(
		boardServiceAddr,
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("cant connect to grpc")
	}
	defer grpcConn.Close()

	// =============================

	profileServiceAddr, ok := os.LookupEnv("PROFILE_SERVICE_ADDR")
	if !ok {
		log.Fatalf("")
	}

	lis, err := net.Listen("tcp", profileServiceAddr)
	if err != nil {
		log.Fatalln("cant listen port", err)
	}

	// =============================

	server := grpc.NewServer()

	boardService := protoBoard.NewBoardClient(grpcConn)

	handler := profile.NewService(profileStorage, avatarStorage, boardService)

	protoProfile.RegisterProfileServer(server, handler)

	fmt.Println("starting server at :9082")
	server.Serve(lis)
}