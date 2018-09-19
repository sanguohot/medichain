package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sanguohot/medichain/service"
	"github.com/sanguohot/medichain/util"
	"github.com/sanguohot/medichain/zap"
	uberZap "go.uber.org/zap"
	"io/ioutil"
	"net/http"
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
		zap.Logger.Error(
			err.Error(),
			uberZap.String("remote", c.Request.RemoteAddr),
			uberZap.String("method", c.Request.Method),
			uberZap.String("url", c.Request.RequestURI),
			uberZap.Any("form", c.Request.Form),
			uberZap.Any("param", c.Params),
		)
	} else {
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
	orgUuidStr := c.PostForm("orgUuid")
	fileType := c.PostForm("fileType")
	fileDesc := c.PostForm("fileDesc")
	sha256Hash := c.PostForm("sha256Hash")
	file, _ , err := c.Request.FormFile("file")
	if err != nil {
		DoJsonResponse(c, err, nil)
	}
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		DoJsonResponse(c, err, nil)
	}
	err, obj := service.AddFile(ownerUuidStr, orgUuidStr, addressStr, password, fileType, fileDesc, fileBytes, sha256Hash)
	DoJsonResponse(c, err, obj)
}

func AddFileSignHandler(c *gin.Context) {
	addressStr := c.PostForm("address")
	password := c.PostForm("password")
	fileUuidStr := c.Param("fileUuid")
	keccak256Hash := c.PostForm("keccak256Hash")
	err, obj := service.AddFileSign(fileUuidStr, addressStr, password, keccak256Hash)
	DoJsonResponse(c, err, obj)
}

func GetFileHandler(c *gin.Context) {
	fileUuidStr := c.Param("fileUuid")
	err, file := service.GetFile(fileUuidStr)
	if err != nil {
		DoJsonResponse(c, err, nil)
	}
	c.Header("content-disposition", `attachment; filename=` + fileUuidStr)
	// charset=utf-8
	// application/octet-stream
	defaultDetect := util.DetectContentType(file)
	zap.Sugar.Infof("file %s content type auto detect ===> %s", fileUuidStr, defaultDetect)
	c.Data(http.StatusOK, defaultDetect, file)
	//c.Data(http.StatusOK, "text/plain;charset=utf8;gbk", file)
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
	orgUuidStr := c.Query("orgUuid")
	fromTimeStr := c.Query("fromTime")
	toTimeStr := c.Query("toTime")
	err, list := service.GetFileAddLogList(idCartNo, orgUuidStr, fromTimeStr, toTimeStr, startStr, limitStr)
	DoJsonResponse(c, err, list)
}

func GetFileAddLogDetailHandler(c *gin.Context) {
	fileUuid := c.Param("fileUuid")
	err, detail := service.GetFileAddLogDetail(fileUuid)
	DoJsonResponse(c, err, detail)
}