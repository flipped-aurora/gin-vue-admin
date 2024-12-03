// 自动生成模板CliMainorder
package xiao

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// 订单总表 结构体  CliMainorder
type CliMainorder struct {
	global.GVA_MODEL
	Address string           `json:"address" form:"address" gorm:"primarykey;column:address;comment:用户地址;size:100;"`       //用户地址
	Num     *int             `json:"num" form:"num" gorm:"default:0;column:num;comment:当前订单;"`                             //当前订单
	Amount  *decimal.Decimal `json:"amount" form:"amount" gorm:"type:decimal(14,4);default:0;column:amount;comment:订单总额;"` //订单总额
	Desc    string           `json:"desc" form:"desc" gorm:"default:备注;column:desc;comment:文本备注;"`                       //文本备注
	Status  string           `json:"status" form:"status" gorm:"default:0;column:status;comment:订单状态;"`
	Padian  *decimal.Decimal `json:"padian" form:"padian" gorm:"type:decimal(14,4);default:0;column:padian;comment:金额;"`
	Descnum *decimal.Decimal `json:"descnum" form:"descnum" gorm:"type:decimal(14,4);default:0;column:descnum;comment:金额备注;"` //金额备注
}

// TableName 订单总表 CliMainorder自定义表名 cli_mainorder
func (CliMainorder) TableName() string {
	return "cli_mainorder"
}

// NewCliMainorder 订单总表
func NewCliMainorder(address string) *CliMainorder {
	return &CliMainorder{Address: address}
}

// 查询订单总表方法
func (m *CliMainorder) GetCliMainorder(tx *gorm.DB) (info *CliMainorder, err error) {
	err = tx.Where("address = ?", m.Address).First(&info).Error
	if info != nil {
		return info, err
	}
	return info, err
}
