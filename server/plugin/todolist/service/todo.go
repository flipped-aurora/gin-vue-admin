package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/todolist/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/todolist/model/request"
)

var TodoListService = new(todoListService)

type todoListService struct{}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateTodoList
//@description: 创建TodoList

func (todoList *todoListService) CreateTodoList(e model.TodoList) (err error) {
	err = global.GVA_DB.Create(&e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteFileChunk
//@description: 删除TodoList

func (todoList *todoListService) DeleteTodoList(e model.TodoList) (err error) {
	err = global.GVA_DB.Delete(&e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateTodoList
//@description: 更新TodoList

func (todoList *todoListService) UpdateTodoList(e *model.TodoList) (err error) {
	err = global.GVA_DB.Save(e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetTodoList
//@description: 获取TodoList信息

func (todoList *todoListService) GetTodoList(id uint) (customer model.TodoList, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&customer).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetCustomerInfoList
//@description: 分页获取TodoList列表

func (todoList *todoListService) GetCustomerInfoList(info request.TodoListSearch) (list []model.TodoList, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.TodoList{})
	var CustomerList []model.TodoList
	err = db.Limit(limit).Offset(offset).Find(&CustomerList).Error
	return CustomerList, total, err
}
