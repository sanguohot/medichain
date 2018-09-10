package chain

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/google/uuid"
	"math/big"
	"medichain/contracts/medi"
	"medichain/datacenter"
	"medichain/etc"
	"medichain/util"
	"strings"
)

var addFileEventName = "onAddFile"
func ChainGetFileAddLogList(fromBlock int64, toBlock int64) (error, []datacenter.FileAddLog) {
	client, err := GetEthDialClient()
	if err != nil {
		return err, nil
	}
	err, address := GetAddressFromCns(etc.ContractFilesData)
	if err != nil {
		return err, nil
	}
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(fromBlock),
		ToBlock:   big.NewInt(toBlock),
		Addresses: []common.Address{
			*address,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		return err, nil
	}
	contractAbi, err := abi.JSON(strings.NewReader(string(medi.FilesDataABI)))
	if err != nil {
		return err, nil
	}
	var addFileTopic *common.Hash
	for _, item := range contractAbi.Events {
		if item.Name == addFileEventName {
			hash := item.Id()
			addFileTopic = &hash
		}
	}
	if addFileTopic == nil {
		return util.ErrMatchLogNotFound, nil
	}
	return rangeLogsToGetFileAddLogList(contractAbi, *address, logs, *addFileTopic)
}

func rangeLogsToGetFileAddLogList(contractAbi abi.ABI, address common.Address, logs []types.Log, topic common.Hash) (error, []datacenter.FileAddLog) {
	fmt.Printf("FilesData contract %s logs ===> %d\n", address.Hex(), len(logs))
	fl := []datacenter.FileAddLog{}
	for _, log := range logs {
		found := false
		for _, item := range log.Topics {
			if item == topic {
				found = true
				break
			}
		}
		if !found {
			continue
		}
		err, f := parseEventToFileAddLog(contractAbi, address, log)
		if err != nil {
			return err, nil
		}
		fl = append(fl, *f)
	}
	if len(fl) == 0 {
		return util.ErrMatchLogNotFound, nil
	}
	return nil, fl
}
func parseEventToFileAddLog(contractAbi abi.ABI, address common.Address, log types.Log) (error, *datacenter.FileAddLog) {
	var event struct {
		Uuid [16]byte
		OwnerUuid [16]byte
		UploaderUuid [16]byte
		OrgUuid [16]byte
		FileType [32]byte
		FileDesc [4][32]byte
		Keccak256Hash [32]byte
		Sha256Hash [32]byte
		R [32]byte
		S [32]byte
		V uint8
		Time *big.Int
	}
	fmt.Println("TxHash ===>", log.TxHash.Hex())
	// 参数一是事件参数构造的对象
	// 参数二是事件函数名，不是合约函数名
	// 参数三是事件的数据
	err := contractAbi.Unpack(&event, addFileEventName, log.Data)
	if err != nil {
		return err, nil
	}
	fileUuid := uuid.UUID(event.Uuid)
	ownerUuid := uuid.UUID(event.OwnerUuid)
	uploaderUuid := uuid.UUID(event.UploaderUuid)
	orgUuid := uuid.UUID(event.OrgUuid)
	fileTypeHash := common.Hash(event.FileType)
	fileDesc := util.Bytes32_4ToString(event.FileDesc)
	keccak256Hash := common.Hash(event.Keccak256Hash)
	sha256Hash := common.Hash(event.Sha256Hash)
	r := common.Hash(event.R)
	s := common.Hash(event.S)
	signature := util.RSVtoSig(r, s, event.V)
	signers := uploaderUuid
	f := datacenter.FileAddLog {
		FileUuid: fileUuid.String(),
		OwnerUuid: ownerUuid.String(),
		UploaderUuid: uploaderUuid.String(),
		OrgUuid: orgUuid.String(),
		FileTypeHash: fileTypeHash.Hex(),
		FileType: etc.FileTypeMap[fileTypeHash],
		FileDesc: fileDesc,
		Keccak256Hash: keccak256Hash.Hex(),
		Sha256Hash: sha256Hash.Hex(),
		CreateTime: event.Time.Uint64(),
		Signature: hexutil.Encode(signature)[2:],
		Signer: signers.String(),
		BlockNum: log.BlockNumber,
		BlockHash: log.BlockHash.Hex(),
		TransactionHash: log.TxHash.Hex(),
		ContractAddress: address.Hex(),
	}
	return nil, &f
}
func ChainGetFileAddLogListAll() (error, []datacenter.FileAddLog) {
	client, err := GetEthDialClient()
	if err != nil {
		return err, nil
	}
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return err, nil
	}
	err, max := datacenter.SqliteGetFileAddLogMaxBlockNum()
	if err != nil {
		return err, nil
	}
	if max >= header.Number.Uint64() {
		return nil, []datacenter.FileAddLog{}
	}
	return ChainGetFileAddLogList(int64(max), header.Number.Int64())
}