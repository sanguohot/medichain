package main

import (
	"medichain/file"
	"fmt"
	"medichain/util"
	"io/ioutil"
	"bytes"
)

func main() {
	var err error
	bigDataItemList, err := file.GetFilesInFolder(file.FolderCode)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(len(bigDataItemList), bigDataItemList[9].FileCode)
	_, err = file.CreateFolderInBigDataCenter(file.FolderName)
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
	err = file.UploadToBigDataCenter(fileBytesUp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fileBytesDown, err := file.DownloadFromBigDataCenter(util.Md5(fileBytesUp))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("\ndownload file length %d", len(fileBytesDown))
	fmt.Println("\nbytes.Equal(fileBytesUp, fileBytesDown) ===>", bytes.Equal(fileBytesUp, fileBytesDown))
	fmt.Println(string(fileBytesDown))
}