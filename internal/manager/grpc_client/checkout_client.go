package grpc_client

import (
	"context"

	"github.com/Tiger-Coders/tigerlily-bff/internal/helper"
	"github.com/Tiger-Coders/tigerlily-bff/internal/pkg/constants"
	"github.com/Tiger-Coders/tigerlily-bff/internal/pkg/logger"
	"github.com/Tiger-Coders/tigerlily-payment/api/rpc"
	"google.golang.org/grpc"
)

type GRPCCheckoutClient struct {
	logs   logger.Logger
	conn   *grpc.ClientConn
	Client rpc.CheckoutServiceClient
}

// Returns a new instance of GRPCCheckoutClient{}...
func NewGRPCCheckoutClient(conn *grpc.ClientConn) *GRPCCheckoutClient {
	return &GRPCCheckoutClient{
		conn:   conn,
		Client: rpc.NewCheckoutServiceClient(conn),
		logs:   *logger.NewLogger(),
	}
}

// Execute() checks the given params and decides which payment strategy to execute
func (g GRPCCheckoutClient) Execute(ctx context.Context, _type, req interface{}) (resp interface{}, err error) {
	switch _type {
	case constants.STRIPE_CHECKOUT_SESSION:
		g.logs.InfoLogger.Printf("executing stripe checkout -> %+v", _type)
		resp, err = g.stripeCheckoutSession(ctx, req)
	case constants.TEST_STRATEGY:
		g.logs.InfoLogger.Printf("executing test checkout -> %+v", _type)
		resp, err = g.customCheckout(ctx, req)
	}
	return
}

// Stripe Checkout method for GRPC Checkout Client...
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

// Custom method for GRPC Checkout Client...
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
