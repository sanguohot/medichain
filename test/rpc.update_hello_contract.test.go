package main

import (
	"fmt"
	"log"
	"github.com/sanguohot/go-ethereum/accounts/abi/bind"
	"github.com/sanguohot/go-ethereum/crypto"
	"github.com/sanguohot/go-ethereum/common"
	"github.com/sanguohot/go-ethereum/ethclient"
	hello "medichain/contracts/hello" // for demo
	"math/rand"
	"time"
	"math/big"
)

func main() {
	rand.Seed(time.Now().Unix())
	client, err := ethclient.Dial("http://10.6.250.56:8545")
	if err != nil {
		log.Fatal(err)
	}
	privateKey, err := crypto.HexToECDSA("bcec428d5205abe0f0cc8a734083908d9eb8563e31f943d760786edf42ad67dd")
	if err != nil {
		log.Fatal("111", err)
	}

	//gasPrice, err := client.SuggestGasPrice(context.Background())
	//if err != nil {
	//	log.Fatal("333", err)
	//}
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(rand.Int63n(100000000))
	auth.Value = nil     // in wei
	auth.GasLimit = uint64(1000000) // in units
	auth.GasPrice = nil
	auth.From = common.HexToAddress("0x64fa644d2a694681bd6addd6c5e36cccd8dcdde3")

	address := common.HexToAddress("0xca21167a870cf8b9618d259af454c6d00b30b028")
	instance, err := hello.NewHello(address, client)
	if err != nil {
		log.Fatal("444", err)
	}

	tx, err := instance.SaySomethingElse(auth, "hehe")
	if err != nil {
		log.Fatal("555", err)
	}

	fmt.Printf("\ntx sent: %s", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870

	result, err := instance.Speak(nil)
	if err != nil {
		log.Fatal("666", err)
	}

	fmt.Println("\nspeak ===>", string(result[:])) // "bar"
}