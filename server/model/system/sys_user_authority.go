package system

type SysUserAuthority struct {
	SysUserId      uint   `gorm:"column:sys_user_id"`
	SysAuthorityId string `gorm:"column:sys_authority_authority_id"`
}

func (s *SysUserAuthority) TableName() string {
	return "sys_user_authority"
}
