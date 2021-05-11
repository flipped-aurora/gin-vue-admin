package request

import "github.com/eyotang/game-api-admin/server/model"

type SysDictionarySearch struct {
	model.SysDictionary
	PageInfo
}
