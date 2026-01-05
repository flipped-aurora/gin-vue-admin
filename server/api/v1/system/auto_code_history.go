// Package system 提供系统模块的API控制器
//
// 本文件实现了自动代码生成历史记录的API接口，包括：
// - First: 获取历史记录的元数据信息
// - Delete: 删除历史记录
// - RollBack: 回滚代码生成操作
// - GetList: 分页查询历史记录列表
//
// 功能说明：
// API层负责处理HTTP请求，包括参数绑定、参数验证、调用服务层、返回响应等。
// 本文件遵循gin-vue-admin的分层架构：API -> Service -> Model。
package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	common "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	request "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// AutoCodeHistoryApi 自动代码生成历史记录API控制器
//
// 功能说明：
// 封装了所有与代码生成历史记录相关的HTTP接口处理逻辑。
//
// 设计模式：
// 采用空结构体实现，所有方法都是值接收者，避免不必要的内存分配。
type AutoCodeHistoryApi struct{}

// First 获取代码生成历史记录的元数据信息
//
// 功能说明：
// 根据历史记录ID查询该次代码生成操作的原始请求参数（JSON格式），
// 用于重新生成代码或查看历史配置。
//
// 请求方式：POST
// 路由：/autoCode/getMeta
// 权限：需要ApiKeyAuth认证
//
// 请求参数：
//
//	{
//	  "id": 1  // 历史记录ID
//	}
//
// 响应示例：
//
//	{
//	  "code": 0,
//	  "data": {
//	    "meta": "{\"table\":\"sys_user\",\"structName\":\"SysUser\",...}"
//	  },
//	  "msg": "获取成功"
//	}
//
// 使用场景：
// - 查看历史代码生成的配置信息
// - 基于历史配置重新生成代码
//
// @Tags      AutoCode
// @Summary   获取meta信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById                                            true  "请求参数"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "获取meta信息"
// @Router    /autoCode/getMeta [post]
func (a *AutoCodeHistoryApi) First(c *gin.Context) {
	// 步骤1: 绑定请求参数
	var info common.GetById
	err := c.ShouldBindJSON(&info)
	if err != nil {
		// 参数绑定失败，返回错误信息
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 步骤2: 调用服务层获取数据
	data, err := autoCodeHistoryService.First(c.Request.Context(), info)
	if err != nil {
		// 服务层返回错误，返回错误信息
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 步骤3: 返回成功响应
	response.OkWithDetailed(gin.H{"meta": data}, "获取成功", c)
}

// Delete 删除代码生成历史记录
//
// 功能说明：
// 从数据库中永久删除指定的代码生成历史记录。
// 注意：此操作只删除数据库记录，不会删除已生成的文件或回滚代码。
// 如果需要撤销代码生成，应该使用RollBack接口。
//
// 请求方式：POST
// 路由：/autoCode/delSysHistory
// 权限：需要ApiKeyAuth认证
//
// 请求参数：
//
//	{
//	  "id": 1  // 历史记录ID
//	}
//
// 响应示例：
//
//	{
//	  "code": 0,
//	  "data": {},
//	  "msg": "删除成功"
//	}
//
// 注意事项：
// - 删除操作不可恢复
// - 建议先确认记录不再需要后再删除
//
// @Tags      AutoCode
// @Summary   删除回滚记录
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById                true  "请求参数"
// @Success   200   {object}  response.Response{msg=string}  "删除回滚记录"
// @Router    /autoCode/delSysHistory [post]
func (a *AutoCodeHistoryApi) Delete(c *gin.Context) {
	// 步骤1: 绑定请求参数
	var info common.GetById
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 步骤2: 调用服务层删除记录
	err = autoCodeHistoryService.Delete(c.Request.Context(), info)
	if err != nil {
		// 删除失败，记录日志并返回错误
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}

	// 步骤3: 返回成功响应
	response.OkWithMessage("删除成功", c)
}

// RollBack 回滚代码生成操作
//
// 功能说明：
// 这是代码生成回滚的核心接口，用于撤销一次代码生成操作。
// 回滚操作包括：
// 1. 删除导出的模板（如果存在）
// 2. 删除生成的API接口（可选）
// 3. 删除生成的菜单（可选）
// 4. 删除数据库表（可选）
// 5. 回滚注入到现有文件的代码
// 6. 移动或删除生成的文件
// 7. 标记历史记录为已回滚状态
//
// 请求方式：POST
// 路由：/autoCode/rollback
// 权限：需要ApiKeyAuth认证
//
// 请求参数：
//
//	{
//	  "id": 1,           // 历史记录ID
//	  "deleteApi": true, // 是否删除API接口
//	  "deleteMenu": true,// 是否删除菜单
//	  "deleteTable": false // 是否删除数据表
//	}
//
// 响应示例：
//
//	{
//	  "code": 0,
//	  "data": {},
//	  "msg": "回滚成功"
//	}
//
// 使用场景：
// - 代码生成后发现配置错误，需要撤销
// - 测试代码生成功能后清理
// - 重新生成代码前先回滚旧代码
//
// 注意事项：
// - 回滚操作会删除生成的文件，请谨慎操作
// - 如果deleteTable为true，数据表将被删除，数据会丢失
//
// @Tags      AutoCode
// @Summary   回滚自动生成代码
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.SysAutoHistoryRollBack             true  "请求参数"
// @Success   200   {object}  response.Response{msg=string}  "回滚自动生成代码"
// @Router    /autoCode/rollback [post]
func (a *AutoCodeHistoryApi) RollBack(c *gin.Context) {
	// 步骤1: 绑定请求参数
	var info request.SysAutoHistoryRollBack
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 步骤2: 调用服务层执行回滚
	err = autoCodeHistoryService.RollBack(c.Request.Context(), info)
	if err != nil {
		// 回滚失败，返回错误信息
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 步骤3: 返回成功响应
	response.OkWithMessage("回滚成功", c)
}

// GetList 分页查询代码生成历史记录列表
//
// 功能说明：
// 分页查询所有代码生成历史记录，支持分页参数，按更新时间倒序排列。
// 用于在管理界面展示历史记录列表，支持分页浏览。
//
// 请求方式：POST
// 路由：/autoCode/getSysHistory
// 权限：需要ApiKeyAuth认证
//
// 请求参数：
//
//	{
//	  "page": 1,      // 页码，从1开始
//	  "pageSize": 20, // 每页大小，最大100
//	  "keyword": ""   // 搜索关键字（可选）
//	}
//
// 响应示例：
//
//	{
//	  "code": 0,
//	  "data": {
//	    "list": [
//	      {
//	        "id": 1,
//	        "table": "sys_user",
//	        "structName": "SysUser",
//	        "createdAt": "2024-01-01T00:00:00Z",
//	        ...
//	      }
//	    ],
//	    "total": 100,
//	    "page": 1,
//	    "pageSize": 20
//	  },
//	  "msg": "获取成功"
//	}
//
// 使用场景：
// - 代码生成历史记录管理页面
// - 查看历史生成记录
// - 选择记录进行回滚操作
//
// @Tags      AutoCode
// @Summary   查询回滚记录
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      common.PageInfo                                true  "请求参数"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "查询回滚记录,返回包括列表,总数,页码,每页数量"
// @Router    /autoCode/getSysHistory [post]
func (a *AutoCodeHistoryApi) GetList(c *gin.Context) {
	// 步骤1: 绑定请求参数
	var info common.PageInfo
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 步骤2: 调用服务层获取列表数据
	list, total, err := autoCodeHistoryService.GetList(c.Request.Context(), info)
	if err != nil {
		// 查询失败，记录日志并返回错误
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	// 步骤3: 构建分页结果并返回
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     info.Page,
		PageSize: info.PageSize,
	}, "获取成功", c)
}
