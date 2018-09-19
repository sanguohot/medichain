package main

import (
	"bytes"
	"github.com/sanguohot/medichain/datacenter"
	"github.com/sanguohot/medichain/util"
	"github.com/sanguohot/medichain/zap"
	"io/ioutil"
)

func main() {
	var err error
	bigDataItemList, err := datacenter.GetFilesInFolder(datacenter.FolderCode)
	if err != nil {
		_, err2 := datacenter.CreateFolderInBigDataCenter(datacenter.FolderName)
		if err2 != nil {
			zap.Logger.Fatal(err2.Error())
		}
		zap.Logger.Fatal(err.Error())
	}
	zap.Sugar.Infof("bigDataItemList size %d", len(bigDataItemList))
	fileName := "e:/evan/test.txt"
	fileBytesUp, err := ioutil.ReadFile(fileName)
	if err != nil {
		zap.Logger.Fatal(err.Error())
	}
	sha256Hash := util.Sha256Hash(fileBytesUp)
	zap.Sugar.Infof("bigDataItemList size %d", len(bigDataItemList))
	err = datacenter.UploadToBigDataCenter(fileBytesUp)
	if err != nil {
		zap.Logger.Fatal(err.Error())
	}
	zap.Sugar.Infof("upload success %s %s", fileName, sha256Hash.Hex())
	fileBytesDown, err := datacenter.DownloadFromBigDataCenter(sha256Hash)
	if err != nil {
		zap.Logger.Fatal(err.Error())
	}
	zap.Sugar.Infof("\ndownload file length %d", len(fileBytesDown))
	zap.Sugar.Infof("\nbytes.Equal(fileBytesUp, fileBytesDown)=%v", bytes.Equal(fileBytesUp, fileBytesDown))
	zap.Sugar.Info(string(fileBytesDown))
}