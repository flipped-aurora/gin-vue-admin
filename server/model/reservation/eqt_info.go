// 自动生成模板EqtInfo
package reservation

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// EqtInfo 结构体
type EqtInfo struct {
	global.GVA_MODEL
	EqtClass       *int       `json:"eqtClass" form:"eqtClass" gorm:"column:eqt_class;comment:设备类型;size:19;"`
	EqtCardNo      string     `json:"eqtCardNo" form:"eqtCardNo" gorm:"column:eqt_card_no;comment:设备卡片码;size:191;"`
	EqtSn          string     `json:"eqtSn" form:"eqtSn" gorm:"column:eqt_sn;comment:SN号;size:191;"`
	EqtName        string     `json:"eqtName" form:"eqtName" gorm:"column:eqt_name;comment:设备名;size:191;"`
	EqtModel       string     `json:"eqtModel" form:"eqtModel" gorm:"column:eqt_model;comment:设备型号;size:191;"`
	EqtStatus      *int       `json:"eqtStatus" form:"eqtStatus" gorm:"column:eqt_status;comment:设备状态，0开放预约，1关闭预约;"`
	EqtStockStatus *int       `json:"eqtStockStatus" form:"eqtStockStatus" gorm:"column:eqt_stock_status;comment:设备库存状态，0是在库，1出库;size:19;"`
	CurrentDept    *int       `json:"currentDept" form:"currentDept" gorm:"column:current_dept;comment:当前科室;size:19;"`
	ServiceTime    *time.Time `json:"serviceTime" form:"serviceTime" gorm:"column:service_time;comment:;"`
	CreatedBy      uint       `gorm:"column:created_by;comment:创建者"`
	UpdatedBy      uint       `gorm:"column:updated_by;comment:更新者"`
	DeletedBy      uint       `gorm:"column:deleted_by;comment:删除者"`
}

// TableName EqtInfo 表名
func (EqtInfo) TableName() string {
	return "eqt_info"
}
