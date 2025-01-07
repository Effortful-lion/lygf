package router

import (
	"lygf/backend/logger"

	"github.com/gin-gonic/gin"
)

// 注册路由 ：负责为每个路由注册处理函数



func SetupRouter(mode string) *gin.Engine {
	// 设置gin框架的运行模式
	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// 1. 创建gin的空白实例
	r := gin.New()
	r.Use(logger.GinLogger(),logger.GinRecovery(true))	// 使用自定义的中间件
	// 2. 注册路由
	r.GET("/",func(ctx *gin.Context) {
		ctx.JSON(200,gin.H{
			"msg":"ok",
		})
	})
	
	// 返回实例
	return r
}