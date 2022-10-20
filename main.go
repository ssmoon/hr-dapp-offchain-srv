package main

import (
	"fmt"
	"hr-dapp/srv/pkg/conf"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// conf.InitEnv()
	// conf.InitDatabase()
	// cache.Init()

	r := gin.Default()
	//r.Use(ext.Auth())
	//routes.InitializeRoutes(r)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run(fmt.Sprintf(":%d", conf.Conf().Server.HTTPPort))
}
