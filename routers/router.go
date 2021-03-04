package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/rexwsd/blog/api/v1"
	"github.com/rexwsd/blog/bootstrap"
	"github.com/rexwsd/blog/config"
)

//路由初始化
func Run() {
	r := bootstrap.BootInstance().Router
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/api/v1")
	{
		//获取用户
		apiv1.GET("/users", v1.GetUser)
		//添加用户
		apiv1.POST("/users", v1.AddUser)
		//编辑用户
		apiv1.PUT("/users/:id", v1.EditUser)
		//删除用户
		apiv1.DELETE("/user/:id", v1.DelUser)
	}
	err := r.Run(config.Host + ":" + config.Port)
	if err != nil {
		panic(err)
	}
}
