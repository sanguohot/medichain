package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	helloWorldContracts "github.com/sanguohot/medichain/contracts/hello" // for demo
)

func main() {
	client, err := ethclient.Dial("http://10.6.250.55:8546")
	if err != nil {
		fmt.Println("dail http://10.6.250.55:8546 ===>", err)
		return
	}

	privateKey, err := crypto.HexToECDSA("bcec428d5205abe0f0cc8a734083908d9eb8563e31f943d760786edf42ad67dd")
	if err != nil {
		fmt.Println(err)
		return
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("error casting public key to ECDSA")
		return
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		fmt.Println("PendingNonceAt ===>", err)
		return
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println("SuggestGasPrice ===>", err)
		return
	}
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	address, tx, instance, err := helloWorldContracts.DeployHello(auth, client)
	if err != nil {
		fmt.Println("DeployContracts ===>", err)
		return
	}

	fmt.Println(address.Hex())   // 0x5100132E59B4fe81Bc05aEED6839eF6df9DD190F
	fmt.Println(tx.Hash().Hex()) // 0x430b9ed6c77c4a4a38d24059bb52b085b22efd26e3c65c371ec10d28a80bb73d

	_ = instance
}