package main

import (
	"log"
	"net"

	"github.com/sago35/grpczip"
	"github.com/sago35/grpczip/server"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer(grpc.MaxMsgSize(0xFFFFFFFF))

	grpczip.RegisterGrpczipServer(srv, &server.GrpczipServer{})
	srv.Serve(lis)
}
