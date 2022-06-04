package grpc_client

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ZAF07/tigerlily-e-bakery-api-gateway/internal/pkg/constants"
	"github.com/ZAF07/tigerlily-e-bakery-api-gateway/internal/pkg/logger"
	"github.com/ZAF07/tigerlily-e-bakery-inventories/api/rpc"
	"google.golang.org/grpc"
)

type GRPCInventoryClient struct {
	logs logger.Logger
	conn *grpc.ClientConn
}

func NewGRPCInventoryClient(conn *grpc.ClientConn) *GRPCInventoryClient {
	return &GRPCInventoryClient{
		logs: *logger.NewLogger(),
		conn: conn,
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

	// TODO: Put this into a helper
	r := &rpc.GetAllInventoriesReq{}
	a, e := json.Marshal(req)
	if e != nil {
		fmt.Println("err --> ", e)
	}
	er := json.Unmarshal(a, r)
	if err != nil {
		fmt.Println("Cannot unmarchal into req stub -> ", er)
	}

	client := rpc.NewInventoryServiceClient(g.conn)
	resp, err = client.GetAllInventories(ctx, r)
	if err != nil {
		g.logs.ErrorLogger.Printf("[GRPC_MANAGER] Error receiving response from Inventry Service: %+v", err)
	}
	return
}
