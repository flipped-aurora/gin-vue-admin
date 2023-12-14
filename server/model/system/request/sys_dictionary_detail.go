package request

import (
	"kirer.cn/server/model/common/request"
	"kirer.cn/server/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
