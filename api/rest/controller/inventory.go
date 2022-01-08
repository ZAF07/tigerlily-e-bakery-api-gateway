package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ZAF07/tigerlily-e-bakery-api-gateway/internal/pkg/logger"
	"github.com/ZAF07/tigerlily-e-bakery-inventories/api/rpc"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type InventoryApi struct {
	logs logger.Logger
}

func NewInventoryAPI() *InventoryApi {
	return &InventoryApi{
		logs: *logger.NewLogger(),
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

	// Create a connection instance and Dial the GRPC server
	var conn *grpc.ClientConn
	conn, connErr := grpc.Dial(":8000", grpc.WithInsecure())
	if connErr != nil {
		log.Fatalf("cannot connect to GRPC client : %+v", connErr)
	}
	defer conn.Close()

	// Initialises a new GRPC client service stub
	inventoryService := rpc.NewInventoryServiceClient(conn)

	// Construct the request body to pass in GRPC Service method
	req := &rpc.GetAllInventoriesReq{
	Limit: int32(limit),
	Offset: int32(offset),
	}

	// Create an empty context to pass to the service layer (can pass metadata via this channel)
	ctx := context.Background()

	// Invoke the GRPC Service method and wait for response (Unary)
	resp, rErr := inventoryService.GetAllInventories(ctx, req)
	if rErr != nil {
		log.Fatalf("bad response : %+v", rErr)
	}

		c.JSON(http.StatusOK, gin.H{
			"message": "Success!!",
			"status": http.StatusOK,
			"data": resp,
		})
	
	fmt.Println(resp)
}