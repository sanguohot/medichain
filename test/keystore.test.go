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
	privateKey, err := crypto.HexToECDSA("c2b0f36af224863875901b21df43fdaf45743cbfb398ed36f4d70aed809df168")
	if err != nil {
		log.Fatal(err)
	}
	store := keystore.NewKeyStore("./keystore", keystore.LightScryptN, keystore.StandardScryptP)
	fmt.Println(store)
	pass := "123456"
	account, err := store.Find(accounts.Account{Address: common.HexToAddress("0x646829bFA80580b07767b796B4055DB9BDf148b5")})
	if err != nil {
		//log.Fatalln(err)
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