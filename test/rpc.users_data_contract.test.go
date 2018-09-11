package main

import (
	"fmt"
	"github.com/sanguohot/medichain/chain"
	"github.com/sanguohot/medichain/etc"
	"log"
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
	if err, hash := chain.UsersDataAddSuper(*address); err != nil {
		log.Fatal(err)
	}else {
		fmt.Printf("tx sent: %s\n", hash.Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870
	}
	time.Sleep(time.Second * 1)
	size, err := chain.UsersDataGetSuperSize()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("super size ===>", size)
	if size.Uint64() > 0 {
		super, err := chain.UsersDataGetSuperByIndex(big.NewInt(0))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("super 0 ===>",super.Hex())
	}
	//password := "123456"
	//userUuid := uuid.New()
	//_, _, address, err = util.NewWallet(password)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//publicKey, err := util.GetPublicKeyBytes32_2FromStore(*address, password)
	//if err != nil {
	//	log.Fatal(err)
	//}

}