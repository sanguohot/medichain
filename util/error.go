package util

import "errors"

var (
	ErrConfigItemRequire = errors.New("没有找到配置项")
	ErrParamContractNotSupport = errors.New("合约名称不能为空")
	ErrParamCnsAddressShouldNotNil = errors.New("命名合约地址不能为空")
	ErrParamStringShouldNotNil = errors.New("字符串参数不能为空")
	ErrParamStringShouldNotMoreThan128Bytes = errors.New("字符串参数不能为空")
)
