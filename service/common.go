package service

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/sanguohot/medichain/chain"
	"github.com/sanguohot/medichain/util"
	"github.com/sanguohot/medichain/zap"
	"math/big"
	"strconv"
)
var (
	defaultStartStr = "0"
	defaultLimitStr = "10"
)
func transformPagingParamFromStringToBigInt(startStr, limitStr string) (error, *big.Int, *big.Int) {
	if startStr == "" {
		startStr = defaultStartStr
	}
	if limitStr == "" {
		limitStr = defaultLimitStr
	}
	startInt64, err := strconv.ParseInt(startStr, 10, 64)
	if err != nil {
		return err, nil, nil
	}
	limitInt64, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		return err, nil, nil
	}
	if startInt64 < 0 || limitInt64 <= 0 {
		return util.ErrParamPagingInvalid, nil, nil
	}
	return nil, big.NewInt(startInt64), big.NewInt(limitInt64)
}

func transformTimeParamFromStringToUint64(fromTimeStr, toTimeStr string) (error, uint64, uint64) {
	if fromTimeStr == "" {
		fromTimeStr = "0"
	}
	if toTimeStr == "" {
		toTimeStr = "0"
	}
	fromTimeInt64, err := strconv.ParseInt(fromTimeStr, 10, 64)
	if err != nil {
		return err, 0, 0
	}
	toTimeInt64, err := strconv.ParseInt(toTimeStr, 10, 64)
	if err != nil {
		return err, 0, 0
	}
	if fromTimeInt64 < 0 || toTimeInt64 < 0 {
		return util.ErrParamPagingInvalid, 0, 0
	}
	return nil, uint64(fromTimeInt64), uint64(toTimeInt64)
}

func CheckTransactionResult(hashStr string) (bool) {
	hash := common.HexToHash(hashStr)
	emptyHash := common.Hash{}
	if hash == emptyHash {
		zap.Sugar.Error(util.ErrParamdInvalid.Error())
		return false
	}
	if err := chain.CheckReceiptStatusWithoutWait(hash); err!=nil {
		zap.Sugar.Error(err.Error())
		return false
	}
	return true
}