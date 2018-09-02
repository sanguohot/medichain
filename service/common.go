package service

import (
	"math/big"
	"medichain/util"
	"strconv"
)

func TransformPagingParamFromStringToBigInt(startStr string, limitStr string) (error, *big.Int, *big.Int) {
	if startStr == "" {
		startStr = "0"
	}
	if limitStr == "" {
		limitStr = "10"
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
