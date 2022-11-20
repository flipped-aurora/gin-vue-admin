package response

import "github.com/gzpz/golf-sales-system/server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
