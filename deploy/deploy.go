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
	authContract "medichain/contracts/auth"
	"medichain/chain"
	"log"
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

func DeployAllByEasyCnsAddress(cnsAddress common.Address) error {
	if !util.IsValidAndNotZeroAddress(cnsAddress) {
		cnsAddress = common.HexToAddress(etc.GetBcosEasyCnsAddress())
	}
	if !util.IsValidAndNotZeroAddress(cnsAddress) {
		return util.ErrParamCnsAddressShouldNotNil
	}
	fmt.Println("cnsAddress ===>", cnsAddress.Hex())
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
		fmt.Printf("DeployContract %s tx sent: %s\n", key, hash.Hex())
		//time.Sleep(time.Second * 1)
		err, hash = chain.SetAddressToCns(key, *address)
		if err != nil {
			return err
		}
		fmt.Printf("SetAddressToCns %s ===> %s, tx sent: %s\n", key, address.Hex(), hash.Hex())
		//time.Sleep(time.Second * 1)
	}
	name := etc.ContractController
	err, address := chain.GetAddressFromCns(name)
	if err != nil {
		return err
	}
	fmt.Println("ContractController address ===>", address.Hex())
	if err, hash := chain.OrgsDataAddSuper(*address); err != nil {
		return err
	}else {
		fmt.Printf("OrgsDataAddSuper %s, tx sent: %s\n", address.Hex(), hash.Hex())
	}
	if err, hash := chain.UsersDataAddSuper(*address); err != nil {
		return err
	}else {
		fmt.Printf("UsersDataAddSuper %s, tx sent: %s\n", address.Hex(), hash.Hex())
	}
	if err, hash := chain.FilesDataAddSuper(*address); err != nil {
		return err
	}else {
		fmt.Printf("FilesDataAddSuper %s, tx sent: %s\n", address.Hex(), hash.Hex())
	}
	return nil
}

func DeployAllByDefaultEasyCnsAddress()  {
	if err := DeployAllByEasyCnsAddress(common.Address{}); err != nil {
		log.Fatal(err)
	}else {
		fmt.Println("deploy all contracts successfully!")
	}
}