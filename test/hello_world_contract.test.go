package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/sanguohot/go-ethereum/accounts/abi/bind"
	"github.com/sanguohot/go-ethereum/crypto"
	"github.com/sanguohot/go-ethereum/ethclient"

	helloWorldContracts "medichain/contracts/hello" // for demo
)

func main() {
	client, err := ethclient.Dial("http://10.6.250.56:8545")
	if err != nil {
		fmt.Println("dail http://10.6.250.56:8545 ===>", err)
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

	fmt.Println(address.Hex())   // 0x147B8eb97fD247D06C4006D269c90C1908Fb5D54
	fmt.Println(tx.Hash().Hex()) // 0xdae8ba5444eefdc99f4d45cd0c4f24056cba6a02cefbf78066ef9f4188ff7dc0

	_ = instance
}