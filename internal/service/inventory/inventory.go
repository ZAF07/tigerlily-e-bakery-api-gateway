package inventory

import (
	"context"
	"log"
	"net/http"

	"github.com/ZAF07/tigerlily-e-bakery-api-gateway/internal/helper"
	"github.com/ZAF07/tigerlily-e-bakery-api-gateway/internal/manager/grpc_client"
	"github.com/ZAF07/tigerlily-e-bakery-api-gateway/internal/pkg/constants"
	"github.com/ZAF07/tigerlily-e-bakery-api-gateway/internal/pkg/logger"
	"github.com/ZAF07/tigerlily-e-bakery-inventories/api/rpc"
)

type InventoryService struct {
	logs logger.Logger
	hubb *Hub
}

/*
	TODO:
	Create another NewInventoryService W/O a hub
*/
func NewInventoryService(h *Hub) *InventoryService {
	return &InventoryService{
		logs: *logger.NewLogger(),
		hubb: h,
	}
}

// GetAllInventories is the standard HTTP protocol service handler
func (srv InventoryService) GetAllInventories(ctx context.Context, req *rpc.GetAllInventoriesReq) (resp *rpc.GetAllInventoriesResp, err error) {

	grpcClient := grpc_client.NewGRPCClient(constants.INVENTORY_SERVICE)
	res, grpcErr := grpcClient.Strategy.Execute(ctx, constants.GET_INVENTORIES, req)
	if grpcErr != nil {
		srv.logs.ErrorLogger.Printf("[SERVICE] Error getting response from RPC via strategy : %+v", grpcErr)
		return nil, grpcErr
	}

	resp, err = helper.TransformInventoryGetResp(res)
	if err != nil {
		srv.logs.ErrorLogger.Printf("[SERVICE] Error getting transforming response to proper format : %+v", err)
	}

	return
}

// serveWs handles websocket requests from the peer.
func (srv InventoryService) ServeWs(w http.ResponseWriter, r *http.Request) {

	// Here i am upgrading the HTTP connection to a Websocket Protocol connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// Creating a new client connection data structure
	client := &Client{Hub: srv.hubb, Conn: conn, Send: make(chan []byte, 256)}
	// Register a new connection to the hub
	client.Hub.Register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.WritePump()
	go client.ReadPump()
}
