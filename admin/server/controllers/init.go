package controllers

import (
	"errors"

	"company/vpngo/admin/server/common"
	"company/vpngo/admin/server/models"

	"github.com/gin-gonic/gin"
)

var Username string

func AuthRequired(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		models.CommonResult(c, models.UnauthorizedErr)
		c.Abort()
		return
	}

	uid, err := CheckToken(token)
	if err != nil {
		common.CommonResult(c, common.UnauthorizedErr)
		c.Abort()
		return
	}

	Uid = uid
}

func CheckToken(token string) (uint64, error) {
	t, err := common.ParseJWT(models.PriKey, token)
	if err != nil {
		return err
	}
	claims, ok := t.Claims.(*common.MyCustomClaims)
	if !ok || !t.Valid {
		return errors.New("Unauthorized")
	}

	if claims.Username == "" {
		return errors.New("Unauthorized")
	}

	Username = claims.Username
	return nil
}
