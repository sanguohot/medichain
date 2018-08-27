package service

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"github.com/luopotaotao/IdValidator/src/idvalidator"
	"errors"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/accounts"
)

var (
	ErrPwdRequire = errors.New("密码不能为空")
	ErrPublicKeyTransform = errors.New("公钥格式转换失败")
)
type User struct {
	Uuid uuid.UUID
	Address common.Address
	PublicKey ecdsa.PublicKey
	idCartNoHash common.Hash
	AccountUrl accounts.URL
}

func AddUser(orgUuid uuid.UUID, idCartNo string, password string) (error, *User, *common.Hash) {
	if ok, err := idvalidator.Validate(idCartNo); !ok {
		return err, nil, nil
	}
	if password == "" {
		return ErrPwdRequire, nil, nil
	}
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return err, nil, nil
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return ErrPublicKeyTransform, nil, nil
	}
	store := keystore.NewKeyStore("./keystore", keystore.LightScryptN, keystore.StandardScryptP)
	account, err := store.ImportECDSA(privateKey, password)
	if err != nil {
		return err, nil, nil
	}
	userUuid := uuid.New()
	user := User{
		Address: account.Address,
		Uuid: userUuid,
		AccountUrl: account.URL,
		PublicKey: *publicKeyECDSA,
		idCartNoHash: crypto.Keccak256Hash([]byte(password)),
	}

	return nil, &user, nil
}

func GetUserByUuid(uuid uuid.UUID) (error, *User) {
	return nil, nil
}

func GetUserByIdCartNo(idCartNo string) (error, *User) {
	return nil, nil
}

func GetUserByAddress(address common.Address) (error, *User) {
	return nil, nil
}