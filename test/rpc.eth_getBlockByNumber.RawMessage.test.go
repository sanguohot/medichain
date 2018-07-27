package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/rpc"
	"encoding/json"
	"github.com/ethereum/go-ethereum/core/types"
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
	json.Unmarshal(result, &header)
	fmt.Printf("\nheader ===> %s", header.ParentHash.String())
	fmt.Printf("\nheader ===> %s", header.Coinbase.String())
}