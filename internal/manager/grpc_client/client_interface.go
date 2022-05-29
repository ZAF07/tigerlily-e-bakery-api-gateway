package grpc_client

import (
	"context"

	"github.com/ZAF07/tigerlily-e-bakery-inventories/api/rpc"
)

type GRPCClientInterface interface {
	GRPCClient()
	GetAllInventories(ctx context.Context, req *rpc.GetAllInventoriesReq) (resp *rpc.GetAllInventoriesResp, err error)
}
