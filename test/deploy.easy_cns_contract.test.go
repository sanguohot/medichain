package main

import (
	"log"
	"fmt"
	"medichain/deploy"
	"medichain/util"
)

func main() {
	util.InitConfig("../etc/config.json")
	err, address, hash := deploy.DeployEasyCns()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(address.Hex())
	fmt.Println(hash.Hex())
}