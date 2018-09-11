package main

import (
	"fmt"
	"github.com/sanguohot/medichain/chain"
	//"github.com/sanguohot/medichain/etc"
	"log"
	//"time"
	"math/big"
	"github.com/google/uuid"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sanguohot/medichain/util"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/common/hexutil"
	//"github.com/sanguohot/medichain/etc"
	"github.com/sanguohot/medichain/etc"
)

func main() {
	name := etc.ContractController
	err, address := chain.GetAddressFromCns(name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ContractController address ===>", address.Hex())
	if err, hash := chain.FilesDataAddSuper(*address); err != nil {
		log.Fatal(err)
	}else {
		fmt.Printf("tx sent: %s\n", hash.Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870
	}
	return
	//time.Sleep(time.Second * 1)
	owner, err := chain.FilesDataGetOwner()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("files data owner ===>", owner.Hex())
	size, err := chain.FilesDataGetSuperSize()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("super size ===>", size)
	if size.Uint64() > 0 {
		super, err := chain.FilesDataGetSuperByIndex(big.NewInt(0))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("super 0 ===>",super.Hex())
	}
	password := "123456"
	userAddress := common.HexToAddress("0x996ff28238f9083896b075ed78b70b6b4687b42a")
	fileUuid := uuid.New()
	ownerUuid, err := uuid.Parse("6933cd1d-f7cb-4201-b92b-77f58a6d590a")
	if err != nil {
		log.Fatal(err)
	}
	uploaderUuid, err := uuid.Parse("eddd0c10-bf7e-4bb0-b33f-6ad0150b94c8")
	if err != nil {
		log.Fatal(err)
	}
	fileType := common.BytesToHash([]byte("电子病历"))
	fileDesc, err := util.StringToBytes32_4("电子病历/住院病历")
	if err != nil {
		log.Fatal(err)
	}

	keccak256Hash := crypto.Keccak256Hash([]byte("hello world"))
	signBytes, err := util.SignWithHash(keccak256Hash, userAddress, password)
	if err != nil {
		log.Fatal(err)
	}
	r, s, v := util.SigRSV(signBytes)
	sha256Hash := common.HexToHash("b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9")
	publicKeyBytes32_2, err := util.GetPublicKeyBytes32_2FromStore(userAddress, password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("公钥[\"%s\",\"%s\"]\n", hexutil.Encode(publicKeyBytes32_2[0][:]), hexutil.Encode(publicKeyBytes32_2[1][:]))
	fmt.Println("keccak256Hash ===>", keccak256Hash.Hex())
	fmt.Println("sha256Hash ===>", sha256Hash.Hex())
	fmt.Println("r =>", hexutil.Encode(r[:])[:])
	fmt.Println("s =>",hexutil.Encode(s[:])[:])
	fmt.Println("v =>", v)
	fmt.Println("userAddress ===>", userAddress.Hex())
	fmt.Println("uploaderUuid ===>", uploaderUuid)
	err, hash := chain.FilesDataAddFile(fileUuid, ownerUuid, uploaderUuid, fileType, *fileDesc, keccak256Hash, sha256Hash, r, s, v)
	fmt.Println(hash.Hex())
	if err != nil {
		log.Fatal(err)
	}else {
		fmt.Printf("tx sent: %s\n", hash.Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870
	}
}