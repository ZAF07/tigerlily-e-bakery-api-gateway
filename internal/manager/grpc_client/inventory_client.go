package grpc_client

import (
	"context"

	"github.com/ZAF07/tigerlily-e-bakery-api-gateway/internal/helper"
	"github.com/ZAF07/tigerlily-e-bakery-api-gateway/internal/pkg/constants"
	"github.com/ZAF07/tigerlily-e-bakery-api-gateway/internal/pkg/logger"
	"github.com/ZAF07/tigerlily-e-bakery-inventories/api/rpc"
	"google.golang.org/grpc"
)

type GRPCInventoryClient struct {
	logs   logger.Logger
	conn   *grpc.ClientConn
	Client rpc.InventoryServiceClient
}

// Returns a new instance of GRPCInventoryClient{}...
func NewGRPCInventoryClient(conn *grpc.ClientConn) *GRPCInventoryClient {
	return &GRPCInventoryClient{
		conn:   conn,
		Client: rpc.NewInventoryServiceClient(conn),
		logs:   *logger.NewLogger(),
	}
}

func (g GRPCInventoryClient) Execute(ctx context.Context, _type, req interface{}) (resp interface{}, err error) {
	switch _type {
	case constants.GET_INVENTORIES:
		resp, err = g.getAllInventories(ctx, req)
	}
	return
}

func (g GRPCInventoryClient) getAllInventories(ctx context.Context, req interface{}) (resp *rpc.GetAllInventoriesResp, err error) {
	defer g.conn.Close()

	r, tErr := helper.TransformInventoryGetReq(req)
	if tErr != nil {
		return nil, tErr
	}

	resp, err = g.Client.GetAllInventories(ctx, r)
	if err != nil {
		g.logs.ErrorLogger.Printf("[GRPC_MANAGER] Error receiving response from Inventory Service: %+v", err)
	}
	return
}
