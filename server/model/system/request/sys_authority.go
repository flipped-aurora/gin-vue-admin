package request

import "github.com/flipped-aurora/gin-vue-admin/server/model/system"

type AuthorityCopy struct {
	OldAuthorityId uint                `json:"oldAuthorityId"`
	Authority      system.SysAuthority `json:"authority"`
}

func (r *AuthorityCopy) Create() system.SysAuthority {
	return system.SysAuthority{
		AuthorityId:   r.Authority.AuthorityId,
		AuthorityName: r.Authority.AuthorityName,
		DefaultRouter: r.Authority.DefaultRouter,
	}
}
