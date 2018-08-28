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
	err, controllerAddress := GetAddressFromCns(etc.ContractController)
	if err != nil {
		return err, nil
	}
	instance, err := medi.NewController(*controllerAddress, client)
	if err != nil {
		return err, nil
	}
	return nil, instance
}

func AddOrgToCns(uuid [16]byte, publicKey [2][32]byte, name [4][32]byte) (error, *common.Hash) {
	err, instance := GetControllerInstance()
	if err != nil {
		return err, nil
	}
	fromAddressStr := etc.GetBcosOwnerAddress()
	if fromAddressStr == "" {
		return util.ErrConfigItemRequire, nil
	}
	fromAddress := common.HexToAddress(fromAddressStr)
	password := etc.GetBcosOwnerPassword()
	auth, err := util.GetTransactOptsFromStore(fromAddress, password, 0)
	if err != nil {
		return err, nil
	}
	nameHash := util.Bytes32_4Hash(name)
	tx, err := instance.AddOrg(auth, uuid, publicKey, nameHash, name, big.NewInt(time.Now().Unix()))
	if err != nil {
		return err, nil
	}
	hash := tx.Hash()
	return nil, &hash
}