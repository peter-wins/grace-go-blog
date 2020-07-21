package routers

import (
	"github.com/gin-gonic/gin"
)

type Option func(*gin.Engine)

var options = []Option{}

func Include(opts ...Option) {
	options = append(options, opts...)
}

func Init() *gin.Engine {
	r := gin.Default()
	//静态资源
	r.Static("/statics", "./statics")
	// 模板存储路径
	r.LoadHTMLGlob("views/***/**/*")
	for _, opt := range options {
		opt(r)
	}
	return r
}
