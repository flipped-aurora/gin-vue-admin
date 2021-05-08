// 自动生成模板ProductPlugin
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type ProductPlugin struct {
      global.GVA_MODEL
      ProductID  string `json:"productID" form:"productID" gorm:"column:product_id;type:varchar(20);"`
      ProductName  string `json:"productName" form:"productName" gorm:"column:product_name;type:varchar(256);"`
      PluginPath  string `json:"pluginPath" form:"pluginPath" gorm:"column:plugin_path;type:varchar(512);"`
}


