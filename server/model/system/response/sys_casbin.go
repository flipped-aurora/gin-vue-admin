package response

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
)

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
