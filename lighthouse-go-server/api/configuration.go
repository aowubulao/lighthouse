package api

import (
	"github.com/aowubulao/lighouse-go-server/utils"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func ConfigurationRegister(r *gin.Engine) {
	getConfig(r)
	getConfigVersion(r)
	updateConfig(r)
}

func getConfig(r *gin.Engine) {
	r.GET("/api/v1/configuration/:name", func(c *gin.Context) {
		name := c.Params.ByName("name")
		configContent, err := os.ReadFile("data/" + name)
		if err != nil {
			log.Println("Read data file error: ", err.Error())
			returnError(c)
			return
		}
		versionContent, err := os.ReadFile("version/" + name)
		if err != nil {
			log.Println("Read version file error: ", err.Error())
			returnError(c)
			return
		}
		returnOkResult(c, "{\"config\":"+string(configContent)+",\"version\":\""+string(versionContent)+"\"}")
	})
}

func getConfigVersion(r *gin.Engine) {
	r.GET("/api/v1/configuration/:name/version", func(c *gin.Context) {
		name := c.Params.ByName("name")
		versionContent, err := os.ReadFile("version/" + name)
		if err != nil {
			log.Println("Read version file error: ", err.Error())
			returnError(c)
			return
		}
		returnOkResult(c, "{\"version\":\" "+string(versionContent)+"\"}")
	})
}

func updateConfig(r *gin.Engine) {
	r.PUT("/api/v1/configuration/:name", func(c *gin.Context) {
		requestJsonMap, err := utils.GetRequestJsonMap(c)
		if err != nil {
			returnError(c)
			return
		}
		name := c.Params.ByName("name")

		err = utils.WriteFileByPath("data/"+name, requestJsonMap["config"])
		if err != nil {
			log.Println("Write file error: ", err.Error())
			returnError(c)
			return
		}
	})
}
