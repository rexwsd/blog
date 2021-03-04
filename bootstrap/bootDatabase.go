package bootstrap

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gohouse/gorose/v2"
	"github.com/rexwsd/blog/config"
)

func BootDatabase() func(*Booter) {
	return func(booter *Booter) {
		//创建数据库链接
		connection, err := gorose.Open(config.DBConfig)
		if err != nil {
			panic(err)
		}
		booter.Connection = connection
	}
}
