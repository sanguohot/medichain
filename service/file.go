package service

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"github.com/sanguohot/medichain/chain"
	"github.com/sanguohot/medichain/datacenter"
	"github.com/sanguohot/medichain/etc"
	"github.com/sanguohot/medichain/util"
)

type FileAction struct {
	UUID            uuid.UUID `json:"uuid"`
	Keccak256Hash   common.Hash `json:"keccak256Hash"`
	Sha256Hash      common.Hash `json:"sha256Hash"`
	TransactionHash common.Hash `json:"transactionHash"`
}
type FileSignerAction struct {
	UUID uuid.UUID `json:"uuid"`
	//Type string `json:"type"`
}
type FileSignerAndDataAction struct {
	Signers []FileSignerAction `json:"signers"`
	Signatures []string `json:"signatures"`
}
type FileAddSignAction struct {
	Signature string `json:"signature"`
	TransactionHash common.Hash `json:"transactionHash"`
}
func requireKeccak256HashNotExist(hash common.Hash) (error) {
	fileUuid, err := chain.FilesDataGetUuidByKeccak256Hash(hash)
	if err != nil {
		return err
	}
	err = requireFileUuidNotExist(*fileUuid)
	if err != nil {
		return err
	}
	return nil
}
func requireFileUuidNotExist(fileUuid uuid.UUID) error {
	isExist, err := chain.FilesDataIsUuidExist(fileUuid)
	if err != nil {
		return err
	}
	if isExist {
		return util.ErrFileExist
	}
	return nil
}
func requireFileUuidExist(fileUuid uuid.UUID) error {
	isExist, err := chain.FilesDataIsUuidExist(fileUuid)
	if err != nil {
		return err
	}
	if !isExist {
		return util.ErrFileNotExist
	}
	return nil
}
func requireSha256HashNotExist(hash common.Hash) (error) {
	fileUuid, err := chain.FilesDataGetUuidBySha256Hash(hash)
	if err != nil {
		return err
	}
	err = requireFileUuidNotExist(*fileUuid)
	if err != nil {
		return err
	}
	return nil
}
func requireHashNotExist(keccak256Hash, sha256Hash common.Hash) error {
	err := requireKeccak256HashNotExist(keccak256Hash)
	if err != nil {
		return err
	}
	err = requireSha256HashNotExist(sha256Hash)
	if err != nil {
		return err
	}
	return nil
}
func AddFile(ownerUuidStr, orgUuidStr, addressStr, password, fileType, fileDesc string, file []byte, sha256HashStr string) (error, *FileAction) {
	ownerUuid, err := uuid.Parse(ownerUuidStr)
	if err != nil {
		return err, nil
	}
	isExist, err := chain.UsersDataIsUuidExist(ownerUuid)
	if err != nil {
		return err, nil
	}
	if !isExist {
		return util.ErrUserNotExist, nil
	}
	var orgUuid uuid.UUID
	if orgUuidStr != "" {
		orgUuid, err = uuid.Parse(orgUuidStr)
		if err != nil {
			return err, nil
		}
		isExist, err = chain.OrgsDataIsUuidExist(orgUuid)
		if err != nil {
			return err, nil
		}
		if !isExist {
			return util.ErrOrgNotExist, nil
		}
	}

	if !common.IsHexAddress(addressStr) {
		return util.ErrInvalidAddress, nil
	}
	address := common.HexToAddress(addressStr)
	sha256Hash := common.HexToHash(sha256HashStr)
	if sha256Hash != util.Sha256Hash(file) {
		return util.ErrFileUploadNotComplete, nil
	}
	// define them and check file type
	fileTypeHash := crypto.Keccak256Hash([]byte(fileType))
	if etc.FileTypeMap[fileTypeHash] != fileType {
		return fmt.Errorf("%s ===> %s", util.ErrFileTypeNotSupport.Error(), fileType), nil
	}
	keccak256Hash := crypto.Keccak256Hash(file)
	err = requireHashNotExist(keccak256Hash, sha256Hash)
	if err != nil {
		return err, nil
	}
	fileDescBytes32_4, err := util.StringToBytes32_4(fileDesc)
	if err != nil {
		return err, nil
	}
	fileUuid := uuid.New()
	err, txHash := chain.ControllerAddFile(fileUuid, ownerUuid, orgUuid, fileTypeHash, *fileDescBytes32_4, keccak256Hash, sha256Hash, address, password)
	if err != nil {
		return err, nil
	}
	// 上传到大数据库
	err = datacenter.UploadToBigDataCenter(file)
	if err != nil {
		return err, nil
	}
	fileAction := FileAction{
		UUID: fileUuid,
		Keccak256Hash: keccak256Hash,
		Sha256Hash: sha256Hash,
		TransactionHash: *txHash,
	}
	return nil, &fileAction
}

