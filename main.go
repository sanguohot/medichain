package main

import (
	"github.com/gin-gonic/gin"
	"medichain/middle"
	"net/http"
	"time"
)

func main() {
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
		v1.POST("/user/:name/:action", func(c *gin.Context) {
			name := c.Param("name")
			action := c.Param("action")
			c.String(http.StatusOK, "Hello %s, you are %sing", name, action)
		})
	}
	s := &http.Server{
		Addr:           "0.0.0.0:3443:8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}