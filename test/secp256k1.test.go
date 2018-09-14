package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"log"
)

func main()  {
	bytes, err := hexutil.Decode("0x71b9732fa43b52d8c9085c053064b55dccc8785da7f12b3812787ee3def367ba0a8796dafb76fd5a146c226036092b7f5dd59feecb92776f7bb2cf2a48594ae201")
	if err != nil {
		log.Fatal(err)
	}
	pub, err := secp256k1.RecoverPubkey(crypto.Keccak256Hash([]byte("hello")).Bytes(), bytes)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(hexutil.Encode(pub))
}

