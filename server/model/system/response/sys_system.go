package response

import "kirer.cn/server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
