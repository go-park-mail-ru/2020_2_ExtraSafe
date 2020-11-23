package main

import (
	"database/sql"
	"fmt"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/service"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/storage"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/storage/cardsStorage"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/storage/tasksStorage"
	protoBoard "github.com/go-park-mail-ru/2020_2_ExtraSafe/services/proto/board"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
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

	taskStorage := tasksStorage.NewStorage(db)
	cardStorage := cardsStorage.NewStorage(db)
	storage := boardStorage.NewStorage(db, cardStorage, taskStorage)

	lis, err := net.Listen("tcp", ":8083")
	if err != nil {
		log.Fatalln("cant listet port", err)
	}

	server := grpc.NewServer()
	handler := service.NewService(storage)

	protoBoard.RegisterBoardServer(server, handler)

	fmt.Println("starting server at :8083")
	server.Serve(lis)
}