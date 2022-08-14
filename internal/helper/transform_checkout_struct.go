package helper

import (
	"encoding/json"
	"fmt"

	"github.com/Tiger-Coders/tigerlily-payment/api/rpc"
)

// Transforms req into rpc.GetAllInventoriesResp
func TransformCheckoutresp(res interface{}) (resp *rpc.CheckoutResp, err error) {
	resp = &rpc.CheckoutResp{}
	marshal, marshalErr := json.Marshal(res)
	if marshalErr != nil {
		fmt.Println("Error marshaling res --> ", marshalErr)
		err = marshalErr
	}
	unmarshalErr := json.Unmarshal(marshal, resp)
	if err != nil {
		fmt.Println("Error unmarshaling response --> ", unmarshalErr)
		err = unmarshalErr
	}
	return
}

//  Transforms req into rpc.GetAllInventoriesReq{}
func TransformCheckoutReq(res interface{}) (req *rpc.CheckoutReq, err error) {
	req = &rpc.CheckoutReq{}
	marshal, marshalErr := json.Marshal(res)
	if marshalErr != nil {
		fmt.Println("Error marshaling res --> ", marshalErr)
		err = marshalErr
	}
	unmarshalErr := json.Unmarshal(marshal, req)
	if err != nil {
		fmt.Println("Error unmarshaling response --> ", unmarshalErr)
		err = unmarshalErr
	}
	return
}
