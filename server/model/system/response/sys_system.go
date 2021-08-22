package response

import "github.com/flipped-aurora/gin-vue-admin/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
