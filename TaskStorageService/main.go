package main

import (
	"TaskStorageService/initializers"
	"TaskStorageService/proto/server"
	"TaskStorageService/proto/taskpb"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectRedis()
}

func main() {
	db := initializers.ConnectToDB()
	initializers.SyncDatabase(db)
	// Создаем gRPC-сервер
	grpcServer := grpc.NewServer()
	taskServer := &server.TaskServer{DB: db}

	taskpb.RegisterTaskServiceServer(grpcServer, taskServer)
	port := ":" + os.Getenv("GRPC_PORT")
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}

	fmt.Println("gRPC-сервер запущен на порту: ", port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Ошибка запуска gRPC: %v", err)
	}
}
