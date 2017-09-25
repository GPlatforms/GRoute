package models

import (
	"encoding/json"
	"net/http"

	"company/vpngo/server/common"

	"github.com/gin-gonic/gin"
)

type ResultData struct {
	Code    int32       `json:"code"`
	Msg     string      `json:"msg"`
	BaseUrl interface{} `json:"base_url"`
}

var (
	RecordNotFound = "record not found"
)

var (
	SignErr = &ResultData{1102, "sign error", ""}
	TimeErr = &ResultData{1103, "time error", ""}
)

func CommonResult(c *gin.Context, result *ResultData) {
	b, _ := json.Marshal(result)
	aesEnc := common.AesEncrypt{}
	aesData, _ := aesEnc.Encrypt(b, Config.DataSecret)

	c.String(http.StatusOK, string(aesData))
}
