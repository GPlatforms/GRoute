package controllers

import (
	"encoding/json"
	"math"
	"strconv"
	"time"

	"company/vpngo/admin/server/common"
	"company/vpngo/admin/server/models"

	"github.com/gin-gonic/gin"
)

type AdminController struct{}

func (a *AdminController) AdminLogin(cxt *gin.Context) {
	username := cxt.PostForm("username")
	password := cxt.PostForm("password")

	if username == "" || password == "" {
		models.CommonResult(cxt, models.LoginErr)
		return
	}

	adminInfo := new(models.AdminInfo)
	adminInfo.GetAdminInfo(username)
	if password != adminInfo.Password {
		models.CommonResult(cxt, models.LoginErr)
		return
	}

	token := common.JWT(models.PriKey, username, models.Config.TokenExpire)
	adminInfo.Token = token

	result := new(models.ResultData)
	result.Code = 200
	result.Data = adminInfo

	models.CommonResult(cxt, result)
}
