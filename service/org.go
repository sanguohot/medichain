package service

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/sanguohot/medichain/util"
	"github.com/sanguohot/medichain/chain"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type OrgAction struct {
	UUID      uuid.UUID `json:"uuid"`
	Address   common.Address `json:"address"`
	PublicKey string `json:"publicKey"`
	NameHash  common.Hash `json:"nameHash"`
	TransactionHash    common.Hash `json:"transactionHash"`
}
func checkOrgNameHash(hash common.Hash) error {
	orgUuid, err := chain.OrgsDataGetUuidByNameHash(hash)
	if err != nil {
		return err
	}
	isExist, err := chain.OrgsDataIsUuidExist(*orgUuid)
	if err != nil {
		return err
	}
	if isExist {
		return util.ErrOrgExist
	}
	return err
}
func AddOrg(name string, password string) (error, *OrgAction) {
	bytes32_4, err := util.StringToBytes32_4(name)
	if err != nil {
		return err, nil
	}
	hash := util.Bytes32_4Hash(*bytes32_4)
	err = checkOrgNameHash(hash)
	if err != nil {
		return err, nil
	}
	orgUuid := uuid.New()
	_, publicKeyECDSA, address, err := util.NewWallet(password)
	err, txHash := chain.ControllerAddOrg(orgUuid, *bytes32_4, *address, password)
	if err != nil {
		return err, nil
	}
	org := OrgAction{
		Address: *address,
		UUID: orgUuid,
		PublicKey: hexutil.Encode(crypto.FromECDSAPub(publicKeyECDSA))[4:],
		NameHash: hash,
		TransactionHash: *txHash,
	}
	return nil, &org
}