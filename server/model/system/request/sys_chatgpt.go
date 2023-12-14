package request

import (
	"kirer.cn/server/model/common/request"
	"kirer.cn/server/model/system"
)

type ChatGptRequest struct {
	system.ChatGpt
	request.PageInfo
}
