package controller

import (
	"context"
	"net/http"

	"github.com/ZAF07/tigerlily-e-bakery-api-gateway/internal/pkg/logger"
	"github.com/ZAF07/tigerlily-e-bakery-api-gateway/internal/service/checkout"
	"github.com/ZAF07/tigerlily-e-bakery-payment/api/rpc"
	"github.com/gin-gonic/gin"
)

type CheckoutAPI struct {
	logs logger.Logger
}

// SEPERATE LOGIC INTO SERVICE LAYER

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
		a.logs.ErrorLogger.Printf("error binding req struct : %+v", err)
	}
	
	ctx := context.Background()

	// Initialise a new service instance
	service := checkout.NewCheckoutService()

	resp, err := service.Checkout(ctx, req)
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