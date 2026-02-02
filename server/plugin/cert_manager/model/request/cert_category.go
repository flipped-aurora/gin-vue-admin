package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type CertCategorySearch struct {
	Name string `json:"name" form:"name"`
	request.PageInfo
}

type BatchUpdateCategory struct {
	IDs      []uint `json:"ids" binding:"required"`
	Category string `json:"category" binding:"required"`
}
