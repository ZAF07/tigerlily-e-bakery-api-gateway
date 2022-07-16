package controller

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/ZAF07/tigerlily-e-bakery-api-gateway/config"
	"github.com/ZAF07/tigerlily-e-bakery-api-gateway/internal/cache"
	"github.com/ZAF07/tigerlily-e-bakery-api-gateway/internal/manager/grpc_client"
	"github.com/ZAF07/tigerlily-e-bakery-api-gateway/internal/pkg/constants"
	"github.com/ZAF07/tigerlily-e-bakery-api-gateway/internal/pkg/logger"
	"github.com/ZAF07/tigerlily-e-bakery-api-gateway/internal/service/inventory"
	"github.com/ZAF07/tigerlily-e-bakery-inventories/api/rpc"
	"github.com/go-redis/redis/v9"

	"github.com/gin-gonic/gin"
)

type InventoryApi struct {
	logs        logger.Logger
	hubb        *inventory.Hub
	rdb         *redis.Client
	Inventories *config.AppConfig
}

// Returns a new instance of InventoryAPI{}
func NewInventoryAPI(h *inventory.Hub, invs *config.AppConfig) *InventoryApi {
	return &InventoryApi{
		logs:        *logger.NewLogger(),
		hubb:        h,
		rdb:         cache.NewRedisCache(),
		Inventories: invs,
	}
}

// Handler for GET domain/inventory
func (controller InventoryApi) GetAllInventories(c *gin.Context) {

	// Parse the request data
	limit, limitErr := strconv.Atoi(c.Query("limit"))
	if limitErr != nil {
		controller.logs.ErrorLogger.Printf("[CONTROLLER] Error converting limit param into integer : +%v", limitErr)
		log.Fatalf("[CONTROLLER] Error converting limit param into integer : %+v", limitErr)
	}
	offset, offsetErr := strconv.Atoi(c.Query("offset"))
	if offsetErr != nil {
		controller.logs.ErrorLogger.Printf("[CONTROLLER] Error converting offset param into integer : %+v", offsetErr)
		log.Fatalf("Error converting offset params into integer :%v", offsetErr)
	}

	// Construct the request body to pass in GRPC Service method
	req := &rpc.GetAllInventoriesReq{
		Limit:  int32(limit),
		Offset: int32(offset),
	}
	controller.logs.InfoLogger.Printf("[CONTROLLER] Received GET Inventories request with these params Limit: %+v Offset: %+v\n", req.Limit, req.Offset)

	/*
		TODO:
		Refactor to use context cancel if request fails in propagation
	*/
	// Create an empty context to pass to the service layer (can pass metadata via this channel)
	ctx := context.Background()
	grpcClient := grpc_client.NewGRPCClient(controller.Inventories.InventoryServicePort)
	service := inventory.NewInventoryService(&inventory.Hub{}, grpcClient, controller.rdb, controller.Inventories)

	resp, err := service.GetAllInventories(ctx, req)
	if err != nil {
		controller.logs.ErrorLogger.Printf("[CONTROLLER] Error getting response : %+v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "We are facing some trouble, please try again",
			"status":  http.StatusInternalServerError,
			"data":    "",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success!!",
		"status":  http.StatusOK,
		"data":    resp,
	})
}

func (controller InventoryApi) GetAllInventoriesCache(c *gin.Context) {
	controller.logs.InfoLogger.Println("Request for GetAllInventoriesCache")
	service := inventory.NewInventoryService(controller.hubb, grpc_client.NewGRPCClient(constants.INVENTORY_PORT), controller.rdb, controller.Inventories)

	ctx := context.Background()
	resp := service.GetAllInventoriesCache(ctx)

	c.JSON(http.StatusOK, gin.H{
		"message": "Success from cache",
		"status":  http.StatusOK,
		"data":    resp,
	})
}

// WsInventory is the Websocket protocol service handler
func (controller InventoryApi) WsInventory(c *gin.Context) {
	service := inventory.NewInventoryService(controller.hubb, &grpc_client.GRPCClient{}, controller.rdb, controller.Inventories)
	service.ServeWs(c.Writer, c.Request)
}
