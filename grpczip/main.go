package main

import (
	"github.com/lileio/lile"
	"github.com/sago35/grpczip"
	"github.com/sago35/grpczip/grpczip/cmd"
	"github.com/sago35/grpczip/server"
	"google.golang.org/grpc"
)

func main() {
	s := &server.GrpczipServer{}

	lile.Name("grpczip")
	lile.Server(func(g *grpc.Server) {
		grpczip.RegisterGrpczipServer(g, s)
	})

	cmd.Execute()
}
