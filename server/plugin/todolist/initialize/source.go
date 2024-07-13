package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
)

// 自动注册基础数据
func RegisterBaseData() {
	utils.RegisterMenus(
		system.SysBaseMenu{
			Path:      "todolist",
			Name:      "todolistGroup",
			Hidden:    false,
			Component: "view/routerHolder.vue",
			Sort:      1000,
			Meta: system.Meta{
				Title: "TodoList",
				Icon:  "school",
			},
		},
		system.SysBaseMenu{
			Path:      "todolist",
			Name:      "todolist",
			Hidden:    false,
			Component: "plugin/todolist/view/todolist.vue",
			Sort:      0,
			Meta: system.Meta{
				Title: "TodoList",
				Icon:  "school",
			},
		},
	)
	utils.RegisterApis(
		system.SysApi{
			Path:        "/todolist/createTodoList",
			Description: "创建TodoList",
			ApiGroup:    "TodoList",
			Method:      "POST",
		},
		system.SysApi{
			Path:        "/todolist/updateTodoList",
			Description: "更新TodoList",
			ApiGroup:    "TodoList",
			Method:      "DELETE",
		},
		system.SysApi{
			Path:        "/todolist/deleteTodoList",
			Description: "批量删除TodoList",
			ApiGroup:    "TodoList",
			Method:      "DELETE",
		},
		system.SysApi{
			Path:        "/todolist/getTodoList",
			Description: "获取TodoList",
			ApiGroup:    "TodoList",
			Method:      "PUT",
		},
		system.SysApi{
			Path:        "/todolist/getTodoListList",
			Description: "获取TodoList列表",
			ApiGroup:    "TodoList",
			Method:      "POST",
		},
	)
}
