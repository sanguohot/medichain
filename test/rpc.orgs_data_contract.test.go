package main

import (
	"fmt"
	"github.com/sanguohot/medichain/chain"
	"github.com/sanguohot/medichain/etc"
	"log"
	//"github.com/google/uuid"
	//"github.com/sanguohot/medichain/util"
	"time"
	"math/big"
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
	time.Sleep(time.Second * 1)
	size, err := chain.OrgsDataGetSuperSize()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("super size ===>", size)
	if size.Uint64() > 0 {
		super, err := chain.OrgsDataGetSuperByIndex(big.NewInt(0))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("super 0 ===>",super.Hex())
	}
	//password := "123456"
	//orgUuid := uuid.New()
	//_, _, address, err = util.NewWallet(password)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//publicKey, err := util.GetPublicKeyBytes32_2FromStore(*address, password)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//name = "清华大学"
	//nameBytes32_4, err := util.StringToBytes32_4(name)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//if err, hash := chain.OrgsDataAddOrg(orgUuid, *address, *publicKey, *nameBytes32_4); err != nil {
	//	log.Fatal(err)
	//}else {
	//	fmt.Printf("tx sent: %s\n", hash.Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870
	//}
}