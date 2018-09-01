package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	//"github.com/ethereum/go-ethereum/common"
	"encoding/hex"
	"medichain/util"
)

func main() {
	privateKey, err := crypto.HexToECDSA("7aaf3e2786ff4b38f4aceb6f86ff4a3670206376087d4bd0f041f91e61412e66")
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	data := []byte("hello world")
	hash := crypto.Keccak256Hash(data)
	//hash = common.HexToHash("0x1c8aff950685c2ed4bc3174f3472287b56d9517b9c948127319a09a7a36deac8")
	fmt.Println(hash.Hex())
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
	fmt.Println("r =>", hexutil.Encode(r[:])[:])
	fmt.Println("s =>",hexutil.Encode(s[:])[:])
	fmt.Println("v =>", v)
	signatureNoRecoverID := signature[:len(signature)-1] // remove recovery id
	verified := crypto.VerifySignature(publicKeyBytes, hash.Bytes(), signatureNoRecoverID)
	fmt.Println(verified) // true
}