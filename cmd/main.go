package main

import (
	"fmt"
	"lygf/backend/dao/mysql"
	"lygf/backend/dao/redis"
	"lygf/backend/logger"
	"lygf/backend/router"
	"lygf/backend/setting"

	"go.uber.org/zap"
)

// 启动服务，项目的入口文件
func main() {

	// 读取配置文件
	err := setting.Init()
	if err != nil {
		panic(err)
	}

	// 初始化日志
	err = logger.Init(setting.Conf.LogConfig, setting.Conf.Mode)
	if err != nil {
		panic(err)
	}

	// 初始化mysql
	err = mysql.Init(setting.Conf.MysqlConfig)
	if err != nil {
		panic(err)
	}

	// 初始化redis
	err = redis.Init(setting.Conf.RedisConfig)
	if err != nil {
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