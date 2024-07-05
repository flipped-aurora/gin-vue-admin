package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/todolist/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/todolist/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/todolist/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var TodoListApi = new(todoListApi)

type todoListApi struct{}

// CreateTodoList
// @Tags      TodoList
// @Summary   创建todoList
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      model.TodoList            true  "创建todoList"
// @Success   200   {object}  response.Response{msg=string}  "创建todoList"
// @Router    /todoList/createTodoList [post]
func (e *todoListApi) CreateTodoList(c *gin.Context) {
	var todoList model.TodoList
	err := c.ShouldBindJSON(&todoList)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = service.TodoListService.CreateTodoList(todoList)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteTodoList
// @Tags      TodoList
// @Summary   删除TodoList
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      model.TodoList            true  "TodoListID"
// @Success   200   {object}  response.Response{msg=string}  "删除TodoList"
// @Router    /customer/customer [delete]
func (e *todoListApi) DeleteTodoList(c *gin.Context) {
	var customer model.TodoList
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = service.TodoListService.DeleteTodoList(customer)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// UpdateTodoList
// @Tags      TodoList
// @Summary   更新TodoList信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      model.TodoList            true  "TodoListID, TodoList信息"
// @Success   200   {object}  response.Response{msg=string}  "更新TodoList信息"
// @Router    /customer/customer [put]
func (e *todoListApi) UpdateTodoList(c *gin.Context) {
	var customer model.TodoList
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = service.TodoListService.UpdateTodoList(&customer)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// GetTodoList
// @Tags      TodoList
// @Summary   获取单一TodoList信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     model.TodoList                                                true  "TodoListID"
// @Success   200   {object}  response.Response{data=exampleRes.TodoListResponse,msg=string}  "获取单一TodoList信息,返回包括TodoList详情"
// @Router    /customer/customer [get]
func (e *todoListApi) GetTodoList(c *gin.Context) {
	var customer model.TodoList
	err := c.ShouldBindQuery(&customer)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	data, err := service.TodoListService.GetTodoList(customer.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(data, "获取成功", c)
}

// GetTodoListList
// @Tags      TodoList
// @Summary   分页获取权限TodoList列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.TodoListSearch                                        true  "页码, 每页大小"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取权限TodoList列表,返回包括列表,总数,页码,每页数量"
// @Router    /customer/customerList [get]
func (e *todoListApi) GetTodoListList(c *gin.Context) {
	var pageInfo request.TodoListSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	customerList, total, err := service.TodoListService.GetCustomerInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     customerList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}
