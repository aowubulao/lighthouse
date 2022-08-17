package api

import (
	"github.com/gin-gonic/gin"
)

func returnCommon(c *gin.Context, msg string, code int, data string) {
	c.JSON(code, gin.H{
		"message": msg,
		"data":    data,
	})
}

func returnError(c *gin.Context) {
	returnCommon(c, "error", 500, "")
}

func returnOk(c *gin.Context) {
	returnCommon(c, "ok", 200, "")
}

func returnOkResult(c *gin.Context, data string) {
	returnCommon(c, "ok", 200, data)
}
