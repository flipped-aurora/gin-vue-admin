package request

// Casbin info structure
type CasbinInfo struct {
	Path   string `json:"path"`
	Method string `json:"method"`
}

// Casbin structure for input parameters
type CasbinInReceive struct {
	AuthorityId string       `json:"authorityId"`
	CasbinInfos []CasbinInfo `json:"casbinInfos"`
}
