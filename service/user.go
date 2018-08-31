package service

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"github.com/luopotaotao/IdValidator/src/idvalidator"
	"medichain/util"
	"medichain/chain"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type UserAction struct {
	Uuid uuid.UUID
	Address common.Address
	PublicKey string
	IdCartNoHash common.Hash
	TxHash common.Hash
}

func AddUser(orgUuidStr string, idCartNo string, password string) (error, *UserAction) {
	if ok, err := idvalidator.Validate(idCartNo); !ok {
		return err, nil
	}
	if password == "" {
		return util.ErrPwdRequire, nil
	}
	orgUuid := uuid.UUID{}
	if orgUuidStr != "" {
		orgUuid, err := uuid.Parse(orgUuidStr)
		if err != nil {
			return err, nil
		}
		isExist, err := chain.OrgsDataIsUuidExist(orgUuid)
		if err != nil {
			return err, nil
		}
		if !isExist {
			return util.ErrOrgNotExist, nil
		}
	}
	idCartNoHash := crypto.Keccak256Hash([]byte(idCartNo))
	userUuid, err := chain.UsersDataGetUuidByIdCartNoHash(idCartNoHash)
	if err != nil {
		return err, nil
	}
	isExist, err := chain.UsersDataIsUuidExist(*userUuid)
	if err != nil {
		return err, nil
	}
	if isExist {
		return util.ErrUserExist, nil
	}
	_, publicKeyECDSA, address, err := util.NewWallet(password)
	userUuidNew := uuid.New()
	err, txHash := chain.ControllerAddUser(userUuidNew, orgUuid, idCartNoHash, *address, password)
	if err != nil {
		return err, nil
	}
	user := UserAction{
		Address: *address,
		Uuid: userUuidNew,
		PublicKey: hexutil.Encode(crypto.FromECDSAPub(publicKeyECDSA))[4:],
		IdCartNoHash: idCartNoHash,
		TxHash: *txHash,
	}
	return nil, &user
}
