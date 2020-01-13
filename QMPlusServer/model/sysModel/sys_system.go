package sysModel

import (
	"gin-vue-admin/config"
	"gin-vue-admin/tools"
)

type System struct {
	Config config.Config
}

func (s *System)GetSystemConfig()(err error,conf config.Config){
	return nil,config.GinVueAdminconfig
}

func (s *System)SetSystemConfig()(err error){
	confs:= tools.StructToMap(s.Config)
	for k,v:= range confs {
		config.VTool.Set(k,v)
	}
	err = config.VTool.WriteConfig()
	return err
}