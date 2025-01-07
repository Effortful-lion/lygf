package main

import (
	"fmt"
	"lygf/backend/dao/mysql"
	"lygf/backend/dao/redis"
	"lygf/backend/logger"
	//model "lygf/backend/model/mysql"
	"lygf/backend/router"
	"lygf/backend/setting"

	"go.uber.org/zap"
)

// 启动服务，项目的入口文件
func main() {

	// 读取配置文件
	err := setting.Init()
	if err != nil {
		zap.L().Error("配置文件读取失败", zap.Error(err))
		panic(err)
	}

	// 初始化日志
	err = logger.Init(setting.Conf.LogConfig, setting.Conf.Mode)
	if err != nil {
		zap.L().Error("日志初始化失败", zap.Error(err))
		panic(err)
	}
	//通常用于将日志缓冲区中的内容（如果有的话）刷新到对应的输出目标（比如文件、标准输出等）
	//避免出现日志丢失等情况，保证日志记录的完整性
	defer zap.L().Sync()

	// 初始化mysql
	err = mysql.Init(setting.Conf.MysqlConfig)
	if err != nil {
		zap.L().Error("mysql init fail")
		panic(err)
	}

	// 初始化mysql表（只在需要时初始化一次就可以了）
	//model.Init()

	// 初始化redis
	err = redis.Init(setting.Conf.RedisConfig)
	if err != nil {
		zap.L().Error("redis init fail")
		panic(err)
	}

	// 注册路由
	r := router.SetupRouter(setting.Conf.Mode)

	// 启动服务
	if err := r.Run(fmt.Sprintf("127.0.0.1:%d",setting.Conf.AppConfig.Port)); err != nil{
		zap.L().Error("启动失败",zap.String("msg",err.Error()))
		return
	}

}