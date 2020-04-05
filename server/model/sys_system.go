package model

import (
	"gin-vue-admin/config"
	"gin-vue-admin/global"
	"gin-vue-admin/utils"
)

//配置文件结构体
type System struct {
	Config config.Server
}

//读取配置文件
func (s *System) GetSystemConfig() (err error, conf config.Server) {
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
