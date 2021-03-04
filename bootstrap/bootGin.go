package bootstrap

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	def_config "github.com/rexwsd/blog/config"
	"github.com/rexwsd/blog/middleware"
)

func BootGin() func(*Booter) {
	return func(booter *Booter) {
		gin.SetMode(def_config.RunMode)
		booter.Router = gin.Default()
		config := cors.DefaultConfig()
		config.AllowAllOrigins = def_config.AllowAllOrigins
		// 配置允许 OPTIONS 请求, 默认是没有的
		config.AddAllowMethods("OPTIONS")
		config.AddAllowHeaders("Authorization")
		// 配置静态目录
		booter.Router.Static("static", "静态文件路径")
		booter.Router.Use(cors.New(config)).Use(middleware.Cros())
	}
}
