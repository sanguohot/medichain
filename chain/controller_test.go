package chain

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/sanguohot/medichain/etc"
	"github.com/sanguohot/medichain/util"
	"testing"
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
	password := "12345678"
	if err, _, _ := GetControllerInstanceAndAuth(address, password); err != nil {
		t.Log(etc.TEST_OK)
	}else {
		t.Error(etc.TEST_FAIL)
	}
}

func TestControllerGetOrgByUuid(t *testing.T)  {
	orgUuid := uuid.New()
	if err, address, _, _, _, _ := ControllerGetOrgByUuid(orgUuid); err != nil {
		t.Error(etc.TEST_FAIL)
	}else if util.IsValidAndNotZeroAddress(address) {
		t.Error(etc.TEST_FAIL)
	}else {
		t.Log(etc.TEST_OK)
	}
}

func TestControllerGetUserByUuid(t *testing.T)  {
	userUuid := uuid.New()
	if err, address, _, _, _, _ := ControllerGetUserByUuid(userUuid); err != nil {
		t.Error(etc.TEST_FAIL)
	}else if util.IsValidAndNotZeroAddress(address) {
		t.Error(etc.TEST_FAIL)
	}else {
		t.Log(etc.TEST_OK)
	}
}

func TestControllerGetFileByUuid(t *testing.T)  {
	fileUuid := uuid.New()
	emptyHash := common.Hash{}
	if err, _, _, _, _, _, keccak256Hash, _ := ControllerGetFileByUuid(fileUuid); err != nil {
		t.Error(etc.TEST_FAIL)
	}else if common.Hash(keccak256Hash) == emptyHash {
		t.Log(etc.TEST_OK)
	}else {
		t.Error(etc.TEST_FAIL)
	}
}