package sysModel

import (
	"gin-vue-admin/config"
	"gin-vue-admin/tools"
)

//配置文件结构体
type System struct {
	Config config.Config
}

//读取配置文件
func (s *System) GetSystemConfig() (err error, conf config.Config) {
	return nil, config.GinVueAdminconfig
}

//设置配置文件
func (s *System) SetSystemConfig() (err error) {
	confs := tools.StructToMap(s.Config)
	for k, v := range confs {
		config.VTool.Set(k, v)
	}
	err = config.VTool.WriteConfig()
	return err
}
