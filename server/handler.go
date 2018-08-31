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

func DoJsonResponse(c *gin.Context, err error, p interface{})  {
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Code": "FAIL",
			"Message": err.Error(),
		})
	}else {
		c.JSON(http.StatusOK, gin.H{
			"Code": "SUCC",
			"Data": p,
		})
	}
}
func AddUserHandler(c *gin.Context) {
	idCartNo := c.PostForm("idCartNo")
	orgUuidStr := c.PostForm("orgUuid")
	password := c.PostForm("password")
	err, obj := service.AddUser(orgUuidStr, idCartNo, password)
	DoJsonResponse(c, err, obj)
}

func AddOrgHandler(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	err, obj := service.AddOrg(name, password)
	DoJsonResponse(c, err, obj)
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
		DoJsonResponse(c, err, nil)
	}
	err, obj := service.AddFile(ownerUuidStr, addressStr, password, fileType, fileDesc, fileBytes, sha256Hash)
	DoJsonResponse(c, err, obj)
}

func AddFileSignHandler(c *gin.Context) {
	addressStr := c.PostForm("address")
	password := c.PostForm("password")
	fileUuidStr := c.Param("fileUuid")
	keccak256Hash := c.PostForm("keccak256Hash")
	err, hash := service.AddFileSign(fileUuidStr, addressStr, password, keccak256Hash)
	DoJsonResponse(c, err, hash.Hex())
}

func GetFileHandler(c *gin.Context) {
	fileUuidStr := c.Param("fileUuid")
	err, file := service.GetFile(fileUuidStr)
	if err != nil {
		DoJsonResponse(c, err, nil)
	}
	c.Header("content-disposition", `attachment; filename=` + fileUuidStr)
	c.Data(http.StatusOK, http.DetectContentType(file), file)
}

func GetFileSignerAndDataListHandler(c *gin.Context) {
	fileUuidStr := c.Param("fileUuid")
	startStr := c.Query("start")
	limitStr := c.Query("limit")
	err, list := service.GetFileSignerAndDataList(fileUuidStr, startStr, limitStr)
	DoJsonResponse(c, err, list)
}

func GetFileAddLogListHandler(c *gin.Context) {
	startStr := c.Query("start")
	limitStr := c.Query("limit")
	idCartNo := c.Query("idCartNo")
	err, list := service.GetFileAddLogList(idCartNo, startStr, limitStr)
	//if err != nil {
	//	c.String(http.StatusBadRequest, err.Error())
	//}
	//c.JSON(http.StatusOK, list)
	DoJsonResponse(c, err, list)
}