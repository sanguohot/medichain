package main

import (
	"fmt"
	"log"
	"encoding/hex"

	"medichain/util"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

func main() {
	signHex := "0666b163dcaf5264fc5574edf800a5c8638ad7a5e263a96c4f74202944b9a1d96bc72958dac5bf3c5bc4521fc9f57dac6e2147105ae3d20be4fdc45e99f9409100"
	signBytes, err := hex.DecodeString(signHex)
	if err != nil {
		log.Fatal(err)
	}
	r, s, v := util.SigRSV(signBytes)
	fmt.Println("r =>", hexutil.Encode(r[:])[:])
	fmt.Println("s =>",hexutil.Encode(s[:])[:])
	fmt.Println("v =>", v)
}