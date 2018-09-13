package service

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"github.com/sanguohot/medichain/chain"
	"github.com/sanguohot/medichain/datacenter"
	"github.com/sanguohot/medichain/etc"
	"github.com/sanguohot/medichain/util"
)
type FileAddLogSimpleItem struct {
	FileUUID string `json:"FileUuid"`
	OwnerUUID string `json:"OwnerUuid"`
	OrgUUID string `json:"OrgUuid"`
	FileType string `json:"FileType"`
	CreateTime uint64 `json:"CreateTime"`
}
type FileAddLogSimpleAction struct {
	Total uint64 `json:"Total"`
	List []FileAddLogSimpleItem `json:"List"`
}
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

func GetFileAddLogList(idCartNo, orgUuidStr, startStr, limitStr string) (error, *FileAddLogSimpleAction) {
	if orgUuidStr != ""{
		orgUuid, err := uuid.Parse(orgUuidStr)
		if err != nil {
			return err, nil
		}
		exist, err := chain.OrgsDataIsUuidExist(orgUuid)
		if err != nil {
			return err, nil
		}
		if !exist {
			return util.ErrOrgNotExist, nil
		}
	}

	var (
		ownerUuidStr string
		fl []datacenter.FileAddLog
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
	err, total := datacenter.SqliteGetFileAddLogTotal("", orgUuidStr, ownerUuidStr)
	if err != nil {
		return err, nil
	}
	if total > 0 {
		err, fl = datacenter.SqliteGetFileAddLogList("", orgUuidStr, ownerUuidStr, startBig.Uint64(), limitBig.Uint64())
		if err != nil {
			return err, nil
		}
	}

	return nil, wrapperToFileAddLogSimpleAction(total, fl)
}
func wrapperToFileAddLogSimpleAction(total uint64, list []datacenter.FileAddLog) *FileAddLogSimpleAction {
	wrapper := FileAddLogSimpleAction{}
	wrapper.Total = total
	for _, item := range list {
		wrapper.List = append(wrapper.List, FileAddLogSimpleItem{
			FileUUID: item.FileUuid,
			OwnerUUID: item.OwnerUuid,
			OrgUUID: item.OrgUuid,
			FileType: etc.FileTypeMap[common.HexToHash(item.FileTypeHash)],
			CreateTime: item.CreateTime,
		})
	}
	return &wrapper
}

func GetFileAddLogDetail(fileUuidStr string) (error, *datacenter.FileAddLog) {
	_, err := uuid.Parse(fileUuidStr)
	if err != nil {
		return err, nil
	}
	err, fl := datacenter.SqliteGetFileAddLogList(fileUuidStr, "", "", 0, 0)
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