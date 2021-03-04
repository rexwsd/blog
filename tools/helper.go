package tools

import (
	"github.com/go-ini/ini"
	"github.com/rexwsd/blog/config"
	"os"
)

/*
	获取项目所在路径
*/
func ProjectPath() string {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return path
}
