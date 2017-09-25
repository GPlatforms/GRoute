package models

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResultData struct {
	Code int32       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

var (
	RecordNotFound = "record not found"
	nullData       = struct{}{}
)

var (
	DataBaseErr     = &ResultData{1001, "", nullData}
	UnauthorizedErr = &ResultData{1100, "登录失效,请重新登录", nullData}
	LoginErr        = &ResultData{1102, "用户名或密码错误", nullData}
	AppIdErr        = &ResultData{1103, "app_id出错", nullData}
	DNSUrlErr       = &ResultData{1103, "dns_url传值有误", nullData}
)

func CommonResult(c *gin.Context, result *ResultData) {
	if result.Data == nil {
		result.Data = nullData
	}

	c.JSON(http.StatusOK, result)
}

func RightResult(c *gin.Context, data interface{}) {
	var result ResultData
	result.Code = 200
	if data == nil {
		result.Data = nullData
	} else {
		result.Data = data
	}

	c.JSON(http.StatusOK, result)
}
