package request

import (
	"github.com/flipped-aurora/gin-vue-admin/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
