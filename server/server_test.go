package server

import (
	"os"
	"testing"

	"google.golang.org/grpc"

	"github.com/lileio/lile"
	"github.com/sago35/grpczip"
)

var s = GrpczipServer{}
var cli grpczip.GrpczipClient

func TestMain(m *testing.M) {
	impl := func(g *grpc.Server) {
		grpczip.RegisterGrpczipServer(g, s)
	}

	gs := grpc.NewServer()
	impl(gs)

	addr, serve := lile.NewTestServer(gs)
	go serve()

	cli = grpczip.NewGrpczipClient(lile.TestConn(addr))

	os.Exit(m.Run())
}
