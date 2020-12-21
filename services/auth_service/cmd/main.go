package main

import (
	"fmt"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/auth_service/internal/service"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/auth_service/internal/sessionsStorage"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/proto/auth"
	protoBoard "github.com/go-park-mail-ru/2020_2_ExtraSafe/services/proto/board"
	protoProfile "github.com/go-park-mail-ru/2020_2_ExtraSafe/services/proto/profile"
	"github.com/tarantool/go-tarantool"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)
/*
func init() {
//	if err := os.get
	if err := godotenv.Load("../../../config.env"); err != nil {
	//if err := godotenv.Load("config.env"); err != nil {
		log.Print("No .env file found")
	}
}
*/
func main() {
	// =============================
	userName:= os.Getenv("TABUTASK_SESSIONS_USER")
	//userName, ok := os.LookupEnv("TABUTASK_SESSIONS_USER")
	/*if !ok {
		log.Fatalf("Cannot get env parameter")
	}*/

	addr := os.Getenv("TABUTASK_SESSIONS_ADDR")
	/*addr, ok := os.LookupEnv("TABUTASK_SESSIONS_ADDR")
	if !ok {
		log.Fatalf("Cannot get env parameter")
	}
*/
	fmt.Println(addr)
	tConn, err := tarantool.Connect(addr, tarantool.Opts{ User: userName })
	if err != nil {
		fmt.Println("Connection refused", err)
		return
	}

	defer tConn.Close()

	authStorage := sessionsStorage.NewStorage(tConn)

	// =============================

	boardServiceAddr := os.Getenv("BOARDS_SERVICE_ADDR")
	profileServiceAddr := os.Getenv("PROFILE_SERVICE_ADDR")
	authServiceAddr := os.Getenv("AUTH_SERVICE_ADDR")
	/*boardServiceAddr, ok := os.LookupEnv("BOARDS_SERVICE_ADDR")
	if !ok {
		log.Fatalf("Cannot get env parameter")
	}

	profileServiceAddr, ok := os.LookupEnv("PROFILE_SERVICE_ADDR")
	if !ok {
		log.Fatalf("Cannot get env parameter")
	}

	authServiceAddr, ok := os.LookupEnv("AUTH_SERVICE_ADDR")
	if !ok {
		log.Fatalf("Cannot get env parameter")
	}
*/
	lis, err := net.Listen("tcp", authServiceAddr)
	//lis, err := net.Listen("tcp", "0.0.0.0:9081")
	if err != nil {
		log.Fatalln("cant listet port", err)
	}

	// =============================

	grpcConnProfile, err := grpc.Dial(
		profileServiceAddr,
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("cant connect to grpc")
	}
	defer grpcConnProfile.Close()

	// =============================

	grpcConnBoard, err := grpc.Dial(
		boardServiceAddr,
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("cant connect to grpc")
	}
	defer grpcConnBoard.Close()

	// =============================

	server := grpc.NewServer()

	boardService := protoBoard.NewBoardClient(grpcConnBoard)
	profileService := protoProfile.NewProfileClient(grpcConnProfile)

	handler := auth.NewService(authStorage, profileService, boardService)

	protoAuth.RegisterAuthServer(server, handler)

	fmt.Println(authServiceAddr)

	fmt.Println("starting server at : ", lis.Addr())
	err = server.Serve(lis)
	if err != nil {
		log.Fatalln("Serve auth", err)
	}
}