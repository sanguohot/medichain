package chain

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/sanguohot/medichain/etc"
	"github.com/sanguohot/medichain/util"
	"github.com/sanguohot/medichain/contracts/medi"
	"math/big"
	"time"
	"github.com/google/uuid"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func GetUsersDataInstance() (error, *medi.UsersData) {
	client, err := GetEthDialClient()
	if err != nil {
		return err, nil
	}
	err, address := GetAddressFromCns(etc.ContractUsersData)
	if err != nil {
		return err, nil
	}
	instance, err := medi.NewUsersData(*address, client)
	if err != nil {
		return err, nil
	}
	return nil, instance
}

func GetUsersDataInstanceAndAuth() (error, *medi.UsersData, *bind.TransactOpts) {
	err, instance := GetUsersDataInstance()
	if err != nil {
		return err, nil, nil
	}
	auth, err := util.GetDefaultTransactOptsFromStore()
	if err != nil {
		return err, nil, nil
	}
	return nil, instance, auth
}

func UsersDataAddSuper(address common.Address) (error, *common.Hash) {
	err, instance, auth := GetUsersDataInstanceAndAuth()
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

func UsersDataAddUser(uuid [16]byte, orgUuid [16]byte, userAddress common.Address, publicKey [2][32]byte, idCartNoHash common.Hash) (error, *common.Hash) {
	err, instance, auth := GetUsersDataInstanceAndAuth()
	if err != nil {
		return err, nil
	}
	tx, err := instance.AddUser(auth, uuid, userAddress, orgUuid, publicKey, idCartNoHash, big.NewInt(time.Now().Unix()))
	if err != nil {
		return err, nil
	}
	hash := tx.Hash()
	if err := CheckReceiptStatus(hash); err!=nil {
		return err, nil
	}
	return nil, &hash
}

func UsersDataDelUser(uuid [16]byte) (error, *common.Hash) {
	err, instance, auth := GetUsersDataInstanceAndAuth()
	if err != nil {
		return err, nil
	}
	tx, err := instance.DelUser(auth, uuid)
	if err != nil {
		return err, nil
	}
	hash := tx.Hash()
	if err := CheckReceiptStatus(hash); err!=nil {
		return err, nil
	}
	return nil, &hash
}
func UsersDataIsUuidExist(uuid uuid.UUID) (bool, error) {
	err, instance := GetUsersDataInstance()
	if err != nil {
		return false, err
	}
	return instance.IsUuidExist(nil, uuid)
}
func UsersDataGetUuidByIdCartNoHash(hash common.Hash) (*uuid.UUID, error) {
	err, instance := GetUsersDataInstance()
	if err != nil {
		return nil, err
	}
	bytes16, err := instance.GetUuidByIdCartNoHash(nil, hash)
	if err != nil {
		return nil, err
	}
	userUuid := uuid.UUID(bytes16)
	return &userUuid, nil
}
func UsersDataGetSuperSize() (*big.Int, error) {
	err, instance := GetUsersDataInstance()
	if err != nil {
		return nil, err
	}
	size, err := instance.GetSuperSize(nil)
	if err != nil {
		return nil, err
	}
	return size, nil
}
func UsersDataGetSuperByIndex(index *big.Int) (*common.Address, error) {
	err, instance := GetUsersDataInstance()
	if err != nil {
		return nil, err
	}
	address, err := instance.GetSuperByIndex(nil, index)
	if err != nil {
		return nil, err
	}
	return &address, nil
}
func UsersDataGetOwner() (*common.Address, error) {
	err, instance := GetUsersDataInstance()
	if err != nil {
		return nil, err
	}
	address, err := instance.GetOwner(nil)
	if err != nil {
		return nil, err
	}
	return &address, nil
}