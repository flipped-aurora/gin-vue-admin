// 自动生成模板NestInfo
package NestInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// NestInfo 结构体
type NestInfo struct {
	global.GVA_MODEL
	Nestid               string `json:"nestid" form:"nestid" gorm:"column:nestid;comment:;"`
	NestName             string `json:"nestName" form:"nestName" gorm:"column:nestname;comment:;"`
	NestLocation         string `json:"nestLocation" form:"nestLocation" gorm:"column:nest_location;comment:;"`
	AircraftVideoPushURL string `json:"aircraftVideoPushURL" form:"aircraftVideoPushURL" gorm:"column:aircraft_video_push_url;comment:;"`
	AircraftVideoURL     string `json:"aircraftVideoURL" form:"aircraftVideoURL" gorm:"column:aircraft_video_url;comment:;"`
	NestVideoURL         string `json:"nestVideoURL" form:"nestVideoURL" gorm:"column:nest_video_url;comment:;"`
	CreatedBy            uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy            uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy            uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName NestInfo 表名
func (NestInfo) TableName() string {
	return "nest_info"
}
