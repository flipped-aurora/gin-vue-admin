package request

type SysAuthorityBtnReq struct {
	MenuID      uint   `json:"menuID"`
	AuthorityId uint   `json:"authorityId"`
	Selected    []uint `json:"selected"`
}
