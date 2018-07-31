package main

import (
	"fmt"
	"log"
	"github.com/sanguohot/go-ethereum/common"
	"github.com/sanguohot/go-ethereum/ethclient"
	hello "medichain/contracts/hello" // for demo
)

func main() {
	client, err := ethclient.Dial("http://10.6.250.56:8545")
	if err != nil {
		log.Fatal(err)
		return
	}

	address := common.HexToAddress("0xca21167a870cf8b9618d259af454c6d00b30b028")
	instance, err := hello.NewHello(address, client)
	if err != nil {
		log.Fatal(err)
		return
	}

	result, err := instance.Speak(nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(result) // "1.0"
}