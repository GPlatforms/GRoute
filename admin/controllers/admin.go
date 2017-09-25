package controllers

import (
	// "fmt"

	"company/vpngo/admin/models"

	"github.com/gin-gonic/gin"
)

type AdminController struct{}

func (a *AdminController) AdminHome(cxt *gin.Context) {
	cxt.Redirect(302, "/admin/login")
}

func (a *AdminController) AdminLoginHtml(cxt *gin.Context) {
	username, _ := cxt.Cookie("username")
	if username == "" {
		cxt.HTML(200, "login.html", nil)
		return
	}

	cxt.Redirect(302, "/admin/index")
}

func (a *AdminController) AdminLogoutHtml(cxt *gin.Context) {
	cxt.SetCookie("username", "", 0, "/", "localhost", false, false)

	cxt.Redirect(302, "/admin/login")
}

func (a *AdminController) AdminIndexHtml(cxt *gin.Context) {
	cxt.HTML(200, "index.html", gin.H{"username": Username})
}

func (a *AdminController) AdminLogin(cxt *gin.Context) {
	username := cxt.PostForm("username")
	password := cxt.PostForm("password")
	if username == "" || password == "" {
		cxt.Redirect(302, "/admin/login")
		return
	}

	adminInfo := new(models.AdminInfo)
	adminInfo.GetAdminInfo(username)
	if password != adminInfo.Password {
		cxt.Redirect(302, "/admin/login")
		return
	}

	cxt.SetCookie("username", adminInfo.Username, 1111, "/", "localhost", false, false)
	cxt.Redirect(302, "/admin/index")
}
