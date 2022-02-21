package request

type SysAuthorityBtnReq struct {
	MenuID      uint   `json:"menuID"`
	AuthorityId string `json:"authorityId"`
	Selected    []uint `json:"selected"`
}
