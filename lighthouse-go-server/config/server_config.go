package config

import (
	"encoding/json"
	"os"
)

var port string
var webPassword string
var token string
var aes string

func init() {
	configJson, _ := os.ReadFile("config.json")
	var configMap = make(map[string]string)
	_ = json.Unmarshal(configJson, &configMap)
	port = configMap["port"]
	webPassword = configMap["web_password"]
	token = configMap["token"]
	aes = configMap["aes"]
}

func GetPort() string {
	return port
}

func GetWebPassword() string {
	return webPassword
}

func GetToken() string {
	return token
}

func GetAes() string {
	return aes
}
