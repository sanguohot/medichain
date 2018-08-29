package service

import (
	"github.com/ethereum/go-ethereum/common"
	"crypto"
	"github.com/google/uuid"
	"medichain/util"
	"medichain/chain"
)

type OrgAction struct {
	Uuid uuid.UUID
	Address common.Address
	PublicKey crypto.PublicKey
	nameHash common.Hash
	txHash common.Hash
}

func AddOrg(name string, password string) (error, *OrgAction) {
	bytes32_4, err := util.StringToBytes32_4(name)
	if err != nil {
		return err, nil
	}
	hash := util.Bytes32_4Hash(*bytes32_4)
	orgUuid, err := chain.OrgsDataGetUuidByNameHash(hash)
	if err != nil {
		return err, nil
	}
	isExist, err := chain.OrgsDataIsUuidExist(*orgUuid)
	if err != nil {
		return err, nil
	}
	if isExist {
		return util.ErrOrgExist, nil
	}
	orgUuidNew := uuid.New()
	_, publicKeyECDSA, address, err := util.NewWallet(password)
	err, txHash := chain.ControllerAddOrg(orgUuidNew, *bytes32_4, *address, password)
	if err != nil {
		return err, nil
	}
	org := OrgAction{
		Address: *address,
		Uuid: orgUuidNew,
		PublicKey: *publicKeyECDSA,
		nameHash: hash,
		txHash: *txHash,
	}
	return nil, &org
}