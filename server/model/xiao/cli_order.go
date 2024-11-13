// 自动生成模板CliOrder
package xiao

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// 订单详情 结构体  CliOrder
type CliOrder struct {
	global.GVA_MODEL
	Address  string           `json:"address" form:"address" gorm:"primarykey;column:address;comment:用户地址;size:100;"`          //用户地址
	Amount   *decimal.Decimal `json:"amount" form:"amount" gorm:"type:decimal(14,4);default:0;column:amount;comment:订单金额;"`    //订单金额
	Num      *int             `json:"num" form:"num" gorm:"default:0;column:num;comment:结算次数;"`                                //结算次数
	Text     string           `json:"text" form:"text" gorm:"column:text;comment:订单说明;size:10;"`                               //订单说明
	Status   string           `json:"status" form:"status" gorm:"default:正常;column:status;comment:订单状态;size:10;"`              //订单状态
	Desc     string           `json:"desc" form:"desc" gorm:"default:备注;column:desc;comment:文本备注;"`                            //文本备注
	Descnum  *decimal.Decimal `json:"descnum" form:"descnum" gorm:"type:decimal(14,4);default:0;column:descnum;comment:金额备注;"` //累计结算
	Todaynum *decimal.Decimal `json:"todaynum" form:"todaynum" gorm:"type:decimal(14,4);default:0;column:todaynum;comment:今日结算;"`
}

// TableName 订单详情 CliOrder自定义表名 cli_orders
func (CliOrder) TableName() string {
	return "cli_orders"
}

// NewCliOrder
func NewOrder(address string) *CliOrder {
	return &CliOrder{
		Address: address,
	}
}
func (cliorder *CliOrder) GetCliOrder(tx *gorm.DB) (info *CliOrder, err error) {
	err = tx.Model(cliorder).Where("address = ?", cliorder.Address).First(&info).Error
	return
}

func (cliorder *CliOrder) GetCliAllOrder(tx *gorm.DB) (info []*CliOrder, err error) {
	err = tx.Model(cliorder).Where("address = ?", cliorder.Address).Find(&info).Error
	return
}
