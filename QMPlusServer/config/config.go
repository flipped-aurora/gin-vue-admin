package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	MysqlAdmin   MysqlAdmin
	Qiniu        Qiniu
	CasbinConfig CasbinConfig
}
type CasbinConfig struct {
	ModelPath string // casbin model地址配置
}

type MysqlAdmin struct { // mysql admin 数据库配置
	Username string
	Password string
	Path     string
	Dbname   string
	Config   string
}

type Qiniu struct { // 七牛 密钥配置
	AccessKey string
	SecretKey string
}

var GinVueAdminconfig Config

func init() {
	v := viper.New()
	v.SetConfigName("config")           //  设置配置文件名 (不带后缀)
	v.AddConfigPath("./static/config/") // 第一个搜索路径
	v.SetConfigType("json")
	err := v.ReadInConfig() // 搜索路径，并读取配置数据
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	if err := v.Unmarshal(&GinVueAdminconfig); err != nil {
		fmt.Println(err)
	}
}
