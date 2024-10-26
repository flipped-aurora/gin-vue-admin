// 自动生成模板CliMainwith
package xiao

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/shopspring/decimal"
)

// 提币总表 结构体  CliMainwith
type CliMainwith struct {
	global.GVA_MODEL
	Address  string           `json:"address" form:"address" gorm:"column:address;comment:用户地址;size:100;"`                        //用户地址
	Num      *int             `json:"num" form:"num" gorm:"default:0;column:num;comment:提币次数;"`                                   //提币次数
	Withable *decimal.Decimal `json:"withable" form:"withable" gorm:"type:decimal(14,4);default:0;column:withable;comment:可提金额;"` //可提金额
	Withed   *decimal.Decimal `json:"withed" form:"withed" gorm:"type:decimal(14,4);default:0;column:withed;comment:已提金额;"`       //已提金额
	Total    *decimal.Decimal `json:"total" form:"total" gorm:"type:decimal(14,4);default:0;column:total;comment:累计总额;"`          //累计总额
	Desc     string           `json:"desc" form:"desc" gorm:"default:备注;column:desc;comment:文本备注;"`                               //文本备注
	Descnum  *decimal.Decimal `json:"descnum" form:"descnum" gorm:"type:decimal(14,4);default:0;column:descnum;comment:金额备注;"`    //金额备注
}

// TableName 提币总表 CliMainwith自定义表名 cli_mainwith
func (CliMainwith) TableName() string {
	return "cli_mainwith"
}

// NewCliMainwith 创建CliMainwith
func NewCliMainwith(address string) *CliMainwith {
	return &CliMainwith{Address: address}

}
