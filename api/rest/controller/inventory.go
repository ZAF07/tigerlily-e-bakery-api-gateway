package controller

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/ZAF07/tigerlily-e-bakery-api-gateway/internal/pkg/logger"
	"github.com/ZAF07/tigerlily-e-bakery-api-gateway/internal/service/inventory"
	"github.com/ZAF07/tigerlily-e-bakery-inventories/api/rpc"

	"github.com/gin-gonic/gin"
)

type InventoryApi struct {
	logs logger.Logger
	hubb *inventory.Hub
}

func NewInventoryAPI(h *inventory.Hub) *InventoryApi {
	return &InventoryApi{
		logs: *logger.NewLogger(),
		hubb: h,
	}
}

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

	// Check if request is synchronus or asynchronus (is synchronus, use websocket, else GRPC)

	// Construct the request body to pass in GRPC Service method
	req := &rpc.GetAllInventoriesReq{
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	// Create an empty context to pass to the service layer (can pass metadata via this channel)
	ctx := context.Background()

	service := inventory.NewInventoryService(inventory.NewHub())

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

// WsInventory is the Websocket protocol service handler
func (controller InventoryApi) WsInventory(c *gin.Context) {
	service := inventory.NewInventoryService(controller.hubb)
	service.ServeWs(c.Writer, c.Request)
}
