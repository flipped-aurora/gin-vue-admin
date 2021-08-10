package system

type SysUseAuthority struct {
	SysUserId               uint   `gorm:"column:sys_user_id"`
	SysAuthorityAuthorityId string `gorm:"column:sys_authority_authority_id"`
}

func (s *SysUseAuthority) TableName() string {
	return "sys_user_authority"
}
