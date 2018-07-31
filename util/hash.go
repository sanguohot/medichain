package util

import (
	"github.com/sanguohot/go-ethereum/crypto/sha3"
	"github.com/sanguohot/go-ethereum/rlp"
	"github.com/sanguohot/go-ethereum/common"
)

func RlpHash(x interface{}) (h common.Hash) {
	hw := sha3.NewKeccak256()
	rlp.Encode(hw, x)
	hw.Sum(h[:0])
	return h
}
