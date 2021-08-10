package response

import "gin-vue-admin/model/system"

type SysAuthorityResponse struct {
	Authority system.SysAuthority `json:"authority"`
}

type SysAuthorityCopyResponse struct {
	Authority      system.SysAuthority `json:"authority"`
	OldAuthorityId string              `json:"oldAuthorityId"` // 旧角色ID
}
