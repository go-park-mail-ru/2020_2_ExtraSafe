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


func main() {
	dbAddr := os.Getenv("TABUTASK_DB_ADDR")
	dbPort := os.Getenv("TABUTASK_DB_PORT")
	driverName:= os.Getenv("TABUTASK_BOARDS_DRIVER")
	userName:= os.Getenv("TABUTASK_BOARDS_USER")
	password:= os.Getenv("TABUTASK_BOARDS_PASSWORD")
	dbName:= os.Getenv("TABUTASK_BOARDS_NAME")

	connections := strings.Join([]string{"host=",dbAddr, "port=",  dbPort, "user=", userName, "password=", password, "dbname=", dbName, "sslmode=disable"}, " ")
	db, err := sql.Open(driverName, connections)
	if err != nil {
		log.Fatalf("Cannot connect to database", err)
	}

	db.SetMaxIdleConns(3)
	db.SetMaxOpenConns(10)

	err = db.Ping()
	if err != nil {
		log.Fatalf("Cannot connect to database", err)
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

	lis, err := net.Listen("tcp", boardServiceAddr)
	if err != nil {
		log.Fatalln("cant listen port", err)
	}

	server := grpc.NewServer()

	profileService := protoProfile.NewProfileClient(grpcConn)

	handler := service.NewService(bStorage, fileStorage, profileService)

	protoBoard.RegisterBoardServer(server, handler)

	fmt.Println("starting server at : ", lis.Addr())

	err = server.Serve(lis)
	if err != nil {
		log.Fatalln("Serve auth", err)
	}
}