package main

import (
	"context"
	"fmt"
	"github.com/sanguohot/go-ethereum/ethclient"
	"math/big"
	"github.com/sanguohot/go-ethereum/common/hexutil"
	"github.com/sanguohot/go-ethereum/common"
)

func main() {
	client, err := ethclient.Dial("http://10.6.250.56:8545")
	if err != nil {
		fmt.Println("dail http://10.6.250.56:8545 ===> %s", err)
		return
	}
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		fmt.Println("client.HeaderByNumber ===> %s", err)
		return
	}

	fmt.Println(header.Number.String()) // 5671744
	blockNumber := big.NewInt(628)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		fmt.Println("client.BlockByNumber ===> %s", err)
		return
	}
	fmt.Println(block.Number().Uint64())     // 5671744
	fmt.Println(block.Time().Uint64())       // 1527211625
	fmt.Println(block.Difficulty().Uint64()) // 3217000136609065
	fmt.Println(block.TxHash().Hex())          // 0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9
	// block hash计算不对
	fmt.Println(block.Hash().Hex())
	fmt.Println(hexutil.Encode(block.Extra()))
	fmt.Println(len(block.Transactions()))   // 144
	count, err := client.TransactionCount(context.Background(), common.HexToHash("0xfa9f2281e202f74c5c78228c22c89d8a9b4a39fb6e33bd57b4de677e991dc028"))
	if err != nil {
		fmt.Println("client.TransactionCount ===> %s", err)
		return
	}

	fmt.Println(count) // 144
}