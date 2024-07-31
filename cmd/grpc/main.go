package main

import (
	"log"
	"net"

	"github.com/abyanb4/table-link/proto"
	"github.com/abyanb4/table-link/resources/db"
	"github.com/abyanb4/table-link/server"
	"google.golang.org/grpc"
)

func main() {
	db.InitDB("user=root dbname=golang sslmode=disable password=password1234")

	lis, err := net.Listen("tcp", ":33341")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterUserServiceServer(s, &server.AuthServiceServer{})

	log.Println("Server is running on port 33341")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
