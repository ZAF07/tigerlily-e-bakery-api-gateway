package strategy

import (
	"context"

	"github.com/ZAF07/tigerlily-e-bakery-payment/api/rpc"
)

type StripeBasicCheckout struct {}

func NewStripeBasicStrategy() *StripeBasicCheckout {
	return &StripeBasicCheckout{}	
}

func (s StripeBasicCheckout) Checkout(ctx context.Context, req*rpc.CheckoutReq, client rpc.CheckoutServiceClient) (resp *rpc.CheckoutResp, err error) {
	resp, err = client.StripeCheckoutSession(ctx, req)
	if err != nil {
		return nil, err
	}
	return
}