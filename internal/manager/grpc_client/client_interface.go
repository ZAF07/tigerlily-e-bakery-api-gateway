package grpc_client

import (
	"context"
)

type GRPCClientInterface interface {
	// Strategy method for each GRPC client. Switches between _type to decide which client strategy to run. Calls *rpc client service method
	Execute(ctx context.Context, _type, req interface{}) (resp interface{}, err error)
}
