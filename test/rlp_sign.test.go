package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"github.com/sanguohot/medichain/util"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
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
	hash := util.RlpHash(data)
	fmt.Println(hash.Hex()) // 5cf75b333ed060a53faa521b3ca125e910fa7b3ea861c9f47c38e756612d91e2

	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hexutil.Encode(signature)[2:]) // e940084bb1940cf0ebd2a56c7b27c1b6988658a1aebccf2116ea203dc475eb8b2b7e30d2b40167fe01a35891c2b6f7181d4720135d4a459f7b2c1ee9df90a40701
	r, s, v := util.SigRSV(signature)
	fmt.Println("r =>", hexutil.Encode(r[:])[:])
	fmt.Println("s =>",hexutil.Encode(s[:])[:])
	fmt.Println("v =>", v)
	signatureNoRecoverID := signature[:len(signature)-1] // remove recovery id
	verified := crypto.VerifySignature(publicKeyBytes, hash.Bytes(), signatureNoRecoverID)
	fmt.Println(verified) // true
}