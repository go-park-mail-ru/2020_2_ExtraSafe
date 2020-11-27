package main

import (
	"database/sql"
	"fmt"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/service"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/storage"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/storage/cardsStorage"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/storage/checklistStorage"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/storage/commentStorage"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/storage/tagStorage"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/storage/tasksStorage"
	protoBoard "github.com/go-park-mail-ru/2020_2_ExtraSafe/services/proto/board"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	db, err := sql.Open("postgres", "user=tabutask_admin password=1221 dbname=tabutask_db")
	//FIXME new DB
	//db, err := sql.Open("postgres", "user=tabutask_admin password=1221 dbname=tabutask_boards")
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

	taskStorage := tasksStorage.NewStorage(db)
	cardStorage := cardsStorage.NewStorage(db)
	tagsStorage := tagStorage.NewStorage(db)
	commentsStorage := commentStorage.NewStorage(db)
	checklistsStorage := checklistStorage.NewStorage(db)
	boardStorage := storage.NewStorage(db, cardStorage, taskStorage, tagsStorage, commentsStorage, checklistsStorage)

	lis, err := net.Listen("tcp", ":8083")
	if err != nil {
		log.Fatalln("cant listen port", err)
	}

	server := grpc.NewServer()
	handler := service.NewService(boardStorage)

	protoBoard.RegisterBoardServer(server, handler)

	fmt.Println("starting server at :8083")
	server.Serve(lis)
}