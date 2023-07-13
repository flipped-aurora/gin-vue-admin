// 自动生成模板NestAirline
package NestAirlinePkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// NestAirline 结构体
type NestAirline struct {
	global.GVA_MODEL
	Missionid             string `json:"missionid" form:"missionid" gorm:"column:missionid;comment:;"`
	Name                  string `json:"name" form:"name" gorm:"column:name;comment:;"`
	Type                  string `json:"type" form:"type" gorm:"column:type;comment:;"`
	AutoFlightSpeed       *int   `json:"autoFlightSpeed" form:"autoFlightSpeed" gorm:"column:auto_flight_speed;comment:;"`
	GotoFirstWaypointMode *int   `json:"gotoFirstWaypointMode" form:"gotoFirstWaypointMode" gorm:"column:goto_first_waypoint_mode;comment:;"`
	FinishAction          *int   `json:"finishAction" form:"finishAction" gorm:"column:finish_action;comment:;"`
	FlightPathMode        *int   `json:"flightPathMode" form:"flightPathMode" gorm:"column:flight_path_mode;comment:;"`
	HeadingMode           *int   `json:"headingMode" form:"headingMode" gorm:"column:heading_mode;comment:;"`
	Param                 string `json:"param" form:"param" gorm:"column:param;comment:;type:longtext"`
	Safealt               string `json:"safealt" form:"safealt" gorm:"column:safealt;comment:;"`
	Kml                   string `json:"kml" form:"kml" gorm:"column:kml;comment:;"`
	Gps                   string `json:"gps" form:"gps" gorm:"column:gps;comment:;"`
	Station               string `json:"station" form:"station" gorm:"column:station;comment:;"`
	ClearHomeLocation     string `json:"clearHomeLocation" form:"clearHomeLocation" gorm:"column:clear_home_location;comment:;"`
	Producer              string `json:"producer" form:"producer" gorm:"column:producer;comment:;"`
	ProductionUnit        string `json:"productionUnit" form:"productionUnit" gorm:"column:production_unit;comment:;"`
	IsActive              string `json:"isActive" form:"isActive" gorm:"column:is_active;comment:;"`
	FixedReturnPoint      string `json:"fixedReturnPoint" form:"fixedReturnPoint" gorm:"column:fixed_return_point;comment:;"`
	CreatedBy             uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy             uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy             uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName NestAirline 表名
func (NestAirline) TableName() string {
	return "nest_airline"
}
