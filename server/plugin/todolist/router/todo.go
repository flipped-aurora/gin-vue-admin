package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/todolist/api"
	"github.com/gin-gonic/gin"
)

var TodoListRouter = new(todoListRouter)

type todoListRouter struct{}

func (s *todoListRouter) InitRouter(router *gin.RouterGroup, routerPub *gin.RouterGroup) {
	tRouter := router.Group("todolist").Use(middleware.OperationRecord())
	tRouterWithoutAuth := routerPub.Group("todolist")
	todoListApi := api.TodoListApi
	{
		tRouter.POST("createTodoList", todoListApi.CreateTodoList)   // 创建todoList
		tRouter.PUT("updateTodoList", todoListApi.UpdateTodoList)    // 更新todoList
		tRouter.DELETE("deleteTodoList", todoListApi.DeleteTodoList) // 删除todoList
	}
	{
		tRouterWithoutAuth.GET("getTodoList", todoListApi.GetTodoList)         // 获取单一todoList信息
		tRouterWithoutAuth.GET("getTodoListList", todoListApi.GetTodoListList) // 获取todoList列表
	}
}
