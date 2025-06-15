package main

import (
	"fmt"

	"github.com/DavidLee0620/ExchangeApp/Exchangeapp_backend/config"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitConfig()
	type Info struct {
		message string
	}

	fmt.Println(config.AppConfig.App.Port)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(config.AppConfig.App.Port) // listen and serve on 0.0.0.0:8080
}
