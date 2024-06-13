package temp

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// cowsHardStatus表 结构体  CowsHardStatus
type Temp struct {
	global.GVA_MODEL
	SerialNo string `json:"serialNo" form:"serialNo" gorm:"column:serial_no;comment:cows_hard_aiboard.serial_no;size:32;"` //cows_hard_aiboard.serial_no
}

// TableName cowsHardStatus表 CowsHardStatus自定义表名 cows_hard_status
func (Temp) TableName() string {
	return "temp"
}
