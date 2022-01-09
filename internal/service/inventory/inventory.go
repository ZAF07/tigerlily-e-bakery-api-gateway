package inventory

import (
	"context"
	"log"

	"github.com/ZAF07/tigerlily-e-bakery-api-gateway/internal/pkg/logger"
	"github.com/ZAF07/tigerlily-e-bakery-inventories/api/rpc"
	"google.golang.org/grpc"
)

type InventoryService struct {
	logs logger.Logger
}

func NewInventoryService() *InventoryService {
	return &InventoryService{
		logs: *logger.NewLogger(),
	}
}

func (srv InventoryService) GetAllInventories(ctx context.Context, req *rpc.GetAllInventoriesReq) (resp *rpc.GetAllInventoriesResp, err error) {

	// Create a connection instance and Dial the GRPC server
	var conn *grpc.ClientConn
	conn, connErr := grpc.Dial(":8000", grpc.WithInsecure())
	if connErr != nil {
		srv.logs.ErrorLogger.Printf(" [SERVICE] Cannot connect to GRPC server")
		log.Fatalf("cannot connect to GRPC server : %+v", connErr)
	}
	defer conn.Close()

	// Initialises a new GRPC client service stub
	inventoryService := rpc.NewInventoryServiceClient(conn)

	resp, err = inventoryService.GetAllInventories(ctx, req)
	if err != nil {
		srv.logs.ErrorLogger.Printf("[SERVICE] Error getting response from RPC server : %+v", err)
	}
	
	return 
}
