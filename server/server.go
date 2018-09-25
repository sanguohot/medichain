package server

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sanguohot/medichain/etc"
	_ "github.com/sanguohot/medichain/timer"
	"github.com/sanguohot/medichain/zap"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func init()  {
	InitServer()
}

func InitServer() {
	//r := gin.Default()
	if etc.ServerTypeIsProdOrPre() {
		gin.SetMode(gin.ReleaseMode)
	}else {
		gin.SetMode(gin.DebugMode)
	}
	r := gin.New()
	r.Use(gin.Recovery())
	// 默认设置logger，但启用logger会导致吞吐量大幅度降低，具体查看benmark/wrk/output.txt
	if os.Getenv("GIN_LOG") != "off" {
		r.Use(gin.Logger())
	}
	r.MaxMultipartMemory = 100 << 20 // 100 MB
	r.GET("/ping", PongHandler)
	r.GET("/ping_sleep_one_sec", PingSleepOneSecHandler)
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
		v1.GET("/transation_result/:hash", CheckTransactionResultHandler)
	}
	s := &http.Server{
		Addr:           fmt.Sprintf("%s:%d", etc.GetServerHostAddress(), etc.GetServerHostPort()),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1 MB
	}
	zap.Sugar.Infof("server listening on %s", s.Addr)
	if etc.GetServerTlsEnable() {
		pool := x509.NewCertPool()
		caCertPath := etc.GetServerPkiCa()
		caCrt, err := ioutil.ReadFile(caCertPath)
		if err != nil {
			zap.Logger.Fatal(err.Error())
		}
		pool.AppendCertsFromPEM(caCrt)
		tlsConfig := &tls.Config{
			ClientCAs:  pool,
		}
		if etc.GetServerTlsVerifyPeer() {
			tlsConfig.ClientAuth = tls.RequireAndVerifyClientCert
		}else {
			tlsConfig.ClientAuth = tls.NoClientCert
		}
		if err := s.ListenAndServeTLS(etc.GetServerPkiCert(), etc.GetServerPkiKey()); err != nil {
			zap.Logger.Fatal(err.Error())
		}
	}else {
		if err := s.ListenAndServe(); err != nil {
			zap.Logger.Fatal(err.Error())
		}
	}
}
