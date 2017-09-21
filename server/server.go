package main

import (
	"company/vpngo/server/controllers"
	"company/vpngo/server/models"

	"github.com/gin-gonic/gin"
)

var configInfoController = &controllers.ConfigInfoController{}

func main() {
	models.InitConfig()

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())
	if models.Config.Mode != "prod" {
		router.Use(gin.Logger())
	}

	router.GET("/groute/v1/config", configInfoController.GetConfigInfo)

	err := router.Run(":" + models.Config.Port)
	if err != nil {
		panic(err)
	}
}
