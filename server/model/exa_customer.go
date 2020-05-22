package model

import (
	"github.com/jinzhu/gorm"
)

type ExaCustomer struct {
	gorm.Model
	CustomerName       string  `json:"customerName" form:"customerName"`
	CustomerPhoneData  string  `json:"customerPhoneData" form:"customerPhoneData"`
	SysUserID          uint    `json:"sysUserId" form:"sysUserId"`
	SysUserAuthorityID string  `json:"sysUserAuthorityID" form:"sysUserAuthorityID"`
	SysUser            SysUser `json:"sysUser" form:"sysUser"`
}
