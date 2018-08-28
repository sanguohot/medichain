package util

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"math/rand"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	//"bytes"
	"math/big"
	"medichain/etc"
	"time"
	"bytes"
)

func GetKeyJsonFromStore(address common.Address, password string) ([]byte, error) {
	store := keystore.NewKeyStore(etc.GetBcosKeystore(), keystore.LightScryptN, keystore.StandardScryptP)
	account, err := store.Find(accounts.Account{Address: address})
	if err != nil {
		return nil, err
	}
	keyBytes, err := store.Export(account, password, password)
	if err != nil {
		return nil, err
	}
	return  keyBytes, nil
}

func GetTransactOptsFromStore(address common.Address, password string, nonce uint64) (*bind.TransactOpts, error) {
	rand.Seed(time.Now().Unix())
	keyBytes, err := GetKeyJsonFromStore(address, password)
	if err != nil {
		return nil, err
	}
	auth, err := bind.NewTransactor(bytes.NewReader(keyBytes), password)
	if err != nil {
		return nil, err
	}
	auth.Value = nil
	auth.GasLimit = uint64(1000000) // in units
	if nonce > 0 {
		auth.Nonce = big.NewInt(int64(nonce))
	}else {
		auth.Nonce = big.NewInt(rand.Int63n(100000000)) // 也是fisco-bcos的randomid字段，这里按照web3.js计算随机数
	}
	auth.GasPrice = nil
	auth.From = address
	auth.BlockLimit = uint64(750)
	return  auth, nil
}