package request

// SysDepartmentSearch 部门查询(树形返回,默认不分页)
type SysDepartmentSearch struct {
	Name string `json:"name" form:"name"` // 按部门名称模糊查询
}

// SetDepartmentUsers 全量覆盖部门成员(反向分配)
type SetDepartmentUsers struct {
	DeptId  uint   `json:"deptId" form:"deptId"`   // 部门ID
	UserIds []uint `json:"userIds" form:"userIds"` // 用户ID列表
}
