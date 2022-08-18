package utils

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func GetRequestJsonMap(c *gin.Context) (map[string]string, error) {
	var requestMap = make(map[string]string)
	err := json.NewDecoder(c.Request.Body).Decode(&requestMap)
	return requestMap, err
}
