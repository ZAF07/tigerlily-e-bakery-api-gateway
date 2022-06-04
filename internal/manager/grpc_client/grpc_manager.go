package grpc_client

import (
	"log"

	"github.com/ZAF07/tigerlily-e-bakery-api-gateway/internal/pkg/constants"
	"github.com/ZAF07/tigerlily-e-bakery-api-gateway/internal/pkg/logger"
	inventory_rpc "github.com/ZAF07/tigerlily-e-bakery-inventories/api/rpc"
	payment_rpc "github.com/ZAF07/tigerlily-e-bakery-payment/api/rpc"
	"google.golang.org/grpc"
)

type GRPCClient struct {
	InventoryReq *inventory_rpc.GetAllInventoriesReq
	PaymentReq   *payment_rpc.CheckoutReq

	Strategy GRPCClientInterface
}

func NewGRPCClient(_type string) *GRPCClient {
	client := &GRPCClient{}
	var conn *grpc.ClientConn
	conn, connErr := grpc.Dial(":8000", grpc.WithInsecure())
	if connErr != nil {
		logs := logger.NewLogger()
		logs.ErrorLogger.Printf(" [SERVICE] Cannot connect to GRPC server")
		log.Fatalf("cannot connect to GRPC server: %+v", connErr)
	}
	switch _type {
	case constants.INVENTORY_SERVICE:
		client.Strategy = NewGRPCInventoryClient(conn)
	}
	return client
}
