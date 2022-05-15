package controller

import (
	"context"
	"net/http"

	"github.com/ZAF07/tigerlily-e-bakery-api-gateway/internal/pkg/constants"
	"github.com/ZAF07/tigerlily-e-bakery-api-gateway/internal/pkg/logger"
	"github.com/ZAF07/tigerlily-e-bakery-api-gateway/internal/service/checkout"
	"github.com/ZAF07/tigerlily-e-bakery-api-gateway/internal/strategy"
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
	
	if req.PaymentType == "" {
		c.JSON(http.StatusBadRequest, 
		gin.H{
			"message": "Missing payment type",
			"status": http.StatusBadRequest,
		})
		return 
	}

	ctx := context.Background()

	var service checkout.CheckoutService
	switch req.PaymentType {
	case constants.STRIPE_CHECKOUT_SESSION:
		// Initialise a new service instance
		service = *checkout.NewCheckoutService(strategy.NewStripeBasicStrategy())
	case constants.TEST_STRATEGY:
		service = *checkout.NewCheckoutService(strategy.NewTestStrategy())
	}

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
		"message": "Success",
		"status": http.StatusOK,
		"data": resp,
	})
}