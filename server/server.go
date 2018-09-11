package server

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sanguohot/medichain/etc"
	"github.com/sanguohot/medichain/server/middle"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func init()  {
	InitServer()
}

func InitServer() {
	r := gin.Default()
	r.MaxMultipartMemory = 100 << 20 // 100 MB
	r.Use(middle.UserAuthHandler)
	r.Use(middle.FileAuthHandler)
	r.GET("/ping", PongHandler)
	v1 := r.Group("/api/v1")
	{
		v1.POST("/user", AddUserHandler)
		v1.POST("/org", AddOrgHandler)
		v1.POST("/file", AddFileHandler)
		v1.POST("/file/:fileUuid/signature", AddFileSignHandler)
		v1.GET("/file/:fileUuid/signature", GetFileSignerAndDataListHandler)
		v1.GET("/fileAddLog", GetFileAddLogListHandler)
		v1.GET("/file/:fileUuid", GetFileHandler)
		v1.GET("/file_log", GetFileAddLogListHandler)
		v1.GET("/file_log/:fileUuid", GetFileAddLogDetailHandler)
	}
	s := &http.Server{
		Addr:           fmt.Sprintf("%s:%d", etc.GetServerHostAddress(), etc.GetServerHostPort()),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1 MB
	}
	if etc.GetServerTlsEnable() {
		pool := x509.NewCertPool()
		caCertPath := etc.GetServerPkiCa()
		caCrt, err := ioutil.ReadFile(caCertPath)
		if err != nil {
			log.Fatal(err)
		}
		pool.AppendCertsFromPEM(caCrt)
		tlsConfig := &tls.Config{
			ClientCAs:  pool,
		}
		if etc.GetServerTlsVerifyPeer() {
			tlsConfig.ClientAuth = tls.RequireAndVerifyClientCert
		}
		if err := s.ListenAndServeTLS(etc.GetServerPkiCert(), etc.GetServerPkiKey()); err != nil {
			log.Fatal(err)
		}
	}else {
		if err := s.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}
}
