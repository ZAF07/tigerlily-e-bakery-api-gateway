package checkout

import (
	"context"

	"github.com/Tiger-Coders/tigerlily-payment/api/rpc"
)

type CheckotService interface {
	Checkout(ctx context.Context, req *rpc.CheckoutReq) (resp *rpc.CheckoutResp, err error)
}
