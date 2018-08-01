package main

import (
	"fmt"
	"log"
	"github.com/sanguohot/go-ethereum/accounts/abi/bind"
	"github.com/sanguohot/go-ethereum/crypto"
	"github.com/sanguohot/go-ethereum/ethclient"
	hello "medichain/contracts/hello" // for demo
	"math/rand"
	"time"
	"math/big"
	"crypto/ecdsa"
	"github.com/sanguohot/go-ethereum/common"
)

func main() {
	rand.Seed(time.Now().Unix())
	client, err := ethclient.Dial("http://10.6.250.56:8545")
	if err != nil {
		log.Fatal(err)
		return
	}
	privateKey, err := crypto.HexToECDSA("bcec428d5205abe0f0cc8a734083908d9eb8563e31f943d760786edf42ad67dd")
	if err != nil {
		log.Fatal("crypto.HexToECDSA error ===>", err)
		return
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
		return
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	toAddress := common.HexToAddress("0xB818715eb048286A608B5E9851877AD7A30a41A0")
	auth := bind.NewKeyedTransactor(privateKey) // 设置私钥和签名方法
	auth.Nonce = big.NewInt(rand.Int63n(100000000)) // 也是fisco-bcos的randomid字段，这里按照web3.js计算随机数
	auth.Value = nil
	auth.GasLimit = uint64(1000000)
	auth.BlockLimit = uint64(740) // 区别go-etherneum，似乎是fisco-bcos判断了范围？
	auth.GasPrice = nil
	auth.From = fromAddress
	instance, err := hello.NewHello(toAddress, client)
	if err != nil {
		log.Fatal("hello.NewHello error ===>", err)
		return
	}
	result, err := instance.Speak(nil) // 调用Hello.sol 的 speak方法
	if err != nil {
		log.Fatal("instance.Speak error ===>", err)
		return
	}
	fmt.Println("\nfirst speak ===>", string(result[:])) //
	tx, err := instance.SaySomethingElse(auth, "━((*′д｀)爻(′д｀*))━!!!!") // 调用Hello.sol 的 SaySomethingElse方法
	if err != nil {
		log.Fatal("instance.SaySomethingElse", err)
		return
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870
	time.Sleep(3 * time.Second)
	result, err = instance.Speak(nil)
	if err != nil {
		log.Fatal("instance.Speak error ===>", err)
		return
	}
	fmt.Println("\nlast speak ===>", string(result[:])) // "bar"
}