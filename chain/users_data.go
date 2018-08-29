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

func GetUsersDataInstance() (error, *medi.UsersData) {
	url := fmt.Sprintf("http://%s:%d", etc.GetBcosHostAddress(), etc.GetBcosHostRpcPort())
	client, err := ethclient.Dial(url)
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
		return nil, nil
	}
	bytes16, err := instance.GetUuidByIdCartNoHash(nil, hash)
	userUuid := uuid.UUID(bytes16)
	return &userUuid, nil
}