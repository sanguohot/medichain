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

func GetFilesDataInstance() (error, *medi.FilesData) {
	client, err := GetEthDialClient()
	if err != nil {
		return err, nil
	}
	err, address := GetAddressFromCns(etc.ContractFilesData)
	if err != nil {
		return err, nil
	}
	instance, err := medi.NewFilesData(*address, client)
	if err != nil {
		return err, nil
	}
	return nil, instance
}

func GetFilesDataInstanceAndAuth() (error, *medi.FilesData, *bind.TransactOpts) {
	err, instance := GetFilesDataInstance()
	if err != nil {
		return err, nil, nil
	}
	auth, err := util.GetDefaultTransactOptsFromStore()
	if err != nil {
		return err, nil, nil
	}
	return nil, instance, auth
}

func FilesDataAddSuper(address common.Address) (error, *common.Hash) {
	err, instance, auth := GetFilesDataInstanceAndAuth()
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

func FilesDataAddFile(uuid [16]byte, ownerUuid [16]byte, uploaderUuid [16]byte, fileType common.Hash, fileDesc [4][32]byte, keccak256Hash common.Hash,
	sha256Hash common.Hash, r [32]byte, s [32]byte, v uint8, address common.Address, password string) (error, *common.Hash) {
	err, instance, auth := GetFilesDataInstanceAndAuth()
	if err != nil {
		return err, nil
	}
	tx, err := instance.AddFile(auth, uuid, ownerUuid, uploaderUuid, fileType, fileDesc, keccak256Hash, sha256Hash, r, s, v, big.NewInt(time.Now().Unix()))
	if err != nil {
		return err, nil
	}
	hash := tx.Hash()
	return nil, &hash
}

func FilesDataDelFile(uuid [16]byte) (error, *common.Hash) {
	err, instance, auth := GetFilesDataInstanceAndAuth()
	if err != nil {
		return err, nil
	}
	tx, err := instance.DelFile(auth, uuid)
	if err != nil {
		return err, nil
	}
	hash := tx.Hash()
	return nil, &hash
}
func FilesDataIsUuidExist(uuid uuid.UUID) (bool, error) {
	err, instance := GetFilesDataInstance()
	if err != nil {
		return false, err
	}
	return instance.IsUuidExist(nil, uuid)
}
func FilesDataGetUuidByKeccak256Hash(hash common.Hash) (*uuid.UUID, error) {
	err, instance := GetFilesDataInstance()
	if err != nil {
		return nil, err
	}
	bytes16, err := instance.GetUuidByKeccak256Hash(nil, hash)
	if err != nil {
		return nil, err
	}
	fileUuid := uuid.UUID(bytes16)
	return &fileUuid, nil
}
func FilesDataGetUuidBySha256Hash(hash common.Hash) (*uuid.UUID, error) {
	err, instance := GetFilesDataInstance()
	if err != nil {
		return nil, err
	}
	bytes16, err := instance.GetUuidBySha256Hash(nil, hash)
	if err != nil {
		return nil, err
	}
	fileUuid := uuid.UUID(bytes16)
	return &fileUuid, nil
}
func FilesDataGetSha256Hash(fileUuid uuid.UUID) (*common.Hash, error) {
	err, instance := GetFilesDataInstance()
	if err != nil {
		return nil, err
	}
	bytes32, err := instance.GetSha256Hash(nil, fileUuid)
	if err != nil {
		return nil, err
	}
	hash := common.Hash(bytes32)
	return &hash, nil
}
func FilesDataGetKeccak256Hash(fileUuid uuid.UUID) (*common.Hash, error) {
	err, instance := GetFilesDataInstance()
	if err != nil {
		return nil, err
	}
	bytes32, err := instance.GetKeccak256Hash(nil, fileUuid)
	if err != nil {
		return nil, err
	}
	hash := common.Hash(bytes32)
	return &hash, nil
}
func FilesDataGetFileSignerSize(fileUuid uuid.UUID) (*big.Int, error) {
	err, instance := GetFilesDataInstance()
	if err != nil {
		return nil, err
	}
	size, err := instance.GetFileSignerSize(nil, fileUuid)
	if err != nil {
		return nil, err
	}
	return size, nil
}