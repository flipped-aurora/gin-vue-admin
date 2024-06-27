// 自动生成模板SysExportTemplate
package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 导出模板 结构体  SysExportTemplate
type SysExportTemplate struct {
	global.GVA_MODEL
	DBName       string         `json:"dbName" form:"dbName" gorm:"column:db_name;comment:数据库名称;"`               //数据库名称
	Name         string         `json:"name" form:"name" gorm:"column:name;comment:模板名称;"`                       //模板名称
	TableName    string         `json:"tableName" form:"tableName" gorm:"column:table_name;comment:表名称;"`        //表名称
	TemplateID   string         `json:"templateID" form:"templateID" gorm:"column:template_id;comment:模板标识;"`    //模板标识
	TemplateInfo string         `json:"templateInfo" form:"templateInfo" gorm:"column:template_info;type:text;"` //模板信息
	Limit        *int           `json:"limit" form:"limit" gorm:"column:limit;comment:导出限制"`
	Order        string         `json:"order" form:"order" gorm:"column:order;comment:排序"`
	Conditions   []Condition    `json:"conditions" form:"conditions" gorm:"foreignKey:TemplateID;references:TemplateID;comment:条件"`
	JoinTemplate []JoinTemplate `json:"joinTemplate" form:"joinTemplate" gorm:"foreignKey:TemplateID;references:TemplateID;comment:关联"`
}

type JoinTemplate struct {
	global.GVA_MODEL
	TemplateID string `json:"templateID" form:"templateID" gorm:"column:template_id;comment:模板标识"`
	JOINS      string `json:"joins" form:"joins" gorm:"column:joins;comment:关联"`
	Table      string `json:"table" form:"table" gorm:"column:table;comment:关联表"`
	ON         string `json:"on" form:"on" gorm:"column:on;comment:关联条件"`
}

func (JoinTemplate) TableName() string {
	return "sys_export_template_join"
}

type Condition struct {
	global.GVA_MODEL
	TemplateID string `json:"templateID" form:"templateID" gorm:"column:template_id;comment:模板标识"`
	From       string `json:"from" form:"from" gorm:"column:from;comment:条件取的key"`
	Column     string `json:"column" form:"column" gorm:"column:column;comment:作为查询条件的字段"`
	Operator   string `json:"operator" form:"operator" gorm:"column:operator;comment:操作符"`
}

func (Condition) TableName() string {
	return "sys_export_template_condition"
}
