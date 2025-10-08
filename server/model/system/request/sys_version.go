package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
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
	DictIds     []uint `json:"dictIds"`                        // 选中的字典ID列表
}

// ImportVersionRequest 导入版本请求结构体
type ImportVersionRequest struct {
	VersionInfo      VersionInfo            `json:"version" binding:"required"` // 版本信息
	ExportMenu       []system.SysBaseMenu   `json:"menus"`                      // 菜单数据，直接复用SysBaseMenu
	ExportApi        []system.SysApi        `json:"apis"`                       // API数据，直接复用SysApi
	ExportDictionary []system.SysDictionary `json:"dictionaries"`               // 字典数据，直接复用SysDictionary
}

// VersionInfo 版本信息结构体
type VersionInfo struct {
	Name        string `json:"name" binding:"required"`        // 版本名称
	Code        string `json:"code" binding:"required"`        // 版本号
	Description string `json:"description"`                    // 版本描述
	ExportTime  string `json:"exportTime"`                     // 导出时间
}
