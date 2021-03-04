package main

import (
	"github.com/rexwsd/blog/bootstrap"
	"github.com/rexwsd/blog/routers"
)

func main() {
	// 驱动基本服务
	booter := bootstrap.BootInstance()
	// 驱动数据库
	booter.Use(bootstrap.BootDatabase())
	// 驱动gin框架
	booter.Use(bootstrap.BootGin())

	//启动web服务
	routers.Run()
}
