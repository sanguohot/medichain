package service

import (
	"medichain/chain"
	"medichain/datacenter"
	"github.com/ethereum/go-ethereum/crypto"
)

func SetFileAddLogList() error {
	err, list := chain.ChainGetFileAddLogListAll()
	if err != nil {
		return err
	}
	err = datacenter.SqliteSetFileAddLogList(list)
	if err != nil {
		return err
	}
	return nil
}

func GetFileAddLogList(idCartNo, startStr, limitStr string) (error, []datacenter.FileAddLog) {
	var (
		ownerUuidStr string
	)
	if idCartNo != "" {
		uuid, err := chain.FilesDataGetUuidByKeccak256Hash(crypto.Keccak256Hash([]byte(idCartNo)))
		if err != nil {
			return err, nil
		}
		ownerUuidStr = uuid.String()
	}
	return datacenter.SqliteGetFileAddLogList(ownerUuidStr)
}
