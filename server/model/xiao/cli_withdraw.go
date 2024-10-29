// 自动生成模板CliWithdraw
package xiao

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// 提币详情 结构体  CliWithdraw
type CliWithdraw struct {
	global.GVA_MODEL
	Address string           `json:"address" form:"address" gorm:"column:address;comment:用户地址;size:100;"`                     //用户地址
	Amount  *decimal.Decimal `json:"amount" form:"amount" gorm:"type:decimal(14,4);default:0;column:amount;comment:提币金额;"`    //提币金额
	Text    string           `json:"text" form:"text" gorm:"column:text;comment:提币说明;size:10;"`                               //提币说明
	Status  string           `json:"status" form:"status" gorm:"default:自动通过;column:status;comment:审核状态;size:10;"`            //审核状态
	Desc    string           `json:"desc" form:"desc" gorm:"default:备注;column:desc;comment:文本备注;"`                            //文本备注
	Descnum *decimal.Decimal `json:"descnum" form:"descnum" gorm:"type:decimal(14,4);default:0;column:descnum;comment:金额备注;"` //金额备注
}

// TableName 提币详情 CliWithdraw自定义表名 cli_withdraw
func (CliWithdraw) TableName() string {
	return "cli_withdraw"
}

// NewCliWithdraw 创建CliWithdraw
func NewCliWithdraw(address string) *CliWithdraw {
	return &CliWithdraw{
		Address: address,
	}
}

// 查询提币记录
func (cliwith *CliWithdraw) GetCliWithdraw(tx *gorm.DB) ([]*CliWithdraw, error) {
	var cliwithdraw []*CliWithdraw
	err := tx.Model(&CliWithdraw{}).Where("address = ?", cliwith.Address).Find(&cliwithdraw).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return cliwithdraw, nil
}
