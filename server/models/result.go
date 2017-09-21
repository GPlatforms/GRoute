package models

import (
	"net/http"

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
	c.JSON(http.StatusOK, result)
}
