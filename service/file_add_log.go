package service

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"medichain/chain"
	"medichain/datacenter"
	"medichain/etc"
	"medichain/util"
)
type FileAddLogSimpleAction struct {
	FileUuid	string
	OwnerUuid	string
	OrgUuid	string
	FileType	string
	CreateTime	uint64
}
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

func GetFileAddLogList(idCartNo, startStr, limitStr string) (error, []FileAddLogSimpleAction) {
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
	err, fl := datacenter.SqliteGetFileAddLogList("", ownerUuidStr, startBig.Uint64(), limitBig.Uint64())
	if err != nil {
		return err, nil
	}
	return nil, wrapperToFileAddLogSimpleAction(fl)
}
func wrapperToFileAddLogSimpleAction(list []datacenter.FileAddLog) []FileAddLogSimpleAction {
	var wrapperList []FileAddLogSimpleAction
	for _, item := range list {
		wrapperList = append(wrapperList, FileAddLogSimpleAction{
			FileUuid: item.FileUuid,
			OwnerUuid: item.OwnerUuid,
			OrgUuid: item.OrgUuid,
			FileType: etc.FileTypeMap[common.HexToHash(item.FileTypeHash)],
			CreateTime: item.CreateTime,
		})
	}
	return wrapperList
}

func GetFileAddLogDetail(fileUuidStr string) (error, *datacenter.FileAddLog) {
	_, err := uuid.Parse(fileUuidStr)
	if err != nil {
		return err, nil
	}
	err, fl := datacenter.SqliteGetFileAddLogList(fileUuidStr, "", 0, 0)
	if err != nil {
		return err, nil
	}
	if len(fl) == 0 {
		return util.ErrFileNotExist, nil
	}
	detail := fl[0]
	detail.FileType = ""
	return nil, &fl[0]
}