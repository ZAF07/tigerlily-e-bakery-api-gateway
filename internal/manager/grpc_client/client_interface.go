package grpc_client

import (
	"context"
)

type GRPCClientInterface interface {
	Execute(ctx context.Context, _type, req interface{}) (resp interface{}, err error)
}
