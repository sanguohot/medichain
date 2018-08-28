package deploy

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"medichain/contracts/medi" // for demo
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"medichain/util"
	"context"
	"github.com/ethereum/go-ethereum/core/types"
)

var (
	ErrConfigItemRequire = errors.New("没有找到配置项")
	ErrParamContractNotSupport = errors.New("合约名称不能为空")
	ErrParamCnsAddressShouldNotNil = errors.New("命名合约地址不能为空")
	ContractUsersData = "UsersData"
	ContractFilesData = "FilesData"
	ContractOrgsData = "OrgsData"
	ContractController = "Controller"
	ContractEasyCns = "EasyCns"
	ContractMap = map[string]bool{ContractUsersData:true, ContractFilesData:true, ContractOrgsData:true, ContractController:true, ContractEasyCns:true}
)

func DeployContract(cnsAddress *common.Address, contract string) (error, *common.Address, *common.Hash) {
	if contract == "" || !ContractMap[contract] {
		return ErrParamContractNotSupport, nil, nil
	}
	if contract!=ContractEasyCns && cnsAddress==nil {
		return ErrParamCnsAddressShouldNotNil, nil, nil
	}
	url := fmt.Sprintf("http://%s:%d", util.Config.GetString("bcos.host.address"), util.Config.GetInt("bcos.host.rpc_port"))
	client, err := ethclient.Dial(url)
	if err != nil {
		return err, nil, nil
	}
	fromAddressStr := util.Config.GetString("bcos.owner.address")
	if fromAddressStr == "" {
		return ErrConfigItemRequire, nil, nil
	}
	fromAddress := common.HexToAddress(fromAddressStr)
	password := util.Config.GetString("bcos.owner.password")
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return err, nil, nil
	}
	auth, err := util.GetTransactOptsFromStore(fromAddress, password, nonce)
	if err != nil {
		return err, nil, nil
	}

	var (
		address common.Address
		tx *types.Transaction
	)
	switch contract {
	case ContractEasyCns:
		address, tx, _, err = medi.DeployEasyCns(auth, client)
	case ContractOrgsData:
		address, tx, _, err = medi.DeployOrgsData(auth, client, *cnsAddress)
	case ContractUsersData:
		address, tx, _, err = medi.DeployUsersData(auth, client, *cnsAddress)
	case ContractFilesData:
		address, tx, _, err = medi.DeployFilesData(auth, client, *cnsAddress)
	case ContractController:
		address, tx, _, err = medi.DeployController(auth, client, *cnsAddress)
	}
	hash := tx.Hash()
	return nil, &address, &hash
}