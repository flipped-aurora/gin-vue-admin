package request

import (
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
