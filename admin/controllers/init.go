package controllers

import (
	"errors"
	"fmt"

	"company/vpngo/admin/common"
	"company/vpngo/admin/models"

	"github.com/gin-gonic/gin"
)

var Username string

func AuthRequired(c *gin.Context) {
	username, err := c.Cookie("username")
	fmt.Println("user:", username, err)
	if username == "" {
		c.Redirect(302, "/admin/login")
		return
	}

	Username = username
	// token := c.Request.Header.Get("Authorization")
	// if token == "" {
	// 	models.CommonResult(c, models.UnauthorizedErr)
	// 	c.Abort()
	// 	return
	// }

	// err := CheckToken(token)
	// if err != nil {
	// 	models.CommonResult(c, models.UnauthorizedErr)
	// 	c.Abort()
	// 	return
	// }
}

func CheckToken(token string) error {
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
