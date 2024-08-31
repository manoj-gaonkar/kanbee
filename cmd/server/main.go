package main

import (
	"log"
	"net"

	"github.com/nrssi/kanbee/internal/services"
	kbp "github.com/nrssi/kanbee/internal/services/kanban"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	addr := "localhost:50051"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen : %v", err)
	}
	grpcserver := grpc.NewServer()
	kbp.RegisterKanbanServiceServer(grpcserver, &services.KanbeeServiceServer{})
	reflection.Register(grpcserver)
	log.Println("gRPC server is running on ", addr)
	if err := grpcserver.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
