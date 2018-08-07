package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/rpc"
	"encoding/json"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/common"
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
		fmt.Println("eth_getBlockByNumber error ===>", err)
		return
	}
	fmt.Printf("CallContext ===> %s", result)
	type txExtraInfo struct {
		BlockNumber *string         `json:"blockNumber,omitempty"`
		BlockHash   *common.Hash    `json:"blockHash,omitempty"`
		From        *common.Address `json:"from,omitempty"`
	}
	type rpcTransaction struct {
		tx *types.Transaction
		txExtraInfo
	}
	//func (tx *rpcTransaction) UnmarshalJSON(msg []byte) error {
	//	if err := json.Unmarshal(msg, &tx.tx); err != nil {
	//	return err
	//}
	//	return json.Unmarshal(msg, &tx.txExtraInfo)
	//}
	type rpcBlock struct {
		Hash         common.Hash      `json:"hash"`
		Transactions []rpcTransaction `json:"transactions"`
		UncleHashes  []common.Hash    `json:"uncles"`
	}
	var body rpcBlock
	var header *types.Header
	err = json.Unmarshal(result, &header)
	if err != nil {
		fmt.Println("\nparse header error ===>", err)
		return
	}
	err = json.Unmarshal(result, &body)
	if err != nil {
		fmt.Println("\nparse block error ===>", err)
		return
	}
	fmt.Printf("\nheader ===> %s", header.ParentHash.String())
	fmt.Printf("\nheader ===> %s", header.Coinbase.String())
	fmt.Printf("\nheader ===> %s", header.Number)
	fmt.Printf("\nblock ===> %s", body.Hash.String())
	fmt.Println("\nTransactions ===>", body.Transactions[0].BlockNumber)     // 5671744
}