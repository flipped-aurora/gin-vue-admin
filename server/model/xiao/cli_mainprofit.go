// 自动生成模板CliMainprofit
package xiao

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/shopspring/decimal"
)

// 结算总表 结构体  CliMainprofit
type CliMainprofit struct {
	global.GVA_MODEL
	Address  string           `json:"address" form:"address" gorm:"primarykey;column:address;comment:用户地址;size:100;"`             //用户地址
	Num      *int             `json:"num" form:"num" gorm:"default:0;column:num;comment:结算次数;"`                                   //结算次数
	Static   *decimal.Decimal `json:"static" form:"static" gorm:"type:decimal(14,4);default:0;column:static;comment:静态收益;"`       //静态收益
	Pull     *decimal.Decimal `json:"pull" form:"pull" gorm:"type:decimal(14,4);default:0;column:pull;comment:直推收益;"`             //直推收益
	Indirect *decimal.Decimal `json:"indirect" form:"indirect" gorm:"type:decimal(14,4);default:0;column:indirect;comment:间接收益;"` //间接收益
	Team     *decimal.Decimal `json:"team" form:"team" gorm:"type:decimal(14,4);default:0;column:team;comment:团队收益;"`             //团队收益
	Amount   *decimal.Decimal `json:"amount" form:"amount" gorm:"type:decimal(14,4);default:0;column:amount;comment:累积金额;"`       //累积金额
	Desc     string           `json:"desc" form:"desc" gorm:"column:desc;comment:文本备注;"`                                          //文本备注
	Descnum  *decimal.Decimal `json:"descnum" form:"descnum" gorm:"type:decimal(14,4);default:0;column:descnum;comment:金额备注;"`    //金额备注
}

// TableName 结算总表 CliMainprofit自定义表名 cli_mainprofit
func (CliMainprofit) TableName() string {
	return "cli_mainprofit"
}

// NewCliMainprofit 实例化CliMainprofit
func NewCliMainprofit(address string) *CliMainprofit {
	return &CliMainprofit{Address: address}
}
