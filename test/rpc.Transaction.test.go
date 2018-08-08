package main

import (
	"context"
	"fmt"
	"log"
	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("http://10.6.250.56:8545")
	if err != nil {
		log.Fatal(err)
		return
	}

	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	// 这里是transactionsRoot
	fmt.Println("block.TxHash() ===>", block.TxHash().Hex())
	for _, tx := range block.Transactions() {
		// 注意这里与之前的hash计算不一致, 明天考虑是否需要修改
		// Transaction.Hash ===> It uniquely identifies the transaction.
		// FrontierSigner.Hash ===> It does not uniquely identify the transaction.
		// as fisco-bcos use blockLimit to generate a tx hash, and does not return it the rpc call
		// we can not get the same one by types.HomesteadSigner{}.Hash(tx)
		fmt.Println("does not uniquely:types.HomesteadSigner{}.Hash(tx) ===>", types.HomesteadSigner{}.Hash(tx).Hex())
		fmt.Println("uniquely:tx.Hash() ===>", tx.Hash().Hex())        // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
		fmt.Println("tx.data.AccountNonce ===>", tx.Nonce())
		fmt.Println("tx.data.Price ===>", tx.GasPrice())
		fmt.Println("tx.data.GasLimit ===>", tx.Gas())
		fmt.Println("tx.data.BlockLimit ===>", tx.BlockLimit())
		fmt.Println("tx.data.Recipient ===>", tx.To().Hex())
		fmt.Println("tx.data.Amount ===>", tx.Value())
		fmt.Println("tx.data.Payload ===>", tx.Data())
		msg, err := tx.AsMessage(types.HomesteadSigner{})
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println("from ===>", msg.From().Hex())
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
			return
		}

		fmt.Println(receipt.Status) // 1
	}

	blockHash := common.HexToHash("0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9")
	count, err := client.TransactionCount(context.Background(), blockHash)
	if err != nil {
		log.Fatal(err)
	}

	for idx := uint(0); idx < count; idx++ {
		tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(tx.Hash().Hex()) // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
	}

	txHash := common.HexToHash("0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2")
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tx.Hash().Hex()) // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
	fmt.Println(isPending)       // false
}