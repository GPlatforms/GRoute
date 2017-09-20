package main

import (
	"company/vpngo/server/controllers"
	"company/vpngo/server/models"

	"github.com/gin-gonic/gin"
)

var dnsInfoController = &controllers.DNSInfoController{}

func main() {
	models.InitConfig()

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/api/v1/app/config/dns_info", dnsInfoController.GetDNSInfo)

	err := router.Run(":" + models.Config.Port)
	if err != nil {
		panic(err)
	}
}
