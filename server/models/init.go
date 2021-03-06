package models

import (
	"fmt"
	"log"

	"company/vpngo/server/common"
	"company/vpngo/server/logger"

	"github.com/happyEgg/wlog"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	GORM      *gorm.DB
	Config    *common.Config
	ErrLogger *wlog.WLogger
)

func InitConfig() {
	Config = common.InitConfig("./conf.json")

	dbInit()
	log.Println("init finish")
	ErrLogger = logger.ErrorDiary()
}

func dbInit() {
	var err error
	value := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=Local",
		Config.Base.Mysql.User, Config.Base.Mysql.Password, Config.Base.Mysql.Host, Config.Base.Mysql.Name)

	GORM, err = gorm.Open("mysql", value)
	if err != nil {
		log.Fatalln("mysql connect failed:", err)
	}

	GORM.DB().SetMaxIdleConns(100)
	GORM.DB().SetMaxOpenConns(200)

	GORM.SingularTable(true)
}
