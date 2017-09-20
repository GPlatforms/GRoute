package controllers

import (
	"encoding/json"
	"log"
	"math"
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
	dnsInfo.GetDNSInfo(appId)

	value := make([]string, 0, 10)
	err := json.Unmarshal([]byte(dnsInfo.DnsUrl), &value)
	log.Println(err)
	data := map[string]interface{}{"base_url": value}

	result := new(models.ResultData)
	result.Code = 200
	result.Msg = "success"
	result.Data = data

	models.CommonResult(cxt, result)
}
