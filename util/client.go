package util

import (
	"net/http"
	"time"
	"net"
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