package config

import (
	"github.com/go-ini/ini"
	"github.com/joho/godotenv"
	"os"
)

var AllowAllOrigins bool //跨域设置
var Env string           //运行环境
var RunMode string       //运行模式
var Host string          //web地址
var Port string          //web端口

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	Env = os.Getenv("ENV")
	var app = getConfig("app")
	var cfg, _ = app.GetSection("cors")
	AllowAllOrigins = cfg.Key("ALLOW_ALL_ORIGINS").MustBool(false)
	RunMode = app.Section("").Key("RUN_MODE").MustString("debug")
	Host = app.Section("server").Key("HTTP_HOST").MustString("")
	Port = app.Section("server").Key("HTTP_PORT").MustString("8000")
}

/*
	获取配置文件
*/
func getConfig(configName string) *ini.File {
	cfg, err := ini.Load("env/" + Env + "/" + configName + ".ini")
	if err != nil {
		panic(err)
	}
	return cfg
}
