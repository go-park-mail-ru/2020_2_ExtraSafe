package main

import (
	"fmt"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/auth_servise/internal/service"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/auth_servise/internal/storage"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/proto/auth"
	protoBoard "github.com/go-park-mail-ru/2020_2_ExtraSafe/services/proto/board"
	protoProfile "github.com/go-park-mail-ru/2020_2_ExtraSafe/services/proto/profile"
	"github.com/tarantool/go-tarantool"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	// =============================
	tConn, err := tarantool.Connect("127.0.0.1:3301", tarantool.Opts{ User: "guest" })
	defer tConn.Close()
	if err != nil {
		fmt.Println("Connection refused")
	}

	authStorage := storage.NewStorage(tConn)

	// =============================

	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalln("cant listet port", err)
	}

	// =============================

	grcpConn, err := grpc.Dial(
		"127.0.0.1:8082",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("cant connect to grpc")
	}
	defer grcpConn.Close()

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

	server := grpc.NewServer()

	boardService := protoBoard.NewBoardClient(grpcConn)
	profileService := protoProfile.NewProfileClient(grcpConn)

	handler := auth.NewService(authStorage, profileService, boardService)

	protoAuth.RegisterAuthServer(server, handler)

	fmt.Println("starting server at :8081")
	server.Serve(lis)
}