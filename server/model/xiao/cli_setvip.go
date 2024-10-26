// 自动生成模板CliSetvip
package xiao

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/shopspring/decimal"
)

// 团队设置 结构体  CliSetvip
type CliSetvip struct {
	global.GVA_MODEL
	Levels  string           `json:"levels" form:"levels" gorm:"column:levels;comment:Vip级别;size:10;"`                        //Vip级别
	Team    string           `json:"team" form:"team" gorm:"column:team;comment:团队要求;size:20;"`                               //团队要求
	Come    *decimal.Decimal `json:"come" form:"come" gorm:"type:decimal(14,4);column:come;comment:业绩要求;"`                    //业绩要求
	Rate    *decimal.Decimal `json:"rate" form:"rate" gorm:"type:decimal(14,4);column:rate;comment:收益比率;"`                    //收益比率
	Desc    string           `json:"desc" form:"desc" gorm:"column:desc;comment:文本备注;"`                                       //文本备注
	Descnum *decimal.Decimal `json:"descnum" form:"descnum" gorm:"type:decimal(14,4);default:0;column:descnum;comment:金额备注;"` //金额备注
}

// TableName 团队设置 CliSetvip自定义表名 cli_setvip
func (CliSetvip) TableName() string {
	return "cli_setvip"
}
