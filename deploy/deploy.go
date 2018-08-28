package deploy

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"medichain/contracts/medi" // for demo
	"github.com/ethereum/go-ethereum/common"
	"medichain/util"
	"medichain/etc"
	"context"
	"github.com/ethereum/go-ethereum/core/types"
)

func DeployContract(cnsAddress *common.Address, contract string) (error, *common.Address, *common.Hash) {
	if contract == "" || !etc.ContractMap[contract] {
		return util.ErrParamContractNotSupport, nil, nil
	}
	if contract!=etc.ContractEasyCns && cnsAddress==nil {
		return util.ErrParamCnsAddressShouldNotNil, nil, nil
	}
	url := fmt.Sprintf("http://%s:%d", etc.GetBcosHostAddress(), etc.GetBcosHostRpcPort())
	client, err := ethclient.Dial(url)
	if err != nil {
		return err, nil, nil
	}
	fromAddressStr := etc.GetBcosOwnerAddress()
	if fromAddressStr == "" {
		return util.ErrConfigItemRequire, nil, nil
	}
	fromAddress := common.HexToAddress(fromAddressStr)
	password := etc.GetBcosOwnerPassword()
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
	case etc.ContractEasyCns:
		address, tx, _, err = medi.DeployEasyCns(auth, client)
	case etc.ContractOrgsData:
		address, tx, _, err = medi.DeployOrgsData(auth, client, *cnsAddress)
	case etc.ContractUsersData:
		address, tx, _, err = medi.DeployUsersData(auth, client, *cnsAddress)
	case etc.ContractFilesData:
		address, tx, _, err = medi.DeployFilesData(auth, client, *cnsAddress)
	case etc.ContractController:
		address, tx, _, err = medi.DeployController(auth, client, *cnsAddress)
	}
	hash := tx.Hash()
	return nil, &address, &hash
}