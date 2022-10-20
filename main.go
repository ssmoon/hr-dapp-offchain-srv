package main

import (
	"fmt"
	"hr-dapp/srv/pkg/conf"
	"hr-dapp/srv/pkg/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	conf.InitEnv()
	conf.InitDatabase()

	r := gin.Default()
	routes.InitializeRoutes(r)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run(fmt.Sprintf(":%d", conf.Conf().Server.HTTPPort))
}
