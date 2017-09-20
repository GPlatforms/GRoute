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
	SignErr = &ResultData{1102, "sign error", nullData}
	TimeErr = &ResultData{1103, "time error", nullData}
)

func CommonResult(c *gin.Context, result *ResultData) {
	if result.Data == nil {
		result.Data = nullData
	}

	c.JSON(http.StatusOK, result)
}
