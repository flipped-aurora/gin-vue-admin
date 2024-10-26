// 自动生成模板CliProfit
package xiao

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/shopspring/decimal"
)

// 结算详情 结构体  CliProfit
type CliProfit struct {
	global.GVA_MODEL
	Address string           `json:"address" form:"address" gorm:"primarykey;column:address;comment:用户地址;size:100;"`          //用户地址
	Amount  *decimal.Decimal `json:"amount" form:"amount" gorm:"type:decimal(14,4);default:0;column:amount;comment:结算金额;"`    //结算金额
	Text    string           `json:"text" form:"text" gorm:"column:text;comment:结算说明;size:10;"`                               //结算说明
	Status  string           `json:"status" form:"status" gorm:"default:正常;column:status;comment:状态;"`                        //当前状态
	Desc    string           `json:"desc" form:"desc" gorm:"default:备注;column:desc;comment:文本备注;"`                            //文本备注
	Descnum *decimal.Decimal `json:"descnum" form:"descnum" gorm:"type:decimal(14,4);default:0;column:descnum;comment:金额备注;"` //金额备注
}

// TableName 结算详情 CliProfit自定义表名 cli_profit
func (CliProfit) TableName() string {
	return "cli_profit"
}
