package checkout

import (
	"context"

	"github.com/ZAF07/tigerlily-e-bakery-api-gateway/internal/pkg/constants"
	"github.com/ZAF07/tigerlily-e-bakery-api-gateway/internal/pkg/logger"
	"github.com/ZAF07/tigerlily-e-bakery-payment/api/rpc"
	"google.golang.org/grpc"
)

type CheckoutService struct {
	logs logger.Logger
}

func NewCheckoutService() *CheckoutService {
	return &CheckoutService{
		logs: *logger.NewLogger(),
	}
}

func (srv CheckoutService) Checkout(ctx context.Context, req *rpc.CheckoutReq) (resp *rpc.CheckoutResp, err error) {
	// Initialise a GRPC Server
	var conn * grpc.ClientConn
	checkoutService := rpc.NewCheckoutServiceClient(conn)

	// Dial the GRPC SERVER
	conn, connErr := grpc.Dial(":8001", grpc.WithInsecure())
	if connErr != nil {
		srv.logs.ErrorLogger.Printf("[CONTROLLER] Error dialing GRPC server : %+v", connErr)
	}
	defer conn.Close()

	switch req.PaymentType {
		case constants.STRIPE_CHECKOUT_SESSION:
			resp, err = checkoutService.StripeCheckoutSession(ctx, req)
			if err != nil {
				srv.logs.ErrorLogger.Printf("[CONTROLLER] Bad response from GRPC. Don't forget to add enums proto for error codes : %+v", err)
		
			resp = &rpc.CheckoutResp{
				Success: false,
			}
			return		
	}
		
	}
	resp = &rpc.CheckoutResp{
		Success: false,
		Message: "Missing checkout payment provider, Please specify a payment provider",
	}
	return
}