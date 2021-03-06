package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/sha3"
)

func main() {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	//privateKey, err := crypto.HexToECDSA("0c31d730db5e4c8e75364cbaab0db79346bf7b9aca355381c783edd1bc479135")
	//if err != nil {
	//	log.Fatal(err)
	//}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println("私钥", len(hexutil.Encode(privateKeyBytes)[2:]), hexutil.Encode(privateKeyBytes)[2:]) // 0xfad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("公钥", len(hexutil.Encode(publicKeyBytes)[4:]), hexutil.Encode(publicKeyBytes)[4:]) // 0x049a7df67f79246283fdc93af76d4f8cdd62c4886e8cd870944e817dd0b97934fdd7719d0810951e03418205868a5c1b40b192451367f28e0088dd75e15de40c05
	fmt.Printf("公钥[\"0x%s\",\"0x%s\"]\n", hexutil.Encode(publicKeyBytes)[4:68], hexutil.Encode(publicKeyBytes)[68:])

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("地址", address) // 0x96216849c49358B10257cb55b28eA603c874b05E

	hash := sha3.NewKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Println("校验和地址", hexutil.Encode(hash.Sum(nil)[12:])) // 0x96216849c49358b10257cb55b28ea603c874b05e
}