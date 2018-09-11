package chain

import (
	"fmt"
	"github.com/sanguohot/medichain/etc"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"context"
	"github.com/sanguohot/medichain/util"
	"time"
)

func GetEthDialClient() (*ethclient.Client, error) {
	url := fmt.Sprintf("http://%s:%d", etc.GetBcosHostAddress(), etc.GetBcosHostRpcPort())
	return ethclient.Dial(url)
}

func CheckReceiptStatus(txHash common.Hash) error {
	// 一秒出块，需要延时一秒以上才能查询到交易结果
	time.Sleep(time.Millisecond * 1100)
	client, err := GetEthDialClient()
	if err != nil {
		return err
	}
	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		return err
	}
	fmt.Printf("%s receipt logs len %d, status %d\n", txHash.Hex(), len(receipt.Logs), receipt.Status)
	if len(receipt.Logs) <= 0 {
		return util.ErrContractExeFail
	}
	return nil
}