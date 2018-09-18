package chain

import (
	"fmt"
	"github.com/sanguohot/medichain/etc"
	"testing"
	"time"
)

func TestGetCnsInstance(t *testing.T)  {
	if err, _ := GetCnsInstance(); err != nil {
		t.Error(etc.TEST_FAIL)
	}else {
		t.Log(etc.TEST_OK)
	}
}
func TestGetAddressFromCns(t *testing.T)  {
	if err, _ := GetAddressFromCns(etc.ContractController); err != nil {
		t.Error(etc.TEST_FAIL)
	}else {
		t.Log(etc.TEST_OK)
	}
	if err, _ := GetAddressFromCns(fmt.Sprintf("%d", time.Now().Unix())); err != nil {
		t.Log(etc.TEST_OK)
	}else {
		t.Error(etc.TEST_FAIL)
	}
}