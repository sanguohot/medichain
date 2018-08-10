package file

import (
	"os"
	"io/ioutil"
	"fmt"
	"net/http"
	"errors"
)

var (
	serverAddr string = "10.0.253.8:9998"
	uploadPath string = ""
	downloadPath string = "/ylyx/downloadByMd5"
)
func UploadToBigDataCenter() {
	file, err := os.Open("./filename")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	url := fmt.Sprintf("http://%s%s", serverAddr, uploadPath)
	res, err := http.Post(url, "binary/octet-stream", file)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	message, _ := ioutil.ReadAll(res.Body)
	fmt.Printf(string(message))
}

func DownloadFromBigDataCenter(md5Code string) ([]byte, error) {
	if md5Code == "" {
		return nil, errors.New("md5Code不能为空")
	}
	url := fmt.Sprintf("http://%s%s?md5=%s", serverAddr, downloadPath, md5Code)
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	//fmt.Println(string(body))
	return body, nil
}