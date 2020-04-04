package model

import (
	"gin-vue-admin/global"
	"gin-vue-admin/init"
	"gin-vue-admin/utils"
)

//配置文件结构体
type System struct {
	Config init.Config
}

//读取配置文件
func (s *System) GetSystemConfig() (err error, conf init.Config) {
	return nil, global.GVA_CONFIG
}

//设置配置文件
func (s *System) SetSystemConfig() (err error) {
	confs := utils.StructToMap(s.Config)
	for k, v := range confs {
		global.GVA_VP.Set(k, v)
	}
	err = global.GVA_VP.WriteConfig()
	return err
}
