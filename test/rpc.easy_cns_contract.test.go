package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"medichain/chain"
	"medichain/etc"
	"log"
)

func main() {
	name := etc.ContractController
	if err, address := chain.GetAddressFromCns(name); err != nil {
		log.Fatal(err)
	}else {
		fmt.Printf("%s ===> %s\n", name, address.Hex())
	}
	if err, hash := chain.SetAddressToCns(name, common.HexToAddress("0xA757d8d58ce9BCFDF74C2a09f716C127F2B86384")); err != nil {
		log.Fatal(err)
	}else {
		fmt.Printf("tx sent: %s\n", hash.Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870
	}
}