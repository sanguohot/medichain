package main

import (
	"medichain/util"
	"github.com/ethereum/go-ethereum/core/types"
	"encoding/json"
	"log"
	"fmt"
	//"github.com/ethereum/go-ethereum/rlp"
	//"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	//"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/common/hexutil"
)
var (
	headerJson = `{"parentHash":"0xde34cbcfc8b3babf2d0e9543a659e5cefd128f4fdeb3b2e285e700fe98b57b44","sha3Uncles":"0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347","miner":"0x0000000000000000000000000000000000000000","stateRoot":"0x124480dded8d5ac5b31d271460e95d14a07d92d5c1fe7d9380d0f6979f0b4c74","transactionsRoot":"0xe462d3d274b03f570890748790a0aacce8bd28984d5160e26c5bd4914688b3f0","receiptsRoot":"0xfd285caee836d754c499a6eb8e49784429f9c207e61f8f5cc6c351c9138d33bd","logsBloom":"0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000101000000000000000000000008000000000000000000000000000000000000000800000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","difficulty":"0x1","number":"0x2ea","gasLimit":"0x77359400","gasUsed":"0x1ca3f6f","timestamp":"0x165853e549b","extraData":"0xd98097312e312e302b2b302d524c696e75782f672b2b2f496e74","mixHash":"0x0000000000000000000000000000000000000000000000000000000000000000","nonce":"0x0000000000000000","hash":"0x8616976fb6d91992aa577b6cd222ab6fd5e55d5885b384b8d21a93c579fe1150"}`
	nodeList = []string{
		"f8c6b840d93d0e32eb100548afc4648783cb3f4292bf5533ac04305ae0aa1093c608be955e4db5974d3acca7c55afbd60ebe5765671201dd2c2fc2c0f2a75ec2a15b0205",
		"b840d7eedefce10828f70ec165cc4c1a5dd6bf52ae85f67237f033e97b35c5b6043886d612a4f669f453d64e662b53b493c17193595324b4a7f450509e9935d2ce84",
		"b8407709710c5e07db7261b2dedb0cc3d0e5b294cba22b5b6dfcfca3af7e744c061d3227a726850bd606e7d62e29cb4f8a4259043a68bb18f9977fe2fdc28fc5fd4a",
		}
)

func main() {
	var h types.Header
	if err := json.Unmarshal([]byte(headerJson), &h); err != nil {
		log.Fatal(err)
	}
	//var tmp [3][]byte
	//var dst []byte
	//for i, item := range nodeList{
	//	src, err := rlp.EncodeToBytes(item)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	fmt.Println("item ===>", item, hexutil.Encode(src))
	//
	//	tmp[i] = src
	//}
	//dst = util.BytesCombine(tmp[0], tmp[1], tmp[2])
	//fmt.Println("rlp three and combine ===>", hexutil.Encode(dst))
	//dst, err := rlp.EncodeToBytes(dst)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("rlp last ===>", hexutil.Encode(dst))
	nodeByte0, _ := hexutil.Decode(nodeList[0])
	nodeByte1, _ := hexutil.Decode(nodeList[1])
	nodeByte2, _ := hexutil.Decode(nodeList[2])
	genIndexBig := big.NewInt(0x1)
	hash := util.RlpHash([]interface{}{
		h.ParentHash,
		h.UncleHash,
		h.Coinbase,
		h.Root,
		h.TxHash,
		h.ReceiptHash,
		h.Bloom,
		h.Difficulty,
		h.Number,
		h.GasLimit,
		h.GasUsed,
		h.Time,
		h.Extra,
		genIndexBig,
		//dst,
		//h.MixDigest,
		//h.Nonce,
		nodeByte0,
		nodeByte1,
		nodeByte2,
	})
	fmt.Println(hash.Hex())
}