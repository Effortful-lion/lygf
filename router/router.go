package router

import (
	"lygf/backend/controller"
	"lygf/backend/logger"
	"lygf/backend/middleware"

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

	// 测试路由
	r.GET("/",func(ctx *gin.Context) {
		ctx.JSON(200,gin.H{
			"msg":"ok",
		})
	})

	v1 := r.Group("/api/v1")

	// 用户登录注册管理
	{
		// 发送验证码
		v1.POST("/auth/code",controller.SendEmailCode)
		// 用户注册
		v1.POST("/auth/register",controller.UserRegister)
		// 用户登录
		v1.POST("/auth/login",controller.UserLogin)
		// 提供 前后端 系统时间校准（使得token中的过期时间和系统时间可比较）
		v1.GET("/time",controller.GetTime)
	}


	// 使用 鉴权中间件：对 非公开资源/操作请求需要登录（查看商品、价格等可以；修改地址、加入购物车等操作需要身份验证）
	v1.Use(middleware.JWTAuthMiddleware())

	// 用户信息管理
	{
		// 获取用户信息
		v1.GET("/users",controller.GetUserInfo)
		// 获取用户信息（用于回显编辑信息）
		v1.GET("/users/edit",controller.GetUserInfoForEdit)
		// 更新用户信息
		v1.PUT("/users",controller.UpdateUserInfo)
		// 用户注销
		v1.DELETE("/users",controller.UserDelete)
		// 用户退出
		v1.POST("/auth/logout",controller.UserLogout)
		
	}

	// 首页接口
	{
		// 获取分类列表: 其实前端应该可以写死，就不需要了
		v1.GET("home/categories",controller.GetCategories)
		// 根据两个分类条件获取所有商铺列表
		v1.GET("home/shop/:cateID1/:cateId2",controller.GetShops)
	}


	
	// 返回实例
	return r
}