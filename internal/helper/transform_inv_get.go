package helper

import (
	"encoding/json"
	"fmt"

	"github.com/ZAF07/tigerlily-e-bakery-inventories/api/rpc"
)

// Transforms req into rpc.GetAllInventoriesResp
func TransformInventoryGetResp(res interface{}) (resp *rpc.GetAllInventoriesResp, err error) {
	resp = &rpc.GetAllInventoriesResp{}
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
func TransformInventoryGetReq(res interface{}) (req *rpc.GetAllInventoriesReq, err error) {
	req = &rpc.GetAllInventoriesReq{}
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
