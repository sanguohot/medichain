package chain

import (
	"github.com/ethereum/go-ethereum/common"
	"fmt"
	"medichain/etc"
	"github.com/ethereum/go-ethereum/ethclient"
	"medichain/util"
	"medichain/contracts/medi"
	"time"
	"math/big"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func GetControllerInstance() (error, *medi.Controller) {
	url := fmt.Sprintf("http://%s:%d", etc.GetBcosHostAddress(), etc.GetBcosHostRpcPort())
	client, err := ethclient.Dial(url)
	if err != nil {
		return err, nil
	}
	err, address := GetAddressFromCns(etc.ContractController)
	if err != nil {
		return err, nil
	}
	instance, err := medi.NewController(*address, client)
	if err != nil {
		return err, nil
	}
	return nil, instance
}

func GetControllerInstanceAndAuth(address common.Address, password string) (error, *medi.Controller, *bind.TransactOpts) {
	err, instance := GetControllerInstance()
	if err != nil {
		return err, nil, nil
	}
	auth, err := util.GetTransactOptsFromStore(address, password, 0)
	if err != nil {
		return err, nil, nil
	}
	return nil, instance, auth
}

func ControllerAddOrg(uuid [16]byte, name [4][32]byte, address common.Address, password string) (error, *common.Hash) {
	err, instance, auth := GetControllerInstanceAndAuth(address, password)
	if err != nil {
		return err, nil
	}
	nameHash := util.Bytes32_4Hash(name)
	publicKey, err := util.GetPublicKeyBytes32_2FromStore(address, password)
	if err != nil {
		return err, nil
	}
	tx, err := instance.AddOrg(auth, uuid, *publicKey, nameHash, name, big.NewInt(time.Now().Unix()))
	if err != nil {
		return err, nil
	}
	hash := tx.Hash()
	return nil, &hash
}

func ControllerAddUser(uuid [16]byte, orgUuid [16]byte, idCartNoHash common.Hash, address common.Address, password string) (error, *common.Hash) {
	err, instance, auth := GetControllerInstanceAndAuth(address, password)
	if err != nil {
		return err, nil
	}
	publicKey, err := util.GetPublicKeyBytes32_2FromStore(address, password)
	if err != nil {
		return err, nil
	}
	tx, err := instance.AddUser(auth, uuid, orgUuid, *publicKey, idCartNoHash, big.NewInt(time.Now().Unix()))
	if err != nil {
		return err, nil
	}
	hash := tx.Hash()
	return nil, &hash
}

func ControllerAddFile(uuid, ownerUuid [16]byte, fileType common.Hash, fileDesc [4][32]byte, keccak256Hash, sha256Hash common.Hash, address common.Address, password string) (error, *common.Hash) {
	err, instance, auth := GetControllerInstanceAndAuth(address, password)
	if err != nil {
		return err, nil
	}
	signBytes, err := util.SignWithHash(keccak256Hash, address, password)
	if err != nil {
		return err, nil
	}
	r, s, v := util.SigRSV(signBytes)
	tx, err := instance.AddFile(auth, uuid, ownerUuid, fileType, fileDesc, keccak256Hash, sha256Hash, r, s, v, big.NewInt(time.Now().Unix()))
	if err != nil {
		return err, nil
	}
	hash := tx.Hash()
	return nil, &hash
}

func ControllerAddSign(fileUuid [16]byte, keccak256Hash common.Hash, address common.Address, password string) (error, *common.Hash) {
	err, instance, auth := GetControllerInstanceAndAuth(address, password)
	if err != nil {
		return err, nil
	}
	signBytes, err := util.SignWithHash(keccak256Hash, address, password)
	if err != nil {
		return err, nil
	}
	r, s, v := util.SigRSV(signBytes)
	tx, err := instance.AddSign(auth, fileUuid, r, s, v)
	if err != nil {
		return err, nil
	}
	hash := tx.Hash()
	return nil, &hash
}

func ControllerGetFileSignersAndDataByUuid(fileUuid [16]byte, start *big.Int, limit *big.Int) (error, [][16]byte, [][32]byte, [][32]byte, []uint8) {
	err, instance := GetControllerInstance()
	if err != nil {
		return err, nil, nil, nil, nil
	}
	idl, rl, sl, vl, err := instance.GetFileSignersAndDataByUuid(nil, fileUuid, start, limit)
	if err != nil {
		return err, nil, nil, nil, nil
	}
	return nil, idl, rl, sl, vl
}

func ControllerGetOrgByUuid(orgUuid [16]byte) (error, common.Address, [2][32]byte, [32]byte, [4][32]byte, *big.Int) {
	err, instance := GetControllerInstance()
	if err != nil {
		return err, common.Address{}, [2][32]byte{}, [32]byte{}, [4][32]byte{}, nil
	}
	address, publicKey, nameHash, name, time, err := instance.GetOrgByUuid(nil, orgUuid)
	if err != nil {
		return err, common.Address{}, [2][32]byte{}, [32]byte{}, [4][32]byte{}, nil
	}
	return nil, address, publicKey, nameHash, name, time
}

func ControllerGetUserByUuid(userUuid [16]byte) (error, common.Address, [16]byte, [2][32]byte, [32]byte, *big.Int) {
	err, instance := GetControllerInstance()
	if err != nil {
		return err, common.Address{}, [16]byte{}, [2][32]byte{}, [32]byte{}, nil
	}
	address, orgUuid, publicKey, idCartNoHash, time, err := instance.GetUserByUuid(nil, userUuid)
	if err != nil {
		return err, common.Address{}, [16]byte{}, [2][32]byte{}, [32]byte{}, nil
	}
	return nil, address, orgUuid, publicKey, idCartNoHash, time
}

func ControllerGetFileByUuid(fileUuid [16]byte) (error, [16]byte, [16]byte, [32]byte, [4][32]byte, [32]byte, [32]byte, *big.Int) {
	err, instance := GetControllerInstance()
	if err != nil {
		return err, [16]byte{}, [16]byte{}, [32]byte{}, [4][32]byte{}, [32]byte{}, [32]byte{}, nil
	}
	ownerUuid, uploaderUuid, fileType, fileDesc, sha256Hash, keccak256Hash, time, err := instance.GetFileByUuid(nil, fileUuid)
	if err != nil {
		return err, [16]byte{}, [16]byte{}, [32]byte{}, [4][32]byte{}, [32]byte{}, [32]byte{}, nil
	}
	return nil, ownerUuid, uploaderUuid, fileType, fileDesc, sha256Hash, keccak256Hash, time
}