package grpc_client

import (
	"context"

	"github.com/ZAF07/tigerlily-e-bakery-api-gateway/internal/helper"
	"github.com/ZAF07/tigerlily-e-bakery-api-gateway/internal/pkg/constants"
	"github.com/ZAF07/tigerlily-e-bakery-api-gateway/internal/pkg/logger"
	"github.com/ZAF07/tigerlily-e-bakery-payment/api/rpc"
	"google.golang.org/grpc"
)

type GRPCCheckoutClient struct {
	logs   logger.Logger
	conn   *grpc.ClientConn
	Client rpc.CheckoutServiceClient
}

// Returns a new instance of GRPCInventoryClient{}...
func NewGRPCCheckoutClient(conn *grpc.ClientConn) *GRPCCheckoutClient {
	return &GRPCCheckoutClient{
		conn:   conn,
		Client: rpc.NewCheckoutServiceClient(conn),
		logs:   *logger.NewLogger(),
	}
}

func (g GRPCCheckoutClient) Execute(ctx context.Context, _type, req interface{}) (resp interface{}, err error) {
	switch _type {
	case constants.STRIPE_CHECKOUT_SESSION:
		resp, err = g.stripeCheckoutSession(ctx, req)
	case constants.TEST_STRATEGY:
		resp, err = g.customCheckout(ctx, req)
	}
	return
}

func (g GRPCCheckoutClient) stripeCheckoutSession(ctx context.Context, req interface{}) (resp *rpc.CheckoutResp, err error) {
	defer g.conn.Close()

	r, tErr := helper.TransformCheckoutReq(req)
	if tErr != nil {
		return nil, tErr
	}

	resp, err = g.Client.StripeCheckoutSession(ctx, r)
	if err != nil {
		g.logs.ErrorLogger.Printf("[GRPC_MANAGER] Error receiving response from Checkout Service for StripeCheckoutSession: %+v", err)
	}
	return
}

func (g GRPCCheckoutClient) customCheckout(ctx context.Context, req interface{}) (resp *rpc.CheckoutResp, err error) {
	defer g.conn.Close()

	r, tErr := helper.TransformCheckoutReq(req)
	if tErr != nil {
		g.logs.ErrorLogger.Printf("[GRPC_MANAGER] Error transforming interface to rpc.CheckoutReq: %+v\n", tErr)
	}

	resp, err = g.Client.CustomCheckout(ctx, r)
	if err != nil {
		g.logs.ErrorLogger.Printf("[GRPC_MANAGER] Error receiving response from Checkout Service for CustomCheckout: %+v\n", err)
	}
	return
}
