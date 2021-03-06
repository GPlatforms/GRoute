package models

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/url"

	"company/vpngo/admin/common"
	"company/vpngo/admin/logger"

	"github.com/happyEgg/wlog"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	GORM      *gorm.DB
	Config    *common.Config
	ErrLogger *wlog.WLogger
	PriKey    []byte
)

func InitConfig() {
	var err error
	Config = common.InitConfig("./conf.json")

	common.GenRsaKey(1024, "private.pem")
	PriKey, err = ioutil.ReadFile("private.pem")
	if err != nil {
		log.Fatalln("private.pem file read failed:", err)
	}

	dbInit()
	ErrLogger = logger.ErrorDiary()
}

func dbInit() {
	var err error
	value := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=%s",
		Config.Base.Mysql.User, Config.Base.Mysql.Password, Config.Base.Mysql.Host, Config.Base.Mysql.Name, url.QueryEscape("Asia/Shanghai"))

	GORM, err = gorm.Open("mysql", value)
	if err != nil {
		log.Fatalln("mysql connect failed:", err)
	}

	GORM.DB().SetMaxIdleConns(100)
	GORM.DB().SetMaxOpenConns(200)

	GORM.SingularTable(true)
}
