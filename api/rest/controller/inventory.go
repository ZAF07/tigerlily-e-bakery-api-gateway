package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"

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
		// Create a connection instance and Dial the GRPC server
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot connect to GRPC client : %+v", err)
	}
	defer conn.Close()

	// Initialises a new GRPC client service stub
	inventoryService := rpc.NewInventoryServiceClient(conn)

	// Construct the request body to pass in GRPC Service method
	req := &rpc.GetAllInventoriesReq{
	Limit: 10,
	Offset: 0,
	}
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