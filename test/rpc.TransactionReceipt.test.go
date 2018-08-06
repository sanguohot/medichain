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
	}

	receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash("0x66b3193fbb8dd91a98be0f608463f54f6e0db7fdb6af7d34a3f647ae74f341e3"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(receipt.Status) // 1
}