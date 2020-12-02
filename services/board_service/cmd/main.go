package main

import (
	"database/sql"
	"fmt"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/boardStorage"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/boardStorage/cardsStorage"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/boardStorage/checklistStorage"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/boardStorage/commentStorage"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/boardStorage/tagStorage"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/boardStorage/tasksStorage"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/boardStorage/attachmentStorage"
	fStorage "github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/fileStorage"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/service"
	protoBoard "github.com/go-park-mail-ru/2020_2_ExtraSafe/services/proto/board"
	protoProfile "github.com/go-park-mail-ru/2020_2_ExtraSafe/services/proto/profile"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	db, err := sql.Open("postgres", "user=tabutask_admin password=1221 dbname=tabutask_boards")
	if err != nil {
		return
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

	grpcConn, err := grpc.Dial(
		"127.0.0.1:9082",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("cant connect to grpc")
	}
	defer grpcConn.Close()

	// =============================

	lis, err := net.Listen("tcp", ":9083")
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