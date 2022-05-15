package strategy

import (
	"context"

	"github.com/ZAF07/tigerlily-e-bakery-payment/api/rpc"
)

type StripeCheckoutSession struct {}

func (s StripeCheckoutSession) Checkout(ctx context.Context, req*rpc.CheckoutReq, client rpc.CheckoutServiceClient) (resp *rpc.CheckoutResp, err error) {
	resp, err = client.StripeCheckoutSession(ctx, req)
	if err != nil {
		return nil, err
	}
	return
}