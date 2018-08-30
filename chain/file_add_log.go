package chain

import (
	"medichain/contracts/medi"
	"fmt"
	"medichain/etc"
	"medichain/datacenter"
	"github.com/ethereum/go-ethereum"
	"math/big"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"strings"
	"github.com/ethereum/go-ethereum/common"
	"context"
	"github.com/google/uuid"
	"medichain/util"
	"github.com/ethereum/go-ethereum/common/hexutil"
)


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
	fl := []datacenter.FileAddLog{}
	fmt.Println("logs count ===>", len(logs))
	for i, vLog := range logs {
		var event struct {
			uuid [16]byte
			ownerUuid [16]byte
			uploaderUuid [16]byte
			fileType [32]byte
			fileDesc [4][32]byte
			keccak256Hash [32]byte
			sha256Hash [32]byte
			r [32]byte
			s [32]byte
			v uint8
			time uint
		}
		// 参数一是事件参数构造的对象
		// 参数二是事件函数名，不是合约函数名
		// 参数三是事件的数据
		err := contractAbi.Unpack(&event, "onAddFile", vLog.Data)
		if err != nil {
			return err, nil
		}
		fileUuid := uuid.UUID(event.uuid)
		ownerUuid := uuid.UUID(event.ownerUuid)
		uploaderUuid := uuid.UUID(event.uploaderUuid)
		fileType := common.Hash(event.fileType)
		fileDesc := util.Bytes32_4ToString(event.fileDesc)
		keccak256Hash := common.Hash(event.keccak256Hash)
		sha256Hash := common.Hash(event.sha256Hash)
		r := common.Hash(event.r)
		s := common.Hash(event.s)
		signature := util.RSVtoSig(r, s, event.v)
		signers := uploaderUuid
		fmt.Println(event)
		f := datacenter.FileAddLog {
			FileUuid: fileUuid.String(),
			OwnerUuid: ownerUuid.String(),
			UploaderUuid: uploaderUuid.String(),
			FileType: fileType.Hex(),
			FileDesc: fileDesc,
			Keccak256Hash: keccak256Hash.Hex(),
			Sha256Hash: sha256Hash.Hex(),
			CreateTime: event.time,
			Signature: hexutil.Encode(signature)[2:],
			Signer: signers.String(),
			BlockNum: vLog.BlockNumber,
			BlockHash: vLog.BlockHash.Hex(),
			TransactionHash: vLog.TxHash.Hex(),
			ContractAddress: address.Hex(),
		}
		fl[i] = f
	}

	return nil, fl
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
	return ChainGetFileAddLogList(0, header.Number.Int64())
}