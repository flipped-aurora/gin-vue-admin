package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type SysTimedTaskSearch struct {
	request.PageInfo
	Name         string `json:"name" form:"name"`
	ExecutorType string `json:"executorType" form:"executorType"`
	Enabled      *bool  `json:"enabled" form:"enabled"`
}

type SysTimedTaskLogSearch struct {
	request.PageInfo
	TaskId uint   `json:"taskId" form:"taskId"`
	Status string `json:"status" form:"status"`
}

type ToggleTimedTask struct {
	ID      uint `json:"ID" binding:"required"`
	Enabled bool `json:"enabled"`
}

type TriggerTimedTask struct {
	ID uint `json:"ID" binding:"required"`
}
