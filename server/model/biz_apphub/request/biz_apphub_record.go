package request

import "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"

type BizAppHubRecord struct {
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}
