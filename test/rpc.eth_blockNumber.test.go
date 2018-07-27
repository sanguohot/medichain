package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/rpc"
)

func main() {
	client, err := rpc.Dial("http://10.6.250.57:8547")
	if err != nil {
		fmt.Println("dail http://10.6.250.57:8547 ===>", err)
		return
	}
	var result *rpc.BlockNumber
	err = client.CallContext(context.Background(), &result, "eth_blockNumber", "latest", true)
	if err != nil {
		fmt.Println("CallContext ===>", err)
		return
	}
	fmt.Printf("CallContext ===> %s", result)
}