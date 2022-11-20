package response

import (
	"github.com/gzpz/golf-sales-system/server/model/system/request"
)

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
