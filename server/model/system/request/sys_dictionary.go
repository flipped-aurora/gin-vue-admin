package request

import (
	"gin-vue-admin/model/common/request"
	"gin-vue-admin/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}
