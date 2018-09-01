package main

import (
	"log"
	"fmt"
	"medichain/deploy"
	"medichain/etc"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	etc.InitConfig("")
	cnsAddress := common.HexToAddress(etc.GetBcosEasyCnsAddress())
	err, address, hash := deploy.DeployContract(&cnsAddress, etc.ContractAuth)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(address.Hex())
	fmt.Println(hash.Hex())
}