package controllers

import (
	"encoding/json"
	"math"
	"strconv"
	"strings"
	"time"

	"company/vpngo/admin/server/common"
	"company/vpngo/admin/server/models"

	"github.com/gin-gonic/gin"
)

type DNSInfoController struct {
	DNSInfo []*models.DNSInfo `json:"dns_info"`
	Page    int               `json:"page"`
	Count   int64             `json:"count"`
}

func (d *DNSInfoController) GetDNSInfoList(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	if page == 1 {
		page = 0
	}

	dnsInfo := new(models.DNSInfo)
	count := dnsInfo.Count()

	dnsInfoList, err := dnsInfo.GetDNSInfoList(page, 20)
	if err != nil {
		dataErr := new(models.ResultData)
		*dataErr = *models.DataBaseErr
		dataErr.Msg = err.Error()
		models.CommonResult(c, dataErr)
		return
	}

	returnData := new(DNSInfoController)
	returnData.DNSInfo = dnsInfoList
	returnData.Page = page
	returnData.Count = count

	models.RightResult(c, returnData)
}

func (d *DNSInfoController) AddDNSInfo(c *gin.Context) {
	app_id, _ := strconv.ParseInt(c.PostForm("app_id"), 10, 64)
	dnsUrl := c.PostForm("dns_url")

	if app_id == 0 {
		models.CommonResult(c, models.AppIdErr)
		return
	}

	var dnsArr []string
	err := json.Unmarshal([]byte(dnsUrl), &dnsArr)
	if err != nil {
		models.CommonResult(c, models.DNSUrlErr)
		return
	}

	dnsInfo := new(models.DNSInfo)
	err = dnsInfo.AddDNSInfo(app_id, dnsUrl)
	if err != nil {
		dataErr := new(models.ResultData)
		*dataErr = *models.DataBaseErr
		dataErr.Msg = err.Error()
		models.CommonResult(c, dataErr)
		return
	}

	count := dnsInfo.Count()
	dnsInfoList, err := dnsInfo.GetDNSInfoList(0, 20)
	if err != nil {
		dataErr := new(models.ResultData)
		*dataErr = *models.DataBaseErr
		dataErr.Msg = err.Error()
		models.CommonResult(c, dataErr)
		return
	}

	returnData := new(DNSInfoController)
	returnData.DNSInfo = dnsInfoList
	returnData.Page = 1
	returnData.Count = count

	models.RightResult(c, returnData)
}

func (d *DNSInfoController) UpdateDNSInfo(c *gin.Context) {
	app_id, _ := strconv.ParseInt(c.PostForm("app_id"), 10, 64)
	dnsUrl := c.PostForm("dns_url")

	if app_id == 0 {
		models.CommonResult(c, models.AppIdErr)
		return
	}

	var dnsArr []string
	err := json.Unmarshal([]byte(dnsUrl), &dnsArr)
	if err != nil {
		models.CommonResult(c, models.DNSUrlErr)
		return
	}

	dnsInfo := new(models.DNSInfo)
	err = dnsInfo.ModifyDNSInfo(app_id, dnsUrl)
	if err != nil {
		dataErr := new(models.ResultData)
		*dataErr = *models.DataBaseErr
		dataErr.Msg = err.Error()
		models.CommonResult(c, dataErr)
		return
	}

	count := dnsInfo.Count()
	dnsInfoList, err := dnsInfo.GetDNSInfoList(0, 20)
	if err != nil {
		dataErr := new(models.ResultData)
		*dataErr = *models.DataBaseErr
		dataErr.Msg = err.Error()
		models.CommonResult(c, dataErr)
		return
	}

	returnData := new(DNSInfoController)
	returnData.DNSInfo = dnsInfoList
	returnData.Page = 1
	returnData.Count = count

	models.RightResult(c, returnData)
}

func (d *DNSInfoController) DeleteDNSInfo(c *gin.Context) {
	appIdStr := c.PostForm("app_id")

	appIdStrArr := strings.Split(appIdStr, ",")
	appIdArr := make([]int64, 0, len(appIdStrArr))
	for _, v := range appIdStrArr {
		id, _ := strconv.ParseInt(v, 10, 64)
		if id == 0 {
			continue
		}
		appIdArr = append(appIdArr, id)
	}

	if len(appIdArr) == 0 {
		models.CommonResult(c, models.AppIdErr)
		return
	}

	dnsInfo := new(models.DNSInfo)
	dnsInfo.DeleteDNSInfo(appIdArr, dnsUrl)

	count := dnsInfo.Count()
	dnsInfoList, err := dnsInfo.GetDNSInfoList(0, 20)
	if err != nil {
		dataErr := new(models.ResultData)
		*dataErr = *models.DataBaseErr
		dataErr.Msg = err.Error()
		models.CommonResult(c, dataErr)
		return
	}

	returnData := new(DNSInfoController)
	returnData.DNSInfo = dnsInfoList
	returnData.Page = 1
	returnData.Count = count

	models.RightResult(c, returnData)
}
