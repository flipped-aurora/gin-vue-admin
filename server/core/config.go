package core

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/utils"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const defaultConfigFile = "config.yaml"

//开发环境配置
const devConfigFile = "config-dev.yaml"

func init() {
	v := viper.New()
	if ok, _ := utils.PathExists(devConfigFile); ok {
		fmt.Printf("use dev config %v \n", devConfigFile)
		v.SetConfigFile(devConfigFile)
	} else {
		v.SetConfigFile(defaultConfigFile)
	}
	err := v.ReadInConfig()
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
