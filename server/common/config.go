package common

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"
)

var RUNMODE string

type Config struct {
	Port      string
	AppSecret string
	Mode      string
	Base      *DbBase
}

type DBConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
type DbBase struct {
	Mysql *DBConfig `json:"mysql"`
}

type baseConfig struct {
	Port      int     `json:"port"`
	AppSecret string  `json:"app_secret"`
	Mode      string  `json:"mode"`
	Dev       *DbBase `json:"dev"`
	Prod      *DbBase `json:"prod"`
}

func InitConfig(path string) *Config {
	body, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln("conf.json file not found")
	}

	baseConfig := new(baseConfig)
	config := new(Config)

	err = json.Unmarshal(body, baseConfig)
	if err != nil {
		log.Fatalln("json format error")
	}

	config.Port = strconv.Itoa(baseConfig.Port)
	config.AppSecret = baseConfig.AppSecret
	config.Mode = baseConfig.Mode
	if baseConfig.Mode == "dev" {
		RUNMODE = "dev"
		config.Base = baseConfig.Dev
	} else {
		RUNMODE = "prod"
		config.Base = baseConfig.Prod
	}

	return config
}
