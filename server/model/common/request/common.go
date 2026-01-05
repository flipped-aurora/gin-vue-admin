// Package request 提供通用的请求参数结构体定义
// 这些结构体用于统一处理API请求中的常见参数，如分页、ID查询等
//
// 功能说明：
// - PageInfo: 分页查询通用参数结构
// - GetById: 根据ID查询的通用参数结构
// - IdsReq: 批量ID查询的通用参数结构
// - GetAuthorityId: 根据角色ID查询的通用参数结构
// - Empty: 空请求体结构
//
// 使用场景：
// 这些结构体在整个项目的API层广泛使用，提供统一的参数验证和处理机制
package request

import (
	"gorm.io/gorm"
)

// PageInfo 分页查询通用输入参数结构体
//
// 功能说明：
// 用于接收前端传递的分页查询参数，支持页码、每页大小和关键字搜索
//
// 字段说明：
// - Page: 当前页码，从1开始计数
// - PageSize: 每页显示的记录数，最大限制为100条
// - Keyword: 搜索关键字，用于模糊查询
//
// 使用示例：
//
//	// 前端请求示例
//	// GET /api/v1/users?page=1&pageSize=20&keyword=张三
//
//	// 后端使用示例
//	var pageInfo request.PageInfo
//	if err := c.ShouldBindQuery(&pageInfo); err != nil {
//	    // 处理错误
//	}
//	// 在GORM查询中使用
//	db.Scopes(pageInfo.Paginate()).Find(&users)
type PageInfo struct {
	Page     int    `json:"page" form:"page"`         // 页码，从1开始，默认为1
	PageSize int    `json:"pageSize" form:"pageSize"` // 每页大小，范围1-100，默认为10
	Keyword  string `json:"keyword" form:"keyword"`   // 关键字，用于搜索过滤
}

// Paginate 返回一个GORM作用域函数，用于实现分页查询
//
// 功能说明：
// 该方法返回一个闭包函数，该函数可以应用到GORM查询链中，自动处理分页逻辑
//
// 实现细节：
// 1. 自动校验和修正页码：如果页码小于等于0，则设置为1
// 2. 自动校验和修正每页大小：
//   - 如果大于100，则限制为100（防止单次查询数据量过大）
//   - 如果小于等于0，则设置为默认值10
//
// 3. 计算偏移量：offset = (page - 1) * pageSize
// 4. 应用GORM的Offset和Limit方法
//
// 返回值：
// 返回一个函数，该函数接收*gorm.DB并返回应用了分页的*gorm.DB
//
// 使用示例：
//
//	var users []User
//	var pageInfo request.PageInfo
//	pageInfo.Page = 2
//	pageInfo.PageSize = 20
//
//	// 方式1：直接使用
//	db.Scopes(pageInfo.Paginate()).Find(&users)
//
//	// 方式2：结合其他查询条件
//	db.Where("status = ?", "active").
//	   Scopes(pageInfo.Paginate()).
//	   Order("created_at desc").
//	   Find(&users)
//
// 注意事项：
// - 该方法不会执行查询，只是构建查询链
// - 需要在最后调用Find、First等方法才会真正执行查询
func (r *PageInfo) Paginate() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// 页码校验：确保页码至少为1
		if r.Page <= 0 {
			r.Page = 1
		}
		// 每页大小校验和限制
		switch {
		case r.PageSize > 100:
			// 防止单次查询数据量过大，限制最大为100条
			r.PageSize = 100
		case r.PageSize <= 0:
			// 如果未指定或指定为0，使用默认值10
			r.PageSize = 10
		}
		// 计算偏移量：例如第2页，每页20条，则offset = (2-1) * 20 = 20
		offset := (r.Page - 1) * r.PageSize
		// 应用分页：跳过offset条记录，限制返回pageSize条记录
		return db.Offset(offset).Limit(r.PageSize)
	}
}

// GetById 根据ID查询的通用参数结构体
//
// 功能说明：
// 用于接收前端传递的单个ID参数，用于查询、删除、更新等操作
//
// 字段说明：
// - ID: 记录的主键ID，支持JSON和表单两种绑定方式
//
// 使用示例：
//
//	// 前端请求示例
//	// GET /api/v1/user/123
//	// 或 POST /api/v1/user/delete
//	// Body: {"id": 123}
//
//	// 后端使用示例
//	var req request.GetById
//	if err := c.ShouldBindJSON(&req); err != nil {
//	    // 处理错误
//	}
//	userID := req.Uint() // 转换为uint类型用于数据库查询
//	db.Where("id = ?", userID).First(&user)
type GetById struct {
	ID int `json:"id" form:"id"` // 主键ID，支持JSON和表单绑定
}

// Uint 将ID转换为uint类型
//
// 功能说明：
// 由于数据库主键通常使用uint类型，该方法提供类型转换功能
//
// 返回值：
// 返回uint类型的ID值
//
// 使用示例：
//
//	var req request.GetById
//	req.ID = 123
//	userID := req.Uint() // 返回 uint(123)
//	db.Where("id = ?", userID).First(&user)
func (r *GetById) Uint() uint {
	return uint(r.ID)
}

// IdsReq 批量ID查询的通用参数结构体
//
// 功能说明：
// 用于接收前端传递的多个ID参数，用于批量查询、批量删除等操作
//
// 字段说明：
// - Ids: ID数组，支持JSON和表单两种绑定方式
//
// 使用示例：
//
//	// 前端请求示例
//	// POST /api/v1/users/batch-delete
//	// Body: {"ids": [1, 2, 3, 4, 5]}
//
//	// 后端使用示例
//	var req request.IdsReq
//	if err := c.ShouldBindJSON(&req); err != nil {
//	    // 处理错误
//	}
//	// 批量删除
//	db.Where("id IN ?", req.Ids).Delete(&User{})
//
//	// 批量查询
//	var users []User
//	db.Where("id IN ?", req.Ids).Find(&users)
type IdsReq struct {
	Ids []int `json:"ids" form:"ids"` // ID数组，支持批量操作
}

// GetAuthorityId 根据角色ID查询的通用参数结构体
//
// 功能说明：
// 专门用于权限管理模块，接收角色ID参数进行角色相关的查询操作
//
// 字段说明：
// - AuthorityId: 角色/权限ID，使用uint类型，支持JSON和表单绑定
//
// 使用示例：
//
//	// 前端请求示例
//	// GET /api/v1/authority/1
//	// 或 POST /api/v1/authority/users
//	// Body: {"authorityId": 1}
//
//	// 后端使用示例
//	var req request.GetAuthorityId
//	if err := c.ShouldBindJSON(&req); err != nil {
//	    // 处理错误
//	}
//	// 查询该角色下的所有用户
//	var users []User
//	db.Where("authority_id = ?", req.AuthorityId).Find(&users)
type GetAuthorityId struct {
	AuthorityId uint `json:"authorityId" form:"authorityId"` // 角色ID，用于权限相关查询
}

// Empty 空请求体结构
//
// 功能说明：
// 用于不需要任何参数的API接口，作为占位符使用
//
// 使用场景：
// - 某些接口不需要接收任何参数，但为了保持代码一致性使用此结构
// - 可以用于表示"无参数"的语义
//
// 使用示例：
//
//	// 后端使用示例
//	func SomeAPI(c *gin.Context) {
//	    var req request.Empty
//	    // 不需要绑定参数，直接处理业务逻辑
//	    // ...
//	}
type Empty struct{}
