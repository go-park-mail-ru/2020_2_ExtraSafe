package main

import (
	"database/sql"
	"fmt"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/profile_service/internal/imgstorage"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/profile_service/internal/service"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/profile_service/internal/userstorage"
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
	driverName := os.Getenv("TABUTASK_USERS_DRIVER")
	userName := os.Getenv("TABUTASK_USERS_USER")
	password := os.Getenv("TABUTASK_USERS_PASSWORD")
	dbName := os.Getenv("TABUTASK_USERS_NAME")

	connections := strings.Join([]string{"host=", dbAddr, "port=", dbPort, "user=", userName, "password=", password, "dbname=", dbName, "sslmode=disable"}, " ")
	db, err := sql.Open(driverName, connections)
	if err != nil {
		log.Fatalln("Cannot connect to database", err)
	}

	db.SetMaxIdleConns(3)
	db.SetMaxOpenConns(10)

	err = db.Ping()
	if err != nil {
		log.Fatalln("Cannot ping to database", err)
	}

	profileStorage := userstorage.NewStorage(db)
	avatarStorage := imgstorage.NewStorage()

	// =============================

	boardServiceAddr := os.Getenv("BOARDS_SERVICE_ADDR")

	grpcConn, err := grpc.Dial(
		boardServiceAddr,
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("cant connect to grpc")
	}
	defer grpcConn.Close()

	// =============================
	profileServiceAddr := os.Getenv("PROFILE_SERVICE_ADDR")

	lis, err := net.Listen("tcp", profileServiceAddr)
	if err != nil {
		log.Fatalln("cant listen port", err)
	}

	// =============================

	server := grpc.NewServer()

	boardService := protoBoard.NewBoardClient(grpcConn)

	handler := profile.NewService(profileStorage, avatarStorage, boardService)

	protoProfile.RegisterProfileServer(server, handler)

	fmt.Println("starting server at : ", lis.Addr())

	err = server.Serve(lis)
	if err != nil {
		log.Fatalln("Serve auth", err)
	}
}
