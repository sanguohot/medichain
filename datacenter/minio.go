package datacenter

import (
	"bytes"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/minio/minio-go"
	"github.com/sanguohot/medichain/etc"
	"github.com/sanguohot/medichain/util"
	"github.com/sanguohot/medichain/zap"
	"net/http"
)

const (
	DEFAULT_BUCKET = "medichain"
	DEFAULT_LOCATION = "us-east-1"
)

func SaveToMinio(key string, value []byte) error {
	if !etc.GetMinioEnable() {
		zap.Sugar.Infof("minio is not enable")
		return nil
	}
	// Initialize minio client object.
	minioClient, err := minio.New(fmt.Sprintf("%s:%d", etc.GetMinioAddress(), etc.GetMinioPort()), etc.GetMinioAccessKey(), etc.GetMinioSecretKey(), etc.GetMinioSecure())
	if err != nil {
		return err
	}
	// Make a new bucket called mymusic.
	bucketName := DEFAULT_BUCKET
	location := DEFAULT_LOCATION
	exists, err := minioClient.BucketExists(bucketName)
	if err != nil {
		return err
	}
	if !exists {
		err = minioClient.MakeBucket(bucketName, location)
		if err != nil {
			return err
		}
		zap.Sugar.Infof("Successfully created %s", bucketName)
	}
	// Upload the zip file with FPutObject
	n, err := minioClient.PutObject(bucketName, key, bytes.NewReader(value), int64(len(value)), minio.PutObjectOptions{ContentType: http.DetectContentType(value)})
	if err != nil {
		return err
	}

	zap.Sugar.Infof("Successfully uploaded %s of size %d", key, n)
	return nil
}

func SaveKeystoreToMinio(account common.Address, password string) {
	go func() {
		value, err := util.GetKeyJsonFromStore(account, password)
		if err != nil {
			zap.Logger.Error(err.Error())
			return
		}
		if err := SaveToMinio(account.Hex(), value); err != nil {
			zap.Logger.Error(err.Error())
			return
		}
	}()

}