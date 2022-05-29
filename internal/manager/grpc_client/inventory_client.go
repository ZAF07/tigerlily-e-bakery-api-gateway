package grpc_client

import (
	"context"
	"log"

	"github.com/ZAF07/tigerlily-e-bakery-api-gateway/internal/pkg/logger"
	"github.com/ZAF07/tigerlily-e-bakery-inventories/api/rpc"
	"google.golang.org/grpc"
)

type GRPCInventoryClient struct {
	logs logger.Logger
}

func NewGRPCInventoryClient() *GRPCInventoryClient {
	return &GRPCInventoryClient{
		logs: *logger.NewLogger(),
	}
}

func (g GRPCInventoryClient) GetAllInventories(ctx context.Context, req *rpc.GetAllInventoriesReq) (resp *rpc.GetAllInventoriesResp, err error) {
	var conn *grpc.ClientConn
	conn, connErr := grpc.Dial(":8000", grpc.WithInsecure())
	if connErr != nil {
		g.logs.ErrorLogger.Printf(" [SERVICE] Cannot connect to GRPC server")
		log.Fatalf("cannot connect to GRPC server: %+v", connErr)
	}
	defer conn.Close()

	client := rpc.NewInventoryServiceClient(conn)
	resp, err = client.GetAllInventories(ctx, req)
	if err != nil {
		g.logs.ErrorLogger.Printf("[GRPC_MANAGER] Error receiving response from Inventry Service: %+v", err)
	}
	return
}

func (g GRPCInventoryClient) GRPCClient() {}
