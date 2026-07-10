package system

// SysUserPosition 是 sysUser 和 sysPosition 的连接表(用户多岗位)
type SysUserPosition struct {
	SysUserId     uint `gorm:"column:sys_user_id"`
	SysPositionId uint `gorm:"column:sys_position_id"`
}

func (s *SysUserPosition) TableName() string {
	return "sys_user_positions"
}
