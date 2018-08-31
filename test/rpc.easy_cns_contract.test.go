package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"medichain/chain"
	"medichain/etc"
	"log"
	"time"
)

func main() {
	name := etc.ContractController
	if err, hash := chain.SetAddressToCns(name, common.HexToAddress("0xd91763C4185531050fD690824Fba1d998c634823")); err != nil {
		log.Fatal(err)
	}else {
		fmt.Printf("tx sent: %s\n", hash.Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870
	}
	time.Sleep(time.Second * 1)
	if err, address := chain.GetAddressFromCns(name); err != nil {
		log.Fatal(err)
	}else {
		fmt.Printf("%s ===> %s\n", name, address.Hex())
	}
}