package api

import (
	"github.com/aowubulao/lighouse-go-server/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// HttpInterceptor 拦截器
func HttpInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		aesByte, err := utils.DecryptByAes(c.GetHeader("web-token"))
		if err != nil {
			returnCommon(c, "forbidden", http.StatusUnauthorized, "")
			return
		}
		aesToken := string(aesByte)
		splitList := strings.Split(aesToken, "?")
		if len(splitList) != 2 {
			returnCommon(c, "forbidden", http.StatusUnauthorized, "")
			return
		}
		timestampStr := splitList[1]
		timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
		if (time.Now().UnixNano()/1e6)-timestamp > (30 * 1000 * 60 * 30) {
			returnCommon(c, "token expired", http.StatusUnauthorized, "")
			return
		}
	}
}
