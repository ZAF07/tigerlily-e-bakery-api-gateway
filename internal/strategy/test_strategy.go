package strategy

import (
	"context"
	"fmt"

	"github.com/ZAF07/tigerlily-e-bakery-payment/api/rpc"
)

type TestStrategy struct{}

// NOT IN USE
func NewTestStrategy() *TestStrategy {
	return &TestStrategy{}
}

//  NOT IN USE
func (s TestStrategy) Checkout(ctx context.Context, req *rpc.CheckoutReq, client rpc.CheckoutServiceClient) (resp *rpc.CheckoutResp, err error) {
	fmt.Println("YUP WE HIT THE TEST STRATEGY")
	resp, err = client.CustomCheckout(ctx, req)
	if err != nil {
		return nil, err
	}
	return
}
