package chain

import (
	"fmt"
	"medichain/etc"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetEthDialClient() (*ethclient.Client, error) {
	url := fmt.Sprintf("http://%s:%d", etc.GetBcosHostAddress(), etc.GetBcosHostRpcPort())
	return ethclient.Dial(url)
}