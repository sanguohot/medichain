package main

import (
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"fmt"
	"go-ethereum-change/common/hexutil"
)



func main() {
	privateKey, err := crypto.HexToECDSA("8b019a2745a97fab47b48373a0e2f790df52c9d05ce8714f29ec15e02c3e9d0f")
	if err != nil {
		log.Fatal(err)
	}
	store := keystore.NewKeyStore("e:/evan/goland/src/medichain/key", keystore.LightScryptN, keystore.StandardScryptP)
	fmt.Println(store)
	pass := "123456"
	account, err := store.ImportECDSA(privateKey, pass)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(account)
	keyBytes, err := store.Export(account, pass, pass)
	if err != nil {
		log.Fatalln(err)
	}
	key, err := keystore.DecryptKey(keyBytes, pass)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(key.Address.Hex())
	fmt.Println(hexutil.Encode(crypto.FromECDSA(key.PrivateKey))[2:])
}