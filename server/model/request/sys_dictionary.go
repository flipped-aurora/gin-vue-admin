package request

import "github.com/eyotang/gin-vue-admin/server/model"

type SysDictionarySearch struct {
	model.SysDictionary
	PageInfo
}
