package main

import (
	"fmt"
	"medichain/chain"
	"log"
	"github.com/ethereum/go-ethereum/common"
	"medichain/util"
)

func main() {
	hash := common.HexToHash("0x407851f41ef389ddc9a563186867fc3ebe4765866c73b8eed73164b54d01d23b")
	address := common.HexToAddress("0x1d306aDe962d9f20590be45e68DB6246da73811D")
	password := "123456"
	signBytes, err := util.SignWithHash(hash, address, password)
	if err != nil {
		log.Fatal(err)
	}
	r, s, v := util.SigRSV(signBytes)
	r = common.HexToHash("0x8fb8efb0b2854dd0f1ff66437cf3803f6e764b4258f83977da965cac13094b65")
	s = common.HexToHash("0x3980644f21196035b497b7955c8dde7a6f5681a03505e50355df19c3aa7168c8")
	v = 27
	err, address01 := chain.AuthVerifyWithoutPrefix(hash, v, r, s)
	if err != nil {
		log.Fatal(err)
	}
	err, address02 := chain.AuthVerifyWithPrefix(hash, v, r, s)
	fmt.Println(address01.Hex(), address02.Hex())
}