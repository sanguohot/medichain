package main

import (
	"github.com/sanguohot/medichain/datacenter"
	"fmt"
	"github.com/sanguohot/medichain/util"
	"io/ioutil"
	"bytes"
	"log"
)

func main() {
	var err error
	bigDataItemList, err := datacenter.GetFilesInFolder(datacenter.FolderCode)
	if err != nil {
		_, err2 := datacenter.CreateFolderInBigDataCenter(datacenter.FolderName)
		if err2 != nil {
			log.Fatal(err2)
		}
		log.Fatal(err)
	}
	fmt.Println("bigDataItemList size ===>", len(bigDataItemList))

	fileName := "e:/evan/test.txt"
	fileBytesUp, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
		return
	}
	sha256Hash := util.Sha256Hash(fileBytesUp)
	fmt.Println("begin upload ===>", fileName, sha256Hash.Hex())
	err = datacenter.UploadToBigDataCenter(fileBytesUp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("upload success ===>", fileName, sha256Hash.Hex())
	fileBytesDown, err := datacenter.DownloadFromBigDataCenter(sha256Hash)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("\ndownload file length %d", len(fileBytesDown))
	fmt.Println("\nbytes.Equal(fileBytesUp, fileBytesDown) ===>", bytes.Equal(fileBytesUp, fileBytesDown))
	fmt.Println(string(fileBytesDown))
}