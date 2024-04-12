package main

import (
	"log"

	pb "github.com/magento1982/grpc-blog/api/auth_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()
	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}

	client := pb.NewAuthServiceClient(conn)

	callSayHello(client)

}
