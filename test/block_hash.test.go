package main

import (
	"medichain/util"
	"github.com/ethereum/go-ethereum/core/types"
	"encoding/json"
	"log"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
)
var (
	headerJson = `{"parentHash":"0xadeebda50b9265870bb905e67049eadaaffcfbcd29c70f04df5950c3879437e6","sha3Uncles":"0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347","miner":"0x0000000000000000000000000000000000000000","stateRoot":"0x24f77b012e239413a11541d4c191d9032859a5c47d866e40d75468030f31d300","transactionsRoot":"0x4cf3279d19ee56a6fc09e6c19612ab314c989a9e0bf447f8bca490a3bb7ebfea","receiptsRoot":"0x5d3f260b063b9093f1179f9c13d8aa7e8c947f5596868217d6b6bfe25cd6de5b","logsBloom":"0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000101000000000000000000000008000000000000000000000000000000000000000800000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","difficulty":"0x1","number":"0x2e9","gasLimit":"0x77359400","gasUsed":"0x1ca3f6f","timestamp":"0x165844414a7","extraData":"0xd98097312e312e302b2b302d524c696e75782f672b2b2f496e74","mixHash":"0x0000000000000000000000000000000000000000000000000000000000000000","nonce":"0x0000000000000000","hash":"0x8750e7ec375fb0ebafbcd180dd013aba616e78d6d7bd8af29e4b5028a1bfafce"}`
	nodeList = []string{
		"0xd93d0e32eb100548afc4648783cb3f4292bf5533ac04305ae0aa1093c608be955e4db5974d3acca7c55afbd60ebe5765671201dd2c2fc2c0f2a75ec2a15b0205",
		"0xd7eedefce10828f70ec165cc4c1a5dd6bf52ae85f67237f033e97b35c5b6043886d612a4f669f453d64e662b53b493c17193595324b4a7f450509e9935d2ce84",
		"0x7709710c5e07db7261b2dedb0cc3d0e5b294cba22b5b6dfcfca3af7e744c061d3227a726850bd606e7d62e29cb4f8a4259043a68bb18f9977fe2fdc28fc5fd4a",
		}
)

func main() {
	var h types.Header
	if err := json.Unmarshal([]byte(headerJson), &h); err != nil {
		log.Fatal(err)
	}
	node1Byte, err := hexutil.Decode(nodeList[0])
	if err != nil {
		log.Fatal(err)
	}
	node2Byte, err := hexutil.Decode(nodeList[1])
	if err != nil {
		log.Fatal(err)
	}
	node3Byte, err := hexutil.Decode(nodeList[2])
	if err != nil {
		log.Fatal(err)
	}

	//node1Byte := []byte(nodeList[0])
	//node2Byte := []byte(nodeList[1])
	//node3Byte := []byte(nodeList[2])
	genIndexByte, err := hexutil.DecodeBig("0x1")
	if err != nil {
		log.Fatal(err)
	}
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
		genIndexByte,
		h.MixDigest,
		h.Nonce,
		node1Byte,
		node2Byte,
		node3Byte,
	})
	fmt.Println(hash.Hex())
}