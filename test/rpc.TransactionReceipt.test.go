package main

import (
	"context"
	"fmt"
	"log"
	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
	"strings"
	"github.com/sanguohot/medichain/contracts/hello"
)

func main() {
	client, err := ethclient.Dial("http://10.6.250.55:8546")
	if err != nil {
		log.Fatal(err)
		return
	}

	receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash("0xb4ec11e3a487dcbb64437566ad314b1648744c8ad9804008cb8d71e4449ddf8a"))
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("receipt.Logs ===>", receipt.Logs)
	fmt.Println("receipt.Status ===>", receipt.Status)
	fmt.Println("receipt.TxHash ===>", receipt.TxHash.Hex())
	contractAbi, err := abi.JSON(strings.NewReader(string(hello.HelloABI)))
	if err != nil {
		log.Fatal("abi.JSON", err)
		return
	}
	for _, vLog := range receipt.Logs {
		var event struct {
			NewSaying   string
		}
		// 参数一是事件参数构造的对象
		// 参数二是事件函数名，不是合约函数名
		// 参数三是事件的数据
		err := contractAbi.Unpack(&event, "onSaySomethingElse", vLog.Data)
		if err != nil {
			log.Fatal("contractAbi.Unpack", err)
			return
		}

		fmt.Println("event onSaySomethingElse===>", event)

		var topics [4]string
		for i := range vLog.Topics {
			topics[i] = vLog.Topics[i].Hex()
		}

		fmt.Println("topics ===>", topics) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4
	}

	eventSignature := []byte("onSaySomethingElse(string)")
	hash := crypto.Keccak256Hash(eventSignature)
	fmt.Println("eventSignature hash ===>", hash.Hex())
}