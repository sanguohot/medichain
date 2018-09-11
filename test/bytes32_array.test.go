package main

import (
	"github.com/ethereum/go-ethereum/common"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/sanguohot/medichain/util"
	"log"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	two := [2][32]byte{}
	two[0] = common.HexToHash("8ec34f461212f6bbfd759f400b5a80679ecd56f496148ad2c4669d9e1127965b")
	two[1] = common.HexToHash("e94cbd7230c091e5bfa310596f354d1339a122e7335bdf799f504b3c1fcb07a5")
	var one []byte
	one = util.BytesCombine(two[0][:], two[1][:])
	fmt.Println("0x8ec34f461212f6bbfd759f400b5a80679ecd56f496148ad2c4669d9e1127965be94cbd7230c091e5bfa310596f354d1339a122e7335bdf799f504b3c1fcb07a5 ===>", crypto.Keccak256Hash(one).Hex())
	fmt.Println(len(one), hexutil.Encode(one))
	three, err := util.StringToBytes32_4("广西中医一附院")
	if err != nil {
		log.Fatal(err)
	}
	tmp := []byte("广西中医一附院")
	four := util.Bytes32_4ToString(*three)
	fmt.Printf("[\"0x%s\", \"0x%s\", \"0x%s\", \"0x%s\"]\n", common.Bytes2Hex(three[0][:]), common.Bytes2Hex(three[1][:]), common.Bytes2Hex(three[2][:]), common.Bytes2Hex(three[3][:]))
	fmt.Println("广西中医一附院 ===>", crypto.Keccak256Hash(tmp).Hex())
	fmt.Println("广西中医一附院 ===>", util.Bytes32_4Hash(*three).Hex())
	fmt.Println(four, string(tmp))
}