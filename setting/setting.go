package setting

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// 使用viper读取配置文件，并监控配置文件变化
// 作用：
// 1. 定义全局配置变量
// 2. 读取配置文件到配置变量
// 3. 监控配置文件变化

var Conf = new(Config)

type Config struct {
	*AppConfig       `mapstructure:"app"`
	*LogConfig       `mapstructure:"log"`
	*MysqlConfig     `mapstructure:"mysql"`
	*RedisConfig     `mapstructure:"redis"`
}

type AppConfig struct {
	Name string `mapstructure:"name"`	// 项目名称
	Mode string `mapstructure:"mode"`	// 运行模式
	Version string `mapstructure:"version"` 	// 版本
	Port int `mapstructure:"port"`	// 运行端口
	StartTime string `mapstructure:"start_time"`	// 启动时间
	MachineID int `mapstructure:"machine_id"`	// 机器ID
	*LogConfig `mapstructure:"log"`	// 日志配置
	*MysqlConfig `mapstructure:"mysql"`	// mysql配置
	*RedisConfig `mapstructure:"redis"`	// redis配置
}

type LogConfig struct {
	Level string `mapstructure:"level"`
	Filename string `mapstructure:"filename"`
	MaxSize int `mapstructure:"max_size"`
	MaxAge int `mapstructure:"max_age"`
	MaxBackups int `mapstructure:"max_backups"`
}

type MysqlConfig struct {
    Host string `mapstructure:"host"`
    User string `mapstructure:"user"`
    Password string `mapstructure:"password"`
    DB string `mapstructure:"dbname"`
    Port int `mapstructure:"port"`
    MaxOpenConns int `mapstructure:"max_open_conns"`
    MaxIdleConns int `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port string `mapstructure:"port"`
	DB int `mapstructure:"db"`
	PoolSize int `mapstructure:"pool_size"`
}

func Init()(err error){
	
	//指定配置文件
	// viper.SetConfigName("config")
	// viper.SetConfigType("yaml")
	// viper.AddConfigPath("./conf/")	// 这里的路径指向的是程序的运行目录，比如：D:\GolandCode\src\lygf\backend
	viper.SetConfigFile("./conf/config.yaml")
	err = viper.ReadInConfig()		//从配置文件中读取配置项
	if err != nil{
		// 文件未找到，读取失败
		zap.L().Error("读取配置文件失败", zap.Error(err))
		return
	}
	// 反序列化到配置变量的结构体中
	if err := viper.Unmarshal(Conf);err != nil {
		zap.L().Error("viper.Unmarshal failed", zap.Error(err))
		return err
	}

	// 其实air已经可以监控了

	// // 监控配置文件变化 并 对变化做处理
	// viper.WatchConfig()	
	// viper.OnConfigChange(func(in fsnotify.Event) {
	// 	time.Sleep(1 * time.Second) // 延迟1秒，可根据实际情况调整
	// 	if err := viper.Unmarshal(Conf); err!= nil {
	// 		zap.L().Error("viper.Unmarshal failed", zap.Error(err))
	// 		// 可以在这里添加更完善的错误处理逻辑，比如尝试重新读取等操作
	// 	}
		
	// 	fmt.Println("配置文件被修改")
	// })


	return
}
