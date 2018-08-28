package main

import (
	"fmt"
	"log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"medichain/contracts/medi" // for demo
	"math/rand"
	"time"
	"math/big"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"medichain/chain"
)

func main() {
	rand.Seed(time.Now().Unix())
	client, err := ethclient.Dial("http://10.6.250.56:8545")
	if err != nil {
		log.Fatal(err)
	}
	privateKey, err := crypto.HexToECDSA("7aaf3e2786ff4b38f4aceb6f86ff4a3670206376087d4bd0f041f91e61412e66")
	if err != nil {
		log.Fatal("crypto.HexToECDSA error ===>", err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	toAddress := common.HexToAddress("0x5c249bd450Ce5f7f61bf9e98877394792DFd589E")
	auth := bind.NewKeyedTransactor(privateKey) // 设置私钥和签名方法
	auth.Nonce = big.NewInt(rand.Int63n(100000000)) // 也是fisco-bcos的randomid字段，这里按照web3.js计算随机数
	auth.Value = nil
	auth.GasLimit = uint64(1000000)
	//auth.BlockLimit = uint64(740) // 区别go-etherneum，似乎是fisco-bcos判断了范围？
	auth.GasPrice = nil
	auth.From = fromAddress
	instance, err := medi.NewEasyCns(toAddress, client)
	if err != nil {
		log.Fatal("medi.NewEasyCns error ===>", err)
		return
	}
	addr, _ := instance.Get(nil, "test")
	fmt.Println(addr.Hex())
	tx, err := instance.Set(auth, "test", common.HexToAddress("0x17eEA60EfADaBE8616FCCbA1Bc3fa0B154407C06"))
	if err != nil {
		log.Fatal("medi.NewEasyCns", err)
		return
	}
	fmt.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870


	if err, hash := chain.SetAddressToCns("test", common.HexToAddress("0x98fB1108eBC0FA3A65d3C22eCB34e1F075f9ab5D")); err != nil {
		fmt.Println(err)
	}else {
		fmt.Printf("\ntx sent: %s", hash.Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870
	}
}