package checkout

import (
	"context"

	"github.com/ZAF07/tigerlily-e-bakery-payment/api/rpc"
)

type CheckotService interface {
	Checkout(ctx context.Context, req *rpc.CheckoutReq) (resp *rpc.CheckoutResp, err error)
}
