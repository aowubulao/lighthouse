package api

import (
	"github.com/aowubulao/lighouse-go-server/config"
	"github.com/aowubulao/lighouse-go-server/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func LoginRegister(r *gin.Engine) {
	loginApi(r)
}

func loginApi(r *gin.Engine) {
	r.POST("/api/v1/login", func(c *gin.Context) {
		requestJsonMap, err := utils.GetRequestJsonMap(c)
		if err != nil {
			returnError(c)
			return
		}
		reqPasswd := requestJsonMap["password"]
		if reqPasswd != config.GetWebPassword() {
			returnCommon(c, "password wrong", http.StatusForbidden, "")
		} else {
			returnOkResult(c, generateTmpToken(config.GetWebPassword()))
		}
	})
}

func generateTmpToken(password string) string {
	timestamp := time.Now().UnixNano() / 1e6
	str := password + "?" + strconv.FormatInt(timestamp, 10)
	tmpToken, _ := utils.EncryptByAes([]byte(str))
	return tmpToken
}
