package request

import "gin-vue-admin/model"

type ParamGame struct {
	ID string `uri:"id" binding:"required"`
}

type CreateConnection struct {
	ParamGame
	Host string `json:"host" form:"host" binding:"required"`
	Port uint   `json:"port" form:"port" binding:"required"`
}

type CloseConnection struct {
	ParamGame
	model.ConnectionToken
}

type GameRequest struct {
	ParamGame
	model.ConnectionToken
	Data []byte `json:"data" binding:"required"`
}
