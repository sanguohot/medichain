package util

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"math/rand"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	//"bytes"
	"math/big"
	"medichain/etc"
	"time"
	"bytes"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"encoding/hex"
)

func GetKeyJsonFromStore(address common.Address, password string) ([]byte, error) {
	store := keystore.NewKeyStore(etc.GetBcosKeystore(), keystore.LightScryptN, keystore.StandardScryptP)
	account, err := store.Find(accounts.Account{Address: address})
	if err != nil {
		return nil, err
	}
	keyBytes, err := store.Export(account, password, password)
	if err != nil {
		return nil, err
	}
	return  keyBytes, nil
}

func GetTransactOptsFromStore(address common.Address, password string, nonce uint64) (*bind.TransactOpts, error) {
	if !IsValidAndNotZeroAddress(address) {
		return nil, ErrInvalidOrZeroAddress
	}
	rand.Seed(time.Now().Unix())
	keyBytes, err := GetKeyJsonFromStore(address, password)
	if err != nil {
		return nil, err
	}
	auth, err := bind.NewTransactor(bytes.NewReader(keyBytes), password)
	if err != nil {
		return nil, err
	}
	auth.Value = nil
	auth.GasLimit = uint64(300000000) // in units
	if nonce > 0 {
		auth.Nonce = big.NewInt(int64(nonce))
	}else {
		auth.Nonce = big.NewInt(rand.Int63n(100000000)) // 也是fisco-bcos的randomid字段，这里按照web3.js计算随机数
	}
	auth.GasPrice = nil
	auth.From = address
	//auth.BlockLimit = uint64(750)
	return  auth, nil
}

func GetDefaultTransactOptsFromStore() (*bind.TransactOpts, error) {
	fromAddressStr := etc.GetBcosOwnerAddress()
	if fromAddressStr == "" {
		return nil, ErrConfigItemRequire
	}
	fromAddress := common.HexToAddress(fromAddressStr)
	password := etc.GetBcosOwnerPassword()
	return GetTransactOptsFromStore(fromAddress, password, 0)
}

func GetPublicKeyFromStore(address common.Address, password string) (*ecdsa.PublicKey, error) {
	keyBytes, err := GetKeyJsonFromStore(address, password)
	if err != nil {
		return nil, err
	}
	key, err := keystore.DecryptKey(keyBytes, password)
	if err != nil {
		return nil, err
	}
	publicKey := key.PrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, err
	}
	return publicKeyECDSA, nil
}

func GetPublicKeyBytesFromStore(address common.Address, password string) ([]byte, error) {
	pub, err := GetPublicKeyFromStore(address, password)
	if err != nil {
		return nil, err
	}
	return crypto.FromECDSAPub(pub), nil
}

func GetPublicKeyBytes32_2FromStore(address common.Address, password string) (*[2][32]byte, error) {
	pub, err := GetPublicKeyBytesFromStore(address, password)
	// 64字节公钥加上固定首字节
	if len(pub) != 65 {
		return nil, ErrPublicKeyShouldEqualTo64Bytes
	}
	if err != nil {
		return nil, err
	}
	return BytesToBytes32_2(pub[1:])
}

func NewWallet(password string) (*ecdsa.PrivateKey, *ecdsa.PublicKey, *common.Address, error) {
	if password == "" {
		return nil, nil, nil, ErrPwdRequire
	}
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, nil, nil, err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, nil, nil, ErrPublicKeyTransform
	}
	store := keystore.NewKeyStore(etc.GetBcosKeystore(), keystore.LightScryptN, keystore.StandardScryptP)
	account, err := store.ImportECDSA(privateKey, password)
	if err != nil {
		return nil, nil, nil, err
	}

	return privateKey, publicKeyECDSA, &account.Address, nil
}

func Keccak256SignWithContent(content []byte, address common.Address, password string) ([]byte, error) {
	return SignWithHash(crypto.Keccak256Hash(content), address, password)
}
func Sha256SignWithContent(content []byte, address common.Address, password string) ([]byte, error) {
	return SignWithHash(Sha256Hash(content), address, password)
}
func SignWithHash(hash common.Hash, address common.Address, password string) ([]byte, error) {
	store := keystore.NewKeyStore(etc.GetBcosKeystore(), keystore.LightScryptN, keystore.StandardScryptP)
	account, err := store.Find(accounts.Account{Address: address})
	if err != nil {
		return nil, err
	}
	signature, err := store.SignHashWithPassphrase(account, password, hash.Bytes())
	if err != nil {
		return nil, err
	}
	return hex.DecodeString(hexutil.Encode(signature)[2:])
}