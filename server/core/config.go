package core

import (
	"fmt"
	"gin-vue-admin/global"
	_ "gin-vue-admin/packfile"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/spf13/pflag"
)

const defaultConfigFile = "config.yaml"

func init() {
	pflag.StringP("configFile","c", "", "choose config file.")
	pflag.Parse()

	// 优先级: 命令行 > 环境变量 > 默认值
	v := viper.New()
	v.BindPFlags(pflag.CommandLine)
	v.SetEnvPrefix("gva")
	v.BindEnv("configFile") // GVA_CONFIGFILE

	configFile := v.GetString("configFile")
	if configFile == ""{
		configFile = defaultConfigFile
	}

	v.SetConfigFile(configFile)
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
