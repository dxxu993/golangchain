package conf

import (
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type AppConf struct {
	OpenaiKey               string `toml:"OpenaiKey"`
	AppRootPath             string
	MySqlDB                 MySqlDB `toml:"MySqlDB"`
	WechatMpOpenID          string  `toml:"WechatMpOpenID"`
	WechatMpAppSecret       string  `toml:"WechatMpAppSecret"`
	NewUserLimitGptReplyNum int64   `toml:"NewUserLimitGptReplyNum"`
}

type MySqlDB struct {
	Host     string
	Port     string
	UserName string
	Password string
	Database string
}

var RelRootPath = "../"

func InitAppConf() *AppConf {
	rootPath, err := filepath.Abs(RelRootPath)
	if err != nil {
		panic("get root path fail")
	}
	config := &AppConf{}
	configPath := filepath.Join(rootPath, "./conf/app.toml")
	_, err1 := toml.DecodeFile(configPath, config)
	if err1 != nil {
		panic("read app conf fail:" + err1.Error())
	}
	config.AppRootPath = rootPath
	return config
}
