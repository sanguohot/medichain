package main

import (
	"context"
	"fmt"
	"log"
	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("http://10.6.250.56:8545")
	if err != nil {
		log.Fatal(err)
		return
	}

	receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash("0x91f2b940396b446d3e542031b55f1bd3944dc397c78ba4f788f640ec8612b6dc"))
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(receipt.Logs)
	fmt.Println(receipt.Status)
	fmt.Println(receipt.TxHash.Hex())
}