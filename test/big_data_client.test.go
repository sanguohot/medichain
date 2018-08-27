package main

import (
	"medichain/data"
	"fmt"
	"medichain/util"
	"io/ioutil"
	"bytes"
)

func main() {
	var err error
	bigDataItemList, err := data.GetFilesInFolder(data.FolderCode)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("bigDataItemList size ===>", len(bigDataItemList))
	_, err = data.CreateFolderInBigDataCenter(data.FolderName)
	if err != nil {
		fmt.Println(err)
		//return
	}
	fileName := "e:/evan/test.txt"
	fileBytesUp, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = data.UploadToBigDataCenter(fileBytesUp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fileBytesDown, err := data.DownloadFromBigDataCenter(util.Md5(fileBytesUp))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("\ndownload file length %d", len(fileBytesDown))
	fmt.Println("\nbytes.Equal(fileBytesUp, fileBytesDown) ===>", bytes.Equal(fileBytesUp, fileBytesDown))
	fmt.Println(string(fileBytesDown))
}