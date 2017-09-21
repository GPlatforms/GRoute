package main

import (
	"company/vpngo/admin/server/controllers"
	"company/vpngo/admin/server/models"

	"github.com/gin-gonic/gin"
)

var (
	adminController   = &controllers.AdminController{}
	dnsInfoController = &controllers.DNSInfoController{}
)

func main() {
	models.InitConfig()

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())
	if models.Config.Mode != "prod" {
		router.Use(gin.Logger())
	}

	noAuth := router.Group("/admin/api/v1/app")
	{
		noAuth.POST("/login", adminController.AdminLogin)
	}

	auth := router.Group("/admin/api/v1/app")
	auth.Use(controllers.AuthRequired)
	{
		auth.GET("/config", dnsInfoController.GetDNSInfoList)
		auth.POST("/config", dnsInfoController.AddDNSInfo)
		auth.PATCH("/config", dnsInfoController.UpdateDNSInfo)
		auth.DELETE("/config", dnsInfoController.DeleteDNSInfo)
	}

	err := router.Run(":" + models.Config.Port)
	if err != nil {
		panic(err)
	}
}
