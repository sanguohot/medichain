package main

import (
	"fmt"
	"log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	ethereum "github.com/ethereum/go-ethereum"
	hello "medichain/contracts/hello"
	"context"
	"math/big"
	"strings"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

func main() {
	client, err := ethclient.Dial("http://10.6.250.56:8545")
	if err != nil {
		log.Fatal(err)
		return
	}
	// 0xb349Eba018bFA9d89Da90829629D39668F6653A2
	// 0xca21167a870cf8b9618d259af454c6d00b30b028
	// 0xB818715eb048286A608B5E9851877AD7A30a41A0
	contractAddress := common.HexToAddress("0xB818715eb048286A608B5E9851877AD7A30a41A0")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(640),
		ToBlock:   big.NewInt(647),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal("client.FilterLogs", err)
		return
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(hello.HelloABI)))
	if err != nil {
		log.Fatal("abi.JSON", err)
		return
	}
	fmt.Println("logs count ===>", len(logs))
	for _, vLog := range logs {
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

		fmt.Println("event ===>", event)

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