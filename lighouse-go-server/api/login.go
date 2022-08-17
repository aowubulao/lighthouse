package api

import (
	"encoding/json"
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
		var requestMap = make(map[string]string)
		err := json.NewDecoder(c.Request.Body).Decode(&requestMap)
		if err != nil {
			returnError(c)
			return
		}
		reqPasswd := requestMap["password"]
		if reqPasswd != config.GetWebPassword() {
			returnCommon(c, "password wrong", http.StatusForbidden, "")
		} else {
			returnOkResult(c, generateTmpToken(config.GetWebPassword()))
		}
	})
}

func generateTmpToken(password string) string {
	timestamp := time.Now().UnixNano() / 1e6
	str := password + strconv.FormatInt(timestamp, 10)
	tmpToken, _ := utils.EncryptByAes([]byte(str))
	return tmpToken
}
