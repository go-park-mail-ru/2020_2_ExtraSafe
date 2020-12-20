package main

import (
	"database/sql"
	"fmt"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/boardStorage"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/boardStorage/attachmentStorage"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/boardStorage/cardsStorage"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/boardStorage/checklistStorage"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/boardStorage/commentStorage"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/boardStorage/tagStorage"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/boardStorage/tasksStorage"
	fStorage "github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/fileStorage"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/service"
	protoBoard "github.com/go-park-mail-ru/2020_2_ExtraSafe/services/proto/board"
	protoProfile "github.com/go-park-mail-ru/2020_2_ExtraSafe/services/proto/profile"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"strings"
)
/*
func init() {
	if err := godotenv.Load("../../../config.env"); err != nil {
		log.Print("No .env file found")
	}
}
*/
func main() {
	driverName:= os.Getenv("TABUTASK_BOARDS_DRIVER")
	userName:= os.Getenv("TABUTASK_BOARDS_USER")
	password:= os.Getenv("TABUTASK_BOARDS_PASSWORD")
	dbName:= os.Getenv("TABUTASK_BOARDS_NAME")
	/*driverName, ok := os.LookupEnv("TABUTASK_BOARDS_DRIVER")
	if !ok {
		log.Fatalf("Cannot find driver name")
	}

	userName, ok := os.LookupEnv("TABUTASK_BOARDS_USER")
	if !ok {
		log.Fatalf("Cannot find user name")
	}

	password, ok := os.LookupEnv("TABUTASK_BOARDS_PASSWORD")
	if !ok {
		log.Fatalf("Cannot find password")
	}

	dbName, ok := os.LookupEnv("TABUTASK_BOARDS_NAME")
	if !ok {
		log.Fatalf("Cannot find db name")
	}*/

	connections := strings.Join([]string{"user=", userName, "password=", password, "dbname=", dbName}, " ")
	db, err := sql.Open(driverName, connections)
	if err != nil {
		log.Fatalf("Cannot connect to database", err)
	}

	db.SetMaxIdleConns(3)
	db.SetMaxOpenConns(10)

	err = db.Ping()
	if err != nil {
		return
	}

	taskStorage := tasksStorage.NewStorage(db)
	cardStorage := cardsStorage.NewStorage(db)
	tagsStorage := tagStorage.NewStorage(db)
	commentsStorage := commentStorage.NewStorage(db)
	checklistsStorage := checklistStorage.NewStorage(db)
	attachStorage := attachmentStorage.NewStorage(db)
	bStorage := boardStorage.NewStorage(db, cardStorage, taskStorage, tagsStorage, commentsStorage, checklistsStorage, attachStorage)

	fileStorage := fStorage.NewStorage()

	// =============================

	profileServiceAddr:= os.Getenv("PROFILE_SERVICE_ADDR")
	/*profileServiceAddr, ok := os.LookupEnv("PROFILE_SERVICE_ADDR")
	if !ok {
		log.Fatalf("")
	}*/

	grpcConn, err := grpc.Dial(
		profileServiceAddr,
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("cant connect to grpc")
	}
	defer grpcConn.Close()

	// =============================

	boardServiceAddr:= os.Getenv("BOARDS_SERVICE_ADDR")
	/*boardServiceAddr, ok := os.LookupEnv("BOARDS_SERVICE_ADDR")
	if !ok {
		log.Fatalf("")
	}*/

	lis, err := net.Listen("tcp", boardServiceAddr)
	if err != nil {
		log.Fatalln("cant listen port", err)
	}

	server := grpc.NewServer()

	profileService := protoProfile.NewProfileClient(grpcConn)

	handler := service.NewService(bStorage, fileStorage, profileService)

	protoBoard.RegisterBoardServer(server, handler)

	fmt.Println("starting server at :9083")
	server.Serve(lis)
}