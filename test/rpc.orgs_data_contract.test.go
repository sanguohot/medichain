package main

import (
	"fmt"
	"medichain/chain"
	"medichain/etc"
	"log"
	"github.com/google/uuid"
	"github.com/ethereum/go-ethereum/common"
	"medichain/util"
)

func main() {
	name := etc.ContractController
	err, address := chain.GetAddressFromCns(name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ContractController address ===>", address.Hex())
	if err, hash := chain.OrgsDataAddSuper(*address); err != nil {
		log.Fatal(err)
	}else {
		fmt.Printf("tx sent: %s\n", hash.Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870
	}
	orgUuid := uuid.New()
	publicKey := [2][32]byte{}
	publicKey[0] = common.HexToHash("8ec34f461212f6bbfd759f400b5a80679ecd56f496148ad2c4669d9e1127965b")
	publicKey[1] = common.HexToHash("e94cbd7230c091e5bfa310596f354d1339a122e7335bdf799f504b3c1fcb07a5")
	orgAddress := common.HexToAddress("0x1d306aDe962d9f20590be45e68DB6246da73811D")
	name = "华中科技大学同济医学院附属妇女儿童医疗保健中心"
	nameBytes32_4, err := util.StringToBytes32_4(name)
	if err != nil {
		log.Fatal(err)
	}
	if err, hash := chain.OrgsDataAddOrg(orgUuid, orgAddress, publicKey, *nameBytes32_4); err != nil {
		log.Fatal(err)
	}else {
		fmt.Printf("tx sent: %s\n", hash.Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870
	}
}