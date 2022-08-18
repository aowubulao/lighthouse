package api

import "github.com/gin-gonic/gin"

func ConfigurationRegister(r *gin.Engine) {
	updateConfig(r)
}

func updateConfig(r *gin.Engine) {
	r.PUT("/api/v1/configuration/", func(c *gin.Context) {

	})
}
