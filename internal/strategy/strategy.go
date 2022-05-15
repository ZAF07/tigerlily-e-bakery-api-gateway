package strategy

import (
	"context"

	"github.com/ZAF07/tigerlily-e-bakery-payment/api/rpc"
)

type Strategy interface {
	Checkout(ctx context.Context, req *rpc.CheckoutReq, client rpc.CheckoutServiceClient) (resp *rpc.CheckoutResp, err error)
}