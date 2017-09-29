package server

import (
	"testing"

	"github.com/sago35/grpczip"
	"github.com/stretchr/testify/assert"
	context "golang.org/x/net/context"
)

func TestGrpczip(t *testing.T) {
	ctx := context.Background()
	req := &grpczip.Request{}

	res, err := cli.Grpczip(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
