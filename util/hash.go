package util

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/sha3"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/sanguohot/medichain/zap"
)

func RlpHash(x interface{}) (h common.Hash) {
	hw := sha3.NewKeccak256()
	rlp.Encode(hw, x)
	hw.Sum(h[:0])
	return h
}
func Sha256Hash(input []byte) (common.Hash) {
	hash := sha256.New()
	hash.Write(input)
	return common.BytesToHash(hash.Sum(nil))
}
func Md5(data []byte) string {
	has := md5.Sum(data)
	return  fmt.Sprintf("%x", has) //将[]byte转成16进制
}

func PrintHexUseComma(data []byte)  {
	var str string
	//fmt.Printf("% #x", data)
	//0xed 0x88 0x60 0x27 0xbe 0x89 0x31 0x62 0x60 0xab 0x55 0x21 0x52 0x48 0xb5 0x79 0x98 0x6c 0xe2 0x16 0xce 0xab 0x58 0x0b 0x84 0x51 0x11 0xc5 0x01 0x5b 0xfd 0x52
	for _, item := range data {
		if str == "" {
			str = fmt.Sprintf("0x%02x", item)
		} else {
			str = fmt.Sprintf("%s, 0x%02x", str, item)
		}
	}
	str = fmt.Sprintf("%s,", str)
	//0xed, 0x88, 0x60, 0x27, 0xbe, 0x89, 0x31, 0x62, 0x60, 0xab, 0x55, 0x21, 0x52, 0x48, 0xb5, 0x79, 0x98, 0x6c, 0xe2, 0x16, 0xce, 0xab, 0x58, 0x0b, 0x84, 0x51, 0x11, 0xc5, 0x01, 0x5b, 0xfd, 0x52,
	zap.Sugar.Info(str)
}

func Bytes32_4Hash(input [4][32]byte) common.Hash {
	return crypto.Keccak256Hash(input[0][:], input[1][:], input[2][:], input[3][:])
}