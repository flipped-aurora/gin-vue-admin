package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/system"

type SysAuthorityResponse struct {
	Authority system.SysAuthority `json:"authority" comment:"系统角色详情"`
}

type SysAuthorityCopyResponse struct {
	Authority      system.SysAuthority `json:"authority"  comment:"系统角色详情"`
	OldAuthorityId string              `json:"oldAuthorityId"  comment:"旧角色ID"` // 旧角色ID
}
