package controller

import (
	"context"
	"net/http"

	"github.com/Tiger-Coders/tigerlily-bff/config"
	"github.com/Tiger-Coders/tigerlily-bff/internal/manager/grpc_client"
	"github.com/Tiger-Coders/tigerlily-bff/internal/pkg/logger"
	"github.com/Tiger-Coders/tigerlily-bff/internal/service/checkout"
	"github.com/Tiger-Coders/tigerlily-payment/api/rpc"
	"github.com/gin-gonic/gin"
)

type CheckoutAPI struct {
	logs      logger.Logger
	appConfig *config.AppConfig
}

// Init the DB here (open a connection to the DB) and pass it along to service and repo layer
func NewCheckoutAPI(appConfig *config.AppConfig) *CheckoutAPI {
	return &CheckoutAPI{
		logs:      *logger.NewLogger(),
		appConfig: appConfig,
	}
}

func (a CheckoutAPI) Checkout(c *gin.Context) {
	a.logs.InfoLogger.Println("[CONTROLLER] Checkout API running")

	var req *rpc.CheckoutReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		a.logs.ErrorLogger.Printf("error binding req struct : %+v\n", err)
	}
	a.logs.InfoLogger.Printf("[CONTROLLER] Received request: %+v\n ", req)

	if req.PaymentType == "" {
		c.JSON(http.StatusBadRequest,
			gin.H{
				"message": "Missing payment type",
				"status":  http.StatusBadRequest,
			})
		return
	}

	/*
		TODO:
			Pass a context.Cancel to return as soon as any error occurs
	*/
	ctx := context.Background()
	grpcClient := grpc_client.NewGRPCClient(a.appConfig.PaymentServicePort)
	service := *checkout.NewCheckoutService(grpcClient)
	resp, err := service.Checkout(ctx, req)
	if err != nil {
		a.logs.ErrorLogger.Printf("[CONTROLLER] Bad response from GRPC. Don't forget to add enums proto for error codes : %+v\n", err)

		c.JSON(http.StatusInternalServerError,
			gin.H{
				"message": "Error checkout",
				"status":  http.StatusInternalServerError,
				"data":    resp,
			})

		return
	}

	c.JSON(http.StatusOK,
		gin.H{
			"message": "Success",
			"status":  http.StatusOK,
			"data":    resp,
		})
}
