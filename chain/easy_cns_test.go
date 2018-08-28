package chain

import (
	"testing"
	"github.com/ethereum/go-ethereum/common"
	"fmt"
	"medichain/etc"
)

func init()  {
	etc.InitConfig("../etc/config.json")
}
func TestSetAddressToCns(t *testing.T)  {
	//if err, _ := SetAddressToCns("a", common.HexToAddress("0x0")); err != nil {
	//	t.Error("faild")
	//}else {
	//	t.Log("pass")
	//}
	if err, _ := SetAddressToCns("test", common.HexToAddress("0x98fB1108eBC0FA3A65d3C22eCB34e1F075f9ab5D")); err != nil {
		t.Error("faild")
		fmt.Println(err)
	}else {
		t.Log("pass")
	}
}