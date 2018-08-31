package chain

import (
	"github.com/ethereum/go-ethereum/common"
	"medichain/etc"
	"medichain/util"
	"medichain/contracts/medi"
	"math/big"
	"time"
	"github.com/google/uuid"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func GetOrgsDataInstance() (error, *medi.OrgsData) {
	client, err := GetEthDialClient()
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

func GetOrgsDataInstanceAndAuth() (error, *medi.OrgsData, *bind.TransactOpts) {
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
	if err := CheckReceiptStatus(hash); err!=nil {
		return err, nil
	}
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
	if err := CheckReceiptStatus(hash); err!=nil {
		return err, nil
	}
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
	if err := CheckReceiptStatus(hash); err!=nil {
		return err, nil
	}
	return nil, &hash
}
func OrgsDataIsUuidExist(uuid uuid.UUID) (bool, error) {
	err, instance := GetOrgsDataInstance()
	if err != nil {
		return false, err
	}
	return instance.IsUuidExist(nil, uuid)
}
func OrgsDataGetUuidByNameHash(hash common.Hash) (*uuid.UUID, error) {
	err, instance := GetOrgsDataInstance()
	if err != nil {
		return nil, err
	}
	bytes16, err := instance.GetUuidByNameHash(nil, hash)
	if err != nil {
		return nil, err
	}
	orgUuid := uuid.UUID(bytes16)
	return &orgUuid, nil
}
func OrgsDataGetSuperSize() (*big.Int, error) {
	err, instance := GetOrgsDataInstance()
	if err != nil {
		return nil, err
	}
	size, err := instance.GetSuperSize(nil)
	if err != nil {
		return nil, err
	}
	return size, nil
}
func OrgsDataGetSuperByIndex(index *big.Int) (*common.Address, error) {
	err, instance := GetOrgsDataInstance()
	if err != nil {
		return nil, err
	}
	address, err := instance.GetSuperByIndex(nil, index)
	if err != nil {
		return nil, err
	}
	return &address, nil
}
func OrgsDataGetOwner() (*common.Address, error) {
	err, instance := GetOrgsDataInstance()
	if err != nil {
		return nil, err
	}
	address, err := instance.GetOwner(nil)
	if err != nil {
		return nil, err
	}
	return &address, nil
}