package datacenter

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sanguohot/medichain/etc"
	"github.com/sanguohot/medichain/util"
	"github.com/sanguohot/medichain/zap"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

var (
	// 10.0.253.8:9998
	// 172.16.194.2:9998
	serverAddr string = fmt.Sprintf("%s:%d", etc.GetBigDataHostAddress(), etc.GetBigDataHostPort())
	uploadPath string = "/ylyx/uploadSingle"
	downloadPath string = "/ylyx/downloadBySha"
	createFolderPath string = "/ylyx/createFolder"
	getFilesPath string = "/ylyx/getFolder"
	codeSucc uint8 = 1
	codeFail uint8 = 0
	FolderName string = etc.GetBigDataRootFolderName()
	FolderCode string = etc.GetBigDataRootFolderCode()
)

type BigDataInfoMap struct {
	FolderCode  string    `json:"folderCode,omitempty"`
	FolderName  string    `json:"folderName,omitempty"`
}
type BigDataItem struct {
	FileName  string    `json:"fileName,omitempty"`
	ParentFolder  string    `json:"parentFolder,omitempty"`
	Size  uint64    `json:"size,omitempty"`
	CreateTime  util.JsonTime   `json:"createTime,omitempty"`
	FileCode  string    `json:"fileCode,omitempty"`
	FileSha  []byte  `json:"fileSha,omitempty"`
}

type BigDataResponse struct {
	Msg  string    `json:"msg,omitempty"`
	DataList  []BigDataItem    `json:"dataList,omitempty"`
	Html  string    `json:"html,omitempty"`
	Code  uint8    `json:"returnCode,omitempty"`
	InfoMap  BigDataInfoMap    `json:"infoMap,omitempty"`
}

func getReqWrapper(method, url string, body io.Reader) (*http.Request, error) {
	return http.NewRequest(method, url, body)
}
func doReq(req *http.Request, isJsonRes bool) (*BigDataResponse, []byte, error) {
	res, err := util.NewTimeoutClient().Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, nil, errors.New("resp status:" + fmt.Sprint(res.StatusCode))
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}
	if !isJsonRes {
		return nil, body, nil
	}
	var bigDataResponse *BigDataResponse
	err = json.Unmarshal(body, &bigDataResponse)
	if err != nil {
		return nil, nil, err
	}
	if bigDataResponse.Code == codeFail {
		return nil, nil, errors.New(bigDataResponse.Msg)
	}else if bigDataResponse.Code == codeSucc {
		return bigDataResponse, nil, nil
	}else {
		return nil, nil, errors.New(fmt.Sprintf("unknown result code %d", bigDataResponse.Code))
	}
}
func UploadToBigDataCenter(fileBytes []byte) error {
	if fileBytes == nil || len(fileBytes) == 0 {
		return fmt.Errorf("%s:UploadToBigDataCenter(fileBytes []byte)", util.ErrParamShouldNotNil.Error())
	}
	buf := new(bytes.Buffer)
	w := multipart.NewWriter(buf)
	// do not set 0x
	sha256Hex := util.Sha256Hash(fileBytes).Hex()[2:]
	fw, err := w.CreateFormFile("file", sha256Hex)
	if err != nil {
		return err
	}
	_, err = fw.Write(fileBytes)
	if err != nil {
		return err
	}
	if err := w.WriteField("sha", sha256Hex); err != nil {
		return err
	}
	if err := w.WriteField("folderCode", FolderCode); err != nil {
		return err
	}
	w.Close()
	url := fmt.Sprintf("http://%s%s", serverAddr, uploadPath)
	zap.Sugar.Infof("UploadToBigDataCenter %s", url)
	req, err := getReqWrapper("POST", url, buf)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", w.FormDataContentType())
	if _, _, err = doReq(req, true); err != nil {
		return err
	}
	return nil
}

func DownloadFromBigDataCenter(sha256Hash common.Hash) ([]byte, error) {
	emptyHash := common.Hash{}
	if sha256Hash == emptyHash {
		return nil, fmt.Errorf("%s:DownloadFromBigDataCenter(sha256Hash common.Hash)", util.ErrParamShouldNotNil.Error())
	}
	url := fmt.Sprintf("http://%s%s?sha=%s", serverAddr, downloadPath, sha256Hash.Hex()[2:])
	zap.Sugar.Infof("DownloadFromBigDataCenter %s", url)
	req, err := getReqWrapper("GET", url, nil)
	var body []byte
	if _, body, err = doReq(req, false); err != nil {
		return nil, err
	}
	return body, nil
}

func CreateFolderInBigDataCenter(folderName string) (string, error) {
	if folderName == "" {
		return "", fmt.Errorf("%s:CreateFolderInBigDataCenter(folderName string)", util.ErrParamShouldNotNil.Error())
	}
	url := fmt.Sprintf("http://%s%s?name=%s", serverAddr, createFolderPath, folderName)
	zap.Sugar.Infof("DownloadFromBigDataCenter %s", url)
	req, err := getReqWrapper("POST", url, nil)
	var bigDataResponse *BigDataResponse
	if bigDataResponse, _, err = doReq(req, true); err != nil {
		return "", err
	}
	zap.Sugar.Infof("CreateFolderInBigDataCenter %s", bigDataResponse)
	return bigDataResponse.InfoMap.FolderCode, err
}

func GetFilesInFolder(folderCode string) ([]BigDataItem, error) {
	url := fmt.Sprintf("http://%s%s", serverAddr, getFilesPath)
	if folderCode != "" {
		url = fmt.Sprintf("%s?folderCode=%s", url, folderCode)
	}
	zap.Sugar.Infof("GetFilesInFolder %s", url)
	req, err := getReqWrapper("GET", url, nil)
	var bigDataResponse *BigDataResponse
	if bigDataResponse, _, err = doReq(req, true); err != nil {
		return nil, err
	}
	return bigDataResponse.DataList, nil
}