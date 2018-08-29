package server

import (
	"github.com/gin-gonic/gin"
	"medichain/server/middle"
	"time"
	"net/http"
	"fmt"
	"medichain/etc"
)

func InitServer() {
	// 默认中间件包含logger和recover
	r := gin.Default()
	r.MaxMultipartMemory = 100 << 20 // 100 MiBr
	r.Use(middle.UserAuthHandler)
	r.Use(middle.FileAuthHandler)
	r.GET("/ping", PongHandler)
	v1 := r.Group("/api/v1")
	{
		v1.POST("/user", AddUserHandler)
		v1.POST("/org", AddOrgHandler)
		v1.POST("/file", AddFileHandler)
		v1.GET("/file/:fileUuid", GetFileHandler)
	}
	s := &http.Server{
		Addr:           fmt.Sprintf("%s:%s", etc.GetServerHostAddress(), etc.GetServerHostPort()),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
