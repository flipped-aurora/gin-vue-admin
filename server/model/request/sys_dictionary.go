package request

import "github.com/eyotang/game-proxy/server/model"

type SysDictionarySearch struct {
	model.SysDictionary
	PageInfo
}
