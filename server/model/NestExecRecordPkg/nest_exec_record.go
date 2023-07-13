// 自动生成模板NestExecRecord
package NestExecRecordPkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// NestExecRecord 结构体
type NestExecRecord struct {
	global.GVA_MODEL
	Status       string     `json:"status" form:"status" gorm:"column:status;comment:;"`
	Creator      string     `json:"creator" form:"creator" gorm:"column:creator;comment:;"`
	CopyTo       string     `json:"copyTo" form:"copyTo" gorm:"column:copy_to;comment:;"`
	MissionName  string     `json:"missionName" form:"missionName" gorm:"column:mission_name;comment:;"`
	Missionid    string     `json:"missionid" form:"missionid" gorm:"column:missionid;comment:;"`
	PlanAt       *time.Time `json:"planAt" form:"planAt" gorm:"column:plan_at;comment:;"`
	Type         string     `json:"type" form:"type" gorm:"column:type;comment:;"`
	ExecuteId    string     `json:"executeId" form:"executeId" gorm:"column:execute_id;comment:;"`
	ExecuteAt    string     `json:"executeAt" form:"executeAt" gorm:"column:execute_at;comment:;"`
	FlyInSecond  *int       `json:"flyInSecond" form:"flyInSecond" gorm:"column:fly_in_second;comment:;"`
	PanoramaLink string     `json:"panoramaLink" form:"panoramaLink" gorm:"column:panorama_link;comment:;"`
	CreatedBy    uint       `gorm:"column:created_by;comment:创建者"`
	UpdatedBy    uint       `gorm:"column:updated_by;comment:更新者"`
	DeletedBy    uint       `gorm:"column:deleted_by;comment:删除者"`
}

// TableName NestExecRecord 表名
func (NestExecRecord) TableName() string {
	return "nest_exec_record"
}
