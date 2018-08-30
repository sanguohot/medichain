package chain

import (
	"github.com/ethereum/go-ethereum/common"
	"medichain/etc"
	"medichain/util"
	"medichain/contracts/medi"
)

func GetCnsInstance() (error, *medi.EasyCns) {
	client, err := GetEthDialClient()
	if err != nil {
		return err, nil
	}
	instance, err := medi.NewEasyCns(common.HexToAddress(etc.GetBcosEasyCnsAddress()), client)
	if err != nil {
		return err, nil
	}
	return nil, instance
}

func SetAddressToCns(name string, address common.Address) (error, *common.Hash) {
	err, instance := GetCnsInstance()
	if err != nil {
		return err, nil
	}
	auth, err := util.GetDefaultTransactOptsFromStore()
	if err != nil {
		return err, nil
	}
	tx, err := instance.Set(auth, name, address)
	if err != nil {
		return err, nil
	}
	hash := tx.Hash()
	return nil, &hash
}

func GetAddressFromCns(name string) (error, *common.Address) {
	err, instance := GetCnsInstance()
	if err != nil {
		return err, nil
	}
	address, err := instance.Get(nil, name)
	if err != nil {
		return err, nil
	}
	return nil, &address
}