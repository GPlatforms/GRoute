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

type DNSInfoController struct{}

func (d *DNSInfoController) GetDNSInfo(cxt *gin.Context) {
	timestamp := cxt.Query("timestamp")
	appIdStr := cxt.Query("app_id")
	sign := cxt.Query("sign")

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
	dnsInfo := new(models.DNSInfo)
	err := dnsInfo.GetDNSInfo(appId)
	if err != nil {
		models.ErrLogger.Error("get dns_info error:", appId, err)
	}

	value := make([]string, 0, 10)
	if dnsInfo.DnsUrl != "" {
		err := json.Unmarshal([]byte(dnsInfo.DnsUrl), &value)
		if err != nil {
			models.ErrLogger.Error("json unmarshal error:", appId, err)
		}
	}

	params := make(map[string]interface{})
	if dnsInfo.Params != "" {
		err := json.Unmarshal([]byte(dnsInfo.Params), &params)
		if err != nil {
			models.ErrLogger.Error("json unmarshal error:", appId, dnsInfo.Params, err)
		}
	}

	params["code"] = 200
	params["msg"] = "success"
	params["base_url"] = value

	cxt.JSON(http.StatusOK, params)
}
