package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Create
//@description: 创建工作流
//@param: wk model.SysWorkflow
//@return: error

func Create(wk model.SysWorkflow) error {
	err := global.GVA_DB.Create(&wk).Error
	return err
}
