package response

import "gin-vue-admin/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
