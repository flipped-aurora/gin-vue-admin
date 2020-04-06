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

// @title    GetSystemConfig
// @description   读取配置文件
// @auth                     （2020/04/05  20:22 ）
// @return    err             error
// @return    conf            Server
func (s *System) GetSystemConfig() (err error, conf config.Server) {
	return nil, global.GVA_CONFIG
}

// @title    SetSystemConfig
// @description   set system config, 设置配置文件
// @auth                    （2020/04/05  20:22 ）
// @return    err            error
func (s *System) SetSystemConfig() (err error) {
	confs := utils.StructToMap(s.Config)
	for k, v := range confs {
		global.GVA_VP.Set(k, v)
	}
	err = global.GVA_VP.WriteConfig()
	return err
}
