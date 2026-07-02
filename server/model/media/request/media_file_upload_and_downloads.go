package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AttachmentCategorySearch struct {
	ClassId        int       `json:"classId" form:"classId"`                // 分类id
	Tag            string    `json:"tag" form:"tag"`                        // 文件类型/标签（扩展名，如 jpg/pdf）
	UserID         uint      `json:"userId" form:"userId"`                  // 上传者ID
	StartCreatedAt time.Time `json:"startCreatedAt" form:"startCreatedAt"`  // 创建时间起（含）
	EndCreatedAt   time.Time `json:"endCreatedAt" form:"endCreatedAt"`      // 创建时间止（不含）
	OrderKey       string    `json:"orderKey" form:"orderKey"`              // 排序字段（如 id/size/created_at）
	Desc           bool      `json:"desc" form:"desc"`                      // 是否倒序
	request.PageInfo
}

// FileBatchDeleteReq 批量删除文件请求
type FileBatchDeleteReq struct {
	IDs []uint `json:"ids" binding:"required"` // 文件记录ID列表
}

// ListOssFilesReq 列举 OSS 存储桶内文件请求
type ListOssFilesReq struct {
	Prefix string `json:"prefix" form:"prefix"` // 对象 key 前缀过滤
	Cursor string `json:"cursor" form:"cursor"` // 分页游标（上次返回的 nextCursor）
	Limit  int    `json:"limit" form:"limit"`   // 每页数量（<=0 默认 100）
}
