package chain

import (
	"github.com/ethereum/go-ethereum/common"
	"fmt"
	"medichain/etc"
	"github.com/ethereum/go-ethereum/ethclient"
	"medichain/util"
	"medichain/contracts/medi"
	"math/big"
	"time"
	"github.com/google/uuid"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func GetOrgsDataInstance() (error, *medi.OrgsData) {
	url := fmt.Sprintf("http://%s:%d", etc.GetBcosHostAddress(), etc.GetBcosHostRpcPort())
	client, err := ethclient.Dial(url)
	if err != nil {
		return err, nil
	}
	err, address := GetAddressFromCns(etc.ContractOrgsData)
	if err != nil {
		return err, nil
	}
	instance, err := medi.NewOrgsData(*address, client)
	if err != nil {
		return err, nil
	}
	return nil, instance
}

func GetOrgsDataInstanceAndAuth() (error, *medi.OrgsData,*bind.TransactOpts) {
	err, instance := GetOrgsDataInstance()
	if err != nil {
		return err, nil, nil
	}
	auth, err := util.GetDefaultTransactOptsFromStore()
	if err != nil {
		return err, nil, nil
	}
	return nil, instance, auth
}

func OrgsDataAddSuper(address common.Address) (error, *common.Hash) {
	err, instance, auth := GetOrgsDataInstanceAndAuth()
	if err != nil {
		return err, nil
	}
	tx, err := instance.AddSuper(auth, address)
	if err != nil {
		return err, nil
	}
	hash := tx.Hash()
	return nil, &hash
}

func OrgsDataAddOrg(uuid [16]byte, orgAddress common.Address, publicKey [2][32]byte, name [4][32]byte) (error, *common.Hash) {
	err, instance, auth := GetOrgsDataInstanceAndAuth()
	if err != nil {
		return err, nil
	}
	nameHash := util.Bytes32_4Hash(name)
	tx, err := instance.AddOrg(auth, uuid, orgAddress, publicKey, nameHash, name, big.NewInt(time.Now().Unix()))
	if err != nil {
		return err, nil
	}
	hash := tx.Hash()
	return nil, &hash
}

func OrgsDataDelOrg(uuid [16]byte) (error, *common.Hash) {
	err, instance, auth := GetOrgsDataInstanceAndAuth()
	if err != nil {
		return err, nil
	}
	tx, err := instance.DelOrg(auth, uuid)
	if err != nil {
		return err, nil
	}
	hash := tx.Hash()
	return nil, &hash
}
func OrgsDataIsUuidExist(uuid uuid.UUID) (bool, error) {
	err, instance := GetOrgsDataInstance()
	if err != nil {
		return false, nil
	}
	return instance.IsUuidExist(nil, uuid)
}