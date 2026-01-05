// Package response 提供分页查询的响应结构体
//
// 功能说明：
// 本文件定义了分页查询接口的标准响应格式，包含列表数据、总数、当前页码等信息。
package response

// PageResult 分页查询结果结构体
//
// 功能说明：
// 用于封装分页查询的返回结果，包含数据列表和分页信息。
// 前端可以使用这些信息实现分页组件的显示和交互。
//
// 字段说明：
// - List: 当前页的数据列表，可以是任意类型的数组
// - Total: 总记录数，用于计算总页数
// - Page: 当前页码，从1开始
// - PageSize: 每页显示的记录数
//
// 使用示例：
//
//	// 后端返回
//	list, total, _ := service.GetList(ctx, pageInfo)
//	response.OkWithDetailed(response.PageResult{
//	    List:     list,
//	    Total:    total,
//	    Page:     pageInfo.Page,
//	    PageSize: pageInfo.PageSize,
//	}, "查询成功", c)
//
//	// 前端接收
//	// {
//	//   "code": 0,
//	//   "data": {
//	//     "list": [...],
//	//     "total": 100,
//	//     "page": 1,
//	//     "pageSize": 20
//	//   },
//	//   "msg": "查询成功"
//	// }
type PageResult struct {
	List     interface{} `json:"list"`     // 当前页的数据列表
	Total    int64       `json:"total"`    // 总记录数
	Page     int         `json:"page"`     // 当前页码
	PageSize int         `json:"pageSize"` // 每页大小
}
