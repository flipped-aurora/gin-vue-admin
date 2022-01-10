package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/system"

type SysAPIResponse struct {
	Api system.SysApi `json:"api" comment:"api详情"`
}

type SysAPIListResponse struct {
	Apis []system.SysApi `json:"apis" comment:"api列表"`
}
