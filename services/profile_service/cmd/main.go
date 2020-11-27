package main

import (
	"database/sql"
	"fmt"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/profile_service/internal/service"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/profile_service/internal/storage"
	protoBoard "github.com/go-park-mail-ru/2020_2_ExtraSafe/services/proto/board"
	protoProfile "github.com/go-park-mail-ru/2020_2_ExtraSafe/services/proto/profile"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	// =============================
	db, err := sql.Open("postgres", "user=tabutask_admin password=1221 dbname=tabutask_db")
	//FIXME new DB
	//db, err := sql.Open("postgres", "user=tabutask_admin password=1221 dbname=tabutask_users")
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

	userStorage := storage.NewStorage(db)
	//avatarStorage := imgStorage.NewStorage()

	// =============================

	grpcConn, err := grpc.Dial(
		"127.0.0.1:8083",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("cant connect to grpc")
	}
	defer grpcConn.Close()

	// =============================

	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatalln("cant listet port", err)
	}

	// =============================

	server := grpc.NewServer()

	boardService := protoBoard.NewBoardClient(grpcConn)

	handler := profile.NewService(userStorage, boardService)

	protoProfile.RegisterProfileServer(server, handler)

	fmt.Println("starting server at :8082")
	server.Serve(lis)
}