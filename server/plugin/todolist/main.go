package todolist

import (
	gvaGlobal "github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/todolist/config"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/todolist/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/todolist/initialize"
	"github.com/gin-gonic/gin"
)

type TodoPlugin struct{}

func CreateTodoPlug(prefix string) *TodoPlugin {
	global.GlobalConfig = &config.TodoList{
		Prefix: prefix,
	}
	return &TodoPlugin{}
}

func (*TodoPlugin) Register(group *gin.Engine) {
	if gvaGlobal.GVA_DB == nil {
		return
	}
	initialize.RegisterTables()
	initialize.RegisterRouters(group)
	initialize.RegisterBaseData()
}
