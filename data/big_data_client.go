package data

import (
	"io/ioutil"
	"fmt"
	"net/http"
	"errors"
	"bytes"
	"mime/multipart"
	"medichain/util"
	"io"
	"encoding/json"
)

var (
	// 10.0.253.8:9998
	// 172.16.194.2:9998
	serverAddr string = fmt.Sprintf("%s:%s", util.Config.GetString("big_data.host.address"), util.Config.GetInt("big_data.host.port"))
	uploadPath string = "/ylyx/upload_single"
	downloadPath string = "/ylyx/downloadByMd5"
	createFolderPath string = "/ylyx/createFolder"
	getFilesPath string = "/ylyx/getFolder"
	codeSucc uint8 = 1
	codeFail uint8 = 0
	FolderName string = "medichain"
	FolderCode string = "77bb6e0977bb6e09dir0000000000000"
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
	FileMd5  []byte  `json:"fileMd5,omitempty"`
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
	buf := new(bytes.Buffer)
	w := multipart.NewWriter(buf)
	md5Hex := util.Md5(fileBytes)
	fw, err := w.CreateFormFile("file", md5Hex)
	if err != nil {
		return err
	}
	_, err = fw.Write(fileBytes)
	if err != nil {
		return err
	}
	if err := w.WriteField("md5", md5Hex); err != nil {
		return err
	}
	w.Close()
	url := fmt.Sprintf("http://%s%s", serverAddr, uploadPath)
	fmt.Println("UploadToBigDataCenter ===>", url)
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

func DownloadFromBigDataCenter(md5Code string) ([]byte, error) {
	if md5Code == "" {
		return nil, errors.New("md5Code不能为空")
	}
	url := fmt.Sprintf("http://%s%s?md5=%s", serverAddr, downloadPath, md5Code)
	fmt.Println("DownloadFromBigDataCenter ===>", url)
	req, err := getReqWrapper("GET", url, nil)
	var body []byte
	if _, body, err = doReq(req, false); err != nil {
		return nil, err
	}
	return body, nil
}

func CreateFolderInBigDataCenter(folderName string) (string, error) {
	if folderName == "" {
		return "", errors.New("folderName不能为空")
	}
	url := fmt.Sprintf("http://%s%s?name=%s", serverAddr, createFolderPath, folderName)
	fmt.Println("DownloadFromBigDataCenter ===>", url)
	req, err := getReqWrapper("POST", url, nil)
	var bigDataResponse *BigDataResponse
	if bigDataResponse, _, err = doReq(req, true); err != nil {
		return "", err
	}
	fmt.Println("CreateFolderInBigDataCenter ===>", bigDataResponse)
	return bigDataResponse.InfoMap.FolderCode, err
}

func GetFilesInFolder(folderCode string) ([]BigDataItem, error) {
	url := fmt.Sprintf("http://%s%s", serverAddr, getFilesPath)
	if folderCode != "" {
		url = fmt.Sprintf("%s?folderCode=%s", url, folderCode)
	}
	fmt.Println("getFilesInFolder ===>", url)
	req, err := getReqWrapper("GET", url, nil)
	var bigDataResponse *BigDataResponse
	if bigDataResponse, _, err = doReq(req, true); err != nil {
		return nil, err
	}
	return bigDataResponse.DataList, nil
}