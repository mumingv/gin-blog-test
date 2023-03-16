package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Conf 实例化对象：对应yaml文件
var Conf = new(AppConfig)

// AppConfig App Config Struct
type AppConfig struct {
	Mode         string `mapstructure:"mode"`
	Port         int    `mapstructure:"port"`
	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
}

// MySQLConfig MySQL Config Struct
type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DB           string `mapstructure:"db"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

// LogConfig Log Config
type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

// Init 定义一个初始化函数
// 1. 读取config.yaml
// 2. 文件变更的时候，可以自动识别（main.go，热加载自动识别）
func Init() error {
	// 目的：读取yaml文件 -> Conf
	viper.SetConfigFile("./conf/config.yaml")
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		viper.Unmarshal(&Conf)
	})
	viper.WatchConfig()

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("ReadInConfig failed, err: %v", err))
	}

	err = viper.Unmarshal(&Conf)
	return err
}
