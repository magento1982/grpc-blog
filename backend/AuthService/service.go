package main

import (
	"log"
	"net"

	pb "github.com/magento1982/grpc-blog/api/auth_service"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type authServiceServer struct {
	pb.AuthServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Failed to the start server %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, &authServiceServer{})

	log.Printf("The server start with port %s", port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to start with %v", err)
	}
}
