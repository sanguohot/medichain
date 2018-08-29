package service

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"github.com/luopotaotao/IdValidator/src/idvalidator"
	"medichain/util"
	"crypto/ecdsa"
	"medichain/chain"
)

type User struct {
	Uuid uuid.UUID
	Address common.Address
	PublicKey ecdsa.PublicKey
	idCartNoHash common.Hash
	txHash common.Hash
}

func AddUser(orgUuidStr string, idCartNo string, password string) (error, *User) {
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
	user := User{
		Address: *address,
		Uuid: userUuidNew,
		PublicKey: *publicKeyECDSA,
		idCartNoHash: idCartNoHash,
		txHash: *txHash,
	}
	return nil, &user
}
