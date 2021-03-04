package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/gohouse/gorose/v2"
	"sync"
)

type Booter struct {
	//web引擎
	Router *gin.Engine
	//数据库链接引擎
	Connection *gorose.Engin
}

var b *Booter
var once sync.Once

//单例
func BootInstance() *Booter {
	once.Do(func() {
		b = &Booter{}
	})
	return b
}

//Use 驱动中间件
func (b *Booter) Use(options ...func(*Booter)) {
	for _, option := range options {
		option(b)
	}
}
