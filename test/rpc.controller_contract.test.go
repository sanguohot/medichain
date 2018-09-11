package main

import (
	"fmt"
	"log"
	"github.com/sanguohot/medichain/chain"
	//"github.com/google/uuid"
	"github.com/sanguohot/medichain/util"
	//"hash"
	"time"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"github.com/sanguohot/medichain/service"
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
	fmt.Println("org 天津中医药大学第一附属医院 ===>", id, addressNew.Hex(), common.Hash(b).Hex(), hexutil.Encode(util.Bytes32_2ToBytes(a))[2:], util.Bytes32_4ToString(c), d)

	idCartNo := "828007199605192670"
	id, err = chain.UsersDataGetUuidByIdCartNoHash(crypto.Keccak256Hash([]byte(idCartNo)))
	if err != nil {
		log.Fatal(err)
	}
	err, addressNew, ua, ub, uc, ud := chain.ControllerGetUserByUuid(*id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("user 828007199605192670 ====>", id, addressNew.Hex(), ua, hexutil.Encode(util.Bytes32_2ToBytes(ub))[2:], common.Hash(uc).Hex(), ud)

	size, err := chain.FilesDataGetUuidListSize()
	fmt.Println("file size ===>", size.Uint64())
	fileUuid, err := uuid.Parse("fa84d023-7e54-427d-9d64-c91e2a040dd0")
	if err != nil {
		log.Fatal(err)
	}
	err, fa, fb, fc, fd, fe, ff, fg :=  chain.ControllerGetFileByUuid(fileUuid)
	fmt.Println("fa84d023-7e54-427d-9d64-c91e2a040dd0 ====>", uuid.UUID(fa), uuid.UUID(fb), common.Hash(fc).Hex(), util.Bytes32_4ToString(fd), common.Hash(fe).Hex(), common.Hash(ff).Hex(), fg)
	err, action := service.GetFileSignerAndDataList("fa84d023-7e54-427d-9d64-c91e2a040dd0", "0", "10")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(action)
}