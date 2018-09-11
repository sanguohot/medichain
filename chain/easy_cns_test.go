package chain

import (
	"testing"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sanguohot/medichain/etc"
)

func init()  {
	etc.InitConfig("../etc/config.json")
}
func TestGetCnsInstance(t *testing.T)  {
	if err, _ := GetCnsInstance(); err != nil {
		t.Error(etc.TEST_FAIL)
	}else {
		t.Log(etc.TEST_OK)
	}
}
func TestSetAddressToCns(t *testing.T)  {
	if err, _ := SetAddressToCns("test", common.HexToAddress("0x98fB1108eBC0FA3A65d3C22eCB34e1F075f9ab5D")); err != nil {
		t.Error(etc.TEST_FAIL)
	}else {
		t.Log(etc.TEST_OK)
	}
}
func TestGetAddressFromCns(t *testing.T)  {
	if err, _ := GetAddressFromCns("test"); err != nil {
		t.Error(etc.TEST_FAIL)
	}else {
		t.Log(etc.TEST_OK)
	}
}