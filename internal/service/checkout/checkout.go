package checkout

import (
	"context"

	"github.com/ZAF07/tigerlily-e-bakery-api-gateway/internal/pkg/logger"
	"github.com/ZAF07/tigerlily-e-bakery-api-gateway/internal/strategy"
	"github.com/ZAF07/tigerlily-e-bakery-payment/api/rpc"
	"google.golang.org/grpc"
)

type CheckoutService struct {
	logs logger.Logger
	strategy strategy.Strategy
}

func NewCheckoutService(s strategy.Strategy) *CheckoutService {
	return &CheckoutService{
		logs: *logger.NewLogger(),
		strategy: s,
	}
}

func (srv CheckoutService) Checkout(ctx context.Context, req *rpc.CheckoutReq) (resp *rpc.CheckoutResp, err error) {
	// Initialise a GRPC Server
	var conn * grpc.ClientConn
	
	// Dial the GRPC SERVER
	conn, connErr := grpc.Dial(":8001", grpc.WithInsecure())
	if connErr != nil {
		srv.logs.ErrorLogger.Printf("[CONTROLLER] Error dialing GRPC server : %+v", connErr)
	}
	defer conn.Close()
	
	GRPCcheckoutService := rpc.NewCheckoutServiceClient(conn)
	resp, err = srv.strategy.Checkout(ctx, req, GRPCcheckoutService)
	if err != nil {
		srv.logs.ErrorLogger.Printf("[CONTROLLER] Bad response from GRPC. Don't forget to add enums proto for error codes : %+v", err)

		resp = &rpc.CheckoutResp{
			Success: false,
		}
		return		
	}

	resp = &rpc.CheckoutResp{
		Success: true,
		Message: "Checkout Success",
		StatusUrl: resp.StatusUrl,
	}
	return
}