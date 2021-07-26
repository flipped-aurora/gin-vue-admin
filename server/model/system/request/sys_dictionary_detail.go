package request

import (
	"gin-vue-admin/model/common/request"
	"gin-vue-admin/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
