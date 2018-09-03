package main

import (
	"log"
	//"medichain/etc"
	"github.com/ethereum/go-ethereum/common"
	"medichain/deploy"
	//"fmt"
	"medichain/etc"
	"fmt"
)

func main() {
	cnsAddress := common.HexToAddress(etc.GetBcosEasyCnsAddress())
	err, address, hash := deploy.DeployContract(&cnsAddress, etc.ContractAuth)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(address.Hex())
	fmt.Println(hash.Hex())
	deploy.DeployAllByDefaultEasyCnsAddress(false)
}