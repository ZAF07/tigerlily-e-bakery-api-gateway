package strategy

import (
	"context"

	"github.com/Tiger-Coders/tigerlily-payment/api/rpc"
)

// NOT IN USE
type Strategy interface {
	Checkout(ctx context.Context, req *rpc.CheckoutReq, client rpc.CheckoutServiceClient) (resp *rpc.CheckoutResp, err error)
}
