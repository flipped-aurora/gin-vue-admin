package request

import (
	"github.com/gzpz/golf-sales-system/server/model/common/request"
	"github.com/gzpz/golf-sales-system/server/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
