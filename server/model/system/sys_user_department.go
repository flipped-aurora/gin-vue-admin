package system

// SysUserDepartment 是 sysUser 和 sysDepartment 的连接表(用户多部门归属)
// 由 GORM 依据 SysUser.Departments 的 many2many tag 自动建表,此显式 struct 仅供 service 层直接操作 join 表
type SysUserDepartment struct {
	SysUserId       uint `gorm:"column:sys_user_id"`
	SysDepartmentId uint `gorm:"column:sys_department_id"`
}

func (s *SysUserDepartment) TableName() string {
	return "sys_user_departments"
}
