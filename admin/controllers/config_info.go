package controllers

import (
	"encoding/json"
	"strconv"
	"strings"

	"company/vpngo/admin/models"

	"github.com/gin-gonic/gin"
)

type ConfigInfoController struct {
	ConfigInfo []*models.ConfigInfo
	Page       int

	Count int64
}

func (c *ConfigInfoController) ConfigListHtml(cxt *gin.Context) {
	page, _ := strconv.Atoi(cxt.Query("page"))
	configListInfo := c.GetConfigInfoList(page)

	cxt.HTML(200, "dns-list.html", configListInfo)
}

func (d *ConfigInfoController) GetConfigInfoList(page int) *ConfigInfoController {
	returnData := new(ConfigInfoController)

	if page == 0 {
		page = 1
	}

	configInfo := new(models.ConfigInfo)
	count := configInfo.Count()

	configInfoList, err := configInfo.GetConfigInfoList(page, 20)
	if err != nil {
		return returnData
	}

	returnData.ConfigInfo = configInfoList
	returnData.Page = page
	returnData.Count = count

	return returnData
}

func (d *ConfigInfoController) AddConfigInfo(c *gin.Context) {
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

	configInfo := new(models.ConfigInfo)
	err = configInfo.AddConfigInfo(app_id, dnsUrl)
	if err != nil {
		dataErr := new(models.ResultData)
		*dataErr = *models.DataBaseErr
		dataErr.Msg = err.Error()
		models.CommonResult(c, dataErr)
		return
	}

	count := configInfo.Count()
	configInfoList, err := configInfo.GetConfigInfoList(0, 20)
	if err != nil {
		dataErr := new(models.ResultData)
		*dataErr = *models.DataBaseErr
		dataErr.Msg = err.Error()
		models.CommonResult(c, dataErr)
		return
	}

	returnData := new(ConfigInfoController)
	returnData.ConfigInfo = configInfoList
	returnData.Page = 1
	returnData.Count = count

	models.RightResult(c, returnData)
}

func (d *ConfigInfoController) UpdateConfigInfo(c *gin.Context) {
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

	configInfo := new(models.ConfigInfo)
	err = configInfo.ModifyConfigInfo(app_id, dnsUrl)
	if err != nil {
		dataErr := new(models.ResultData)
		*dataErr = *models.DataBaseErr
		dataErr.Msg = err.Error()
		models.CommonResult(c, dataErr)
		return
	}

	count := configInfo.Count()
	configInfoList, err := configInfo.GetConfigInfoList(0, 20)
	if err != nil {
		dataErr := new(models.ResultData)
		*dataErr = *models.DataBaseErr
		dataErr.Msg = err.Error()
		models.CommonResult(c, dataErr)
		return
	}

	returnData := new(ConfigInfoController)
	returnData.ConfigInfo = configInfoList
	returnData.Page = 1
	returnData.Count = count

	models.RightResult(c, returnData)
}

func (d *ConfigInfoController) DeleteConfigInfo(c *gin.Context) {
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

	configInfo := new(models.ConfigInfo)
	configInfo.DeleteConfigInfo(appIdArr)

	count := configInfo.Count()
	configInfoList, err := configInfo.GetConfigInfoList(0, 20)
	if err != nil {
		dataErr := new(models.ResultData)
		*dataErr = *models.DataBaseErr
		dataErr.Msg = err.Error()
		models.CommonResult(c, dataErr)
		return
	}

	returnData := new(ConfigInfoController)
	returnData.ConfigInfo = configInfoList
	returnData.Page = 1
	returnData.Count = count

	models.RightResult(c, returnData)
}
