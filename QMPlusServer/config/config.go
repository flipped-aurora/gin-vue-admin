package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Admin Admin
}
type Admin struct {
	UserName string
	Password string
	Path     string
	Dbname   string
	Config   string
}

var Dbconfig Config

func init() {
	v := viper.New()
	v.SetConfigName("config")             //  设置配置文件名 (不带后缀)
	v.AddConfigPath("./config/dbconfig/") // 第一个搜索路径
	v.SetConfigType("json")
	err := v.ReadInConfig() // 搜索路径，并读取配置数据
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	if err := v.Unmarshal(&Dbconfig); err != nil {
		fmt.Println(err)
	}
}
