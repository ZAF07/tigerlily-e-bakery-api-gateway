package strategy

import (
	"context"

	"github.com/Tiger-Coders/tigerlily-payment/api/rpc"
)

type StripeBasicCheckout struct{}

// NOT IN USE
func NewStripeBasicStrategy() *StripeBasicCheckout {
	return &StripeBasicCheckout{}
}

func (s StripeBasicCheckout) Checkout(ctx context.Context, req *rpc.CheckoutReq, client rpc.CheckoutServiceClient) (resp *rpc.CheckoutResp, err error) {
	resp, err = client.StripeCheckoutSession(ctx, req)
	if err != nil {
		return nil, err
	}
	return
}
