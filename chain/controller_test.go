package chain

import (
	"testing"
	"medichain/etc"
	"github.com/ethereum/go-ethereum/common"
)

func init()  {
	etc.InitConfig("../etc/config.json")
}
func TestGetControllerInstance(t *testing.T)  {
	if err, _ := GetControllerInstance(); err != nil {
		t.Error(etc.TEST_FAIL)
	}else {
		t.Log(etc.TEST_OK)
	}
}
func TestGetControllerInstanceAndAuth(t *testing.T)  {
	address := common.HexToAddress("0x646829bFA80580b07767b796B4055DB9BDf148b5")
	password := "123456"
	if err, _, _ := GetControllerInstanceAndAuth(address, password); err != nil {
		t.Error(etc.TEST_FAIL)
	}else {
		t.Log(etc.TEST_OK)
	}
	address = common.HexToAddress("0x646829bFA80580b07767b796B4055DB9BDf148b5")
	password = "12345678"
	if err, _, _ := GetControllerInstanceAndAuth(address, password); err != nil {
		t.Log(etc.TEST_OK)
	}else {
		t.Error(etc.TEST_FAIL)
	}
}