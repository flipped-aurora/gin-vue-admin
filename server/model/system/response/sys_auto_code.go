package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/system"

type Db struct {
	Database string `json:"database" gorm:"column:database"`
}

type Table struct {
	TableName string `json:"tableName" gorm:"column:table_name"`
}

type Column struct {
	DataType      string `json:"dataType" gorm:"column:data_type"`
	ColumnName    string `json:"columnName" gorm:"column:column_name"`
	DataTypeLong  string `json:"dataTypeLong" gorm:"column:data_type_long"`
	ColumnComment string `json:"columnComment" gorm:"column:column_comment"`
	PrimaryKey    bool   `json:"primaryKey" gorm:"column:primary_key"`
}

type PluginInfo struct {
	PluginName   string                 `json:"pluginName"`
	PluginType   string                 `json:"pluginType"` // web, server, full
	Apis         []system.SysApi        `json:"apis"`
	Menus        []system.SysBaseMenu   `json:"menus"`
	Dictionaries []system.SysDictionary `json:"dictionaries"`
}
