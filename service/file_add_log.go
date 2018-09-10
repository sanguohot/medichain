package service

import (
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"medichain/chain"
	"medichain/datacenter"
)

func SetFileAddLogList() error {
	err, list := chain.ChainGetFileAddLogListAll()
	if err != nil {
		return err
	}
	fmt.Print("now set file add log list to sqlite...")
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
		uuid, err := chain.UsersDataGetUuidByIdCartNoHash(crypto.Keccak256Hash([]byte(idCartNo)))
		if err != nil {
			return err, nil
		}
		ownerUuidStr = uuid.String()
	}
	err, startBig, limitBig := TransformPagingParamFromStringToBigInt(startStr, limitStr)
	if err != nil {
		return err, nil
	}
	return datacenter.SqliteGetFileAddLogList(ownerUuidStr, startBig.Uint64(), limitBig.Uint64())
}
