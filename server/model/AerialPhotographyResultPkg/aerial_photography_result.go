// 自动生成模板AerialPhotographyResult
package AerialPhotographyResultPkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/datatypes"
	"time"
)

// AerialPhotographyResult 结构体
type AerialPhotographyResult struct {
	global.GVA_MODEL
	Name                  string         `json:"name" form:"name" gorm:"column:name;comment:;"`
	PhotographyCreatetime *time.Time     `json:"photographyCreatetime" form:"photographyCreatetime" gorm:"column:photography_createtime;comment:;"`
	UploadBy              string         `json:"uploadBy" form:"uploadBy" gorm:"column:upload_by;comment:;"`
	Type                  *int           `json:"type" form:"type" gorm:"column:type;comment:0-高清正射 1-三维模型;"`
	AerialPhotographyFile datatypes.JSON `json:"aerialPhotographyFile" form:"aerialPhotographyFile" gorm:"column:aerial_photography_file;comment:;"`
	Status                *int           `json:"status" form:"status" gorm:"column:status;comment:0-上传中 1-解压中 2-已完成 3-异常;"`
	Position              string         `json:"position" form:"position" gorm:"column:position;comment:;"`
	LoadOrNot             *int           `json:"loadOrNot" form:"loadOrNot" gorm:"column:load_or_not;comment:0-加载 1-不加载;"`
	NestIds               string         `json:"nestIds" form:"nestIds" gorm:"column:nest_ids;comment:机巢id集;"`
	CreatedBy             uint           `gorm:"column:created_by;comment:创建者"`
	UpdatedBy             uint           `gorm:"column:updated_by;comment:更新者"`
	DeletedBy             uint           `gorm:"column:deleted_by;comment:删除者"`

	FileUrl *string `gorm:"-"`
}

// TableName AerialPhotographyResult 表名
func (AerialPhotographyResult) TableName() string {
	return "aerial_photography_result"
}
