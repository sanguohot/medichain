package main

import (
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
)



func main() {
	privateKey, err := crypto.HexToECDSA("7aaf3e2786ff4b38f4aceb6f86ff4a3670206376087d4bd0f041f91e61412e66")
	if err != nil {
		log.Fatal(err)
	}
	store := keystore.NewKeyStore("./keystore", keystore.LightScryptN, keystore.StandardScryptP)
	fmt.Println(store)
	pass := "123456"
	account, err := store.Find(accounts.Account{Address: common.HexToAddress("0x1d306aDe962d9f20590be45e68DB6246da73811D")})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("store.Find ===>", account)
	account, err = store.ImportECDSA(privateKey, pass)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("store.ImportECDSA ===>", account)
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