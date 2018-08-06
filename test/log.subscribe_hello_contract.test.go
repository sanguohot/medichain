package main

import (
	"fmt"
	"log"
	"github.com/ethereum/go-ethereum/common"
	//"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	ethereum "github.com/ethereum/go-ethereum"
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
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
	// subscribe an event require websocket channel or ipc
	client, err := rpc.Dial("wss://10.6.250.54:8822")
	if err != nil {
		log.Fatal("ethclient.Dial ===> ", err)
		return
	}
	// 0xb349Eba018bFA9d89Da90829629D39668F6653A2
	// 0xca21167a870cf8b9618d259af454c6d00b30b028
	// 0xB818715eb048286A608B5E9851877AD7A30a41A0
	contractAddress := common.HexToAddress("0xb349Eba018bFA9d89Da90829629D39668F6653A2")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	//sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	sub, err := client.EthSubscribe(context.Background(), logs, "logs", toFilterArg(query))
	if err != nil {
		log.Fatal("client.SubscribeFilterLogs ===> ", err)
		return
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal("chan error ===> ", err)
		case vLog := <-logs:
			fmt.Println(vLog) // pointer to event log
		}
	}
}