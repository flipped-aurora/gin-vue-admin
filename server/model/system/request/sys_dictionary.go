package request

import (
	"github.com/gzpz/golf-sales-system/server/model/common/request"
	"github.com/gzpz/golf-sales-system/server/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}
