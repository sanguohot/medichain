package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"medichain/service"
	"io/ioutil"
	"fmt"
)

func PongHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func AddUserHandler(c *gin.Context) {
	idCartNo := c.PostForm("idCartNo")
	orgUuidStr := c.PostForm("orgUuid")
	password := c.PostForm("password")
	err, user := service.AddUser(orgUuidStr, idCartNo, password)
	if err != nil {
		c.String(http.StatusOK, err.Error())
	}
	c.JSON(http.StatusOK, user)
}

func AddOrgHandler(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	err, org := service.AddOrg(name, password)
	if err != nil {
		c.String(http.StatusOK, err.Error())
	}
	c.JSON(http.StatusOK, org)
}

func AddFileHandler(c *gin.Context) {
	addressStr := c.PostForm("address")
	password := c.PostForm("password")
	ownerUuidStr := c.PostForm("ownerUuid")
	fileType := c.PostForm("fileType")
	fileDesc := c.PostForm("fileDesc")
	sha256Hash := c.PostForm("sha256Hash")
	file, header , err := c.Request.FormFile("file")
	fmt.Println("header.Filename ===>", header.Filename)
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		c.String(http.StatusOK, err.Error())
	}
	err, fileAction := service.AddFile(ownerUuidStr, addressStr, password, fileType, fileDesc, fileBytes, sha256Hash)
	if err != nil {
		c.String(http.StatusOK, err.Error())
	}
	c.JSON(http.StatusOK, fileAction)
}

func AddFileSignHandler(c *gin.Context) {
	addressStr := c.PostForm("address")
	password := c.PostForm("password")
	fileUuidStr := c.Param("fileUuid")
	keccak256Hash := c.PostForm("keccak256Hash")
	err, hash := service.AddFileSign(fileUuidStr, addressStr, password, keccak256Hash)
	if err != nil {
		c.String(http.StatusOK, err.Error())
	}
	c.String(http.StatusOK, hash.Hex())
}

func GetFileHandler(c *gin.Context) {
	fileUuidStr := c.Param("fileUuid")
	err, file := service.GetFile(fileUuidStr)
	if err != nil {
		c.String(http.StatusOK, err.Error())
	}
	c.Header("content-disposition", `attachment; filename=` + fileUuidStr)
	c.Data(http.StatusOK, http.DetectContentType(file), file)
}