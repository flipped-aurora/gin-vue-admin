package response

import "github.com/eyotang/game-api-admin/server/model/request"

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
