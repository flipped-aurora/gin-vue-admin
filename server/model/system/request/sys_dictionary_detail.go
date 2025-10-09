package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
	ParentID *uint `json:"parentID" form:"parentID"` // 父级字典详情ID，用于查询指定父级下的子项
	Level    *int  `json:"level" form:"level"`       // 层级深度，用于查询指定层级的数据
}

// CreateSysDictionaryDetailRequest 创建字典详情请求
type CreateSysDictionaryDetailRequest struct {
	Label           string `json:"label" form:"label" binding:"required"`                     // 展示值
	Value           string `json:"value" form:"value" binding:"required"`                     // 字典值
	Extend          string `json:"extend" form:"extend"`                                      // 扩展值
	Status          *bool  `json:"status" form:"status"`                                      // 启用状态
	Sort            int    `json:"sort" form:"sort"`                                          // 排序标记
	SysDictionaryID int    `json:"sysDictionaryID" form:"sysDictionaryID" binding:"required"` // 关联标记
	ParentID        *uint  `json:"parentID" form:"parentID"`                                  // 父级字典详情ID
}

// UpdateSysDictionaryDetailRequest 更新字典详情请求
type UpdateSysDictionaryDetailRequest struct {
	ID              uint   `json:"ID" form:"ID" binding:"required"`                           // 主键ID
	Label           string `json:"label" form:"label" binding:"required"`                     // 展示值
	Value           string `json:"value" form:"value" binding:"required"`                     // 字典值
	Extend          string `json:"extend" form:"extend"`                                      // 扩展值
	Status          *bool  `json:"status" form:"status"`                                      // 启用状态
	Sort            int    `json:"sort" form:"sort"`                                          // 排序标记
	SysDictionaryID int    `json:"sysDictionaryID" form:"sysDictionaryID" binding:"required"` // 关联标记
	ParentID        *uint  `json:"parentID" form:"parentID"`                                  // 父级字典详情ID
}

// GetDictionaryDetailsByParentRequest 根据父级ID获取字典详情请求
type GetDictionaryDetailsByParentRequest struct {
	SysDictionaryID int   `json:"sysDictionaryID" form:"sysDictionaryID" binding:"required"` // 字典ID
	ParentID        *uint `json:"parentID" form:"parentID"`                                  // 父级字典详情ID，为空时获取顶级
	IncludeChildren bool  `json:"includeChildren" form:"includeChildren"`                    // 是否包含子级数据
}
