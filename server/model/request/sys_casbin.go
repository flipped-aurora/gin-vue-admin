package request

// Casbin info structure
type CasbinInfo struct {
	Path   string `json:"path"`   // 路径
	Method string `json:"method"` // 方法
}

// Casbin structure for input parameters
type CasbinInReceive struct {
	AuthorityId string       `json:"authorityId"` // 权限id
	CasbinInfos []CasbinInfo `json:"casbinInfos"`
}
