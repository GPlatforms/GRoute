package controllers

import (
	"encoding/json"
	"math"
	"net/http"
	"strconv"
	"time"

	"company/vpngo/server/common"
	"company/vpngo/server/models"

	"github.com/gin-gonic/gin"
)

type ConfigInfoController struct{}

func (c *ConfigInfoController) GetConfigInfo(cxt *gin.Context) {
	timestamp := cxt.Query("timestamp")
	appIdStr := cxt.Query("app_id")
	sign := cxt.Query("sign")

	models.ErrLogger.Info("client request url:", cxt.ClientIP(), cxt.Request.RequestURI)

	sa := models.Config.AppSecret + appIdStr + timestamp
	if common.SHA1Sign(sa) != sign {
		models.CommonResult(cxt, models.SignErr)
		return
	}

	stamp, _ := strconv.ParseInt(timestamp, 10, 64)
	nowStamp := time.Now().Unix()
	sub := nowStamp - stamp
	if math.Abs(float64(sub)) > 300 {
		models.CommonResult(cxt, models.TimeErr)
		return
	}

	appId, _ := strconv.ParseInt(appIdStr, 10, 64)
	configInfo := new(models.ConfigInfo)
	err := configInfo.GetConfigInfo(appId)
	if err != nil {
		models.ErrLogger.Error("get config_info error:", appId, err)
	}

	value := make([]string, 0, 10)
	if configInfo.DnsUrl != "" {
		err := json.Unmarshal([]byte(configInfo.DnsUrl), &value)
		if err != nil {
			models.ErrLogger.Error("json unmarshal error:", appId, err)
		}
	}

	params := make(map[string]interface{})
	if configInfo.Params != "" {
		err := json.Unmarshal([]byte(configInfo.Params), &params)
		if err != nil {
			models.ErrLogger.Error("json unmarshal error:", appId, configInfo.Params, err)
		}
	}

	params["code"] = 200
	params["msg"] = "success"
	params["base_url"] = value

	b, _ := json.Marshal(params)

	aesEnc := common.AesEncrypt{}
	aesData, _ := aesEnc.Encrypt(b, models.Config.DataSecret)

	cxt.String(http.StatusOK, aesData)
}
