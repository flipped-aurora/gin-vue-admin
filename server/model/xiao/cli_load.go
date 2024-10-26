// 自动生成模板CliLoad
package xiao

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

// cliLoad表 结构体  CliLoad
type CliLoad struct {
	global.GVA_MODEL
	Address   string           `json:"address" form:"address" gorm:"primarykey;column:address;comment:用户地址;size:100;"`        //用户地址
	Loadtimes *int             `json:"loadtimes" form:"loadtimes" gorm:"default:1;column:loadtimes;comment:登录次数;size:16;"`    //登录次数
	Lasttime  *time.Time       `json:"lasttime" form:"lasttime" gorm:"column:lasttime;comment:最近登录;"`                         //最近登录
	Loadip    string           `json:"loadip" form:"loadip" gorm:"column:loadip;comment:登录IP;size:100;"`                      //登录IP
	Loadaddr  string           `json:"loadaddr" form:"loadaddr" gorm:"column:loadaddr;comment:登录地址;size:200;"`                //登录地址
	Usdt      *decimal.Decimal `json:"usdt" form:"usdt" gorm:"type:decimal(14,4);default:0;column:usdt;comment:地址余额;size:0;"` //地址余额
	Status    string           `json:"status" form:"status" gorm:"default:正常;column:status;comment:当前状态;size:10;"`            //当前状态
	Desc      string           `json:"desc" form:"desc" gorm:"default:备注;column:desc;comment:文本备注;size:200;"`                 //文本备注
	Desnum    *decimal.Decimal `json:"desnum" form:"desnum" gorm:"type:decimal(14,4);default:0;column:desnum;comment:金额备注;"`  //金额备注
}

// TableName cliLoad表 CliLoad自定义表名 cli_load
func (CliLoad) TableName() string {
	return "cli_load"
}

var now = time.Now()

// CheckAddress 查询地址是否存在
func (cliLoad *CliLoad) CheckAddress(tx *gorm.DB) (*CliLoad, error) {
	// 查询地址是否存在
	var cliloadinfo CliLoad
	err := tx.Where("address = ?", cliLoad.Address).First(&cliloadinfo).Error

	return &cliloadinfo, err
}
func Newload(address string) *CliLoad {
	return &CliLoad{
		Address:  address,
		Lasttime: &now,
		Desc:     "未注册",
	}
}
