package system

// SysAuthorityDepartment 角色与部门的连接表(数据权限第5档"自定义部门集"的角色配置)
type SysAuthorityDepartment struct {
	SysAuthorityAuthorityId uint `gorm:"column:sys_authority_authority_id;index"`
	SysDepartmentId         uint `gorm:"column:sys_department_id"`
}

func (s *SysAuthorityDepartment) TableName() string {
	return "sys_authority_departments"
}
