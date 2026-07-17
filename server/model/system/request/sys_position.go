package request

import (
	common "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

// SysPositionSearch 岗位分页查询
type SysPositionSearch struct {
	common.PageInfo
	Name   string `json:"name" form:"name"`     // 岗位名称
	Code   string `json:"code" form:"code"`     // 岗位编码
	Status *bool  `json:"status" form:"status"` // 是否启用
}

// SetPositionUsers 全量覆盖岗位成员(反向分配)
type SetPositionUsers struct {
	PositionId uint   `json:"positionId" form:"positionId"` // 岗位ID
	UserIds    []uint `json:"userIds" form:"userIds"`       // 用户ID列表
}
