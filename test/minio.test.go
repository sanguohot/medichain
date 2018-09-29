package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/sanguohot/medichain/datacenter"
	"time"
)

func main() {
	datacenter.SaveKeystoreToMinio(common.HexToAddress("1d306ade962d9f20590be45e68db6246da73811d"), "123456")
	time.Sleep(time.Second)
}