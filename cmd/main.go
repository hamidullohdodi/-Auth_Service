package main

import (
	"auth_service/config/logger"
	pb "auth_service/genproto/user"
	"auth_service/service"
	"auth_service/storage/postgres"
	"google.golang.org/grpc"
	"net"
)

func main() {
	//cfg := config.Load()

	logger, err := logger.New("debug", "develop", "app.log")
	if err != nil {
		logger.Fatal(err.Error())
	}

	db, err := postgres.ConnectDB()
	if err != nil {
		logger.Fatal(err.Error())
	}
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		logger.Fatal(err.Error())
	}
	server := grpc.NewServer()

	pb.RegisterUserServiceServer(server, service.NewUserService(db))

	err = server.Serve(listener)
	if err != nil {
		logger.Fatal(err.Error())
	}
}
