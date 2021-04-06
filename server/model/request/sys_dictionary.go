package request

import "gin-vue-admin/model"

type SysDictionarySearch struct {
	model.SysDictionary
	PageInfo
}
