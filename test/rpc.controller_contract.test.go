package main

import (
	"fmt"
	"log"
	"medichain/chain"
	//"github.com/google/uuid"
	"medichain/util"
	//"hash"
	"time"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func main() {
	name := "天津中医药大学第一附属医院"
	nameBytes32_4, err := util.StringToBytes32_4(name)
	if err != nil {
		log.Fatal(err)
	}
	//orgUuid := uuid.New()
	//password := "123456"
	//_, _, address, err := util.NewWallet(password)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//err, hash := chain.ControllerAddOrg(orgUuid, *nameBytes32_4, *address, password)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("交易成功 ===>", hash.Hex())
	time.Sleep(time.Second * 1)
	id, err := chain.OrgsDataGetUuidByNameHash(util.Bytes32_4Hash(*nameBytes32_4))
	if err != nil {
		log.Fatal(err)
	}
	err, addressNew, a, b, c, d := chain.ControllerGetOrgByUuid(*id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id, addressNew.Hex(), common.Hash(b).Hex(), hexutil.Encode(util.Bytes32_2ToBytes(a))[2:], util.Bytes32_4ToString(c), d)
}