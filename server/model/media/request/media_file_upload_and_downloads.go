package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AttachmentCategorySearch struct {
	ClassId int `json:"classId" form:"classId"`
	request.PageInfo
}
