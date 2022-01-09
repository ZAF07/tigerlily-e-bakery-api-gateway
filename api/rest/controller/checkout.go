package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ZAF07/tigerlily-e-bakery-api-gateway/internal/pkg/logger"
	"github.com/ZAF07/tigerlily-e-bakery-inventories/api/rpc"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type CheckoutAPI struct {
	logs logger.Logger
}

// Init the DB here (open a connection to the DB) and pass it along to service and repo layer
func NewCheckoutAPI() *CheckoutAPI {
	return &CheckoutAPI{
		logs: *logger.NewLogger(),
	}
}

func (a CheckoutAPI) Checkout(c *gin.Context) {
	a.logs.InfoLogger.Println("[CONTROLLER] Checkout API running")

	var req *rpc.CheckoutReq

	err := c.ShouldBindJSON(&req)
	if err != nil {
		fmt.Printf("error binding req struct : %+v", err)
	}
	fmt.Printf("HERE : %+v", req.CheckoutItems[0])
	ctx := context.Background()


	// Initialise a GRPC Server
	var conn * grpc.ClientConn

	// Dial the GRPC SERVER
	conn, connErr := grpc.Dial(":8001", grpc.WithInsecure())
	if connErr != nil {
		a.logs.ErrorLogger.Printf("[CONTROLLER] Error dialing GRPC server : %+v", connErr)
	}
	defer conn.Close()

	checkoutService := rpc.NewCheckoutServiceClient(conn)
	
	resp, err := checkoutService.Checkout(ctx, req)
	if err != nil {
		a.logs.ErrorLogger.Printf("[CONTROLLER] Bad response from GRPC. Don't forget to add enums proto for error codes : %+v", err)

	c.JSON(http.StatusInternalServerError,
		gin.H{
		"message": "Error checkout",
		"status": http.StatusInternalServerError,
		"data": resp,
	})
	return
	}
	
	c.JSON(http.StatusOK,
	gin.H{
		"message": "Success checkout",
		"status": http.StatusOK,
		"data": resp,
	})
}