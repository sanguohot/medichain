package chain

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/sanguohot/medichain/etc"
	"github.com/sanguohot/medichain/util"
	"github.com/sanguohot/medichain/contracts/medi"
)

func GetCnsInstance() (error, *medi.EasyCns) {
	client, err := GetEthDialClient()
	if err != nil {
		return err, nil
	}
	address := common.HexToAddress(etc.GetBcosEasyCnsAddress())
	if !util.IsValidAndNotZeroAddress(address) {
		return util.ErrInvalidOrZeroAddress, nil
	}
	instance, err := medi.NewEasyCns(address, client)
	if err != nil {
		return err, nil
	}
	return nil, instance
}

func SetAddressToCns(name string, address common.Address) (error, *common.Hash) {
	if name == "" {
		return util.ErrParamShouldNotNil, nil
	}
	if !util.IsValidAndNotZeroAddress(address) {
		return util.ErrInvalidOrZeroAddress, nil
	}
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
	if err := CheckReceiptStatus(hash); err!=nil {
		return err, nil
	}
	return nil, &hash
}

func GetAddressFromCns(name string) (error, *common.Address) {
	if name == "" {
		return util.ErrParamShouldNotNil, nil
	}
	err, instance := GetCnsInstance()
	if err != nil {
		return err, nil
	}
	address, err := instance.Get(nil, name)
	if err != nil {
		return err, nil
	}
	if !util.IsValidAndNotZeroAddress(address) {
		return util.ErrInvalidOrZeroAddress, nil
	}
	return nil, &address
}