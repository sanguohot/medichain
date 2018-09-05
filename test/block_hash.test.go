package main

import (
	"medichain/util"
	"github.com/ethereum/go-ethereum/core/types"
	"encoding/json"
	"log"
	"fmt"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/common/hexutil"
	//"math/big"
	"math/big"
	//"hash"
)
var (
	headerJson = `{"parentHash":"0x74ede5fefb269971ca25246531ca8dfc8a177d4086edbe86f59c6307e3baafcb","sha3Uncles":"0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347","miner":"0x0000000000000000000000000000000000000000","stateRoot":"0xec3ed72a871ea40c64897d8c3d2eac9d4ffff259056a6317f3934e63af8b6456","transactionsRoot":"0x5a1b061005a4bb203393d33f279cc77e568df6fd92f63e0b31b55d2a70139bc4","receiptsRoot":"0xd6f97cc7ba377519dd6c66b0080f574bc58c78d92ced4a63f4ed38c28fe2d182","logsBloom":"0x00000800000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000020000000000000100000000000000000000000000000000000000000000000000000000000000000000000001000010000000000000000000000000000","difficulty":"0x1","number":"0xf","gasLimit":"0x7aef40a00","gasUsed":"0x14e9e","timestamp":"0x165a23e75c9","extraData":"0xdc809a312e332e332b2b623265362a524c696e75782f672b2b2f496e74","mixHash":"0x0000000000000000000000000000000000000000000000000000000000000000","nonce":"0x0000000000000000","hash":"0x2a6ea622caff7c4bbad9874641f078f7e635f9deef0191fb2acad2970984fca2"}`
	nodeList = []string{
		"daac08f0ea5920017787e6fd3afc428cc9c209549b1ad587eed049b156e916bfd887c3fc4b1a9310d7b78658bfac0500a36e0b2cbdddc2249de79a3e84e8fb1f",
		"2eb39f082cfedcb77e13955e554801dafa711366270f4986fbc0183156bbbde28f235b985b45f893157eeec968ea53590b7010637834651902597833f1d97dd9",
		"4479dbdc126779284d3347bad2ab8dd63d62bc6fb520dd76b754ba2833904ca6b1ce761413f2db528cf9f3288d91762abd821dfc9f9f3c232d67be0b6cc3aab1",
		"8166ee9a1c69dfe80064c5052c7ac0cdc1d4cf90b4a0068f714afb298a87c136e081fcd6e1175be91b32a8be9274f2052fd633b35edd5025dea84f965d1296e2",
		"fce75ab695d686ee4e37bb92de92a525108478c178d6fea59f0417c8347694736e01b20790b7ec8114aeb132d1b7f355ce1f7581e585f5b30d96e2a50ec65723",
		"d7f31150535d1ed4c755e21854f761dc809f77a665496cc055231bc69cbb6a38cbe2870486530f993d4eb75043bae25a7e4165cf5298ddf33c1001f750cad036",
		}
)

func main() {
	var h types.Header
	if err := json.Unmarshal([]byte(headerJson), &h); err != nil {
		log.Fatal(err)
	}
	dst := make([][64]byte, len(nodeList))
	for i, item := range nodeList{
		itemHex := fmt.Sprintf("0x%s", item)
		t, err := hexutil.Decode(itemHex)
		if err != nil {
			log.Fatal(err)
		}
		if len(t) != 64 {
			log.Fatal(fmt.Errorf("%s", "节点ID必须是64字节"))
		}
		copy(dst[i][0:64], t[0:64])
	}
	genIndex := big.NewInt(0x2)
	rlpDst, err := rlp.EncodeToBytes(dst)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("node=all, nodebytes=%d, rlpbytes=%d, rlpcode=%s\n", len(dst), len(rlpDst), hexutil.Encode(rlpDst))
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
		genIndex,
		dst,
	})
	fmt.Println("block hash ===>", hash.Hex())
}