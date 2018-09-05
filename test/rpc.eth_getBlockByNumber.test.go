package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common"
	"log"
)

func main() {
	client, err := ethclient.Dial("http://10.6.250.55:8546")
	if err != nil {
		fmt.Println("dail http://10.6.250.55:8846 ===> %s", err)
		return
	}
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		fmt.Println("client.HeaderByNumber ===> %s", err)
		return
	}
	block, err := client.BlockByNumber(context.Background(), header.Number)
	if err != nil {
		fmt.Println("client.BlockByNumber ===> %s", err)
		return
	}
	fmt.Println("block.Number ===>", block.Number().Uint64())     // 5671744
	fmt.Println("block.Time ===>", block.Time().Uint64())       // 1527211625
	fmt.Println("block.Difficulty ===>", block.Difficulty().Uint64()) // 3217000136609065
	fmt.Println("block.TxHash(transactionsRoot) ===>", block.TxHash().Hex())          // 0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9
	// block hash计算不对
	fmt.Println("block.Hash ===>", block.Hash().Hex())
	fmt.Println("block.nodelist 0 length ===>", len(block.NodeList()[0]))
	fmt.Println("block.Extra ===>", hexutil.Encode(block.Extra()))
	fmt.Println("block.Transactions ===>", block.Transactions())   // 144
	count, err := client.TransactionCount(context.Background(), common.HexToHash("0xe2192c151e5060b5b26c4fdf9d0ec887edcfd5da64595acc2dc49948e2b4a45b"))
	if err != nil {
		fmt.Println("client.TransactionCount ===> %s", err)
		return
	}

	fmt.Println("client.TransactionCount ===>", count) // 144

	bj, err := block.Header().MarshalJSON();
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bj))
}