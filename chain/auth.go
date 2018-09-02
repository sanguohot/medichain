package chain

import (
	"github.com/ethereum/go-ethereum/common"
	"medichain/etc"
	"medichain/util"
	authContract "medichain/contracts/auth"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func GetAuthInstance() (error, *authContract.Auth) {
	client, err := GetEthDialClient()
	if err != nil {
		return err, nil
	}
	err, address := GetAddressFromCns(etc.ContractAuth)
	if err != nil {
		return err, nil
	}
	instance, err := authContract.NewAuth(*address, client)
	if err != nil {
		return err, nil
	}
	return nil, instance
}

func GetAuthInstanceAndAuth() (error, *authContract.Auth, *bind.TransactOpts) {
	err, instance := GetAuthInstance()
	if err != nil {
		return err, nil, nil
	}
	auth, err := util.GetDefaultTransactOptsFromStore()
	if err != nil {
		return err, nil, nil
	}
	return nil, instance, auth
}

func AuthVerifySignatureWithoutPrefix(hash [32]byte, v uint8, r [32]byte, s [32]byte) (error, *common.Address) {
	err, instance := GetAuthInstance()
	if err != nil {
		return err, nil
	}
	address, err := instance.VerifySignatureWithoutPrefix(nil, hash, v, r, s)
	if err != nil {
		return err, nil
	}
	return nil, &address
}

func AuthVerifySignatureWithPrefix(hash [32]byte, v uint8, r [32]byte, s [32]byte) (error, *common.Address) {
	err, instance := GetAuthInstance()
	if err != nil {
		return err, nil
	}
	address, err := instance.VerifySignatureWithPrefix(nil, hash, v, r, s)
	if err != nil {
		return err, nil
	}
	return nil, &address
}

func AuthGetKeccak256FromHash(hash [32]byte) (error, *common.Hash) {
	err, instance := GetAuthInstance()
	if err != nil {
		return err, nil
	}
	bytes32, err := instance.GetKeccak256FromHash(nil, hash)
	if err != nil {
		return err, nil
	}
	retHash := common.Hash(bytes32)
	return nil, &retHash
}

func AuthGetSha3FromHash(hash [32]byte) (error, *common.Hash) {
	err, instance := GetAuthInstance()
	if err != nil {
		return err, nil
	}
	bytes32, err := instance.GetSha3FromHash(nil, hash)
	if err != nil {
		return err, nil
	}
	retHash := common.Hash(bytes32)
	return nil, &retHash
}