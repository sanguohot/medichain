package util

import (
	"github.com/ethereum/go-ethereum/crypto/sha3"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/common"
	"crypto/md5"
	"fmt"
)

func RlpHash(x interface{}) (h common.Hash) {
	hw := sha3.NewKeccak256()
	rlp.Encode(hw, x)
	hw.Sum(h[:0])
	return h
}

func Md5(data []byte) string {
	has := md5.Sum(data)
	return  fmt.Sprintf("%x", has) //将[]byte转成16进制
}