func AddFileSign(fileUuidStr, addressStr, password, keccak256HashStr string) (error, *FileAddSignAction) {
	fileUuid, err := uuid.Parse(fileUuidStr)
	if err != nil {
		return err, nil
	}
	isExist, err := chain.FilesDataIsUuidExist(fileUuid)
	if err != nil {
		return err, nil
	}
	if !isExist {
		return util.ErrFileNotExist, nil
	}
	if !common.IsHexAddress(addressStr) {
		return util.ErrInvalidAddress, nil
	}
	address := common.HexToAddress(addressStr)
	keccak256HashFromChain, err := chain.FilesDataGetKeccak256Hash(fileUuid)
	if err != nil {
		return err, nil
	}
	keccak256Hash := common.HexToHash(keccak256HashStr)
	if *keccak256HashFromChain != keccak256Hash {
		return util.ErrFileHashNotMatch, nil
	}
	err, r, s, v, txHash := chain.ControllerAddSign(fileUuid, keccak256Hash, address, password)
	if err != nil {
		return err, nil
	}
	signature := hexutil.Encode(util.RSVtoSig(r, s, v))[2:]
	return nil, &FileAddSignAction{
		Signature: signature,
		TransactionHash: *txHash,
	}
}

func GetFile(fileUuidStr string) (error, []byte) {
	fileUuid, err := uuid.Parse(fileUuidStr)
	if err != nil {
		return err, nil
	}
	err = requireFileUuidExist(fileUuid)
	if err != nil {
		return err, nil
	}
	hash, err := chain.FilesDataGetSha256Hash(fileUuid)
	if err != nil {
		return err, nil
	}
	file, err := datacenter.DownloadFromBigDataCenter(*hash)
	if err != nil {
		return err, nil
	}
	return nil, file
}

func GetFileSignerAndDataList(fileUuidStr string, startStr, limitStr string) (error, *FileSignerAndDataAction) {
	fileUuid, err := uuid.Parse(fileUuidStr)
	if err != nil {
		return err, nil
	}
	err = requireFileUuidExist(fileUuid)
	if err != nil {
		return err, nil
	}
	size, err := chain.FilesDataGetFileSignerSize(fileUuid)
	if err != nil {
		return err, nil
	}
	err, startBig, limitBig := transformPagingParamFromStringToBigInt(startStr, limitStr)
	if err != nil {
		return err, nil
	}
	if size.Cmp(startBig) != 1 {
		return util.ErrFileSignListOutOfIndex, nil
	}
	err, idl, rl, sl, vl := chain.ControllerGetFileSignersAndDataByUuid(fileUuid, startBig, limitBig)
	if err != nil {
		return err, nil
	}
	return nil, getFileSignerAndDataActionByChainData(idl, rl, sl, vl)
}

func getFileSignerAndDataActionByChainData(idl [][16]byte, rl [][32]byte, sl [][32]byte, vl []uint8) *FileSignerAndDataAction {
	var (
		signers []FileSignerAction = make([]FileSignerAction, len(idl))
		signatures []string = make([]string, len(idl))
	)
	for i := 0; i < len(idl); i++ {
		signers[i] = FileSignerAction{
			UUID: idl[i],
		}
		signatures[i] = hexutil.Encode(util.RSVtoSig(rl[i], sl[i], vl[i]))[2:]
	}
	return &FileSignerAndDataAction{
		Signers: signers,
		Signatures: signatures,
	}
}