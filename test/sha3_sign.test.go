package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	//"github.com/ethereum/go-ethereum/common"
	"encoding/hex"
	"github.com/sanguohot/medichain/util"
)

func main() {
	privateKeyStr := "7aaf3e2786ff4b38f4aceb6f86ff4a3670206376087d4bd0f041f91e61412e66"
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("publicKey ===>", hexutil.Encode(publicKeyBytes))
	data := []byte("hello")
	hash := crypto.Keccak256Hash(data)
	fmt.Println("data ===>", string(data))
	fmt.Println("keccak256Hash ===>", hash.Hex())
	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hexutil.Encode(signature)[2:]) // 85c4b5350fd0c3ff39d6f0984fe32cf92fb946a2639fd54b26765528809b7c702a9ff35d165b0f1e5345a111b92f3892b0af0a7ba1770248ef1e602272ec812700
	signBytes, err := hex.DecodeString(hexutil.Encode(signature)[2:])
	if err != nil {
		log.Fatal(err)
	}
	r, s, v := util.SigRSV(signBytes)
	fmt.Println("private key ===>", privateKeyStr)
	fmt.Println("public key ===>", hexutil.Encode(publicKeyBytes)[4:])
	fmt.Println("address ===>", crypto.PubkeyToAddress(*publicKeyECDSA).Hex())
	fmt.Println("r ===>", hexutil.Encode(r[:])[:])
	fmt.Println("s ===>",hexutil.Encode(s[:])[:])
	fmt.Println("v ===>", v)
	signatureNoRecoverID := signature[:len(signature)-1] // remove recovery id
	verified := crypto.VerifySignature(publicKeyBytes, hash.Bytes(), signatureNoRecoverID)
	fmt.Println(verified) // true
}