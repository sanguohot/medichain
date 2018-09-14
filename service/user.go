package service

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"github.com/sanguohot/medichain/chain"
	"github.com/sanguohot/medichain/util"
)

type UserAction struct {
	UUID         uuid.UUID `json:"uuid"`
	Address      common.Address `json:"address"`
	PublicKey    string `json:"publicKey"`
	IDCartNoHash common.Hash `json:"idCartNoHash"`
	TransactionHash       common.Hash `json:"transactionHash"`
}
func AddUser(orgUuidStr string, idCartNo string, password string) (error, *UserAction) {
	if idCartNo == "" {
		return fmt.Errorf("身份证不能为空"), nil
	}
	//if ok, err := idvalidator.Validate(idCartNo); !ok {
	//	return fmt.Errorf("身份证%s", err.Error()), nil
	//}
	if password == "" {
		return util.ErrPwdRequire, nil
	}
	orgUuid := [16]byte{}
	// do not user the following code, because it is not zero to all bytes
	//orgUuid := uuid.UUID{}
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
		UUID: userUuidNew,
		PublicKey: hexutil.Encode(crypto.FromECDSAPub(publicKeyECDSA))[4:],
		IDCartNoHash: idCartNoHash,
		TransactionHash: *txHash,
	}
	return nil, &user
}
