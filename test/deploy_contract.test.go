package main

import (
	"log"
	"fmt"
	"medichain/deploy"
	"medichain/util"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	util.InitConfig("")
	cnsAddress := common.HexToAddress(util.Config.GetString("bcos.easy_cns.address"))
	err, address, hash := deploy.DeployContract(&cnsAddress, deploy.ContractOrgsData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(address.Hex())
	fmt.Println(hash.Hex())
}