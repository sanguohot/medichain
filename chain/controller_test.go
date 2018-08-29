package chain

import (
	"testing"
	"medichain/etc"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"medichain/util"
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
func TestControllerAddOrg(t *testing.T) {
	orgUuid := uuid.New()
	publicKey := [2][32]byte{}
	publicKey[0] = common.HexToHash("0x06ed2e37f17a997fd3256e0d1f41107af2f9b2ed4a96f6e65b357b1dbb7fbb01")
	publicKey[1] = common.HexToHash("0x6872be104937d773d928df94a0966e8e7e88aca75dbd768592dc9c94d5afdda9")
	address := common.HexToAddress("0x646829bFA80580b07767b796B4055DB9BDf148b5")
	password := "123456"
	name := "广西中医一附院"
	nameBytes32_4, err := util.StringToBytes32_4(name)
	if err != nil {
		t.Error(etc.TEST_FAIL)
	}
	if err, _ := ControllerAddOrg(orgUuid, *nameBytes32_4, address, password); err != nil {
		t.Error(etc.TEST_FAIL)
	}else {
		t.Log(etc.TEST_OK)
	}
	password = ""
	if err, _ := ControllerAddOrg(orgUuid, *nameBytes32_4, address, password); err != nil {
		t.Log(etc.TEST_OK)
	}else {
		t.Error(etc.TEST_FAIL)
	}
}