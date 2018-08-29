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

func ControllerAddOrg(uuid [16]byte, name [4][32]byte, address common.Address, password string) (error, *common.Hash) {
	err, instance := GetControllerInstance()
	if err != nil {
		return err, nil
	}
	auth, err := util.GetTransactOptsFromStore(address, password, 0)
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
	err, instance := GetControllerInstance()
	if err != nil {
		return err, nil
	}
	auth, err := util.GetTransactOptsFromStore(address, password, 0)
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

func ControllerAddFile(uuid [16]byte, ownerUuid [16]byte, fileType common.Hash, fileDesc [4][32]byte, keccak256Hash common.Hash,
	sha256Hash common.Hash, r [32]byte, s [32]byte, v uint8, address common.Address, password string) (error, *common.Hash) {
	err, instance := GetControllerInstance()
	if err != nil {
		return err, nil
	}
	auth, err := util.GetTransactOptsFromStore(address, password, 0)
	if err != nil {
		return err, nil
	}
	tx, err := instance.AddFile(auth, uuid, ownerUuid, fileType, fileDesc, keccak256Hash, sha256Hash, r, s, v, big.NewInt(time.Now().Unix()))
	if err != nil {
		return err, nil
	}
	hash := tx.Hash()
	return nil, &hash
}

func ControllerGetOrg()  {

}