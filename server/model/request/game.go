package request

import "gin-vue-admin/model"

type ParamGame struct {
	ID string `uri:"id" binding:"required"`
}

type ParamRequest struct {
	ParamGame
	Name string `uri:"request" binding:"required"`
}

type CreateConnection struct {
	Host string `json:"host" form:"host" binding:"required"`
	Port uint32 `json:"port" form:"port" binding:"required"`
}

type CloseConnection struct {
	model.ConnectionToken
}

type GameRequest struct {
	model.ConnectionToken
	Data interface{} `json:"data" binding:"required"`
}
