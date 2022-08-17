package main

import (
	"github.com/aowubulao/lighouse-go-server/api"
	"github.com/aowubulao/lighouse-go-server/config"
	"github.com/gin-gonic/gin"
)

func main() {
	// === base config ===
	r := gin.Default()
	r.Use(gin.Logger())

	// === controllers ===
	api.LoginRegister(r)

	r.Use(api.HttpInterceptor())

	// === start service ====
	err := r.Run(config.GetPort())
	if err != nil {
		return
	}
}
