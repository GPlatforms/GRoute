package main

import (
	// "net/http"

	"company/vpngo/admin/controllers"
	"company/vpngo/admin/models"

	"github.com/gin-gonic/gin"
)

var (
	adminController      = &controllers.AdminController{}
	configInfoController = &controllers.ConfigInfoController{}
)

func main() {
	models.InitConfig()

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())
	if models.Config.Mode != "prod" {
		router.Use(gin.Logger())
	}

	router.LoadHTMLGlob("web/*.html")

	router.GET("/admin", adminController.AdminHome)
	router.GET("/admin/login", adminController.AdminLoginHtml)
	router.GET("/admin/logout", adminController.AdminLogoutHtml)
	router.POST("/admin/login", adminController.AdminLogin)

	webAuth := router.Group("/admin")
	webAuth.Use(controllers.AuthRequired)
	{
		webAuth.GET("/index", adminController.AdminIndexHtml)
		webAuth.GET("/config_list", configInfoController.ConfigListHtml)
	}

	router.Static("/css", "./web/css")
	router.Static("/images", "./web/images")
	router.Static("/fonts", "./web/fonts")
	router.Static("/js", "./web/js")
	router.Static("/lib", "./web/lib")

	router.Static("/admin/css", "./web/css")
	router.Static("/admin/images", "./web/images")
	router.Static("/admin/fonts", "./web/fonts")
	router.Static("/admin/js", "./web/js")
	router.Static("/admin/lib", "./web/lib")

	router.StaticFile("/welcome.html", "./web/welcome.html")
	router.StaticFile("/dns-add.html", "./web/dns-add.html")
	router.StaticFile("/dns-edit.html", "./web/dns-edit.html")
	router.StaticFile("/dns-list.html", "./web/dns-list.html")

	err := router.Run(":" + models.Config.Port)
	if err != nil {
		panic(err)
	}
}
