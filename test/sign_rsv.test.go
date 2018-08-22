package main

import (
	"fmt"
	"log"
	"encoding/hex"

	"medichain/util"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

func main() {
	signHex := "4c1e4e004b76df03ab0d85eef5ff7f888fa1c22c956abd56219825b993fdd5ed68061b7f04f7a66cba7f8a4bb2ee9b721e7cef3e882b870e733befefe125835001"
	signBytes, err := hex.DecodeString(signHex)
	if err != nil {
		log.Fatal(err)
	}
	r, s, v := util.SigRSV(signBytes)
	fmt.Println("r =>", hexutil.Encode(r[:])[:])
	fmt.Println("s =>",hexutil.Encode(s[:])[:])
	fmt.Println("v =>", v)
}