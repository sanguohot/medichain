package util

import (
	"errors"
	"fmt"
	"github.com/sanguohot/medichain/zap"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func NewTimeoutClient() *http.Client {
	var transport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		TLSHandshakeTimeout:   5 * time.Second,
		ResponseHeaderTimeout: 10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	var client = &http.Client{
		Timeout:   time.Second * 30,
		Transport: transport,
	}
	return client
}

func DownloadFromFileUrl(fileUrl string) ([]byte, error) {
	if fileUrl == "" {
		return nil, ErrParamShouldNotNil
	}
	zap.Sugar.Infof("DownloadFromFileUrl %s", fileUrl)
	req, err := http.NewRequest("GET", fileUrl, nil)
	res, err := NewTimeoutClient().Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, errors.New("resp status:" + fmt.Sprint(res.StatusCode))
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}