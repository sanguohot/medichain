package util

import "errors"

var (
	ErrConfigItemRequire = errors.New("没有找到配置项")
	ErrParamContractNotSupport = errors.New("合约名称不能为空")
	ErrParamCnsAddressShouldNotNil = errors.New("命名合约地址不能为空")
	ErrParamShouldNotNil = errors.New("参数不能为空")
	ErrParamShouldNotMoreThan128Bytes = errors.New("参数不能大于128字节")
	ErrParamShouldNotMoreThan64Bytes = errors.New("参数不能大于64字节")
	ErrPublicKeyShouldEqualTo64Bytes = errors.New("公钥必须为64字节")
	ErrPwdRequire = errors.New("密码不能为空")
	ErrPublicKeyTransform = errors.New("公钥格式转换失败")
	ErrUserExist = errors.New("用户已注册")
	ErrOrgExist = errors.New("机构已注册")
	ErrOrgNotExist = errors.New("机构不存在")
)
