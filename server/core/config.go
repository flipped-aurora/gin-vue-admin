package core

import (
	"fmt"
	"gin-vue-admin/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func init() {
	v := viper.New()
	v.SetConfigName("config") //  设置配置文件名 (不带后缀)
	v.AddConfigPath("./")     // 第一个搜索路径
	v.SetConfigType("json")
	err := v.ReadInConfig() // 搜索路径，并读取配置数据
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.GVA_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.GVA_CONFIG); err != nil {
		fmt.Println(err)
	}
	global.GVA_VP = v
}
