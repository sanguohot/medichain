package main

import (
	"context"
	"fmt"
	"github.com/sanguohot/go-ethereum/rpc"
	"encoding/json"
	"github.com/sanguohot/go-ethereum/core/types"
)

func main() {
	client, err := rpc.Dial("http://10.6.250.56:8545")
	if err != nil {
		fmt.Println("dail http://10.6.250.56:8545 ===>", err)
		return
	}
	var result json.RawMessage
	err = client.CallContext(context.Background(), &result, "eth_getBlockByNumber", "latest", true)
	if err != nil {
		fmt.Println("CallContext error ===>", err)
		return
	}
	fmt.Printf("CallContext ===> %s", result)
	var header *types.Header
	err = json.Unmarshal(result, &header)
	if err != nil {
		fmt.Println("\nCallContext error ===>", err)
		// 注意这里报错是一致的，可能某些数据类型未能得到转化
	}
	fmt.Printf("\nheader ===> %s", header.ParentHash.String())
	fmt.Printf("\nheader ===> %s", header.Coinbase.String())
	fmt.Printf("\nheader ===> %s", header.Number)
}