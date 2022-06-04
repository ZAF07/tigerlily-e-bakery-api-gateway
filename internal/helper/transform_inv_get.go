package helper

import (
	"encoding/json"
	"fmt"

	"github.com/ZAF07/tigerlily-e-bakery-inventories/api/rpc"
)

func TransformInventoryGetResp(res interface{}) (resp *rpc.GetAllInventoriesResp, err error) {
	resp = &rpc.GetAllInventoriesResp{}
	m, me := json.Marshal(res)
	if me != nil {
		fmt.Println("Error marshaling res --> ", me)
		err = me
	}
	une := json.Unmarshal(m, resp)
	if err != nil {
		fmt.Println("Error unmarshaling response --> ", une)
		err = une
	}
	return
}
