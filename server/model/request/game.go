package request

import "github.com/eyotang/game-api-admin/server/model"

type ParamGame struct {
	ID string `uri:"id" binding:"required"`
}

type ParamRequest struct {
	ParamGame
	Name string `uri:"request" binding:"required"`
}

type HeaderRequest struct {
	model.ConnectionToken
}

type CreateConnection struct {
	Host string `json:"host" form:"host" binding:"required"`
	Port uint32 `json:"port" form:"port" binding:"required"`
}

type GameRequest struct {
	Data interface{} `json:"data" binding:"required"`
}

type GameResponse struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}
