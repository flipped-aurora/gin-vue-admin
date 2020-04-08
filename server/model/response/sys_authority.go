package response

import "gin-vue-admin/model"

type SysAuthorityResponse struct {
	Authority model.SysAuthority `json:"authority"`
}
