package inventory

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Tiger-Coders/tigerlily-bff/config"
	"github.com/Tiger-Coders/tigerlily-bff/internal/helper"
	"github.com/Tiger-Coders/tigerlily-bff/internal/manager/grpc_client"
	"github.com/Tiger-Coders/tigerlily-bff/internal/pkg/constants"
	"github.com/Tiger-Coders/tigerlily-bff/internal/pkg/logger"
	rm "github.com/ZAF07/tigerlily-e-bakery-cache/redis-cache-manager"
	"github.com/ZAF07/tigerlily-e-bakery-inventories/api/rpc"
	"github.com/go-redis/redis/v9"
)

type InventoryService struct {
	GRPCClient  *grpc_client.GRPCClient
	logs        logger.Logger
	hubb        *Hub
	cache       rm.Redismanager
	inventories *config.AppConfig
}

/*
	TODO:
	Create another NewInventoryService W/O a hub
*/
func NewInventoryService(h *Hub, grpc *grpc_client.GRPCClient, r *redis.Client, invs *config.AppConfig) *InventoryService {
	return &InventoryService{
		GRPCClient:  grpc,
		logs:        *logger.NewLogger(),
		hubb:        h,
		cache:       rm.NewAdminRedisManager(r),
		inventories: invs,
	}
}

// GetAllInventories is the standard HTTP protocol service handler
func (srv InventoryService) GetAllInventories(ctx context.Context, req *rpc.GetAllInventoriesReq) (resp *rpc.GetAllInventoriesResp, err error) {

	srv.GRPCClient.SetStrategy(grpc_client.NewGRPCInventoryClient(srv.GRPCClient.Conn))
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
		// If cache is populated, if so, get items from cache
		if rErr := srv.cache.AddInventories(ctx, resp.Inventories); rErr != nil {
			srv.logs.ErrorLogger.Printf("Error adding into cache : %+v\n", rErr)
		}
		srv.logs.InfoLogger.Println("SUCCESSSSS CACHE!!!!!!!!!!!")
		time.Sleep(10 * time.Second)
		fmt.Println("WAKE UP!!!!!!!!!!")
	}()

	return
}

func (srv InventoryService) GetAllInventoriesCache(ctx context.Context) *rpc.GetAllInventoriesResp {

	srv.logs.InfoLogger.Println("Current inventory item name in local data file -----> ", srv.inventories.Inventories)
	resp, err := srv.cache.GetAllInventories(ctx, srv.inventories.Inventories)
	if err != nil {
		srv.logs.ErrorLogger.Printf(" [SERVICE] Error getting from cache library: %+v\n", err)
	}
	return resp
}

// ServeWs handles websocket requests from the peer. Upgrades protocol to websocket
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
