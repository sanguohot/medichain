package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"medichain/chain"
	"medichain/etc"
	"log"
)

func main() {
	name := etc.ContractFilesData
	if err, address := chain.GetAddressFromCns(name); err != nil {
		log.Fatal(err)
	}else {
		fmt.Printf("%s ===> %s\n", name, address.Hex())
	}
	if err, hash := chain.SetAddressToCns(name, common.HexToAddress("0xE9348D3180c71aCf42183EFa07F09a84c5Cd2558")); err != nil {
		log.Fatal(err)
	}else {
		fmt.Printf("tx sent: %s\n", hash.Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870
	}
}