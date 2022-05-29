package grpc_client

import (
	"github.com/ZAF07/tigerlily-e-bakery-api-gateway/internal/pkg/constants"
	inventory_rpc "github.com/ZAF07/tigerlily-e-bakery-inventories/api/rpc"
	payment_rpc "github.com/ZAF07/tigerlily-e-bakery-payment/api/rpc"
)

type GRPCClient struct {
	InventoryReq *inventory_rpc.GetAllInventoriesReq
	PaymentReq   *payment_rpc.CheckoutReq

	Strategy GRPCClientInterface
}

func NewGRPCClient(_type string, req interface{}) *GRPCClient {
	client := &GRPCClient{}
	switch _type {
	case constants.GET_INVENTORIES:
		client.Strategy = NewGRPCInventoryClient()
	}
	return client
}
