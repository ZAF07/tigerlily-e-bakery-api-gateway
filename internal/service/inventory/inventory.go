package inventory

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ZAF07/tigerlily-e-bakery-api-gateway/internal/helper"
	"github.com/ZAF07/tigerlily-e-bakery-api-gateway/internal/manager/grpc_client"
	"github.com/ZAF07/tigerlily-e-bakery-api-gateway/internal/pkg/constants"
	"github.com/ZAF07/tigerlily-e-bakery-api-gateway/internal/pkg/logger"
	"github.com/ZAF07/tigerlily-e-bakery-inventories/api/rpc"
)

type InventoryService struct {
	GRPCClient *grpc_client.GRPCClient
	logs       logger.Logger
	hubb       *Hub
}

/*
	TODO:
	Create another NewInventoryService W/O a hub
*/
func NewInventoryService(h *Hub, grpc *grpc_client.GRPCClient) *InventoryService {
	return &InventoryService{
		GRPCClient: grpc,
		logs:       *logger.NewLogger(),
		hubb:       h,
	}
}

// GetAllInventories is the standard HTTP protocol service handler
func (srv InventoryService) GetAllInventories(ctx context.Context, req *rpc.GetAllInventoriesReq) (resp *rpc.GetAllInventoriesResp, err error) {

	srv.GRPCClient.SetStrategy(constants.INVENTORY_SERVICE, srv.GRPCClient.Conn)
	res, resErr := srv.GRPCClient.Strategy.Execute(ctx, constants.GET_INVENTORIES, req)
	if resErr != nil {
		srv.logs.ErrorLogger.Printf("[SERVICE] Error getting response from RPC via strategy : %+v", resErr)
		return nil, resErr
	}

	resp, err = helper.TransformInventoryGetResp(res)
	if err != nil {
		srv.logs.ErrorLogger.Printf("[SERVICE] Error getting transforming response to proper format : %+v", err)
	}

	// Can use this method to run Notification service later on in checkout service
	go func() {
		time.Sleep(10 * time.Second)
		fmt.Println("WAKE UP!!!!!!!!!!")
	}()

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
