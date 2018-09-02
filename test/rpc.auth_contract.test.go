package main

import (
	"fmt"
	"medichain/chain"
	"log"
	"github.com/ethereum/go-ethereum/common"
	"medichain/util"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	hash := common.HexToHash("0x47173285a8d7341e5e972fc677286384f802f8ef42a5ec5f03bbfa254cb01fad")
	prefixStr := fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(hash))
	fmt.Println(prefixStr)
	prefix := []byte(prefixStr)
	address := common.HexToAddress("0x1d306aDe962d9f20590be45e68DB6246da73811D")
	password := "123456"
	signBytes, err := util.SignWithHash(hash, address, password)
	if err != nil {
		log.Fatal(err)
	}
	r, s, v := util.SigRSV(signBytes)
	r = common.HexToHash("0x08c568ded6a3bccf7a56f8975b638f54ca3cac82cb421687f5314fd105df7ba5")
	s = common.HexToHash("0x5a1711a7d9a0a79c3fbfcc4b54e762eed5851d2fb46e3b03af8c230183626479")
	v = 27
	hashWithPrefix := crypto.Keccak256Hash(prefix, hash.Bytes())
	fmt.Println("hashWithPrefix ===>", hashWithPrefix.Hex())
	err, address01 := chain.AuthVerifySignatureWithoutPrefix(hashWithPrefix, v, r, s)
	if err != nil {
		log.Fatal(err)
	}
	err, address02 := chain.AuthVerifySignatureWithPrefix(hash, v, r, s)
	if err != nil {
		log.Fatal(err)
	}
	err, hash01 := chain.AuthGetKeccak256FromHash(hash)
	if err != nil {
		log.Fatal(err)
	}
	err, hash02 := chain.AuthGetSha3FromHash(hash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("AuthVerifySignatureWithoutPrefix ===>", address01.Hex())
	fmt.Println("AuthVerifySignatureWithPrefix ===>", address02.Hex())
	fmt.Println("AuthGetKeccak256FromHash ===>", hash01.Hex())
	fmt.Println("AuthGetSha3FromHash ===>", hash02.Hex())
	fmt.Println(util.RlpHash(util.BytesCombine(prefix, hash.Bytes())).Hex(), crypto.Keccak256Hash(prefix, hash.Bytes()).Hex())
}