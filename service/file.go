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

type FileSignerAndDataAction struct {
	idl []uuid.UUID
	rl []common.Hash
	sl []common.Hash
	vl []uint8
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
	err = requireHashNotExist(keccak256Hash, sha256Hash)
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
	err = requireKeccak256HashNotExist(keccak256Hash)
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
	err = requireFileUuidExist(fileUuid)
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

func GetFileSignerAndDataList(fileUuidStr string, startStr string, limitStr string) (error, *FileSignerAndDataAction) {
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
	err, startBig, limitBig := TransformPagingParamFromStringToBigInt(startStr, limitStr)
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
		idlNew []uuid.UUID
		rlNew []common.Hash
		slNew []common.Hash
	)
	for i := 0; i < len(idl); i++ {
		idlNew[i] = uuid.UUID(idl[i])
		rlNew[i] = common.Hash(rl[i])
		slNew[i] = common.Hash(sl[i])
	}
	return &FileSignerAndDataAction{
		idl: idlNew,
		rl: rlNew,
		sl: slNew,
		vl: vl,
	}
}