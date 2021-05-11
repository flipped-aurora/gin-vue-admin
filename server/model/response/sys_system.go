package response

import "github.com/eyotang/game-api-admin/server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
