package contracts

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sanguohot/medichain/chain"
	authContract "github.com/sanguohot/medichain/contracts/auth"
	"github.com/sanguohot/medichain/contracts/medi" // for demo
	"github.com/sanguohot/medichain/etc"
	"github.com/sanguohot/medichain/util"
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
	case etc.ContractAuth:
		address, tx, _, err = authContract.DeployAuth(auth, client)
	}
	hash := tx.Hash()
	return nil, &address, &hash
}
func DeployEasyCns(forceDeploy bool)  {
	cnsAddress := common.HexToAddress(etc.GetBcosEasyCnsAddress())
	if util.IsValidAndNotZeroAddress(cnsAddress) && !forceDeploy {
		zap.Logger.Fatal(fmt.Errorf("%s:%s, 如果需要重新部署请设置forceDeploy=true", util.ErrContractAlreadyDeploy.Error(), cnsAddress.Hex()).Error())
	}
	err, address, hash := DeployContract(&cnsAddress, etc.ContractEasyCns)
	if err != nil {
		zap.Logger.Fatal(err.Error())
	}
	zap.Sugar.Infof("new easy cns address ===> %s", address.Hex())
	zap.Sugar.Infof("transaction hash ===> %s", hash.Hex())
}
func DeployAllByEasyCnsAddress(cnsAddress common.Address) error {
	if !util.IsValidAndNotZeroAddress(cnsAddress) {
		cnsAddress = common.HexToAddress(etc.GetBcosEasyCnsAddress())
	}
	if !util.IsValidAndNotZeroAddress(cnsAddress) {
		return util.ErrParamCnsAddressShouldNotNil
	}
	zap.Sugar.Infof("cnsAddress ===> %s", cnsAddress.Hex())
	for key, value := range etc.ContractMap {
		if value && key==etc.ContractEasyCns {
			continue
		}
		if key == etc.ContractAuth {
			continue
		}
		err, address, hash := DeployContract(&cnsAddress, key)
		if err != nil {
			return err
		}
		zap.Sugar.Infof("DeployContract %s tx sent: %s", key, hash.Hex())
		err, hash = chain.SetAddressToCns(key, *address)
		if err != nil {
			return err
		}
		zap.Sugar.Infof("SetAddressToCns %s ===> %s, tx sent: %s", key, address.Hex(), hash.Hex())
	}
	name := etc.ContractController
	err, address := chain.GetAddressFromCns(name)
	if err != nil {
		return err
	}
	zap.Sugar.Infof("ContractController address ===> %s", address.Hex())
	if err, hash := chain.OrgsDataAddSuper(*address); err != nil {
		return err
	}else {
		zap.Sugar.Infof("OrgsDataAddSuper %s, tx sent: %s", address.Hex(), hash.Hex())
	}
	if err, hash := chain.UsersDataAddSuper(*address); err != nil {
		return err
	}else {
		zap.Sugar.Infof("UsersDataAddSuper %s, tx sent: %s", address.Hex(), hash.Hex())
	}
	if err, hash := chain.FilesDataAddSuper(*address); err != nil {
		return err
	}else {
		zap.Sugar.Infof("FilesDataAddSuper %s, tx sent: %s", address.Hex(), hash.Hex())
	}
	return nil
}

func DeployAllByDefaultEasyCnsAddress(forceDeploy bool)  {
	size, err := chain.FilesDataGetSuperSize();
	if err != nil {
		//log.Fatal(err)
		// 这里不要退出，可能是合约没有部署或者设置报错abi: unmarshalling empty output
	}else {
		if size.Uint64() > 0 && !forceDeploy {
			zap.Logger.Fatal(fmt.Errorf("%s, 如果需要重新部署请设置forceDeploy=true", util.ErrContractAlreadyDeploy.Error()).Error())
		}
	}

	if err := DeployAllByEasyCnsAddress(common.Address{}); err != nil {
		zap.Logger.Fatal(err.Error())
	}else {
		zap.Logger.Info("deploy all contracts successfully!")
	}
}