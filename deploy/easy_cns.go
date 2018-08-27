package deploy

import (
	"fmt"
	"math/big"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"medichain/util"
	"medichain/contracts/medi" // for demo
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/accounts"
	"bytes"
	"context"
)

var (
	ErrConfigItemRequire = errors.New("没有找到配置项")
)

func DeployEasyCns() (error, *common.Address, *common.Hash) {
	url := fmt.Sprintf("http://%s:%d", util.Config.GetString("bcos.host.address"), util.Config.GetInt("bcos.host.rpc_port"))
	client, err := ethclient.Dial(url)
	if err != nil {
		return err, nil, nil
	}
	fromAddressStr := util.Config.GetString("bcos.owner.address")
	if fromAddressStr == "" {
		return ErrConfigItemRequire, nil, nil
	}
	store := keystore.NewKeyStore("./keystore", keystore.LightScryptN, keystore.StandardScryptP)
	fromAddress := common.HexToAddress(fromAddressStr)
	account, err := store.Find(accounts.Account{Address: fromAddress})
	if err != nil {
		return err, nil, nil
	}
	password := util.Config.GetString("bcos.owner.password")
	keyBytes, err := store.Export(account, password, password)
	if err != nil {
		return err, nil, nil
	}
	// use keystore directly
	auth, err := bind.NewTransactor(bytes.NewReader(keyBytes), password)
	if err != nil {
		return err, nil, nil
	}
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(1000000) // in units
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return err, nil, nil
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasPrice = nil
	auth.From = fromAddress

	address, tx, _, err := medi.DeployEasyCns(auth, client)
	if err != nil {
		return err, nil, nil
	}
	hash := tx.Hash()
	return nil, &address, &hash
}