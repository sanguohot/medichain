package main

import (
	"fmt"
	"log"
	"encoding/hex"

	"medichain/util"

	"github.com/sanguohot/go-ethereum/common/hexutil"
)

func main() {
	signHex := "e940084bb1940cf0ebd2a56c7b27c1b6988658a1aebccf2116ea203dc475eb8b2b7e30d2b40167fe01a35891c2b6f7181d4720135d4a459f7b2c1ee9df90a40701"
	signBytes, err := hex.DecodeString(signHex)
	if err != nil {
		log.Fatal(err)
	}
	r, s, v := util.SigRSV(signBytes)
	fmt.Println("r =>", hexutil.Encode(r[:])[2:])
	fmt.Println("s =>",hexutil.Encode(s[:])[2:])
	fmt.Println("v =>", v)
}