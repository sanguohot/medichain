package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"log"
	"github.com/sanguohot/go-ethereum/crypto"
	"github.com/sanguohot/go-ethereum/crypto/ecies"
	"fmt"
	"github.com/sanguohot/go-ethereum/common/hexutil"
)


func main() {
	privateKeyEcdsa, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKeyEcdsa.Public()
	publicKeyEcdsa, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
		return
	}
	privateKeyEcies := ecies.ImportECDSA(privateKeyEcdsa)
	publicKeyEcies := ecies.ImportECDSAPublic(publicKeyEcdsa)
	fmt.Println("privateKeyEcies", privateKeyEcies)
	fmt.Println("publicKeyEcdsa", publicKeyEcdsa)
	data := "hello"
	dataByte := []byte(data)

	encryptDataByte, err := ecies.Encrypt(rand.Reader, publicKeyEcies, dataByte, nil, nil)
	ecryptDataHex := hexutil.Encode(encryptDataByte)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("encrypt %s ===> %s", data, ecryptDataHex)
	decryptDataByte, err := privateKeyEcies.Decrypt(encryptDataByte, nil, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("\ndecrypt %s ===> %s", ecryptDataHex, string(decryptDataByte))
}