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
		return
	}
	// 0xb349Eba018bFA9d89Da90829629D39668F6653A2
	// 0xca21167a870cf8b9618d259af454c6d00b30b028
	// 0xB818715eb048286A608B5E9851877AD7A30a41A0
	address := common.HexToAddress("0xb349Eba018bFA9d89Da90829629D39668F6653A2")
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