package response

import "github.com/eyotang/game-proxy/server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
