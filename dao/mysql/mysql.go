package mysql

import (
	"fmt"
	"lygf/backend/setting"

	"go.uber.org/zap"
	"gorm.io/driver/mysql" // gorm内置的mysql驱动
	"gorm.io/gorm"
)

// 这里是关于mysql的读取配置，初始化的方法等


var db *gorm.DB // 这里是一个全局变量，用来存储mysql的连接实例

func Init(cfg *setting.MysqlConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		zap.L().Error("连接mysql失败", zap.Error(err))
		return
	}
	// 连接成功后，设置其他配置：
	sqlDB, err := db.DB()
	if err != nil {
		zap.L().Error("db转换为sqlDB失败", zap.Error(err))
		return
	}
	// 设置最大连接数
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	// 设置最大空闲连接数
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	return
}