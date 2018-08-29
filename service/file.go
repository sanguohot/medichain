package service

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/ethereum/go-ethereum/crypto"
	"medichain/util"
	"medichain/chain"
	"medichain/data"
)

type FileAction struct {
	Uuid uuid.UUID
	Keccak256Hash common.Hash
	Sha256Hash common.Hash
	txHash common.Hash
}

func checkKeccak256Hash(hash common.Hash) (error) {
	fileUuid, err := chain.FilesDataGetUuidByKeccak256Hash(hash)
	if err != nil {
		return err
	}
	err = checkFileUuid(*fileUuid)
	if err != nil {
		return err
	}
	return nil
}
func checkFileUuid(fileUuid uuid.UUID) error {
	isExist, err := chain.FilesDataIsUuidExist(fileUuid)
	if err != nil {
		return err
	}
	if isExist {
		return util.ErrFileExist
	}
	return nil
}
func checkSha256Hash(hash common.Hash) (error) {
	fileUuid, err := chain.FilesDataGetUuidBySha256Hash(hash)
	if err != nil {
		return err
	}
	err = checkFileUuid(*fileUuid)
	if err != nil {
		return err
	}
	return nil
}
func checkHash(keccak256Hash, sha256Hash common.Hash) error {
	err := checkKeccak256Hash(keccak256Hash)
	if err != nil {
		return err
	}
	err = checkSha256Hash(sha256Hash)
	if err != nil {
		return err
	}
	return nil
}
func AddFile(ownerUuidStr, addressStr, password, fileType, fileDesc string, file []byte, sha256HashStr string) (error, *FileAction) {
	ownerUuid, err := uuid.Parse(ownerUuidStr)
	if err != nil {
		return err, nil
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
	keccak256Hash := crypto.Keccak256Hash(file)
	err = checkHash(keccak256Hash, sha256Hash)
	if err != nil {
		return err, nil
	}
	// 上传到大数据库
	err = data.UploadToBigDataCenter(file)
	if err != nil {
		return err, nil
	}
	fileDescBytes32_4, err := util.StringToBytes32_4(fileDesc)
	if err != nil {
		return err, nil
	}
	fileUuid := uuid.New()
	err, txHash := chain.ControllerAddFile(fileUuid, ownerUuid, fileTypeHash, *fileDescBytes32_4, keccak256Hash, sha256Hash, address, password)
	if err != nil {
		return err, nil
	}
	fileAction := FileAction{
		Uuid: fileUuid,
		Keccak256Hash: keccak256Hash,
		Sha256Hash: sha256Hash,
		txHash: *txHash,
	}
	return nil, &fileAction
}

func AddFileSign(fileUuidStr, addressStr, password, keccak256HashStr string) (error, *common.Hash) {
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
	err = checkKeccak256Hash(keccak256Hash)
	if err != nil {
		return err, nil
	}
	err, txHash := chain.ControllerAddSign(fileUuid, keccak256Hash, address, password)
	if err != nil {
		return err, nil
	}
	return nil, txHash
}

func GetFile(fileUuidStr string) (error, []byte) {
	fileUuid, err := uuid.Parse(fileUuidStr)
	if err != nil {
		return err, nil
	}
	err = checkFileUuid(fileUuid)
	if err != nil {
		return err, nil
	}
	hash, err := chain.FilesDataGetSha256Hash(fileUuid)
	if err != nil {
		return err, nil
	}
	file, err := data.DownloadFromBigDataCenter(hash.Hex()[2:])
	if err != nil {
		return err, nil
	}
	return nil, file
}