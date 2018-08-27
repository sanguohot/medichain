package server

import (
	"github.com/gin-gonic/gin"
	"medichain/server/middle"
	"time"
	"net/http"
	"medichain/util"
	"medichain/service"
	"fmt"
	"github.com/google/uuid"
)

func InitServer() {
	// 默认中间件包含logger和recover
	r := gin.Default()
	r.Use(middle.UserAuthHandler)
	r.Use(middle.FileAuthHandler)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	v1 := r.Group("/api/v1")
	{
		v1.GET("/user/:name", func(c *gin.Context) {
			name := c.Param("name")
			c.String(http.StatusOK, "Hello %s", name)
		})
		v1.POST("/user", func(c *gin.Context) {
			idCartNo := c.PostForm("idCartNo")
			err, user, _ := service.AddUser(uuid.UUID{}, idCartNo)
			if err != nil {
				c.String(http.StatusOK, err.Error())
			}
			c.JSON(http.StatusOK, user)
		})
	}
	s := &http.Server{
		Addr:           fmt.Sprintf("%s:%s", util.Config.GetString("server.host.address"), util.Config.GetInt("server.host.port")),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
