// 自动生成模板FlyResult
package FlyResultPkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/NestExecRecordPkg"
	"time"
)

// FlyResult 结构体
type FlyResult struct {
	global.GVA_MODEL
	ExecuteId      string                           `json:"executeId" form:"executeId" gorm:"column:execute_id;comment:;unique_index:E_F"`
	FileName       string                           `json:"fileName" form:"fileName" gorm:"column:file_name;comment:;unique_index:E_F"`
	FileOrder      *int                             `json:"fileOrder" form:"fileOrder" gorm:"column:file_order;comment:;"`
	Type           *int                             `json:"type" form:"type" gorm:"column:type;comment:;"`
	Location       string                           `json:"location" form:"location" gorm:"column:location;comment:;"`
	Timestamp      *time.Time                       `json:"timestamp" form:"timestamp" gorm:"column:timestamp;comment:;"`
	Downloaded     *int                             `json:"downloaded" form:"downloaded" gorm:"column:downloaded;comment:;"`
	CreatedBy      uint                             `gorm:"column:created_by;comment:创建者"`
	UpdatedBy      uint                             `gorm:"column:updated_by;comment:更新者"`
	DeletedBy      uint                             `gorm:"column:deleted_by;comment:删除者"`
	NestExecRecord NestExecRecordPkg.NestExecRecord `gorm:"foreignKey:ExecuteId;references:execute_id""`
}

// TableName FlyResult 表名
func (FlyResult) TableName() string {
	return "fly_result"
}
