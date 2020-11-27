package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/boardStorage"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/boardStorage/cardsStorage"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/boardStorage/checklistStorage"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/boardStorage/commentStorage"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/boardStorage/tagStorage"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/boardStorage/tasksStorage"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/service"
	protoBoard "github.com/go-park-mail-ru/2020_2_ExtraSafe/services/proto/board"
	protoProfile "github.com/go-park-mail-ru/2020_2_ExtraSafe/services/proto/profile"
	_ "github.com/lib/pq"
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
	bStorage := boardStorage.NewStorage(db, cardStorage, taskStorage, tagsStorage, commentsStorage, checklistsStorage)

	/*// =============================

	grpcConn := new(grpc.ClientConn)
	go func(conn *grpc.ClientConn) *grpc.ClientConn {
		for  {
			conn, err = grpc.Dial(
				"127.0.0.1:9082",
				grpc.WithInsecure(),
			)
			if err != nil {
				fmt.Println("cant connect to grpc")
				continue
			}
			fmt.Println("Connect")
			return conn
		}
	}(grpcConn)

	defer grpcConn.Close()
	// =============================*/

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

	_, err = profileService.Accounts(context.Background(), &protoProfile.UserID{ID: 2})
	fmt.Println(err)

	handler := service.NewService(bStorage, profileService)

	protoBoard.RegisterBoardServer(server, handler)

	fmt.Println("starting server at :9083")
	//fmt.Println(server.GetServiceInfo())
	server.Serve(lis)
}