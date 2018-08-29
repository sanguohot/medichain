package service

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"github.com/luopotaotao/IdValidator/src/idvalidator"
	"medichain/util"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/accounts"
	"medichain/chain"
	"medichain/etc"
)

type User struct {
	Uuid uuid.UUID
	Address common.Address
	PublicKey ecdsa.PublicKey
	idCartNoHash common.Hash
	AccountUrl accounts.URL
}

func AddUser(orgUuidStr string, idCartNo string, password string) (error, *User, *common.Hash) {
	if ok, err := idvalidator.Validate(idCartNo); !ok {
		return err, nil, nil
	}
	if password == "" {
		return util.ErrPwdRequire, nil, nil
	}
	orgUuid := uuid.UUID{}
	if orgUuidStr != "" {
		orgUuid, err := uuid.Parse(orgUuidStr)
		if err != nil {
			return err, nil, nil
		}
		isExist, err := chain.OrgsDataIsUuidExist(orgUuid)
		if err != nil {
			return err, nil, nil
		}
		if !isExist {
			return util.ErrOrgNotExist, nil, nil
		}
	}
	idCartNoHash := crypto.Keccak256Hash([]byte(idCartNo))
	userUuid, err := chain.UsersDataGetUuidByIdCartNoHash(idCartNoHash)
	if err != nil {
		return err, nil, nil
	}
	isExist, err := chain.UsersDataIsUuidExist(*userUuid)
	if err != nil {
		return err, nil, nil
	}
	if isExist {
		return util.ErrUserExist, nil, nil
	}
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return err, nil, nil
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return util.ErrPublicKeyTransform, nil, nil
	}
	store := keystore.NewKeyStore(etc.GetBcosKeystore(), keystore.LightScryptN, keystore.StandardScryptP)
	account, err := store.ImportECDSA(privateKey, password)
	if err != nil {
		return err, nil, nil
	}
	userUuidNew := uuid.New()
	err, hash := chain.ControllerAddUser(userUuidNew, orgUuid, idCartNoHash, account.Address, password)
	if err != nil {
		return err, nil, nil
	}
	user := User{
		Address: account.Address,
		Uuid: userUuidNew,
		AccountUrl: account.URL,
		PublicKey: *publicKeyECDSA,
		idCartNoHash: crypto.Keccak256Hash([]byte(password)),
	}
	return nil, &user, hash
}
