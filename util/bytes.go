package util

import (
	"bytes"
	"strings"
)

func BytesCombine(pBytes ...[]byte) []byte {
	return bytes.Join(pBytes, []byte(""))
}

func StringToBytes32_4(input string) (*[4][32]byte, error) {
	if input == "" {
		return nil, ErrParamStringShouldNotNil
	}
	bytes := []byte(input)
	if len(bytes) > 128 {
		return nil, ErrParamStringShouldNotMoreThan128Bytes
	}
	bytes32_4 := [4][32]byte{}
	copy(bytes32_4[0][:], bytes[0:32])
	if len(bytes) > 32 {
		copy(bytes32_4[1][:], bytes[32:])
	}
	if len(bytes) > 64 {
		copy(bytes32_4[2][:], bytes[64:])
	}
	if len(bytes) > 96 {
		copy(bytes32_4[3][:], bytes[96:])
	}
	return &bytes32_4, nil
}

func Bytes32_4ToString(input [4][32]byte) string {
	b := BytesCombine(input[0][:], input[1][:], input[2][:], input[3][:])
	// 末尾0是不需要的，直接从bytes trim 0，不要转为字符串后trim space，因为space值为32而不是0，trim不掉
	tb := bytes.TrimRight(b, "\x00")
	return strings.TrimSpace(string(tb))
}
