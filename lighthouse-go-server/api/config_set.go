package api

import (
	"github.com/aowubulao/lighouse-go-server/utils"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func init() {
	getWd, err := os.Getwd()
	if err != nil {
		return
	}
	exist := utils.FileIsExist(getWd + "/data")
	if !exist {
		err := os.Mkdir(getWd+"/data", 0755)
		if err != nil {
			log.Println("Create data dir error: ", err.Error())
			return
		}
	}
	exist2 := utils.FileIsExist(getWd + "/version")
	if !exist2 {
		err := os.Mkdir(getWd+"/version", 0755)
		if err != nil {
			log.Println("Create version dir error: ", err.Error())
			return
		}
	}
}

func ConfigSetRegister(r *gin.Engine) {
	createConfigSet(r)
	getConfigSet(r)
}

func createConfigSet(r *gin.Engine) {
	r.POST("/api/v1/config_set/", func(c *gin.Context) {
		requestJsonMap, err := utils.GetRequestJsonMap(c)
		if err != nil {
			returnError(c)
			return
		}
		configFileName := requestJsonMap["name"]
		configJson := requestJsonMap["config"]
		if configFileName == "" || configJson == "" {
			returnCommon(c, "file or config is null", 400, "")
			return
		}
		fileName := "data/" + configFileName
		exist := utils.FileIsExist(fileName)
		if exist {
			returnCommon(c, "file is exist", 400, "")
			return
		}

		thisFile, err := os.Create(fileName)
		defer func(thisFile *os.File) {
			err := thisFile.Close()
			if err != nil {
			}
		}(thisFile)

		if err != nil {
			log.Println("Create file error: " + err.Error())
			returnError(c)
			return
		}

		err = utils.WriteFileByObj(thisFile, configJson)
		if err != nil {
			log.Println("Write file error: " + err.Error())
			returnError(c)
			return
		}

		returnOkResult(c, "")
	})
}

func getConfigSet(r *gin.Engine) {
	r.GET("/api/v1/config_set/", func(c *gin.Context) {
		files, err := utils.GetDirFiles("data/")
		if err != nil {
			log.Println("Read files error: ", err.Error())
			returnError(c)
			return
		}
		returnOkResult(c, files)
	})
}
