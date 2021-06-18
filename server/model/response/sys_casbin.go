package response

import "github.com/eyotang/game-proxy/server/model/request"

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
