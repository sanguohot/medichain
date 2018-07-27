package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	hello "medichain/contracts/hello" // for demo
)

func main() {
	client, err := ethclient.Dial("http://10.6.250.56:8545")
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0x147B8eb97fD247D06C4006D269c90C1908Fb5D54")
	instance, err := hello.NewHello(address, client)
	if err != nil {
		log.Fatal(err)
	}

	result, err := instance.Speak(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result) // "1.0"
}