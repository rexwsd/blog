package config

import (
	"fmt"
	"github.com/gohouse/gorose/v2"
)

var DBConfig = &gorose.Config{}

func init() {
	cfg := getConfig("database")
	defaultConfig, err := cfg.GetSection("mysql.default")
	if err != nil {
		panic(err)
	}
	DBConfig.Driver = defaultConfig.Key("driver").MustString("mysql")
	DBConfig.Prefix = defaultConfig.Key("prefix").MustString("")
	DBConfig.Dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		defaultConfig.Key("user").MustString(""),
		defaultConfig.Key("password").MustString(""),
		defaultConfig.Key("host").MustString(""),
		defaultConfig.Key("port").MustString(""),
		defaultConfig.Key("database").MustString(""),
	)
}
