package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"

type CroppingRecordPageResult struct {
	response.PageResult
	Quantity int64 `json:"quantity"`
	Margin   int64 `json:"margin"`
}
