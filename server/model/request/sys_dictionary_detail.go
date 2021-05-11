package request

import "github.com/eyotang/gin-vue-admin/server/model"

type SysDictionaryDetailSearch struct {
	model.SysDictionaryDetail
	PageInfo
}
