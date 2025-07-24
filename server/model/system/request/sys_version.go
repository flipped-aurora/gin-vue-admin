package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type SysVersionSearch struct {
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	VersionName    *string     `json:"versionName" form:"versionName"`
	VersionCode    *string     `json:"versionCode" form:"versionCode"`
	request.PageInfo
}

// ExportVersionRequest 导出版本请求结构体
type ExportVersionRequest struct {
	VersionName string `json:"versionName" binding:"required"` // 版本名称
	VersionCode string `json:"versionCode" binding:"required"` // 版本号
	Description string `json:"description"`                    // 版本描述
	MenuIds     []uint `json:"menuIds"`                        // 选中的菜单ID列表
	ApiIds      []uint `json:"apiIds"`                         // 选中的API ID列表
}
