package main

import (
	"fmt"
	"github.com/sanguohot/go-ethereum/common"
	ethereum "github.com/sanguohot/go-ethereum"
	"context"
	"math/big"
	"github.com/sanguohot/go-ethereum/rpc"
	"encoding/json"
	"github.com/sanguohot/go-ethereum/common/hexutil"
)
func toBlockNumArg(number *big.Int) string {
	if number == nil {
		return "latest"
	}
	return hexutil.EncodeBig(number)
}
func toFilterArg(q ethereum.FilterQuery) interface{} {
	arg := map[string]interface{}{
		"fromBlock": toBlockNumArg(q.FromBlock),
		"toBlock":   toBlockNumArg(q.ToBlock),
		"address":   q.Addresses,
		"topics":    q.Topics,
	}
	if q.FromBlock == nil {
		arg["fromBlock"] = "0x0"
	}
	return arg
}
func main() {
	client, err := rpc.Dial("http://10.6.250.56:8545")
	if err != nil {
		fmt.Println("dail http://10.6.250.56:8545 ===>", err)
		return
	}
	// 0xb349Eba018bFA9d89Da90829629D39668F6653A2
	// 0xca21167a870cf8b9618d259af454c6d00b30b028
	// 0xB818715eb048286A608B5E9851877AD7A30a41A0
	contractAddress := common.HexToAddress("0xB818715eb048286A608B5E9851877AD7A30a41A0")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(640),
		ToBlock:   big.NewInt(647),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	var result json.RawMessage
	err = client.CallContext(context.Background(), &result, "eth_getLogs", toFilterArg(query))
	if err != nil {
		fmt.Println("eth_getLogs error ===>", err)
		return
	}
	fmt.Printf("eth_getLogs ===> %s", result)
}