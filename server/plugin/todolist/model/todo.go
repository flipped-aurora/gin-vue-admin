package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

type TodoList struct {
	global.GVA_MODEL
	Prefix      string    `json:"prefix"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	IsDone      bool      `json:"is_done"`
}
