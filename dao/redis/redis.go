package redis

import (
	"fmt"
	"lygf/backend/setting"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

// redis数据库连接


var rdb *redis.Client

// 初始化连接
func Init(cfg *setting.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
        PoolSize: cfg.PoolSize,
	})

    _, err = rdb.Ping().Result()
	if err != nil {
		zap.L().Error("连接redis失败", zap.Error(err))
		return
	}
    return err
}

func GetRedis() *redis.Client {
	return rdb
}

func Close() {
	err := rdb.Close()
	if err != nil {
		zap.L().Error("关闭redis连接失败", zap.Error(err))
	}
}