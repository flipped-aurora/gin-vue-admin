package response

import "github.com/eyotang/gin-vue-admin/server/model/request"

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
